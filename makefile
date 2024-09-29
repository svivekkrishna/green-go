.PHONY: run-cli
run-cli:
	cd cli && go run main.go -f test.yaml

.PHONY: build-cli
build-cli:
	go build -o green-go cli/main.go

.PHONY: build-web
build-web:
	echo "Not implemented yet. go build -o green-go-web web/main.go"

.PHONY: build
build: build-cli build-web
	echo "Done."

.PHONY: clean
clean:
	rm -rf green-go green-go-web