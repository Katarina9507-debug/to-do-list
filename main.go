package main

import (
	"fmt"
	"log"

	"to-do_list/tasks"
)

func main() {
	storage := &tasks.TaskStorage{}

	// создаем слайс задач
	tasksToAdd := []tasks.Task{
		{ID: 1, Name: "Завершить учебный проект", Status: false},
		{ID: 2, Name: "Протестировать код", Status: false},
		{ID: 3, Name: "Обновить документацию", Status: false},
	}

	// циклом добавляем все задачи в хранилище
	for _, task := range tasksToAdd {
		if err := storage.AddTask(task); err != nil {
			log.Printf("Ошибка добавления задачи %d: %v", task.ID, err)
		} else {
			log.Printf("Задача добавлена: [%d] %s", task.ID, task.Name)
		}
	}

	// отмечаем задачу как выполненную, в случае её отсутсвия логируем ошибку
	if err := storage.MarkTaskAsDone(true, 1); err != nil {
		log.Printf("Ошибка отметки задачи: %v", err)
	} else {
		log.Printf("Задача отмечена как выполненная")
	}

	// ищем задачу по ID
	if task, err := storage.FindTaskByID(2); err != nil {
		log.Printf("Ошибка поиска: %v", err)
	} else {
		fmt.Printf("Найдена задача: [%d] %s - %v\n",
			task.ID, task.Name, task.Status)
	}

	// удаляем задачу, в случае её отсутсвия логируем ошибку
	if err := storage.DeleteTask(tasks.Task{ID: 3}); err != nil {
		log.Printf("Ошибка удаления: %v", err)
	} else {
		log.Println("Задача успешно удалена")
	}

	// выводим список существющих задач
	fmt.Println("\nТекущий список задач:")
	storage.ListTasks()
}
