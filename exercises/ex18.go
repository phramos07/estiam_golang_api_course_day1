package main

import "fmt"

type Task struct {
	Title       string // public field
	Description string
	done        bool // private field
}

func NewTask(title string, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
	}
}

func (t *Task) Complete() {
	t.done = true
}

func (t *Task) String() string {
	status := "Not done"
	if t.done {
		status = "Done"
	}

	return fmt.Sprintf("Title: %s\nDescription: %s\nStatus: %s",
		t.Title,
		t.Description,
		status,
	)
}

/*
NewTaskManager() -> create a new taskmanager
AddTask(??) -> add a task to the list.
RemoveTask(??) -> remove a task from the list.
ListTasks()-> display all tasks in the list and print information of each task.
*/

type TaskManager struct {
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) AddTask(t *Task) {
	tm.tasks = append(tm.tasks, t)
}

func (tm *TaskManager) ListTasks() {
	for _, task := range tm.tasks {
		fmt.Printf("%s\n\n", task)
	}
}

func (tm *TaskManager) RemoveTask(t *Task) {
	for i, task := range tm.tasks {
		if task == t {
			tm.tasks = append(tm.tasks[0:i], tm.tasks[i+1:]...)
		}
	}
}

/*
Exercise: Creating a Todo List Manager

Create a simple todo list manager using Go. This program should allow users to add, remove, and list tasks in a todo list.

Create a Task struct:
Define a Task struct with fields like Title, Description, and Done (boolean)

Methods:
NewTask() -> creates a new task
Complete() -> mark task as completed
String() -> print task formatted

Create a TaskManager struct

AddTask(??) -> add a task to the list.
RemoveTask(??) -> remove a task from the list.
ListTasks()-> display all tasks in the list.

Test using main.
*/

func main() {

	manager := NewTaskManager()

	task1 := NewTask("Cleaning", "Perform cleaning of the sheets")
	task2 := NewTask("Golang Course", "Solve the exercises of the course")

	manager.AddTask(task1)
	manager.AddTask(task2)

	task1.Complete()

	manager.ListTasks()

	manager.RemoveTask(task1)

	manager.ListTasks()
}
