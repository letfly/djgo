package djgo

import (
	"fmt"
	"net/http"
)

// Declare variables
var (
	HttpAddr string
	HttpPort int
)

func init() {
	HttpAddr = ""
	HttpPort = 8000
}

func NewApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello") // the w is written to the output to the client
}

// Run interface
func Run() {
	http.HandleFunc("/", NewApp)
	addr := fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("err=%s", err)
	}
}
