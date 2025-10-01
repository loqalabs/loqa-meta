# Loqa MVP Backlog

This backlog tracks the committed work required to reach the MVP milestone (Phase 1 – Foundation).

## Workstream A: Core Runtime ✅
- [x] Runtime bootstrap: service container, configuration loader, tracing wiring ([loqa-core#4](https://github.com/loqalabs/loqa-core/issues/4))
- [x] Message bus integration (NATS) with subject catalog ([loqa-core#5](https://github.com/loqalabs/loqa-core/issues/5))
- [x] Capability registry and node heartbeat protocol ([loqa-core#6](https://github.com/loqalabs/loqa-core/issues/6))
- [x] Local event store (SQLite + LiteFS) proof of concept ([loqa-core#7](https://github.com/loqalabs/loqa-core/issues/7))
- [x] Observability stack deployment scripts (Grafana, Tempo, Prometheus) ([loqa-core#8](https://github.com/loqalabs/loqa-core/issues/8))

## Workstream B: Voice Pipeline ✅
- [x] Integrate Whisper/faster-whisper wrapper with streaming API ([loqa-core#15](https://github.com/loqalabs/loqa-core/issues/15))
- [x] LLM inference harness using llama.cpp/Ollama (quantized 3B + 7B) ([loqa-core#16](https://github.com/loqalabs/loqa-core/issues/16))
- [x] TTS adapter for Kokoro with streaming playback ([loqa-core#17](https://github.com/loqalabs/loqa-core/issues/17))
- [x] Session router orchestrating STT → LLM → TTS chain with per-hop QoS settings ([loqa-core#18](https://github.com/loqalabs/loqa-core/issues/18))
- [x] Latency instrumentation surfacing `loqa.voice_latency_ms` ([loqa-core#19](https://github.com/loqalabs/loqa-core/issues/19))

## Workstream C: Skills & Plugins ✅
- [x] Define `skill.yaml` schema and validation
- [x] WASM host runtime skeleton (wazero-based)
- [x] Reference timer skill (WASM) and smart-home bridge (native)
- [x] Wire skills runtime to publish/subscribe via NATS using reference manifests
- [x] Audit log format and storage for skill invocations
- [x] Capability-based permission model for host calls
- [x] Authoring guide for third-party skills (build, package, permissions)

## Workstream D: Governance & Community ✅
- [x] Publish contribution guide and code of conduct
- [x] Launch RFC process and template
- [x] Stand up Discussions categories (announcements, ideas, support)
- [x] Draft security disclosure policy
- [x] Announce Ambiware Labs + Loqa relationship blog post
- [x] Document local developer setup for skills (TinyGo 0.39+, required tools)

## Workstream E: Release Readiness
- [ ] CI pipelines with lint, format, unit tests
- [ ] Nightly build artifacts (pre-release installers/images)
- [ ] MVP documentation (install guide, quickstart, architecture overview)
- [ ] Internal dogfooding environment (two-node cluster)
- [ ] Alpha tester recruitment and feedback loop

## Tracking & Cadence
- Weekly triage: review board, update statuses, validate priorities
- Monthly checkpoint: reassess scope, adjust timeline based on learnings
- Exit criteria: All backlog items complete, latency SLO < 200ms validated, documentation ready for early adopters
