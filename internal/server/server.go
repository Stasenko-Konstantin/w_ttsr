package server

import (
	"context"
	"encoding/json"
	"github.com/Stasenko-Konstantin/w_ttsr/internal/controller"
	"github.com/Stasenko-Konstantin/w_ttsr/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"strings"
)

type server struct {
	port   string
	pool   *pgxpool.Pool
	server *fiber.App
}

// New - инициализация сервера.
// На вход принимает порт на котором будет запущен сервер и строку подключения к postgres.
func New(port, connstr string) (*server, error) {
	app := fiber.New()
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	pool, err := pgxpool.New(context.Background(), connstr)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to postgres")
	}

	db := repository.NewPgxDb(pool)
	repo := repository.New(db)
	ctrl := controller.New(repo)
	_ = ctrl

	p := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	app.Get("/metrics", func(ctx *fiber.Ctx) error {
		p(ctx.Context())
		return nil
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		tasks, err := ctrl.GetTasks()
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(tasks)
		if err != nil {
			return err
		}
		return c.Send(bytes)
	})

	app.Post("/tasks", func(c *fiber.Ctx) error {
		bytes := c.Body()
		if err := ctrl.SaveTasks(bytes); err != nil {
			return err
		}
		return nil
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		bytes := c.Body()
		if err := ctrl.UpdateTask(bytes); err != nil {
			return err
		}
		return nil
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := ctrl.DeleteTask(id); err != nil {
			return err
		}
		return nil
	})

	return &server{port, pool, app}, nil
}

// Start - запускает сервер.
func (s *server) Start() error {
	defer s.pool.Close()
	return s.server.Listen(s.port)
}
