package repository

import (
	"UDC_management/internal/models"
	"database/sql"

	"gofr.dev/pkg/gofr"
)

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) GetStudents(ctx *gofr.Context) ([]models.Student, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT id, name, reason FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Reason); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (r *StudentRepository) AddStudent(ctx *gofr.Context, student models.Student) error {
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO students (id, name, reason) VALUES (?, ?, ?)", student.ID, student.Name, student.Reason)
	return err
}

func (r *StudentRepository) UpdateStudent(ctx *gofr.Context, student models.Student) error {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE students SET name=?, reason=? WHERE id=?", student.Name, student.Reason, student.ID)
	return err
}

func (r *StudentRepository) DeleteStudent(ctx *gofr.Context, studentID int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM students WHERE id=?", studentID)
	return err
}
