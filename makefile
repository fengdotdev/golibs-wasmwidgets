OUTPUT_DIR=output
MAINGO=cmd/sandbox/main.go
WEBASSETS=assets



# Intructions
# make all


env:
	GOOS=js GOARCH=wasm


all: cleanwasm build

# Build wasm
build:
	mkdir -p $(OUTPUT_DIR)
	GOOS=js GOARCH=wasm go build -o $(OUTPUT_DIR)/main.wasm $(MAINGO)

# Clean output directory
clean:
	rm -rf $(OUTPUT_DIR)

cleanwasm:
	rm -rf $(OUTPUT_DIR)/main.wasm

assets:
	mkdir -p $(OUTPUT_DIR)
	cp -r $(WEBASSETS)/* $(OUTPUT_DIR)

# Build server
serverw:
	GOOS=windows GOARCH=amd64
	go build -o ${OUTPUT_DIR}/server.exe cmd/server/main.go

serverl:
	GOOS=linux GOARCH=amd64
	go build -o ${OUTPUT_DIR}/server cmd/server/main.go

serverm:
	GOOS=darwin GOARCH=amd64
	go build -o ${OUTPUT_DIR}/server cmd/server/main.go

# Run server
live:
	./${OUTPUT_DIR}/server -folderoutput=$(OUTPUT_DIR) -port=9080