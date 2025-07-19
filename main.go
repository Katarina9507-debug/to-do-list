package main

import (
	"fmt"
	"log"

	"to-do_list/infrastructure/db"
	"to-do_list/repository"
	"to-do_list/services"
)

func main() {
	// Подключение к базе
	database, err := db.New("config/config.json")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Conn.Close()

	// Инициализация базы данных
	if err := db.InitializeDB(database.Conn); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Создание  репы и сервиса
	taskRepo := repository.NewTaskRepository(database.Conn)
	taskService := services.NewTaskService(taskRepo)

	if err := taskService.AddTask("Завершить учебный проект"); err != nil {
		log.Printf("Failed to add task: %v", err)
	}

	// Удаление задачи по ID
	if err := taskService.DeleteTask(2); err != nil {
		log.Printf("Failed to delete task: %v", err)
	}

	// Отметка о выполнении задачи
	if err := taskService.MarkTaskAsDone(11); err != nil {
		log.Printf("Failed to mark task as done: %v", err)
	}

	// Вывод списка задач
	tasks, err := taskService.ListTasks()
	if err != nil {
		log.Printf("Failed to list tasks: %v", err)
	}

	for _, task := range tasks {
		fmt.Printf("%d. %s (Выполнено: %t)\n", task.ID, task.Name, task.Status)
	}
}
