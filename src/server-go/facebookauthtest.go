package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	"github.com/dkumor/acmewrapper"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world.\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloServer)

	w, err := acmewrapper.New(acmewrapper.Config{
		Domains: []string{"example.com", "www.example.com"},
		Address: ":443",

		TLSCertFile: "cert.pem",
		TLSKeyFile:  "key.pem",

		RegistrationFile: "user.reg",
		PrivateKeyFile:   "user.pem",

		TOSCallback: acmewrapper.TOSAgree,
	})

	if err != nil {
		log.Fatal("acmewrapper: ", err)
	}

	tlsconfig := w.TLSConfig()

	listener, err := tls.Listen("tcp", ":443", tlsconfig)
	if err != nil {
		log.Fatal("Listener: ", err)
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   mux,
		TLSConfig: tlsconfig,
	}
	server.Serve(listener)
}
