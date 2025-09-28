# TinyGo Skill Template

Starter template for building a WASM skill using TinyGo.

## Contents
- `skill.yaml` – placeholder manifest aligned with [`docs/skills/SPEC.md`](https://github.com/ambiware-labs/loqa-core/blob/main/docs/skills/SPEC.md)
- `src/main.go` – basic TinyGo entrypoint using the `host` helper
- `Makefile` – build + validate targets

## Usage

```bash
git clone https://github.com/ambiware-labs/loqa-core.git skills-lab
cd skills-lab/community/extension-labs/templates/tinygo-skill-template
make build
make validate
```

> Replace the metadata, subjects, and permissions before publishing.
