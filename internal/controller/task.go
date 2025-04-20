package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Stasenko-Konstantin/w_ttsr/internal/domain"
	"github.com/Stasenko-Konstantin/w_ttsr/internal/repository"
	"regexp"
	"strconv"
)

// Task - структура контроллера для [domain.Task].
type Task struct {
	repo *repository.Task
}

// New - конструктор контроллера Task.
// На вход принимает3 репозиторий [repository.Task].
func New(repo *repository.Task) *Task {
	return &Task{repo}
}

// GetTasks - возвращает список всех [domain.Task]
func (t *Task) GetTasks() ([]*domain.Task, error) {
	return t.repo.GetTasks()
}

// SaveTasks - сохраняет [domain.Task].
// На вход принимает сырые байты - список [domain.Task].
func (t *Task) SaveTasks(bytes []byte) error {
	var tasks []*domain.Task
	if err := json.Unmarshal(bytes, &tasks); err != nil {
		return err
	}
	for _, task := range tasks {
		if !domain.IsStatus(task.Status) {
			return fmt.Errorf("invalid task status: %s", task.Status)
		}
	}
	return t.repo.SaveTasks(tasks)
}

// UpdateTask - обновляет [domain.Task].
// На вход принимает сырые байты - один [domain.Task].
func (t *Task) UpdateTask(bytes []byte) error {
	task := &domain.Task{}
	if err := json.Unmarshal(bytes, task); err != nil {
		return err
	}
	if !domain.IsStatus(task.Status) {
		return fmt.Errorf("invalid task status: %s", task.Status)
	}
	return t.repo.UpdateTask(task)
}

// DeleteTask - удаляет [domain.Task].
// На вход принимает и обрабатывает строковый параметр `:id=?`.
func (t *Task) DeleteTask(id string) error {
	r, err := regexp.Compile(`:id=(\d+)`)
	if err != nil {
		return err
	}
	strs := r.FindStringSubmatch(id)
	fmt.Println(strs)
	tId, err := strconv.Atoi(strs[1])
	if err != nil {
		return err
	}
	return t.repo.DeleteTask(tId)
}
