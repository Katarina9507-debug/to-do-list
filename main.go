package main

import (
	"log"

	"to-do_list/tasks"
)

func main() {
	storage := &tasks.TaskStorage{}
	// добавляем задачу, в случае ошибки - логируем
	if err := storage.AddTask(tasks.Task{1, "Завершить учебный проект", false}); err != nil {
		log.Fatal(err)
	}
	log.Println("task added")

	// отмечаем задачу как выполненную, в случае её отсутсвия логируем ошибку
	if err := storage.MarkTaskAsDone(true, 1); err != nil {
		log.Fatal(err)
	}
	log.Println("task mark as done")

	// удаляем задачу, в случае её отсутсвия логируем ошибку
	if err := storage.DeleteTask(tasks.Task{1, "", false}); err != nil {
		log.Fatal(err)
	}
	log.Println("task is deleted")

	// выводим список существющих задач
	storage.ListTasks()
}
