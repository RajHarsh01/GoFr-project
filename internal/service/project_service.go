package service

import (
	"UDC_management/internal/models"
	"UDC_management/internal/repository"

	"gofr.dev/pkg/gofr"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{Repo: repo}
}

func (s *StudentService) GetStudents(ctx *gofr.Context) ([]models.Student, error) {
	return s.Repo.GetStudents(ctx)
}

func (s *StudentService) AddStudent(ctx *gofr.Context, student models.Student) error {
	return s.Repo.AddStudent(ctx, student)
}

func (s *StudentService) UpdateStudent(ctx *gofr.Context, student models.Student) error {
	return s.Repo.UpdateStudent(ctx, student)
}

func (s *StudentService) DeleteStudent(ctx *gofr.Context, studentID int) error {
	return s.Repo.DeleteStudent(ctx, studentID)
}
