<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="assets/brand/keirouter-logo-dark.png">
    <img src="keirouter-logo.png" alt="KeiRouter" width="320">
  </picture>
</p>

<h1 align="center">KeiRouter 🚀</h1>

<p align="center">
  <strong>Your friendly, blazing-fast, self-hostable AI gateway.</strong>
</p>

<p align="center">
  <a href="https://github.com/mydisha/keirouter/actions/workflows/ci.yml"><img src="https://github.com/mydisha/keirouter/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go&logoColor=white" alt="Go 1.26+"></a>
  <a href="https://github.com/mydisha/keirouter/pkgs/container/keirouter"><img src="https://img.shields.io/badge/Docker-GHCR-2496ED?logo=docker&logoColor=white" alt="Docker Image"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT"></a>
</p>

<p align="center">
  <img src="assets/keirouter-social-banner.png" alt="KeiRouter dashboard overview" width="800">
</p>

---

## So, what's the deal? 🤔

You use AI coding tools — Claude Code, Codex, Cursor, Cline, or honestly anything that talks to OpenAI, Anthropic or Gemini. And you know the headaches: a drawer full of API keys, rate limits that hit at the worst time, and a token bill that keeps creeping up.

**KeiRouter is the smart middleman that makes all of that someone else's problem.** Point your tools at one local endpoint and let it handle the boring stuff — routing each request to the right model, failing over the moment a provider taps out, caching repeat questions, and squeezing oversized prompts down *before* they ever cost you a token.

Oh, and it's written in Go. So it's light on memory, boots instantly, ships as one binary plus a static dashboard, and the dashboard is actually nice to look at. ✨

> **Heads up:** KeiRouter is under active development. Check the [CHANGELOG](CHANGELOG.md) for what landed recently, and peek at [Architecture](#-architecture) to see how a request flows.

---

## 📑 What's inside

- [Highlights](#-the-good-stuff)
- [Screenshots](#-screenshots)
- [Quick Start](#-quick-start)
  - [macOS](#-macos) · [Windows](#-windows) · [Linux](#-linux-ubuntu-debian-fedora-) · [Docker / VPS](#-docker--vps)
  - [First login](#first-login)
  - [Everyday commands](#everyday-commands)
  - [Updating](#updating) · [Uninstalling](#uninstalling)
  - [Connect your tools](#connect-your-tools)
- [API Surface](#-api-surface)
- [Token Savings](#-token-savings-where-the-money-hides)
- [Smart Routing (Chains)](#-smart-routing-chains)
- [Semantic Cache](#-semantic-cache)
- [Connect via OAuth](#-skip-the-keys-oauth)
- [Skills](#-skills)
- [Plans & Budgets](#-plans--budgets)
- [Rate Limiting](#-rate-limiting)
  - [High availability (multiple replicas)](#high-availability-multiple-replicas)
- [Guardrails](#-guardrails)
- [Observability](#-observability)
  - [Audit log](#audit-log)
- [Networking: Proxy Pools & Tunnels](#-networking-proxy-pools--tunnels)
- [Import, Export & Backups](#-import-export--backups)
- [Branding & White-Label](#-make-it-yours-branding)
- [CLI Tools Auto-Config](#-cli-tools-auto-config)
- [Usage Portal](#-usage-portal)
- [Supported Providers](#-supported-providers-90)
- [Configuration](#-configuration)
- [Architecture](#-architecture)
- [Security](#-security)
- [Development & Contributing](#-hack-on-it)
- [License](#-license)

---

## ⭐ The good stuff

**Routing & reliability**
- 🔀 **One endpoint, every provider** — Your apps speak OpenAI Chat, the OpenAI Responses API, Anthropic Messages, or native Gemini; KeiRouter quietly translates to whatever you actually want to use.
- 🛡️ **Never stop coding** — Provider down? Rate limited? Fallback chains with priority, round-robin, latency or cost ordering keep you moving like nothing happened.
- 🩺 **Provider health** — Real-traffic telemetry plus synthetic probes, with 24h / 7d history that survives restarts.
- ⚡ **Response cache** — Ask the same thing twice and get it back instantly for $0.00. Exact-match out of the box, true semantic matching with an embeddings API.

**Cost control**
- 💸 **Five token savers** — Two on the way in, three on the way out, all on a live savings dashboard. Details in [Token Savings](#-token-savings-where-the-money-hides).
- 💰 **Budget engine** — Per-key USD and token hard limits, with an auto-cutoff so surprise bills stay fictional.
- 🚦 **Rate limiting** — Per-key RPM, TPM, and concurrency caps via global defaults or reusable plans, with TPM settled against the real usage the provider reports.
- 🧩 **HA-ready** — Run several replicas behind a load balancer: rate limits, concurrency and budgets are shared through Redis at one round trip per request.
- 📈 **Quota tracker** — Upstream limits and local usage for every provider account in one place.

**Safety & governance**
- 🛡️ **Guardrails** — PII, prompt-injection, toxicity, topics, and bias detectors layered global → provider → model → chain → key, with mid-stream output scanning, an ML classifier, an LLM judge, and a fully-offline mode.
- 📜 **Audit log** — Every sign-in and config change, hash-chained so tampering shows, exported as OpenTelemetry logs and JSONL for your SIEM.
- 🔐 **Locked down by default** — AES-256-GCM envelope encryption for credentials, a loopback-only admin API, a password-gated dashboard, HMAC sessions, and SSRF protection on outbound calls.

**Operations & UX**
- 📊 **See everything** — Usage by provider, model, chain and request; TTFT and p95 latency; a live console log; host and process resource charts; Prometheus metrics.
- 🧠 **Skills** — Custom instructions injected into matching chat requests, plus reference skills that teach any agent how to use your gateway.
- 🔧 **CLI auto-config** — Detects 11 coding tools on your machine and writes their config in one click.
- 🌐 **Proxy pools & tunnels** — Route upstream traffic through proxies or a one-click Cloudflare Worker relay, and expose the gateway via Cloudflare Tunnel or Tailscale Funnel.
- 📦 **Bring your setup with you** — Import from 9router (JSON or SQLite) and OmniRoute, plus encrypted portable backups.
- 🎨 **White-label it** — Rebrand the dashboard and portal with your own name, logo, and palette.
- 🖥️ **System tray** — Run it in the background with a tray icon on macOS, Windows and Linux.

<p align="center">
  <img src="assets/keirouter-providers-banner.png" alt="Connect a provider dialog" width="800">
</p>

<p align="center">
  <img src="assets/keirouter-usage-banner.png" alt="Usage analytics" width="800">
</p>

---

## 📸 Screenshots

| | |
|---|---|
| <img src="assets/screenshots/chains.png" alt="Chain builder with route preview" width="400"><br><sub><b>Chains</b> — build fallback routes and preview them live</sub> | <img src="assets/screenshots/health.png" alt="Provider health" width="400"><br><sub><b>Provider health</b> — uptime, fallbacks and p95 latency per upstream</sub> |
| <img src="assets/screenshots/guardrails.png" alt="New guardrail policy dialog" width="400"><br><sub><b>Guardrails</b> — scoped policies with a dry-run tester</sub> | <img src="assets/screenshots/endpoints.png" alt="Endpoints page with quick start" width="400"><br><sub><b>Endpoints</b> — base URL, quick-start snippets and API reference</sub> |
| <img src="assets/screenshots/token-saving.png" alt="Token saving settings" width="400"><br><sub><b>Token saving</b> — toggle each saver independently</sub> | <img src="assets/screenshots/system.png" alt="System resources page" width="400"><br><sub><b>System</b> — CPU, memory and goroutine history for the instance</sub> |

---

## 🚀 Quick Start

Three steps on any OS: **install → start → open the dashboard**. Prebuilt binaries mean no Go, no Node, and no config file.

| Your machine | Install with | Run in the background with |
|---|---|---|
| 🍎 **macOS** | [Homebrew](#-macos) | `brew services start keirouter` |
| 🪟 **Windows 10/11** | [PowerShell one-liner](#-windows) | `keirouter tray` |
| 🐧 **Linux** (Ubuntu, Debian, Fedora, …) | [Shell one-liner](#-linux-ubuntu-debian-fedora-) | `--service` (systemd) |
| 🐳 **Server / VPS** | [Docker Compose](#-docker--vps) | Docker itself |

### 🍎 macOS

```bash
brew tap mydisha/keirouter https://github.com/mydisha/keirouter
brew install keirouter
keirouter start
```

Your browser opens the dashboard. Want it running all the time, even after a reboot?

```bash
brew services start keirouter   # starts now and on every login
```

Prefer a menu-bar icon you can quit any time? Run `keirouter tray` instead.

<details>
<summary>No Homebrew? Use the shell installer</summary>

```bash
curl -fsSL https://raw.githubusercontent.com/mydisha/keirouter/main/scripts/install.sh | bash
```

It downloads the right build for Apple Silicon or Intel, checks its checksum, and installs `keirouter` to `/usr/local/bin` (it may ask for your password once).
</details>

### 🪟 Windows

Open **PowerShell** and run:

```powershell
irm https://raw.githubusercontent.com/mydisha/keirouter/main/scripts/install.ps1 | iex
```

Then open a **new** terminal (so the PATH change applies) and start it:

```powershell
keirouter start
```

Want it in the background? `keirouter tray` puts an icon in the taskbar tray. To start it automatically when you log in, add a shortcut to your Startup folder:

```powershell
$s = (New-Object -ComObject WScript.Shell).CreateShortcut("$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\KeiRouter.lnk"); $s.TargetPath = (Get-Command keirouter).Source; $s.Arguments = "tray --no-browser"; $s.WindowStyle = 7; $s.Save()
```

> Pin a version or change the location with `$env:KEIROUTER_VERSION` / `$env:KEIROUTER_DIR` before running the installer.

### 🐧 Linux (Ubuntu, Debian, Fedora, …)

```bash
curl -fsSL https://raw.githubusercontent.com/mydisha/keirouter/main/scripts/install.sh | bash
keirouter start
```

Works on amd64 and arm64 (including 64-bit Raspberry Pi OS). Only `curl` and `tar` are needed, which every Ubuntu already has.

To keep KeiRouter running in the background and start it on boot, install it as a systemd service instead:

```bash
curl -fsSL https://raw.githubusercontent.com/mydisha/keirouter/main/scripts/install.sh | bash -s -- --service
```

```bash
systemctl --user status keirouter      # is it up?
journalctl --user -u keirouter -f      # follow the logs
systemctl --user restart keirouter     # restart after a config change
```

On a desktop with a system tray, `keirouter tray` works too.

### 🐳 Docker / VPS

The easiest way to run KeiRouter on a server:

```bash
curl -fsSL https://raw.githubusercontent.com/mydisha/keirouter/main/scripts/install.sh | bash -s -- --docker
```

Or by hand with Docker Compose:

```bash
git clone https://github.com/mydisha/keirouter.git
cd keirouter
cp .env.example .env       # set KEIROUTER_MASTER_KEY (openssl rand -base64 32)
docker compose up -d --build
```

Want PostgreSQL instead of SQLite? Layer on `compose.postgres.yaml`. VPS, PostgreSQL, Coolify and reverse-proxy notes live in [deploy/README.md](deploy/README.md).

> **Installed the binary on a server instead?** It only listens on `127.0.0.1` by default, which keeps the admin dashboard private. Reach it through an SSH tunnel: `ssh -L 20180:localhost:20180 you@your-server`, then open http://localhost:20180 on your own machine.

### First login

Open **http://localhost:20180** (`keirouter start` opens it for you) and sign in with the default password **`keirouter`**. It'll nudge you to change that password right away (please do 🙏).

Then grab an API key from the **API keys** page, or mint one from the terminal:

```bash
keirouter bootstrap   # prints a kr_ key, once
```

> **Remote dashboard access:** Plain HTTP works for initial LAN/VPS access, but HTTPS is strongly recommended for production. Behind a TLS-terminating reverse proxy, make sure it overwrites client-supplied `X-Forwarded-Proto` and `Forwarded` headers so KeiRouter can apply the `Secure` session-cookie flag correctly. See the [deployment guide](deploy/README.md).

### Everyday commands

| Command | What it does |
|---|---|
| `keirouter start` | Start the server on `:20180` and open the dashboard |
| `keirouter start --no-browser` | Same, without opening a browser |
| `keirouter tray` | Run in the background with a tray / menu-bar icon |
| `keirouter status` | Is it running, and on which URL? |
| `keirouter bootstrap` | Create an API key and print it once |
| `keirouter version` | Print the version |

### Updating

| Installed with | Update with |
|---|---|
| Homebrew | `brew upgrade keirouter` (then `brew services restart keirouter` if you use it) |
| Shell installer (macOS / Linux) | Run the same one-liner again (then `systemctl --user restart keirouter` if you used `--service`) |
| PowerShell installer | Run the same one-liner again |
| Docker | Run the `--docker` one-liner again, or `git pull && docker compose up -d --build` |

Settings → System also tells you when a new release is out.

### Uninstalling

Your data (database, master key, settings) lives in `~/.keirouter` (`%USERPROFILE%\.keirouter` on Windows). Uninstalling never touches it; delete that folder yourself if you want a clean slate.

<details>
<summary>Uninstall commands</summary>

**Homebrew**

```bash
brew services stop keirouter
brew uninstall keirouter
```

**Shell installer (macOS / Linux)**

```bash
systemctl --user disable --now keirouter 2>/dev/null; rm -f ~/.config/systemd/user/keirouter.service
sudo rm -rf /usr/local/bin/keirouter /usr/local/share/keirouter
```

**Windows (PowerShell)**

```powershell
Remove-Item -Recurse -Force "$env:LOCALAPPDATA\KeiRouter", "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\KeiRouter.lnk" -ErrorAction SilentlyContinue
```

Then remove `%LOCALAPPDATA%\KeiRouter` from your user PATH (Settings → System → About → Advanced system settings → Environment Variables).
</details>

### Building from source

Only needed if you want to hack on KeiRouter or run the latest `main`. You'll need **Go 1.26+** and **Node.js 22** (20.19+ works too):

```bash
curl -fsSL https://raw.githubusercontent.com/mydisha/keirouter/main/scripts/quickstart.sh | bash
```

It clones the repo to `~/keirouter`, installs everything, and starts the backend on `:20180` and a hot-reloading dashboard on **http://localhost:5180**. Already cloned it? Just `make setup`. To build and install a release-style binary from source instead, use `install.sh --source`. More in [Hack on it](#-hack-on-it).

### Connect your tools

In your AI tool of choice, set:

- **Base URL:** `http://localhost:20180/v1` for OpenAI-style clients, or `http://localhost:20180` for Anthropic-style clients like Claude Code
- **API Key:** the `kr_` key from the dashboard or `keirouter bootstrap` (sent as `Authorization: Bearer …` or `x-api-key`)
- **Model:** a provider model like `openai/gpt-4o`, or the name of a fallback [chain](#-smart-routing-chains) you cooked up

That's it. You're routing. The **Endpoints** page has copy-paste snippets for cURL, Python, Node.js and Claude Code, and the **CLI tools** page can [write the config for you](#-cli-tools-auto-config).

---

## 🔌 API Surface

Chat is just the start. KeiRouter proxies the whole buffet behind one key:

| Capability | Endpoint | Dialect |
|---|---|---|
| Chat completions | `POST /v1/chat/completions` | OpenAI |
| Responses API (Codex & friends) | `POST /v1/responses` | OpenAI |
| Messages | `POST /v1/messages` | Anthropic |
| Token counting (answered locally) | `POST /v1/messages/count_tokens` | Anthropic |
| Native Gemini | `POST /v1beta/models/{model}:generateContent` | Gemini |
| Embeddings | `POST /v1/embeddings` | OpenAI |
| Image generation | `POST /v1/images/generations` | OpenAI |
| Image understanding | `POST /v1/images/understanding` | — |
| Video generation (async) | `POST /v1/videos/generations` · `GET /v1/videos/{id}` | — |
| Text-to-speech | `POST /v1/audio/speech` | OpenAI |
| Speech-to-text | `POST /v1/audio/transcriptions` | OpenAI |
| Web search | `POST /v1/search` | — |
| Web fetch | `POST /v1/web/fetch` | — |
| Model discovery | `GET /v1/models` · `/v1/models/info` · `/v1/models/{kind}` | OpenAI |
| Your own key's usage | `GET /v1/keys/me/usage` | — |

`GET /healthz` is the liveness probe, and accidental `/v1/v1/...` paths are folded back to `/v1/...` so a misconfigured base URL still works.

---

## 💸 Token Savings (where the money hides)

Every request runs through a deterministic token-saving pipeline **before** it gets translated to the provider's format — so the savings work the same no matter which provider you land on. There are **five savers**, each independently toggleable from **Settings → Token saving**, and a dashboard that shows you exactly how much you clawed back.

| Saver | Side | The pitch |
|---|---|---|
| **RTK / Slimmer** | Input | Shrinks chunky tool output (diffs, greps, listings, build logs, source files) locally before it ever leaves your machine. |
| **Headroom** | Input | Routes request messages through an external [Headroom](https://github.com/headroomlabs-ai/headroom) proxy for deeper compression. *Fail-open* — if the proxy sneezes, your request sails through untouched. Also sniffs out "phantom savings" so only real wins get counted. |
| **Terse** | Output | Drops a concise-output directive so the model skips the small talk and gives you the goods. |
| **Caveman** | Output | Terse's stronger cousin (Wenyan / 文言文 levels included) — trims output tokens by 65–75%. |
| **Ponytail** | Output | Injects a "lazy senior dev" system prompt (`lite` / `full` / `ultra`) that nudges the model toward the smallest possible change. Stacks on top of Terse or Caveman. |

> Terse and Caveman both inject a system directive, so they're mutually exclusive — pick one. Ponytail happily layers on top of either.

**Pipeline order:** `normalizer → RTK → Headroom → Skills → Terse / Caveman → Ponytail → provider translation`.

The **Usage** page breaks savings down per rule, and you can export a shareable *savings card*.

### Getting Headroom running

Headroom is its own open-source compression proxy. KeiRouter just calls its `/v1/compress` endpoint, so you spin it up locally first. One gotcha worth shouting about: the **`headroom` CLI lives in the Python package** — the npm package is a library only, so `npm install -g headroom-ai` will leave you staring at `command not found`. Don't say we didn't warn you. 😉

```bash
# The clean way: pipx isolates the CLI and sorts out your PATH (needs Python 3.10+)
pipx install "headroom-ai[all]"
pipx ensurepath               # puts headroom on your PATH — then restart your shell

headroom proxy --port 8787    # start the proxy
headroom doctor               # make sure it's actually happy
```

On Ubuntu and `pip` is fighting you (PEP 668 blocks global installs)? Go `--user` and make sure `~/.local/bin` is on your `PATH`:

```bash
pip install --user "headroom-ai[all]"
export PATH="$HOME/.local/bin:$PATH"   # drop this in ~/.zshrc or ~/.bashrc
```

Then over in **Settings → Token saving → Headroom proxy**: flip it on, set the **Proxy URL** to `http://localhost:8787`, and hit **Test connection** to confirm the handshake. Green check? You're golden.

---

## 🧠 Smart Routing (Chains)

Why bet on one model when you can have a backup plan? Build a **chain** on the **Chains** page. Say you name one `coding`:

1. `openai/gpt-4o` — your first pick
2. `deepseek/deepseek-chat` — steps in if the first one rate-limits or face-plants

Then just set your app's model to `chain:coding` (or plain `coding`) and let KeiRouter sweat the failover.

Each chain picks a **routing strategy**:

| Strategy | Behaviour |
|---|---|
| **Priority** | Tries models in the order listed |
| **Round robin** | Spreads requests across the steps |
| **Latency** | Fastest measured step first |
| **Cost** | Cheapest step first |

Add an optional **final fallback** that's tried after every step fails, and watch the live route preview while you edit. Inside a provider, multiple accounts can also rotate round-robin or fill in priority order (**Settings → Routing**).

---

## ⚡ Semantic Cache

Off by default; turn it on in your config:

```yaml
cache:
  enabled: true
  backend: memory            # or "redis" (set redis_url) to persist and share across replicas
  ttl: 1h
  embedding_provider: hash   # exact-match, zero dependencies
  # embedding_provider: api  # true semantic matching via any OpenAI-compatible embeddings API
  similarity_threshold: 0.95
```

With `embedding_provider: api`, paraphrases like *"How do Go channels work?"* and *"Explain Go channels"* hit the same entry. Cache savings show up on the **Overview** page. See [config.example.yaml](config.example.yaml) for every knob.

---

## 🔑 Skip the keys: OAuth

Copy-pasting API keys gets old fast. On the **Providers** page, hit **Connect provider** and pick one from **Sign in with a plan** — sign in once, and KeiRouter refreshes your tokens in the background (proactively, before they expire).

| Provider | Flow |
|---|---|
| **Claude Code** | Anthropic OAuth (PKCE) |
| **OpenAI Codex** | OpenAI OAuth (PKCE, local callback) |
| **GitHub Copilot** | GitHub device flow |
| **Gemini CLI** · **Antigravity** | Google OAuth |
| **Qwen Code** | Device flow (PKCE) |
| **Kimi Coding** | Device flow |
| **Grok CLI (Grok Build)** | Device flow |
| **Kiro** | AWS Builder ID device flow, IAM Identity Center, refresh-token import, or API key |
| **Kilo Code** | Device auth |
| **Qoder** | Device-token flow or Personal Access Token |
| **CodeBuddy** (Tencent) | Browser-poll flow |
| **Kimchi** | Browser callback |
| **Cline** · **ClinePass** · **iFlow** | OAuth |
| **Cursor** · **Command Code** | Token import |

> Several of these are marked **Unofficial** in the dashboard: they reuse a client's own sign-in rather than a public API, so check each provider's terms before you lean on them.

---

## 🧠 Skills

The **Skills** page has two halves:

- **Custom skills** — Standing instructions the gateway adds to the system prompt of matching chat requests. Scope each skill by API key, plan, and model pattern (`claude-*`), or leave it empty to apply everywhere. Skills are prepended as a stable prefix, so provider prompt caches stay warm, and they're never injected twice across retries or fallbacks. Mark a skill *optional* and clients can opt out with the `X-KeiRouter-Skip-Skills: true` header.
- **Reference skills** — Ready-made `SKILL.md` docs in [`skills/`](skills/) that teach Claude, Cursor, ChatGPT or any other agent how to call your KeiRouter endpoint (chat, embeddings, images, TTS, STT, web search, web fetch). Copy the URL or the agent prompt straight from the dashboard.

---

## 📋 Plans & Budgets

Plans are reusable budget policies you can stamp onto any API key — write the rules once, apply them everywhere. Each plan covers:

- **Spend limit** — USD cap (micro-dollar precision, because pennies add up)
- **Token limit** — max token usage
- **RPM / TPM / Concurrency limits** — per assigned key (`0` = unlimited)
- **Reset period** — `daily`, `weekly`, `monthly`, or `total`
- **Allowed models** — wildcard patterns like `claude-*`, `gpt-4*` (empty = everything's fair game)
- **Alert threshold** — get pinged at 1–100% of budget
- **Hard cutoff** — block requests when the budget's gone, or just track and let it ride

Every tenant gets a default plan out of the box. Manage them on the **Plans & budgets** page and assign them on each **API key**.

---

## 🚦 Rate Limiting

Keep your gateway (and your upstream quotas) from getting hammered with per-key RPM, TPM, and concurrency caps. Limits are **off by default**; here's an example for a single-instance setup:

```yaml
limits:
  enabled: true
  backend: memory          # one instance; use "redis" for several (see below)
  default_rpm: 600
  default_tpm: 200000
  default_concurrency: 50
  window: 1m
  cleanup_interval: 1m
```

Those defaults only apply to keys **without** a plan. The moment a key has one, the plan's `rpm_limit`, `tpm_limit`, and `concurrency_limit` take over (`0` = unlimited). You can also flip enforcement at runtime from **Settings → Routing**.

TPM is reserved up front and then settled against the usage the provider actually reports (`reconcile_tokens`), so cache hits, blocked requests and local failures hand their tokens back. `output_reserve_cap` stops clients that always send a huge `max_tokens` from draining TPM up front, and concurrency denials include a `Retry-After` hint.

Separately, per-IP abuse guards (`auth_failures_per_minute`, `login_attempts_per_minute`) stay on even when per-key limits are off.

### High availability (multiple replicas)

Running several KeiRouter replicas behind a load balancer? Point them at the same Postgres and Redis and set `backend: redis`. Every replica then enforces **one** RPM/TPM window, **one** concurrency pool per key and **one** budget, instead of each replica handing out its own allowance.

```yaml
database:
  driver: postgres
  dsn: postgres://keirouter:secret@postgres:5432/keirouter?sslmode=disable
limits:
  enabled: true
  backend: redis
  redis_url: redis://redis:6379/0
  fail_mode: local      # local | open | closed while redis is unreachable
```

How it stays fast:

- **One round trip per request.** A Lua script checks and charges concurrency, RPM and TPM atomically, and one more call releases the slot. That's ~0.2–0.5 ms on a LAN, next to LLM calls that take seconds. Keys without limits never touch Redis.
- **Budgets without a `SUM` per request.** Spend lives in shared Redis counters seeded from the usage table and re-seeded every `budget_counter_ttl`. Increments are merged in the background, so Redis writes scale with the number of active budgets, not with traffic.
- **Crash-safe concurrency.** Slots are leases with a heartbeat. A replica that dies mid-request frees its slots after `lease_ttl` instead of leaking them forever.
- **Survives a Redis outage.** A failed call trips a 5 s circuit breaker, so an outage costs one timeout, not one per request. Requests follow `fail_mode`, and budgets fall back to the usage table.
- **Cluster-ready keys.** Each key's counters share a hash tag, so the scripts also run on Redis Cluster.

Budget edits are broadcast over Redis pub/sub, so every replica drops its cached definitions immediately. `GET /api/settings/limits/status` shows the backend, its health, live leases and fallback counts.

---

## 🛡️ Guardrails

A built-in content-safety layer runs detectors against every request and response. Policies stack **global → provider → model → chain → API key** and merge at request time, so the most specific rule wins.

| Detector | Catches | Default engine | Optional engine |
|---|---|---|---|
| **PII** | Email, phone, credit card, IBAN, IP, URL, **NIK / KK / NPWP / BPJS / Indonesian passport**, US SSN, API keys, passwords, private keys | Native Go (Presidio-compatible, checksum-validated) | Microsoft Presidio HTTP sidecar |
| **Prompt Injection** | Instruction overrides, jailbreak personas, smuggled chat-template tokens, prompt leaks, encoded payloads, indirect injection in tool results (id + en) | Signatures + typo-tolerant intent rules | Prompt Guard 2 classifier · embedding match · LLM judge |
| **Topics** | Allow-list / block-list of topics | Whole-word keyword match with inflections | Embedding similarity |
| **Toxicity** | Profanity, hate, harassment, violence, sexual, self-harm (id + en) | Native weighted lexicon | OpenAI Moderation API |
| **Bias** (outbound, experimental) | Political, gender, ethnic, religious stereotyping in responses | Native bilingual lexicon | — |

- **Actions:** `log_only`, `warn`, `mask` (rewrite it), or `block` (refuse it). Strictest action across detectors wins.
- **PII strategies:** `redact`, `replace`, `mask`, `hash`, `anonymize`, or `block`.
- **Obfuscation-resistant:** every detector reads a normalized copy of the text, so zero-width characters, fullwidth or look-alike letters, leetspeak (`1gn0r3`), censoring (`f*ck`) and spaced-out letters don't slip through.
- **Streaming-safe:** a short hold-back buffer scans assistant output before it is forwarded, so leaked PII is masked and blocked text never reaches the client, even when it is split across chunks.
- **Prompt-injection hardening** (per policy): spotlighting marks tool results as untrusted data, a canary marker in the system prompt blocks responses that leak it, and tool calls are refused in a turn whose tool results look injected.
- **Test before you ship:** every policy has a dry-run tester — nothing reaches a provider.
- **Observability:** every decision shows up in the **Audit log** tab of the Guardrails page (live via SSE); Prometheus serves `keirouter_guardrail_decisions_total` and `keirouter_guardrail_eval_seconds` at `/metrics`.
- **Compliance:** a per-tenant **Allow external detector engines** switch forces every detector offline; `KEIROUTER_GUARDRAILS__AUDIT_RETENTION_DAYS` controls retention; the test endpoint is capped at 10 req/min.

Starter templates ship in the **Start from template** picker (Indonesia PII · Strict safety · Compliance audit · Public chatbot · Coding agent · Alerts-only), and you can export/import policies as a JSON bundle.

Want NER-based PII detection (`PERSON`, `LOCATION`, full multilingual)? Spin up the optional **Presidio sidecar**:

```bash
docker compose -f compose.yaml -f compose.postgres.yaml -f compose.presidio.yaml up -d
```

Then flip any PII policy's engine to `presidio` in the dashboard.

Want a trained classifier on top of the injection signatures? Run the **Prompt Guard sidecar** (Llama Prompt Guard 2, multilingual; the model is gated, so set `HF_TOKEN` after accepting its license, or set `PROMPTGUARD_MODEL=protectai/deberta-v3-base-prompt-injection-v2`):

```bash
docker compose -f compose.yaml -f compose.promptguard.yaml up -d
```

Then turn on **ML classifier** in a prompt-injection policy. Any server speaking the Hugging Face text-classification contract works too: point `KEIROUTER_GUARDRAILS__INJECTION__CLASSIFIER_URL` at it. The **LLM judge** layer needs no setup: pick any model, alias or chain KeiRouter routes to.

---

## 📊 Observability

| Page | What you get |
|---|---|
| **Overview** | Requests, success rate, latency and spend vs. the previous period, top models, provider health, savings |
| **Usage** | Cost, tokens and latency (p50 / p95 / TTFT) by provider, model, chain and individual request |
| **Console** | Live gateway log stream with level filters |
| **Health** | Per-provider and per-model uptime, fallbacks, p95 latency and synthetic probes; 24h / 7d history from persisted snapshots |
| **Quota** | Upstream quota windows and local usage for every account, with auto-refresh |
| **System** | CPU, memory, disk, goroutines and KeiRouter's own RSS / heap over time |

Prometheus metrics live at `/metrics` (loopback-only by default): requests, latency, TTFT, tokens, cost, fallbacks, cache hits, upstream errors, saver activity and guardrail decisions.

### Audit log

Every dashboard sign-in and configuration change is recorded on the **Audit log** page: who did it (a stable per-session id), the action (`keys.create`, `settings.endpoint`, `auth.login`, ...), the resource, the outcome, client IP, request id and W3C trace id. Request bodies are never stored, only the names of the fields that were sent. Reading a full config export or a database backup is audited too.

- **Tamper-evident:** each replica writes an HMAC hash chain keyed from the master key. **Verify integrity** (or `GET /api/audit/verify`) reports any row that was edited, inserted or deleted.
- **OpenTelemetry-native:** set `audit.otlp.endpoint` (or the standard `OTEL_EXPORTER_OTLP_ENDPOINT`) and each event is also shipped as an OTLP log record, using semantic-convention attributes such as `user.id`, `client.address`, `http.route` and `event.name=keirouter.audit.<action>`. Any OTel Collector, Grafana Loki, Datadog, Elastic or Splunk pipeline can ingest it.
- **SIEM export:** **Export JSONL** (or `GET /api/audit/export`) downloads the filtered log as JSON Lines.
- **Never in the way:** events are queued and batch-written off the request path, and a slow collector has its own queue. `audit.retention_days` (default 365) prunes old rows.

```yaml
audit:
  otlp:
    endpoint: http://otel-collector:4318
    headers: { Authorization: "Bearer <token>" }
```

---

## 🌐 Networking: Proxy Pools & Tunnels

- **Proxy pools** — Send upstream traffic for chosen accounts through HTTP proxies or edge relays (Vercel, Cloudflare, Deno). **Deploy relay** pushes a Cloudflare Worker relay for you and watches until it's ready. A global outbound proxy lives in **Settings → Network**.
- **Public tunnels** — Need to reach your gateway from elsewhere? Turn on a **Cloudflare quick tunnel** or **Tailscale Funnel** from the **Endpoints** page.

---

## 📦 Import, Export & Backups

Moving in from another router? **Settings → Import & export** takes:

- **9router** JSON exports or the raw **SQLite database** (with an analyze step first, usage history, and password migration)
- **OmniRoute** JSON exports

Pick **Merge**, **Overwrite**, or **Wipe & replace**. You can also download local or passphrase-encrypted portable backups and restore them later.

---

## 🎨 Make it yours: Branding

Rebrand the admin dashboard *and* the public Usage Portal from **Settings → Branding**:

- **App name** — swap "KeiRouter" for your own
- **Logo & favicon URLs** — your SVG/PNG, your vibe
- **Tagline** — the line on the portal login screen
- **Color palette** — 11 built-ins: `kei` (default), `sage-terra`, `ocean`, `midnight`, `forest`, `rose`, `lavender`, `mono`, `sunset`, `arctic`, `cherry`

Light, dark, and system themes are a click away in the account menu.

---

## 🔧 CLI Tools Auto-Config

The **CLI tools** page detects which coding tools are installed, shows whether each one already points at KeiRouter, and can **write (or remove) the config in one click** — pick an endpoint, an API key and an optional default model. Prefer doing it by hand? Every tool has a manual snippet too.

> Claude Code · Codex CLI · Cline / Roo · GitHub Copilot · Factory Droid · OpenClaw · OpenCode · Kilo Code · Hermes Agent · DeepSeek TUI · jcode

Using OpenCode? The [`@keirouter/opencode-plugin`](@keirouter-opencode-plugin/) discovers your models straight from `GET /v1/models`.

---

## 🌐 Usage Portal

A dedicated, no-admin-required view at `/portal` where teammates keep an eye on their own usage — quota and spend, token usage over time, compression savings, and plan limits. Sign in with an API key, or share a read-only **portal ID** (`key_…`) so nobody has to hand out a secret. Signed in with a key, there's also a **Playground** tab for quick chats. No keys to the kingdom required.

---

## 🌐 Supported Providers (90+)

**🧠 LLM / Chat**

| Category | Providers |
|----------|-----------|
| **Major Cloud** | OpenAI, Anthropic, Google Gemini, Azure OpenAI, Vercel AI Gateway |
| **Free / Free Tier** | OpenRouter, NVIDIA NIM, Ollama (Cloud & Local), Cloudflare Workers AI, BytePlus ModelArk, MiMo Free, Kiro |
| **Sign in with a plan** | Claude Code, OpenAI Codex, GitHub Copilot, Gemini CLI, Antigravity, Cursor IDE, Cline, ClinePass, Kilo Code, Qwen Code, Qoder, iFlow, Kimi Coding, Kimchi, CodeBuddy (Tencent), Grok CLI, Command Code |
| **China / Asia** | DeepSeek, Qwen / Alibaba (incl. Intl), GLM Coding (+ China), Kimi (Moonshot), MiniMax Coding (+ China), Volcengine Ark, Xiaomi MiMo (+ Token Plan), SiliconFlow |
| **Performance** | Groq, Cerebras, SambaNova, DeepInfra |
| **Specialized** | xAI (Grok), Mistral, Perplexity (+ Agent / Responses), Cohere, AI21 Labs, Venice AI |
| **Aggregators & hosts** | Together AI, Fireworks AI, Nebius AI, OpenCode, AgentRouter, Hyperbolic, Featherless, Nscale, Baseten, Nous Research, Scaleway AI |
| **Emerging** | Blackbox AI, Chutes AI, and more |
| **Self-hosted** | vLLM, Ollama Local |
| **Custom** | Any OpenAI- or Anthropic-compatible endpoint (self-hosted, proxy, etc.), with your own model registry and prices |

**🎨 Media & Search**

| Type | Providers |
|------|-----------|
| **Image Generation** | OpenAI, OpenAI Codex, Gemini, xAI, Cloudflare, Fal.ai, Stability AI, Black Forest Labs, Recraft, Topaz, Runway ML, NanoBanana, Venice AI, MiniMax, HuggingFace, SD WebUI, ComfyUI |
| **Video Generation** | xAI |
| **Text-to-Speech** | OpenAI, Gemini, NVIDIA NIM, ElevenLabs, Deepgram, MiniMax, Hyperbolic, AWS Polly, Google TTS, Edge TTS, Inworld, Local Device |
| **Speech-to-Text** | OpenAI Whisper, Groq Whisper, Deepgram, AssemblyAI, Gemini, HuggingFace |
| **Embeddings** | OpenAI, Gemini, Mistral, Together AI, Fireworks AI, Nebius, NVIDIA NIM, GitHub Copilot, vLLM, Venice AI, Voyage AI, Jina AI, OpenRouter |
| **Web Search** | Tavily, Exa, Serper, Brave Search, Perplexity, xAI, Gemini, Kimi, MiniMax, Google PSE, OpenAI |
| **Web Fetch** | Tavily, Exa, Firecrawl, Jina |

Browse the full catalog (and filter by capability) under **Connect provider** in the dashboard.

---

## ⚙️ Configuration

Out of the box, KeiRouter runs on an embedded **SQLite** database — zero config, zero fuss. Running it for a team? Switch to **PostgreSQL**: copy [config.example.yaml](config.example.yaml) and run with `-config`, or use environment variables like `KEIROUTER_SERVER__PORT=8080` (double underscore = nesting). The master key is `KEIROUTER_SECURITY__MASTER_KEY` (Compose maps `KEIROUTER_MASTER_KEY` to it for you). Docker/Coolify examples are in [deploy/README.md](deploy/README.md).

Most day-to-day behaviour — token savers, routing, outbound proxy, branding, backups, update checks — lives in the dashboard under **Settings**, and changes save automatically.

---

## 🛠️ Architecture

Curious what happens after you hit send? Here's the life of a request:

1. **Gateway** — takes your HTTP request and figures out the dialect (OpenAI Chat, Responses, Anthropic, Gemini, …).
2. **Resolve** — turns the model name into targets: a provider model, an alias, or a chain with its strategy.
3. **Budget & limits** — reserves budget and acquires the key's RPM / TPM / concurrency slots.
4. **Guardrails (inbound)** — runs PII / injection / toxicity / topics detectors; block → refuse, mask → rewrite the prompt on the spot.
5. **Cache** — serves a hit straight away, for free.
6. **Pipeline** — runs the [token savers](#-token-savings-where-the-money-hides) and injects [skills](#-skills).
7. **Dispatch** — picks the best provider account and juggles fallbacks.
8. **Connector & Transform** — calls the provider, then translates the answer back into your tool's format.
9. **Guardrails (outbound)** — scans the response (or each stream chunk) for leaked PII, bias, or toxicity; block → cancel, mask → rewrite.
10. **Meter** — logs tokens, cost and savings so the dashboard has something pretty to show.

---

## 🔒 Security

- The admin API (`/api/*`) only answers on loopback by default (`security.bind_loopback_only`). Exposing it? Put it behind a reverse proxy and set a stable master key.
- **Guard your master key with your life** — it's the root of trust for every encrypted credential.
- Credentials sit behind AES-256-GCM envelope encryption; the dashboard uses Argon2id password hashing + HMAC session cookies; outbound requests get SSRF protection.

Found a security issue? Please follow [SECURITY.md](SECURITY.md) instead of opening a public issue. 🙏

---

## 🧑‍💻 Hack on it

```bash
make setup   # first time: installs deps + starts backend (:20180) and dashboard (:5180)
make dev     # after that: just start the servers
make test    # run the backend test suite
make build   # build the backend binary + frontend assets
```

PRs and ideas are always welcome — start with [CONTRIBUTING.md](CONTRIBUTING.md).

---

## 📄 License

MIT — see [LICENSE](LICENSE). Go build something cool. 🛠️
