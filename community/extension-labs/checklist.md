# Extension Publishing Checklist

Use this checklist before submitting a skill to Extension Labs or the marketplace:

- [ ] Manifest validates with `go run ./cmd/loqa-skill validate --file skill.yaml`
- [ ] `metadata.version` follows SemVer and matches your release tag
- [ ] All publish/subscribe subjects are declared and documented
- [ ] Permissions are minimal and justified (include reasoning in README)
- [ ] Provide build instructions and runtime dependencies
- [ ] Attach tests or replay scripts demonstrating core workflows
- [ ] Include LICENSE and changelog files in the package
- [ ] Post a “Show and Tell” entry in Discussions with screenshots/logs
- [ ] (Optional) Prepare marketplace metadata as per RFC-0003 when ready
