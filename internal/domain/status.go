package domain

// Status - тип перечисление для статуса [domain.Task].
type Status string

const (
	New        Status = "new"
	InProgress        = "in_progress"
	Done              = "done"
)

// IsStatus - функция валидации статуса.
func IsStatus(status Status) bool {
	switch status {
	case New, InProgress, Done:
		return true
	}
	return false
}
