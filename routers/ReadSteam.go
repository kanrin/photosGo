package routers


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"../libs"
)

func ReadSteam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("[%s] %s | %s", r.Method, r.Host, r.URL)
	pName := ps.ByName("name")
	b, err := libs.GetOssFile(pName)
	if err != nil {
		http.Error(w, "Not Found", 404)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Write(b)
}