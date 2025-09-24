# Loqa MVP Backlog

This backlog tracks the committed work required to reach the MVP milestone (Phase 1 – Foundation).

## Workstream A: Core Runtime ✅
- [x] Runtime bootstrap: service container, configuration loader, tracing wiring ([loqa-core#4](https://github.com/ambiware-labs/loqa-core/issues/4))
- [x] Message bus integration (NATS) with subject catalog ([loqa-core#5](https://github.com/ambiware-labs/loqa-core/issues/5))
- [x] Capability registry and node heartbeat protocol ([loqa-core#6](https://github.com/ambiware-labs/loqa-core/issues/6))
- [x] Local event store (SQLite + LiteFS) proof of concept ([loqa-core#7](https://github.com/ambiware-labs/loqa-core/issues/7))
- [x] Observability stack deployment scripts (Grafana, Tempo, Prometheus) ([loqa-core#8](https://github.com/ambiware-labs/loqa-core/issues/8))

## Workstream B: Voice Pipeline
- [ ] Integrate Whisper/faster-whisper wrapper with streaming API ([loqa-core#15](https://github.com/ambiware-labs/loqa-core/issues/15))
- [ ] LLM inference harness using llama.cpp/Ollama (quantized 3B + 7B) ([loqa-core#16](https://github.com/ambiware-labs/loqa-core/issues/16))
- [ ] TTS adapter for Kokoro with streaming playback ([loqa-core#17](https://github.com/ambiware-labs/loqa-core/issues/17))
- [ ] Session router orchestrating STT → LLM → TTS chain with per-hop QoS settings ([loqa-core#18](https://github.com/ambiware-labs/loqa-core/issues/18))
- [ ] Latency instrumentation surfacing `loqa.voice_latency_ms` ([loqa-core#19](https://github.com/ambiware-labs/loqa-core/issues/19))

## Workstream C: Skills & Plugins
- [ ] Define `skill.yaml` schema and validation
- [ ] WASM host runtime skeleton (Extism/wasmtime integration)
- [ ] Reference timer skill (WASM) and smart-home bridge (native)
- [ ] Audit log format and storage for skill invocations
- [ ] Capability-based permission model for host calls

## Workstream D: Governance & Community
- [ ] Publish contribution guide and code of conduct (complete ✅)
- [ ] Launch RFC process and template (complete ✅)
- [ ] Stand up Discussions categories (announcements, ideas, support)
- [ ] Draft security disclosure policy
- [ ] Announce Ambiware Labs + Loqa relationship blog post

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
