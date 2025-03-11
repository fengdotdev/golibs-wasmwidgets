OUTPUT_DIR=output
MAINGO=cmd/sandbox/main.go
WEBASSETS=web



# Intructions
# make all


all: clean build

# Build wasm
build:
	mkdir -p $(OUTPUT_DIR)
	GOOS=js GOARCH=wasm go build -o $(OUTPUT_DIR)/main.wasm $(MAINGO)
	cp -r $(WEBASSETS)/* $(OUTPUT_DIR)

# Clean output directory
clean:
	rm -rf $(OUTPUT_DIR)

# Build server
serverw:
	GOOS=windows GOARCH=amd64
	go build -o ${OUTPUT_DIR}/server.exe cmd/server/main.go

serverl:
	GOOS=linux GOARCH=amd64
	go build -o ${OUTPUT_DIR}/server cmd/server/main.go

# Run server
live:
	./${OUTPUT_DIR}/server -folderoutput=$(OUTPUT_DIR) -port=9080