package main



//simple server for wasm
import (
	"fmt"
	"net/http"
)
func main() {

	server := http.FileServer(http.Dir("web"))
	http.Handle("/", server)

}