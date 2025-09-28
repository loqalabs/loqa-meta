.PHONY: build validate clean

build:
	mkdir -p build
	tinygo build -o build/my-skill.wasm -target=wasi ./src

validate:
	go run ./cmd/loqa-skill validate --file skill.yaml

clean:
	rm -rf build
