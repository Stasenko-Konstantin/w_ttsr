package repository

import "github.com/Stasenko-Konstantin/w_ttsr/internal/domain"

// Task - структура репозитория для [domain.Task]
type Task struct {
	db Db
}

// New - конструктор репозитория Task.
// На вход принимает экземпляр интерфейса [repository.Db]
func New(db Db) *Task {
	return &Task{db}
}

// GetTasks - возвращает список всех [domain.Task]
func (t *Task) GetTasks() ([]*domain.Task, error) {
	var tasks []*domain.Task
	if err := t.db.Query(`SELECT * FROM tasks;`); err != nil {
		return nil, err
	}
	for t.db.Next() {
		var task domain.Task
		if err := t.db.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			panic(err)
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

// SaveTasks - сохраняет [domain.Task]
// На вход принимает список [*domain.Task].
func (t *Task) SaveTasks(tasks []*domain.Task) error {
	for _, task := range tasks {
		if err := t.db.Query(`INSERT INTO tasks (id, title, description, status, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6);`,
			task.ID,
			task.Title,
			task.Description,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		); err != nil {
			return err
		}
	}
	return nil
}

// UpdateTask - обновляет [domain.Task]
// На вход принимает [*domain.Task]
func (t *Task) UpdateTask(task *domain.Task) error {
	if err := t.db.Query(`UPDATE tasks SET title = $2, description = $3, status = $4, updated_at = $5 WHERE id = $1;`,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}

// DeleteTask - удаляет [domain.Task]
// На вход принимает числовой идентификатор [domain.Task]
func (t *Task) DeleteTask(taskId int) error {
	if err := t.db.Query(`DELETE FROM tasks WHERE id = $1;`, taskId); err != nil {
		return err
	}
	return nil
}
