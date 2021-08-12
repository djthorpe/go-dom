package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	flagPort = flag.String("port", ":9090", "Port to listen to requests on")
)

func main() {
	flag.Parse()
	fmt.Println("Listening on ", *flagPort)
	if wd, err := os.Getwd(); err != nil {
		fmt.Println("Failed to start server", err)
		os.Exit(-1)
	} else if err := http.ListenAndServe(*flagPort, http.FileServer(http.Dir(wd))); err != nil {
		fmt.Println("Failed to start server", err)
		os.Exit(-1)
	}
}
