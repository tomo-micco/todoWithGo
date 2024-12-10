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

func (useCase *TodoUseCases) GetAll(c context.Context) ([]entities.Todo, error) {
	return useCase.repository.GetAll(c)
}

func (useCase *TodoUseCases) FindById(c context.Context, id int64) (entities.Todo, error) {
	return useCase.repository.FindById(c, id)
}

func (useCase *TodoUseCases) Create(c context.Context, todo entities.Todo) (int64, error) {
	return useCase.repository.Create(c, todo)
}

func (useCase *TodoUseCases) Update(c context.Context, todo entities.Todo) (int64, error) {
	return useCase.repository.Update(c, todo)
}

func (useCase *TodoUseCases) Delete(c context.Context, id int64) (int64, error) {
	return useCase.repository.Delete(c, id)
}
