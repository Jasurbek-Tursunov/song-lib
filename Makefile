# ==============================================================================
# Main

copy-default-env:
	cp example.env .env

run-mock-external:
	go mod tidy
	go run ./cmd/mock-external-api/main.go

run:
	go mod tidy
	go run ./cmd/main.go

build-release:
	go mod tidy
	go build -o song-lib ./cmd/main.go

run-release:
	chmod +x song-lib
	./song-lib