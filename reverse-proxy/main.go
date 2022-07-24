package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)

	log.Print("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	link := r.URL.Query().Get("link")
	isSelf := strings.Contains(link, "http://192.168.192.239:8080/")
	if isSelf {
		link = strings.Replace(link, "http://192.168.192.239:8080/?link=", "", -1)
		parsedLink, _ := url.Parse(link)
		link = parsedLink.Path
	}

	fmt.Printf("link: %s \n", link)

	if link != "" {
		u, _ := url.Parse(link)

		fmt.Printf("[reverse proxy server] %s %s \n", time.Now(), u)

		proxy := httputil.NewSingleHostReverseProxy(u)

		r.URL.Host = u.Host
		r.URL.Scheme = u.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = u.Host

		w.Header().Set("Access-Control-Allow-Origin", "*")
		proxy.ServeHTTP(w, r)
	}

}
