package router

import (
	"copy-s3-object/handler"
	"copy-s3-object/service"

	"github.com/gorilla/mux"
)

func NewWebRouter(s3Service service.S3Service) *mux.Router {
	router := mux.NewRouter()
	s3Handler := handler.NewS3Handler(s3Service)
	router.HandleFunc("/health", s3Handler.GetHealth).Methods("GET")
	router.HandleFunc("/copy", s3Handler.CopyObjects).Methods("POST")

	return router
}
