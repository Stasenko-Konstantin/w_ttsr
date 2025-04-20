package repository

// Db - интерфейс для инъекции зависимостей.
type Db interface {
	Query(query string, args ...any) error
	Scan(args ...any) error
	Next() bool
}
