package main

import (
	"GOLANG/constant"
	"GOLANG/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ownServer := http.NewServeMux()
	ownServer.HandleFunc(constant.AverageEndpoint, handler.HandleAverage)
	fmt.Println(constant.ServerStartupMessage)
	log.Fatal(http.ListenAndServe(":8080", ownServer))

}
