package injection

import (
	"context"
	"encoding/base64"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mydisha/keirouter/backend/internal/guardrails"
)

// Detector implements guardrails.Detector for prompt injection.
type Detector struct{}

// New returns the prompt-injection detector.
func New() *Detector { return &Detector{} }

// Name identifies this detector in policy config and audit logs.
func (Detector) Name() string { return "injection" }

// Inbound scans the user messages (and, when enabled, tool results) for
// injection. The operator's system prompt and earlier assistant turns are
// trusted and never scanned. Hits below the severity threshold produce a
// log-only decision so tuning stays visible in the audit log.
func (Detector) Inbound(_ context.Context, in *guardrails.InboundRequest, p guardrails.Policy) (*guardrails.Decision, error) {
	cfg := p.Injection
	if cfg == nil || !cfg.Enabled {
		return nil, nil
	}
	hits := Scan(in.UserText())
	if cfg.ScanToolResults {
		for _, h := range Scan(in.ToolText()) {
			h.Pattern = "indirect_" + h.Pattern
			hits = append(hits, h)
		}
	}
	if len(hits) == 0 {
		return nil, nil
	}
	threshold := cfg.SeverityThreshold
	if threshold == "" {
		threshold = guardrails.DefaultInjectionSeverity
	}
	severity := mapSeverity(Aggregate(hits))
	if !guardrails.MeetsThreshold(severity, threshold) {
		return &guardrails.Decision{
			Action:   guardrails.ActionLogOnly,
			Severity: severity,
			Reason:   reasonFor(hits),
			Findings: toFindings(hits),
		}, nil
	}
	action := cfg.Action
	if action == "" {
		action = guardrails.DefaultInjectionAction
	}
	return &guardrails.Decision{
		Action:   action,
		Severity: severity,
		Reason:   reasonFor(hits),
		Findings: toFindings(hits),
	}, nil
}

// Outbound: injection detection runs only on the prompt side. Returning nil
// is correct: the engine treats nil as ActionAllow with no findings.
func (Detector) Outbound(_ context.Context, _ *guardrails.OutboundResponse, _ guardrails.Policy) (*guardrails.Decision, error) {
	return nil, nil
}

// Aggregate turns individual hits into one severity. Independent weak
// signals reinforce each other: two different medium patterns, or a medium
// plus two lows, read as an attack; three different lows read as medium.
func Aggregate(hits []Hit) Severity {
	seen := map[string]Severity{}
	for _, h := range hits {
		if rank(h.Severity) > rank(seen[h.Pattern]) {
			seen[h.Pattern] = h.Severity
		}
	}
	var high, medium, low int
	for _, s := range seen {
		switch s {
		case SeverityHigh:
			high++
		case SeverityMedium:
			medium++
		case SeverityLow:
			low++
		}
	}
	switch {
	case high > 0, medium >= 2, medium >= 1 && low >= 2:
		return SeverityHigh
	case medium > 0, low >= 3:
		return SeverityMedium
	case low > 0:
		return SeverityLow
	}
	return ""
}

// Scan returns every hit in the text: signatures on the normalized text,
// intent rules on the word view (plus its rot13 and reversed forms), and the
// same checks on decoded base64 payloads. Offsets index into text.
func Scan(text string) []Hit {
	return scan(text, true)
}

func scan(text string, decode bool) []Hit {
	if strings.TrimSpace(text) == "" {
		return nil
	}
	var hits []Hit
	lower := guardrails.Fold(text, guardrails.FoldOptions{Lower: true})
	for _, sig := range signatures {
		if !containsAny(lower.Text, sig.Any) {
			continue
		}
		for _, span := range sig.Regex.FindAllStringIndex(lower.Text, -1) {
			s, e := lower.Span(span[0], span[1])
			if s < 0 {
				continue
			}
			hits = append(hits, Hit{Pattern: sig.Name, Severity: sig.Severity, Start: s, End: e, Text: text[s:e]})
		}
	}

	words := guardrails.Fold(text, guardrails.FoldOptions{Lower: true, Words: true})
	add := func(h Hit, s, e int) {
		os, oe := words.Span(s, e)
		if os < 0 || overlaps(hits, os, oe) {
			return
		}
		h.Start, h.End, h.Text = os, oe, text[os:oe]
		hits = append(hits, h)
	}
	for _, h := range matchIntents(tokenize(words.Text)) {
		add(h, h.Start, h.End)
	}
	for _, h := range matchIntents(tokenize(rot13(words.Text))) {
		h.Pattern = "rot13_" + h.Pattern
		add(h, h.Start, h.End)
	}
	n := len(words.Text)
	for _, h := range matchIntents(tokenize(reverseRunes(words.Text))) {
		h.Pattern = "reversed_" + h.Pattern
		add(h, n-h.End, n-h.Start)
	}

	if decode {
		for _, blob := range decodedPayloads(text) {
			for _, h := range scan(blob.text, false) {
				h.Pattern = "encoded_" + h.Pattern
				h.Start, h.End = blob.start, blob.end
				h.Text = h.Text + " (base64)"
				hits = append(hits, h)
			}
		}
	}
	return hits
}

func containsAny(s string, lits []string) bool {
	if len(lits) == 0 {
		return true
	}
	for _, l := range lits {
		if strings.Contains(s, l) {
			return true
		}
	}
	return false
}

func overlaps(hits []Hit, s, e int) bool {
	for _, h := range hits {
		if h.Start < e && s < h.End {
			return true
		}
	}
	return false
}

var base64Blob = regexp.MustCompile(`[A-Za-z0-9+/_\-]{24,}={0,2}`)

const (
	maxDecodedBlobs = 16
	maxBlobBytes    = 16 << 10
)

type payload struct {
	text       string
	start, end int
}

// decodedPayloads finds base64 runs that decode to readable text. Binary
// data (images, hashes, keys) fails the readability check and is skipped.
func decodedPayloads(text string) []payload {
	var out []payload
	for _, span := range base64Blob.FindAllStringIndex(text, -1) {
		if len(out) >= maxDecodedBlobs {
			break
		}
		raw := text[span[0]:span[1]]
		if len(raw) > maxBlobBytes {
			continue
		}
		for _, enc := range []*base64.Encoding{base64.StdEncoding, base64.RawStdEncoding, base64.URLEncoding, base64.RawURLEncoding} {
			b, err := enc.DecodeString(raw)
			if err != nil {
				continue
			}
			if readable(b) {
				out = append(out, payload{text: string(b), start: span[0], end: span[1]})
			}
			break
		}
	}
	return out
}

func readable(b []byte) bool {
	if len(b) < 12 || !utf8.Valid(b) {
		return false
	}
	printable, spaces, total := 0, 0, 0
	for _, r := range string(b) {
		total++
		if r == ' ' {
			spaces++
		}
		if unicode.IsPrint(r) || r == '\n' || r == '\t' {
			printable++
		}
	}
	return spaces > 0 && printable*100 >= total*90
}

func mapSeverity(s Severity) guardrails.Severity {
	switch s {
	case SeverityHigh:
		return guardrails.SeverityHigh
	case SeverityMedium:
		return guardrails.SeverityMedium
	case SeverityLow:
		return guardrails.SeverityLow
	}
	return ""
}

// reasonFor builds a stable, human-readable summary: "injection: <names>".
func reasonFor(hits []Hit) string {
	names := make(map[string]struct{}, len(hits))
	for _, h := range hits {
		names[h.Pattern] = struct{}{}
	}
	list := make([]string, 0, len(names))
	for n := range names {
		list = append(list, n)
	}
	sort.Strings(list)
	return "injection: " + strings.Join(list, ", ")
}

func toFindings(hits []Hit) []guardrails.Finding {
	out := make([]guardrails.Finding, 0, len(hits))
	for _, h := range hits {
		out = append(out, guardrails.Finding{
			Entity:   "PROMPT_INJECTION:" + h.Pattern,
			Score:    score(h.Severity),
			Start:    h.Start,
			End:      h.End,
			Original: truncate(h.Text, 96),
		})
	}
	return out
}

func score(s Severity) float64 {
	switch s {
	case SeverityHigh:
		return 0.9
	case SeverityMedium:
		return 0.7
	case SeverityLow:
		return 0.5
	}
	return 0.0
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	for n > 0 && !utf8.RuneStart(s[n]) {
		n--
	}
	return s[:n] + "…"
}
