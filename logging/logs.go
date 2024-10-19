package logging

import (
	"log"
	"net/http"
)

const (
	INFO_LEVEL  = "\033[32mINFO\033[0m"
	ERROR_LEVEL = "\033[31mERROR\033[0m"
)

func ErrorLog(message string) {
	log.Printf("%s %s", ERROR_LEVEL, message)
}
func LogRequest(r *http.Request) {
	log.Printf("%s %s %s %s", INFO_LEVEL, r.Method, r.URL, r.RemoteAddr)
}

func LogResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s", INFO_LEVEL, r.Method, r.URL, r.RemoteAddr)
}
