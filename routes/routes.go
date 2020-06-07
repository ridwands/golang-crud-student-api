package routes

import (
	"student_api/controller"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/students", controller.Studentall).Methods("GET")
	router.HandleFunc("/student/{student_id}", controller.Studentone).Methods("GET")
	router.HandleFunc("/student", controller.StudentPost).Methods("POST")
	router.HandleFunc("/student/{student_id}", controller.StudentUpdate).Methods("PUT")
	router.HandleFunc("/student/{student_id}", controller.StudentDelete).Methods("DELETE")
	return router
}
