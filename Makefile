.PHONY: build dev start test fmt lint clean install docs-dev docs-build

dev :
	go run cmd/main.go

start: 
	./bin/ultrahooks
build:
	go build -o bin/ultrahooks

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -f ultrahooks
	rm -rf .ultrahooks

install:
	go install 

docs-dev:
	bun run dev --prefix docs

docs-build:
	bun run build --prefix docs
