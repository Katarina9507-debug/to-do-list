package tasks

type TaskStorage []Task

type Task struct {
	ID     uint //создаем ID как положительное число
	Name   string
	Status bool
}
