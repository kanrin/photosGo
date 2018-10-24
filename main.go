package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kylelemons/go-gypsy/yaml"
	"net/http"
	"log"
	"./routers"
)


func main() {
	router := httprouter.New()
	router.GET("/", routers.Index)
	router.GET("/photos/:name", routers.PhotoUrl)
	config, _ := yaml.ReadFile("config.yml")
	port, _ := config.Get("port")
	log.Printf("listen port: %s", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}