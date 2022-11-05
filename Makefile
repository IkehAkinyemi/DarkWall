## audio: run the cmd/app application with custom directory
.PHONY: audio
audio:
	@echo 'Add the absolute path to your music directory ↲ '
	@echo 'command: make audio dir=<directory>'
	go run ./cmd/app -dir=${music-dir}


## cmd/app: run the cmd/app application default directory
.PHONY: cmd/app
cmd/app:
	@echo 'Running application...'
	go run ./cmd/app


## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## build/api: build the cmd/api application
.PHONY: build/api
build/app:
	@echo 'Building cmd/app...'
	go build  -o=./bin/app ./cmd/app
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/app ./cmd/app