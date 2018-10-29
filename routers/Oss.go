package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"../libs"
	"log"
)

func Oss(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("[%s] %s | %s", r.Method, r.Host, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	list := libs.GetPhotosOss()
	w.Write(list)
}