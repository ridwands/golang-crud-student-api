# golang-crud-student-api
## Setting Database
Go To infrastructure/mysql.go, adjust it
```
db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/golang_api")
```
## Test CRUD
GET /students
```
curl -X GET http://localhost:1234/students | json_pp 
```
Result
```
{
   "status" : 2200,
   "message" : "success",
   "data" : [
      {
         "student_name" : "Ridwan",
         "student_nim" : "1622009",
         "student_id" : "24"
      }
   ]
}
```
POST /student
```
curl -X POST http://localhost:1234/student --data '{"student_nim" : "1622010", "student_name" : "Arman"}' --header 'Content-Type: application/json'
```
Result
```
{
   "status" : 2200,
   "message" : "success",
   "data" : [
      {
         "student_id" : "",
         "student_name" : "Arman",
         "student_nim" : "1622010"
      }
   ]
}
```
PUT /student, paramater {student_id}
```
curl -X PUT http://localhost:1234/student/24 --data '{"student_nim" : "1622009", "student_name" : "Ridwan Dwi Susilo"}' --header 'Content-Type: application/json' | json_pp
```
Result
```
{
   "data" : [
      {
         "student_name" : "Ridwan Dwi Susilo",
         "student_nim" : "1622009",
         "student_id" : ""
      }
   ],
   "status" : 2200,
   "message" : "success"
}

```
DELETE /student, parameter {student_id}
```
curl -X DELETE http://localhost:1234/student/30 | json_pp 
```
Result
```
{
   "data" : null,
   "message" : "success",
   "status" : 2200
}
```