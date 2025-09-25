.PHONY: lint fmt fmt-check

lint:
	npx --yes markdownlint-cli2 "**/*.md"
