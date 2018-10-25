package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"../libs"
	"log"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("[%s] %s | %s", r.Method, r.Host, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	time.Sleep(time.Duration(5)*time.Second)
	w.Write(libs.GetPhotos(r.Host))
}