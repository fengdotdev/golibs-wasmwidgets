package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "runtime"
)

// Function to open the default browser with a given URL
func openBrowser(url string) {
    var err error

    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }

    if err != nil {
        log.Fatalf("Failed to open browser: %v", err)
    }
}

func main() {
    // Define command-line flags
    folderOutput := flag.String("folderoutput", "./output", "output directory")
    port := flag.Int("port", 8080, "server port")

    // Parse flags
    flag.Parse()

    // Set up the HTTP server to serve static files from the specified directory
    fs := http.FileServer(http.Dir(*folderOutput))
    http.Handle("/", fs)

    // Start the server
    addr := fmt.Sprintf(":%d", *port)
    url := fmt.Sprintf("http://localhost%s", addr)
    fmt.Printf("Serving directory %s at %s\n", *folderOutput, url)

    // Open the browser
    go openBrowser(url)

    log.Fatal(http.ListenAndServe(addr, nil))
}
