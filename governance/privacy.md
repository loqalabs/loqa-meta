# Privacy & Data Principles

Loqa is built for local-first, privacy-preserving ambient intelligence. These principles guide every product decision, including optional value-add services.

## Local-First by Default
- All core functionality runs on hardware you control (workstations, minis, SBCs).
- No cloud services are required to install, run, or extend Loqa.
- Offline operation is a first-class requirement across runtime, skills, and tooling.

## Zero Telemetry Without Consent
- The core runtime does **not** collect usage analytics, fingerprints, or behavioural data.
- Logs stay on the device unless an operator explicitly exports them.
- Optional diagnostics will always be opt-in and transparent about payloads.

## Optional Loqa Cloud (Encrypted)
- Loqa Cloud provides convenience sync for those who opt in.
- All data is end-to-end encrypted before it leaves the device; Ambiware Labs cannot read user content.
- Users can inspect, revoke, or self-host Loqa Cloud components as they become available.

## No Silent Calls Home
- Network requests are explicit and documented (e.g., marketplace downloads that the operator triggers).
- The runtime ships with safeguards to block undeclared egress from skills/add-ons.

## User Agency & Transparency
- Configuration changes, permissions, and updates are auditable.
- Skills must declare subjects, permissions, and external calls in their manifest (see [`docs/skills/SPEC.md`](https://github.com/ambiware-labs/loqa-core/blob/main/docs/skills/SPEC.md)).
- Marketplace submissions undergo safety review to ensure they honour these principles.

## Incident Response
- Security disclosures follow [SECURITY.md](../SECURITY.md).
- Privacy-impacting issues are treated as emergencies and communicated openly.

By contributing to Loqa—core or Studio—you agree to uphold these principles.
