package model

type Student struct {
	Student_id   string `json:"student_id"`
	Student_nim  string `json:"student_nim"`
	Student_name string `json:"student_name"`
}

type Response struct {
	Code    int       `json:"status"`
	Message string    `json:"message"`
	Data    []Student `json:"data"`
}
