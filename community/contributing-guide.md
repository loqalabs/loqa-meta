# Contributing to Loqa

Welcome! This guide highlights easy ways to get involved, how to find beginner-friendly work, and how to propose bigger ideas like RFCs or Loqa Studio add-ons.

## Quick Wins (Great First Contributions)
- Fix typos or clarify docs (`loqa-core/docs`, `loqa-meta/community`).
- Improve skill examples or add new TinyGo snippets (`loqa-core/skills/examples`).
- Share a persona prompt or voice tweak template in Extension Labs.
- Tackle issues labeled `good first issue` or `help wanted` across repositories.

Browse open issues:
- [`loqa-core` issues](https://github.com/ambiware-labs/loqa-core/issues)
- [`loqa-site` issues](https://github.com/ambiware-labs/loqa-site/issues)
- [`loqa-meta` issues](https://github.com/ambiware-labs/loqa-meta/issues)

## Submitting a Skill or Persona
1. Follow the [Extension Labs templates](extension-labs/README.md) and the [skills spec](https://github.com/ambiware-labs/loqa-core/blob/main/docs/skills/SPEC.md).
2. Validate your manifest: `go run ./cmd/loqa-skill validate --file skill.yaml`.
3. Share your work in [Loqa Discussions](https://github.com/ambiware-labs/loqa-core/discussions) under “Show and Tell”.
4. (Optional) Prepare metadata for the upcoming marketplace (see [studio/README.md](../studio/README.md)).

## Proposing an RFC
- Check the roadmap and existing RFCs first.
- Open an RFC issue using the template (`.github/ISSUE_TEMPLATE/rfc.md`).
- Draft the RFC using [rfcs/RFC_TEMPLATE.md](../rfcs/RFC_TEMPLATE.md).
- Request feedback in Discussions; maintainers provide guidance and timelines.

## Studio Add-on Proposals
- Review [Loqa Studio guidelines](../studio/README.md).
- Prepare licensing, pricing (if applicable), and support details.
- Submit a PR to the marketplace repo (coming soon) or coordinate via Discussions until the workflow is live.

## Governance & Roadmap Links
- [Governance docs](../governance)
- [Marketplace RFC](../rfcs/RFC-0003_loqa_marketplace_mvp.md)
- [Value-add roadmap](../roadmap/workstream-f/value_add_roadmap.md)
- [Partner outreach pipeline](../community/outreach/partner_pipeline.md)

## Questions?
- Join [GitHub Discussions](https://github.com/ambiware-labs/loqa-meta/discussions).
- Ping `@ambiware-labs` or email [hello@ambiware.ai](mailto:hello@ambiware.ai).

Thanks for helping shape Loqa—every contribution keeps privacy-focused ambient intelligence in the hands of the community.
