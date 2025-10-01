# Workstream F – Value-Add Offerings Roadmap

Issue reference: [loqa-meta#28](https://github.com/loqalabs/loqa-meta/issues/28)

## Objectives
- Generate sustainable revenue without compromising Loqa’s local-first, open-core promise.
- Deliver optional services/packs that lower adoption friction for households, makers, and businesses.
- Validate demand before investing in large-scale infrastructure.

## Target Segments
| Segment | Needs | Potential offering |
| --- | --- | --- |
| Privacy-first households | Simple setup, family automations, trust | Managed updates, curated skill bundles, Home Hub hardware kit |
| Makers / OSS community | Faster experimentation, extensibility | Premium skill packs (advanced STT models, sensors), extension marketplace |
| Small teams / boutiques | Reliability, support, observability | Managed telemetry dashboards, incident response SLAs, deployments support |

## Candidate Offerings (Wave 0–1)
1. **Managed Update Channel**
   - Nightly bundle verification + stable release cadence.
   - CLI command (`loqa-updater`) pulls signed builds, applies rollback-safe update.
   - Pricing: subscription per cluster.
   - Requirements: signing infrastructure, update agent, support SLA.

2. **Premium Skill Packs**
   - Paid collections (e.g., productivity pack, smart-home pack, ambient wellness pack).
   - Distributed via marketplace with license flag.
   - Requirements: packaging format (RFC-0003), licensing meta, payment integration (phase 2).

3. **Observability & Telemetry Service**
   - Hosted Grafana + alerting tuned for Loqa metrics.
   - Optional secure tunnel from customer nodes (opt-in only).
   - Requirements: remote collector or guidelines for push-gateway, minimal PII policy.

4. **Hardware Starter Kits**
   - Pre-imaged Mac mini / Pi kits with skills installed, nice enclosure.
   - Partnerships with hardware vendors (see [loqa-meta#29](https://github.com/loqalabs/loqa-meta/issues/29)).

## Experiments & Validation
| Experiment | Metric | Status |
| --- | --- | --- |
| Dogfooding cluster update agent | Update success rate, rollback frequency | Pending (#17–#22) |
| Survey community on premium skills | Responses, top requested domains | TODO |
| Observability SaaS pilot with early adopters | DAU, alert acknowledgements | TODO |

## Dependencies
- Marketplace MVP (RFC-0003) for distribution and metadata.
- Nightly bundle automation (#19) and dogfooding learnings.
- Partner outreach pipeline (#29) for hardware collaborations.
- Legal review for SaaS terms & licensing.

## Next Actions
1. Draft product briefs for each offering (scope, pricing hypothesis, risks).
2. Schedule user interviews (target 5 households, 5 makers, 3 teams).
3. Prototype managed update channel in the dogfood cluster.
4. Align with finance/legal on billing infrastructure options.

## Open Questions
- What payment provider fits privacy-first brand without heavy overhead?
- Do we need an offline license model for air-gapped deployments?
- Should premium skill packs ship source-available or binary-only?

## Tracking
Progress will be updated in the [Workstream F project board](../roadmap/MVP_BACKLOG.md) once established.
