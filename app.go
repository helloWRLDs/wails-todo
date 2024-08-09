package main

import (
	"context"
	"fmt"
	"syscall"
	"time"
	"todo/internal/domain"
	"todo/internal/repository"

	"github.com/sirupsen/logrus"
)

// App struct
type App struct {
	ctx  context.Context
	repo *repository.Repository
}

// NewApp creates a new App application struct
func NewApp(repo *repository.Repository) *App {
	return &App{
		repo: repo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) InsertTodo(todo domain.Todo) int64 {
	id, err := a.repo.Insert(a.ctx, &todo)
	if err != nil {
		return -1
	}
	return id
}

func (a *App) DeleteTodo(id int64) {
	if err := a.repo.Delete(a.ctx, id); err != nil {
		logrus.WithError(err).Error("internal err")
	}
}

func (a *App) UpdateTodo(todo domain.Todo) {
	if err := a.repo.Update(a.ctx, todo); err != nil {
		logrus.WithError(err).Error("internal err")
	}
}

func (a *App) ListTodos() []domain.Todo {
	todos, err := a.repo.List(a.ctx)
	if err != nil {
		logrus.WithError(err).Error("internal err")
		return []domain.Todo{}
	}
	return todos
}

func (a *App) GetTodos() []domain.Todo {
	todos := []domain.Todo{
		{ID: 0, Body: "Get Groceries", IsDone: false, CreatedAt: time.Now(), Priority: 1},
		{ID: 1, Body: "Get Breakfast", IsDone: true, CreatedAt: time.Now()},
		{ID: 2, Body: "Fix the car", IsDone: true, CreatedAt: time.Now()},
		{ID: 3, Body: "Go to hospital", IsDone: false, CreatedAt: time.Now()},
		{ID: 4, Body: "Play games", IsDone: false, CreatedAt: time.Now()},
	}
	return todos
}

func (a *App) GetHostName() string {
	var computerName [syscall.MAX_COMPUTERNAME_LENGTH + 1]uint16
	size := uint32(len(computerName))
	err := syscall.GetComputerName(&computerName[0], &size)
	if err != nil {
		return "User"
	}
	return syscall.UTF16ToString(computerName[:])
}
