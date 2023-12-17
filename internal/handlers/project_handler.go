package handlers

import (
	"UDC_management/internal/models"
	"UDC_management/internal/service"
	"net/http"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type StudentHandler struct {
	Service *service.StudentService
}

func (h *StudentHandler) GetStudentsHandler(c *gofr.Context) (interface{}, error) {
	students, err := h.Service.GetStudents(c)
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "Internal Server Error"}
	}

	return students, nil
}

func (h *StudentHandler) AddStudentHandler(c *gofr.Context) (interface{}, error) {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Invalid request payload"}
	}

	if err := h.Service.AddStudent(c, student); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "Internal Server Error"}
	}

	return student, nil
}

func (h *StudentHandler) UpdateStudentHandler(c *gofr.Context) (interface{}, error) {
	studentID, err := strconv.Atoi(c.PathParam("id"))
	if err != nil {
		return nil,
	}

	var updatedStudent models.Student
	if err := c.Bind(&updatedStudent); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Invalid request payload"}
	}

	updatedStudent.ID = studentID

	if err := h.Service.UpdateStudent(c, updatedStudent); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "Internal Server Error"}
	}

	return updatedStudent, nil
}
func (h *StudentHandler) DeleteStudentHandler(c *gofr.Context) (interface{}, error) {
	studentID, err := strconv.Atoi(c.PathParam("id"))
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Invalid student ID"}
	}

	if err := h.Service.DeleteStudent(c, studentID); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "Internal Server Error"}
	}

	students, err := h.Service.GetStudents(c)
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "Internal Server Error"}
	}

	return students, nil
}
