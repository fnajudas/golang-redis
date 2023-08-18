package routes

import (
	"fmt"
	"golangredis/config/mysql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	getData "golangredis/controller/getRedis"
	setData "golangredis/controller/setRedis"
)

const (
	writeRTO = 30
	readRTO  = 30
)

type Routes struct {
	Router  *mux.Router
	Redis   *redis.Client
	SetData *setData.SetRedis
	GetData *getData.Handler
}

func (r *Routes) Run(port string) {
	routes := mux.NewRouter()
	baseURL := os.Getenv("BASE_URL_PATH")

	if len(baseURL) > 0 && baseURL != "/" {
		routes.PathPrefix(baseURL).HandlerFunc(mysql.URLRewriter(routes, baseURL))
	}

	routes.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("Aku sehat karena duit wkwk")
	}).Methods("GET", "OPTIONS")

	routes.HandleFunc("/set", r.SetData.SetData).Methods("POST")
	routes.HandleFunc("/get", r.GetData.GetDataRedis).Methods("GET")

	r.Router = routes
	//C. Serving RESTful HTTP to Clients
	log.Print(fmt.Sprintf("[HTTP SRV] Listening on port :%s", port))
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: writeRTO * time.Second,
		ReadTimeout:  readRTO * time.Second,
	}
	log.Panic(srv.ListenAndServe())
}
