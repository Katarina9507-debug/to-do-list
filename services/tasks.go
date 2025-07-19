package services

import (
	"to-do_list/models"
	"to-do_list/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(name string) error {
	task := models.Task{
		Name:   name,
		Status: false,
	}
	return s.repo.AddTask(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) MarkTaskAsDone(id uint) error {
	return s.repo.MarkTaskAsDone(id)
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	return s.repo.ListTasks()
}
