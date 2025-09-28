.PHONY: package validate clean

package:
	tar -czf my-exec-skill.tar.gz skill.yaml scripts

validate:
	go run ./cmd/loqa-skill validate --file skill.yaml

clean:
	rm -f my-exec-skill.tar.gz
