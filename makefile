run:
	GOOS=js GOARCH=wasm go build ./cmd/sandbox/main.go -o main.wasm
