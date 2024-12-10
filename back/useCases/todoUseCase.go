package useCases

import (
	"context"

	"github.com/tomo-micco/TodoWithGo/databases/entities"
	"github.com/tomo-micco/TodoWithGo/databases/repositories"
)

type TodoUseCases struct {
	repository repositories.TodoRepositoryInterface
}

func NewGetTodoUseCase(repository repositories.TodoRepositoryInterface) *TodoUseCases {
	return &TodoUseCases{repository}
}

func (useCase *TodoUseCases) GetAll(c context.Context) []entities.Todo {
	return useCase.repository.GetAll(c)
}

func (useCase *TodoUseCases) FindById(c context.Context, id int64) entities.Todo {
	return useCase.repository.FindById(c, id)
}
