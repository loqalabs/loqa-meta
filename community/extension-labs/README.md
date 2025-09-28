# Extension Labs

Welcome to Extension Labs — the launchpad for building, sharing, and discovering Loqa skills and adapters.

Our goals:
- Make it easy for creators to ship high-quality skills that respect Loqa’s local-first ethos.
- Showcase community work so operators can quickly find trusted extensions.
- Prepare the ecosystem for the upcoming marketplace MVP.

> First time here? Start with the [Quickstart Workflow](#quickstart-workflow) and the [Skill Templates](templates/) below.

## Quickstart Workflow

1. **Fork or template a starter project**
   - [TinyGo WASM template](templates/tinygo-skill-template/README.md)
   - [Exec adapter template](templates/exec-skill-template/README.md)
2. **Develop locally** using the guidance in [`skills/AUTHORING_GUIDE.md`](https://github.com/ambiware-labs/loqa-core/blob/main/skills/AUTHORING_GUIDE.md) and the formal [`skills spec`](https://github.com/ambiware-labs/loqa-core/blob/main/docs/skills/SPEC.md).
3. **Validate your manifest**
   ```bash
   go run ./cmd/loqa-skill validate --file skill.yaml
   ```
4. **Share your work**
   - Open a “Show and Tell” thread in [GitHub Discussions](https://github.com/ambiware-labs/loqa-core/discussions).
   - Submit a PR to add your template/link to this directory.
   - Prepare for the marketplace by tracking [RFC-0003](https://github.com/ambiware-labs/loqa-meta/blob/main/rfcs/RFC-0003_loqa_marketplace_mvp.md).

## Contributor Showcase

| Skill | Maintainer | Description | Tags |
| --- | --- | --- | --- |
| [timer](https://github.com/ambiware-labs/loqa-core/tree/main/skills/examples/timer) | Ambiware Labs | Simple countdown timer that announces completions. | timers, voice |
| [smart-home-bridge](https://github.com/ambiware-labs/loqa-core/tree/main/skills/examples/smart-home) | Ambiware Labs | Bridges Loqa intents to Home Assistant. | home-automation, voice |

> Have a skill to feature? [Open an issue](https://github.com/ambiware-labs/loqa-meta/issues/new?labels=community&title=Showcase%20submission%3A%20%3Cskill%3E) with links and tags.

## Templates

Starter templates live in the [`templates/`](templates/) subdirectory. Each template includes a README, manifest, and build scripts mirroring the Loqa spec.

- [`templates/tinygo-skill-template`](templates/tinygo-skill-template/README.md)
- [`templates/exec-skill-template`](templates/exec-skill-template/README.md)

Contribute improvements as you discover best practices.

## Best Practices & Reviews

- Follow the [contribution checklist](checklist.md) before publishing.
- Validate permissions and declared subjects to avoid runtime rejections.
- Record test transcripts or scripts to help reviewers verify behaviour.
- Use semantic versioning and changelog entries for every release.

## Roadmap

- Track open tasks in [loqa-meta#27](https://github.com/ambiware-labs/loqa-meta/issues/27).
- Marketplace submission guide will land after RFC-0003 is accepted.
- Looking for collaborators? Drop an idea in the Discussions “Ideas” category.

Let’s build the Loqa ecosystem together.
