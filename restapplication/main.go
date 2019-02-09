package main

import (
	logger "log"
	"net/http"
	"os"
	"pavan/MAD-Assignment-1/dbrepository"
	handlerlib "pavan/MAD-Assignment-1/restapplication/packages/httphandlers"
	"pavan/MAD-Assignment-1/restapplication/restauranthandler"
	mongoutils "pavan/MAD-Assignment-1/utils"

	"time"

	"github.com/gorilla/mux"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}
func main() {
	pingHandler := &handlerlib.PingHandler{}
	dbname := "restaurant"
	mongosession, err := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	if err != nil {
		logger.Fatal("Error in creating mongoSession")
	}
	dbsession := dbrepository.NewMongoRepository(mongosession, dbname)
	handler := restauranthandler.NewRestaurantHandler(dbsession)
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/restaurant/ping", pingHandler)
	h.Handle("/restaurant/", handler)
	h.Handle("/restaurant/{id}", handler)
	h.Handle("/restaurant/", handler)
	logger.Fatal(http.ListenAndServe(":8085", h))
}
