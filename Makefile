run:
	go run ./cmd/uptimebot


build:
	CGO_ENABLED=0 go build -o ./bin/uptimebot ./cmd/uptimebot
