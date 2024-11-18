package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getIpRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s - %s\n", r.RemoteAddr, r.Method, r.URL)
	var remoteAddr = parseRemoteAddr(r.RemoteAddr)
	io.WriteString(w, remoteAddr)
}

func parseRemoteAddr(remoteAddr string) string {
	var addr = strings.Split(remoteAddr, ":")
	return addr[0]
}

func main() {
	http.HandleFunc("/", getIpRoute)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}