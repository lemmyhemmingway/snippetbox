clean:
	go fmt ./cmd/web
	go fmt ./internal/models/
run: clean
	go run ./cmd/web
