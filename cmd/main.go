package main

import (
	"database/sql"
	"log"

	"gofr.dev/pkg/gofr"

	"UDC_management/internal/handlers"
	"UDC_management/internal/repository"
	"UDC_management/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/students")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := &handlers.StudentHandler{Service: studentService}

	app := gofr.New()

	app.GET("/students", studentHandler.GetStudentsHandler)
	app.POST("/student", studentHandler.AddStudentHandler)
	app.PUT("/student/:id", studentHandler.UpdateStudentHandler)
	app.DELETE("/student/:id", studentHandler.DeleteStudentHandler)

	app.Start()
}
