package tasks

import (
	"fmt"
	"slices"

	"to-do_list/errors"
)

func (storage *TaskStorage) AddTask(task Task) error {
	if storage == nil {
		return fmt.Errorf(errors.ErrStorageNil)
	}
	// проверка на существование ID задачи, в случае совпадения ID возвращаем ошибку
	for _, i := range *storage {
		if i.ID == task.ID {
			return &errors.MyErrors{
				Method: "AddTask",
				Code:   409, //ошибка дубликатов
				Msg:    fmt.Sprintf(errors.ErrTaskExists, task.ID),
			}
		}
	}

	*storage = append(*storage, task)

	return nil
}

func (storage *TaskStorage) DeleteTask(task Task) error {
	if storage == nil {
		return fmt.Errorf(errors.ErrStorageNil)
	}

	index, err := storage.GetSliceIndex(task)
	if err != nil {
		return &errors.MyErrors{
			Method: "DeleteTask",
			Code:   404,
			Msg:    fmt.Sprintf(errors.ErrTaskNotFound, task.ID),
		}
	}
	*storage = slices.Delete(*storage, index, index+1)

	return nil
}

func (storage *TaskStorage) GetSliceIndex(task Task) (int, error) {
	if storage == nil {
		return -1, fmt.Errorf(errors.ErrStorageNil)
	}

	for index, item := range *storage {
		if task.ID == item.ID {
			return index, nil
		}
	}

	return -1, fmt.Errorf(errors.ErrIndexNotFound, task.ID)
}

func (storage *TaskStorage) MarkTaskAsDone(isDone bool, taskID uint) error {
	if storage == nil {
		return fmt.Errorf(errors.ErrStorageNil)
	}

	// перебираем слайс задач в хранилище и ищем ID задачи, в случае совпадения меняем Status
	for i := range *storage {
		if (*storage)[i].ID == taskID {
			(*storage)[i].Status = isDone
			return nil
		}
	}
	// если совпадения не находятся, возвращаем ошибку
	return &errors.MyErrors{
		Method: "MarkTaskAsDone",
		Code:   404,
		Msg:    fmt.Sprintf(errors.ErrTaskNotFound, taskID),
	}
}

func (storage *TaskStorage) ListTasks() {
	if storage == nil {
		fmt.Errorf(errors.ErrStorageNil)
		return
	}

	isDone := "Не выполнено"
	for _, task := range *storage {
		if task.Status {
			isDone = "Выполнено"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Name, isDone)
	}
}

// метод поиска задач по ID
func (storage *TaskStorage) FindTaskByID(taskID uint) (Task, error) {
	if storage == nil {
		return Task{}, fmt.Errorf(errors.ErrStorageNil)
	}

	var task Task
	for _, item := range *storage {
		if item.ID == taskID {
			task = Task{item.ID, item.Name, item.Status}
			return task, nil
		}
	}
	return Task{}, &errors.MyErrors{
		Method: "FindTaskByID",
		Code:   404,
		Msg:    fmt.Sprintf(errors.ErrTaskNotFound, taskID),
	}
}
