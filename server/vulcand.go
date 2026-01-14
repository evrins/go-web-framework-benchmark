package server

import (
	"fmt"

	"github.com/vulcand/route"
)

// vulcan
func startVulcan(port int) {
	mux := route.NewMux()
	expr := fmt.Sprintf(`Method("%s") && Path("%s")`, "GET", "/hello")
	mux.HandleFunc(expr, helloHandler)

	listenAndServe(port, mux)
}
