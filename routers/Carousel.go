package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"path"
	"log"
	"../libs"
	"io/ioutil"
)

func Carousel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("[%s] %s | %s", r.Method, r.Host, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	if ps.ByName("name") == "0" {
		w.Write(libs.GetPhotos(r.Host, path.Join("carousel")))
	} else {
		pName := ps.ByName("name")
		b, err := ioutil.ReadFile(path.Join("carousel", pName))
		if err != nil {
			http.Error(w, "Not Found", 404)
		}
		w.Write(b)
	}

}