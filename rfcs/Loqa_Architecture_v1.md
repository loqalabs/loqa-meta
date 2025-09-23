# Loqa - Platform Architecture (Merged v1)

**Status:** Draft v1 (merge of Claude + Codex proposals)  
**Author:** Anna Barnes & collaborators  
**Last updated:** 2025-09-23  
**Scope:** Platform architecture for a local-first, open-core, multimodal ambient intelligence platform. Voice is the first interface; the system is modality-flexible and cluster-ready.

---

## 1. Product North Star

Loqa is a local-first, open-core platform for ambient intelligence. It runs entirely on consumer-grade hardware (Mac minis, used GPUs, Jetsons, Pis), scales by adding nodes, and respects privacy by design. Voice is the initial modality; glanceable/visual and touch UIs are first-class citizens on the roadmap.

### Core Tenets
- Local-first: works offline by default; cloud is opt-in and additive.
- Open-core: core pipeline, SDKs, and reference skills remain FOSS. Optional paid add-ons sustain the project without crippling DIY users.
- Multimodal: voice, glanceable displays, touch, sensors, and actuators.
- Composable and clusterable: add nodes to add capacity or skills; avoid single points of cloud dependency.
- Human-centered and private: no tracking or telemetry by default; fine-grained consent controls.

---

## 2. Goals and Non-Goals

### Goals
- Sub-200 ms perceived latency from end-of-speech to audio playback start (short utterances).
- Modular services for STT, NLU or LLM, skill execution, TTS (Kokoro), and visual output.
- Horizontal scaling via node roles and capability discovery.
- Robust plugin model (WASM-first) with safe capability access.
- Strong observability locally (no cloud egress required).

### Non-Goals (MVP)
- True multi-node model sharding of a single LLM. Scale by replicating full models per node, not parameter sharding.
- Full enterprise fleet management (arrives in Pro or Enterprise tier).
- Over-general orchestration (keep the control plane light until needed).

---

## 3. High-Level Architecture

Loqa is a set of cooperating services connected by a message bus and an optional control plane. Each service may run collocated on one node or be distributed across several.

### Core Services
1. Ingress and session router (`loqa-router`).
2. Wake word and edge audio (on pucks or edge devices).
3. STT (Whisper or faster-whisper).
4. NLU or LLM (llama.cpp or Ollama; quantized models).
5. Skill runtime (WASM sandbox plus containers for heavy skills).
6. TTS (Kokoro).
7. Visual output (e-ink, OLED, HTML dashboards).
8. State and memory (local event store and embeddings when enabled).
9. Observability (OpenTelemetry traces, metrics, logs).

### Platform Services
- Message bus: NATS (with JetStream for durability) as primary; Redis Streams and MQTT as edge adapters or bridges.
- Schema registry: protobuf definitions versioned with Buf; contract tests for evolution.
- Control plane (optional): Serf for membership and tags; lightweight Raft for leader election and simple config consensus.
- AuthN or AuthZ: mTLS for inter-service traffic; capability-based permissions for skills; local user roles for UIs.

---

## 4. Node Roles and Capability Discovery

Each node advertises capabilities (for example `stt.whisper:rt`, `llm.llama3:7b:q4_k_m`, `tts.kokoro:en_US`, `display.eink:2.9in`). The router uses these to place work.

### Typical Roles
- Hub (Mini or workstation): router, LLM, skills, observability stack.
- Inference nodes: STT or LLM or TTS replicas for concurrency or specialization.
- Edge pucks (ESP32 or Pi Zero): wake word, audio capture, low-power sensors.
- Display nodes (e-ink, OLED, tablet): ambient cards for timers, weather, now-playing.

### Discovery
- Serf gossip for membership; nodes join a LAN cluster and publish capability tags.
- Router caches capabilities; health checks via NATS request or reply and load probes.

---

## 5. Message Flow

### Voice Path (Streamed)
1. Wake (edge puck) starts the audio stream to ingress.
2. Ingress publishes `audio.frame` to NATS; STT subscribes with backpressure.
3. STT emits partial or final `stt.text` events; the router opens a session timeline.
4. Router invokes the LLM via `nlu.request` (request or reply) with conversation context and quality tier (latency versus accuracy budget).
5. LLM returns a structured intent payload (JSON or protobuf).
6. Skill runtime executes the intent (local WASM or container); may emit follow-up prompts or device actions.
7. TTS (Kokoro) synthesizes the response; egress streams audio to the originating puck or speaker.
8. Observability: OpenTelemetry spans across all hops; per-hop timing captured.

### Glanceable Path
- Shares the same session backbone. `visual.card.update` events drive displays directly (timer progress, weather, now-playing, reminders), optionally skipping TTS.

---

## 6. Latency Budget (Under 200 ms Target)

### Budget (Typical Short Reply)
- Endpointing (VAD or end-of-speech): 10-25 ms.
- STT (short segment, GPU or NEON acceleration): 40-80 ms.
- LLM first tokens (quantized 3-8B, Apple Silicon or GPU): 50-120 ms.
- Skill dispatch (WASM): 5-20 ms.
- Kokoro first audio chunk: 20-50 ms.
- Bus and serialization (LAN or Thunderbolt): 3-10 ms.

### Tactics
- Stream data end-to-end (no batch waits).
- Adaptive quality tiers fall back to smaller or quantized models on latency spikes.
- Prompt and result caches for hot intents.
- Pin frequent skills and models in RAM; locality hints co-locate LLM and skills to avoid extra hops.
- NUMA-aware placement on x86; throttle or background garbage collection in skill runners.

---

## 7. Data Model and Eventing
- Event envelopes carry `trace_id`, `session_id`, `actor_id`, `capability_tier`, and `privacy_scope`.
- Schemas in protobuf; evolve via Buf with backward-compatible tests.
- Event store: SQLite with LiteFS or DuckDB for session timelines and replay (opt-in retention; privacy defaults to minimal).
- Embeddings (optional and local): `sqlite-vec` or a local vector database; disabled by default and gated by per-user consent.

---

## 8. Skill Plugin System

Goals: safety, portability, hot reload, great developer experience.

### Execution Targets
- WASM (WASI) via wasmtime or Extism for default skills using capability-based host calls (storage, bus publish, timers, device I/O, TTS or STT access gates).
- Containerized skills (Docker or Podman) for heavy dependencies; supervised and isolated.
- Native adapters (Rust, Python, TypeScript SDKs) for high-performance or trusted skills.

### Skill Package
- `skill.yaml` manifest: name, version, permissions, inputs and outputs, intents, example utterances, UI surfaces, resource budgets.
- Runtime bundle: WASM module or container image with optional assets (icons, templates).
- Semantic layer: capture positive or negative examples and payload schemas so the LLM router can learn or route without code changes.

### Lifecycle
- Supervisor per skill type with health checks, crash loops, and rolling upgrades.
- Audit log per invocation with structured events for observability.
- Hot reload and version pinning per household.

---

## 9. Security and Privacy
- Network: mTLS between services, local certificate authority, scheduled certificate rotation.
- Authorization: capability tokens scoped to skills with least privilege by default.
- Data: at-rest encryption for optional memories, redaction policies, per-user consent toggles, privacy scopes per session.
- Telemetry: no default telemetry; diagnostics are user-triggered and remain local unless explicitly shared.

---

## 10. Observability (Local-Only by Default)
- Tracing: OpenTelemetry exported to Tempo or Jaeger locally.
- Metrics: Prometheus with Grafana dashboards (latency, tokens per second, queue depth, QoS downgrades).
- Logs: Loki or Vector into local storage.
- Service-level objectives: p95 end-to-end under 200 ms (short replies), error rate, model warm latency, STT partial lead time.

---

## 11. Deployment Topologies
- Single-node (Starter): all services on one Mac mini.
- Dual-node (Balanced): hub (router, LLM, skills) plus inference node (STT and Kokoro TTS).
- Tri-node (Resilient): add overflow node with mirrored LLM and STT; Serf membership plus three-node NATS cluster.
- Mesh (Advanced): additional Minis or GPUs, display nodes, multiple pucks, MQTT at the edge bridged into NATS, optional 10 GbE or Thunderbolt.

---

## 12. Open-Core Boundary (Transparency)

### Always Open (FOSS)
- Core bus integration, router, SDKs, skill runtime (WASM), reference skills, schemas, single-home dashboards, local observability stack.

### Paid or Pro (Opt-In, Does Not Cripple Core)
- Remote assist relay (end-to-end encrypted), multi-home or fleet management, policy packs (RBAC, audit or compliance export), advanced analytics, premium voices or models (including advanced Kokoro voice packs), managed update channels, certified hardware kits.

### Governance
- Public RFCs, roadmap, and license clarity from day one. Self-hosted parity tests ensure paid tiers are not patching artificial limits.

---

## 13. Initial Tech Choices (Proposed)
- Language: Go (router and services) plus Python (ML glue) plus Rust (WASM host and performance-critical parts).
- Bus: NATS with JetStream; Redis Streams and MQTT as adapters.
- Schemas: protobuf with Buf; gRPC where request-reply fits; NATS subjects for pub or sub.
- Models: faster-whisper; llama.cpp or Ollama (3B-8B quant); Kokoro for TTS.
- WASM: wasmtime plus Extism with capability-based host calls.
- Control plane: Serf plus a small Raft implementation (only when needed).
- Store: SQLite with LiteFS or DuckDB; `sqlite-vec` for embeddings.
- Observability: OpenTelemetry, Prometheus, Grafana, Tempo or Jaeger, Loki.

---

## 14. Risks and Mitigations
- Latency regressions with multi-hop.  
  **Mitigation:** enforce streaming, QoS tiers, locality hints, continuous latency tests, and prompt caching.
- Skill sandbox escapes.  
  **Mitigation:** make WASM the default, use a strict capability model, add fuzzing and security reviews, run skills under separate system users or namespaces.
- Complexity creep.  
  **Mitigation:** start with static configuration, add the control plane only when pain is real, publish a clear MVP boundary.
- Community trust (open-core).  
  **Mitigation:** publish boundaries, governance, and parity tests early; never paywall core or basic extensibility.

---

## 15. Roadmap (MVP to Six Months)

### Weeks 1-4 (MVP)
- NATS bus, router, STT to LLM to Kokoro pipeline on a single node, OpenTelemetry traces.
- Skill SDK v0 and WASM host with two reference skills (timers, smart-home bridge).
- Basic web UI (sessions, transcripts, timers) plus e-ink display proof of concept.

### Weeks 5-8
- Node discovery (Serf), capability tags, dual-node split (STT and Kokoro offload).
- QoS tiers and hot prompt cache; local event store; consent settings.

### Weeks 9-12
- Container skill runner, audit logs, Grafana dashboards.
- Open-core statement, contribution guide, RFC process.

### Months 4-6
- Prototype Pro features (remote relay E2EE, simple fleet view).
- Ambient display widgets pack; WASM skill marketplace alpha.
- Latency hardening and model menu (3B to 8B auto-tiering).

---

## 16. Appendices

### A. Subject Taxonomy (NATS)
`audio.frame`, `stt.text.partial`, `stt.text.final`, `nlu.request`, `nlu.response`, `skill.invoke`, `skill.result`, `tts.request`, `tts.stream`, `visual.card.update`, `ctrl.node.announce`, `ctrl.node.health`, `obs.span`, `obs.metric`, `obs.log`

### B. Quality Tiers
- `qos.fast` (<=200 ms, 3B quant, minimal verbalization).
- `qos.balanced` (<=400 ms, 7B quant, fuller phrasing).
- `qos.rich` (<=1 s, 7B-13B, higher-quality Kokoro voices, richer responses).

### C. Privacy Scopes
- `scope.ephemeral` (no retention).
- `scope.session` (kept for session replay or debug).
- `scope.persistent` (explicit consent only; per actor; exportable).

---

**TL;DR:** Loqa is a platform for private, multimodal, local AI. Start simple (single node, voice), architect for growth (capabilities, clustering, WASM skills), and be transparent about the open-core boundary from day one.
