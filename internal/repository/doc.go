// Package repository предоставляет логику работы с БД
//
// - определяет структуру репозитория Task
//
// - конструктор репозитория Task
//
// - методы чтения/создания/обновления/удаления [domain.Task]
//
// - интерфейс Db для инъекции зависимостей
//
// - структуру PgxDb - экземпляр интерфейса Db; соответствующие методы
package repository
