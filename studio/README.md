# Loqa Studio & Marketplace Guidelines

This document complements [RFC-0003](../rfcs/RFC-0003_loqa_marketplace_mvp.md) and the **Composable Open Core** strategy by describing how creators can publish add-ons through Loqa Studio, and how Ambiware Labs vets submissions to keep the ecosystem safe and delightful.

## What You Can Publish

Loqa Studio accepts extensions that respect Loqa’s local-first principles:

| Category | Examples | Notes |
| --- | --- | --- |
| Persona Packs | Custom assistant personalities, domain-specific prompts, voice tuning bundles | Must declare data sources and safety constraints. |
| Premium Skills / Add-ons | Advanced automations, integrations, analytics dashboards | Built against the [skills spec](https://github.com/loqalabs/loqa-core/blob/main/docs/skills/SPEC.md); may target WASM or exec runtimes. |
| Loqa Cloud Integrations | Optional encrypted sync, shared memory spaces, multi-device orchestration | Must use end-to-end encryption and pass security review. |
| Training & Courses | Workshops, onboarding curricula, certification tracks | Provide syllabus outline and delivery format. |
| Pre-flashed Hardware Kits | Partner-built kits ready to plug into Loqa | Coordinate with Ambiware hardware program (see Workstream F roadmap).

Creators may also submit free/OSS extensions—the marketplace supports both commercial and community listings.

## Submission Workflow

1. **Prepare your package**
   - Follow the Extension Labs [templates](../community/extension-labs/) and the skills specification.
   - Include licensing information, changelog, and usage documentation.
2. **Open a Marketplace PR**
   - Submit metadata to the upcoming `loqa-marketplace` repository (see RFC-0003 for schema).
   - Clearly state pricing (if any), support contact, and dependencies.
3. **Vetting & Safety Review**
   - Ambiware Labs reviews for quality, security, and AI safety considerations (e.g., prompt injection handling, data provenance).
   - Automated checks validate manifests, signatures, and binary integrity.
4. **Approval & Publishing**
   - Once approved, the catalog entry is published and becomes installable via `loqa-skill registry install <name>`.
   - Updates follow the same process; semantic versioning is required.

> Until the community steering group is formed, Ambiware Labs’ core team makes final publication decisions. The process will transition to a shared governance model over time.

## Licensing Flexibility

- We encourage permissive or copyleft licenses for add-ons, but proprietary listings are permitted if they align with Loqa’s ethical guidelines.
- All packages must disclose their license in the marketplace metadata.
- Creators retain revenue from commercial listings, minus any agreed marketplace fees (details forthcoming).

## Quality Guidelines

- **Local-first execution** – Add-ons must default to local processing unless the user explicitly connects to external services.
- **Security** – Secrets handled via environment variables or secure vaults; no plain-text storage.
- **Telemetry** – Opt-in only, with transparent data usage disclosure.
- **Support** – Paid listings must document response expectations and support channels.
- **Accessibility** – Provide concise documentation, example commands, and fallback behaviours.

## Monetization & Revenue Share

- Pricing tiers (indicative): free, pay-once, subscription, bundle.
- Revenue share policy will be defined with early partners; creators will receive transparent terms before launch.
- Co-marketing opportunities are coordinated through the [partner outreach pipeline](../community/outreach/partner_pipeline.md).

## Next Steps & Resources

- Track marketplace implementation in [RFC-0003](../rfcs/RFC-0003_loqa_marketplace_mvp.md) and related issues.
- Use the [Extension Publishing Checklist](../community/extension-labs/checklist.md) before submission.
- Join the upcoming Discussions thread for marketplace planning to share feedback.

Loqa Studio exists to help creators thrive while keeping the core platform open and user-sovereign. Let’s build responsibly together.
