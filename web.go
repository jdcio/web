package web

import (
	"crypto/tls"
	"github.com/jdcio/sudoless"
	"github.com/jdcio/tlsconfig"
	"log"
	"net"
	"net/http"
	"time"
)

var p80, p443 net.Listener
var certs []tls.Certificate
var err error

func Load(certGlob string, userName string) {
	p80, err = sudoless.Port(80)
	if err != nil {
		log.Printf("Failed to open port %v", err)
	}
	p443, err = sudoless.Port(443)
	if err != nil {
		log.Printf("Failed to open port %v", err)
	}
	certs = sudoless.Certs(certGlob)
	sudoless.DropPrivileges(userName)
}

func Serve(handler http.Handler) {
	bootHTTP(p80)
	bootHTTPS(handler, p443, certs)
}

func bootHTTP(p80 net.Listener) {
	httpSrv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + req.Host + req.URL.String()
			http.Redirect(w, req, url, http.StatusMovedPermanently)
		}),
	}
	go func() { log.Fatal(httpSrv.Serve(p80)) }()
}

func bootHTTPS(handler http.Handler, p443 net.Listener, certs []tls.Certificate) {
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsconfig.Secure(certs),
		Handler:      handler,
	}
	log.Println(srv.ServeTLS(p443, "", ""))
}
