package main

import ( 
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Go First API Testing...")
	fmt.Println(`GO
API 
TESTING
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil{
		log.Error(err)
	}

}