package repository

import (
	"database/sql"
	"fmt"

	"to-do_list/errors"
	"to-do_list/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) AddTask(task models.Task) error {
	_, err := r.db.Exec(
		"INSERT INTO tasks_list (name, status) VALUES ($1, $2)",
		task.Name,
		task.Status,
	)
	if err != nil {
		return &errors.MyErrors{
			Method: "AddTask",
			Code:   400,
			Msg:    fmt.Sprintf(errors.ErrInsertingTask, task.ID),
		}
	}

	return nil
}

func (r *TaskRepository) DeleteTask(id uint) error {
	_, err := r.db.Exec("DELETE FROM tasks_list WHERE id = $1", id)
	return err
}

func (r *TaskRepository) MarkTaskAsDone(id uint) error {
	_, err := r.db.Exec("UPDATE tasks_list SET status = true WHERE id = $1", id)
	return err
}

func (r *TaskRepository) ListTasks() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT id, name, status FROM tasks_list")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
