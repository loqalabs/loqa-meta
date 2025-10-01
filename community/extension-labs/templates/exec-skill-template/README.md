# Exec Skill Template

Template for skills that call an external executable or script instead of WASM modules.

## Contents
- `skill.yaml` – manifest configured with `runtime.mode: exec`
- `scripts/handler.sh` – example shell handler reading JSON from stdin
- `Makefile` – helper targets for packaging and validation

## Usage

```bash
git clone https://github.com/loqalabs/loqa-core.git skills-lab
cd skills-lab/community/extension-labs/templates/exec-skill-template
make package
make validate
```

The runtime will invoke your script with the event payload on stdin. Ensure the manifest permissions and subjects align with the behaviour.
