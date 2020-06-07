package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"student_api/infrastructure"
	"student_api/model"

	"github.com/gorilla/mux"
)

func Studentall(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	var data []model.Student
	var response model.Response

	db := infrastructure.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM student_table")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&student.Student_id, &student.Student_nim, &student.Student_name); err != nil {
			log.Fatal(err.Error())

		} else {
			data = append(data, student)
		}
	}

	if data == nil {
		response.Code = 2300
		response.Message = "success"
		response.Data = data
	} else {
		response.Code = 2200
		response.Message = "success"
		response.Data = data
	}

	fmt.Println(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func Studentone(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	var data []model.Student
	var response model.Response

	db := infrastructure.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	student_id := vars["student_id"]

	rows, err := db.Query("SELECT * FROM student_table WHERE student_id=?", student_id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&student.Student_id, &student.Student_nim, &student.Student_name); err != nil {
			log.Fatal(err.Error())

		} else {
			data = append(data, student)
		}
	}

	if data == nil {
		response.Code = 2300
		response.Message = "success"
		response.Data = data
	} else {
		response.Code = 2200
		response.Message = "success"
		response.Data = data
	}

	fmt.Println(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func StudentPost(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	var data []model.Student
	var response model.Response

	db := infrastructure.Connect()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO student_table (student_nim, student_name) values (?,?)",
		student.Student_nim, student.Student_name)

	if err != nil {
		panic(err)
	}

	data = append(data, student)

	response.Code = 2200
	response.Message = "success"
	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func StudentUpdate(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	var data []model.Student
	var response model.Response

	db := infrastructure.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	student_id := vars["student_id"]

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("UPDATE student_table set student_name = ?, student_nim = ? where student_id = ?", student.Student_name, student.Student_nim, student_id)

	if err != nil {
		panic(err)
	}

	data = append(data, student)

	response.Code = 2200
	response.Message = "success"
	response.Data = data
	log.Println("Successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func StudentDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	var response model.Response

	db := infrastructure.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	student_id := vars["student_id"]

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM student_table WHERE student_id=?", student_id)

	if err != nil {
		panic(err)
	}

	response.Code = 2200
	response.Message = "success"
	log.Println("Successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
