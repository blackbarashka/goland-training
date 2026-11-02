package service

import (
	"context"
	"task_list/internal/features/tasks"
	"task_list/internal/features/tasks/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) Create(ctx context.Context, title, description string) (*tasks.Task, error) {
	newTask := tasks.NewTask(title, description)
	err := s.repo.Create(ctx, &newTask)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}
func (s *TaskService) Get(ctx context.Context, id string) (*tasks.Task, error) {
	task, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) List(ctx context.Context) ([]*tasks.Task, error) {
	tasksList, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return tasksList, nil
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *TaskService) Update(ctx context.Context, id string, title, description string, completed bool) error {
	taskToUpdate := tasks.Task{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	err := s.repo.Update(ctx, id, &taskToUpdate)
	if err != nil {
		return err
	}
	return nil

}
