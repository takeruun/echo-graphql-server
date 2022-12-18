package usecases

import (
	"app/database"
	"app/entity"
	"app/graph/model"
	"app/services"
	"context"
	"errors"
	"strconv"
)

type TodoUsecase interface {
	FindAll(ctx context.Context) (todos []*entity.Todo, err error)
	Create(ctx context.Context, createParams *model.CreateTodo) (todo *entity.Todo, err error)
	Show(ctx context.Context, todoId int) (todo *entity.Todo, err error)
	Edit(ctx context.Context, updateParams *model.UpdateTodo) (todo *entity.Todo, err error)
	Delete(ctx context.Context, todoId int) error
}

type todoUsecase struct {
	todoRepo      database.TodoRepository
	cookieService services.CookieService
	jwtService    services.JwtService
}

func NewTodoUsecase(todoRepo database.TodoRepository, cookieService services.CookieService, jwtService services.JwtService) TodoUsecase {
	return &todoUsecase{
		todoRepo:      todoRepo,
		cookieService: cookieService,
		jwtService:    jwtService,
	}
}

func (tu *todoUsecase) FindAll(ctx context.Context) (todos []*entity.Todo, err error) {
	accessToken, err := tu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return nil, err
	}

	auth, err := tu.jwtService.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	todos, err = tu.todoRepo.FindAll(uint64(auth.Uid))

	return
}

func (tu *todoUsecase) Create(ctx context.Context, createParams *model.CreateTodo) (todo *entity.Todo, err error) {
	accessToken, err := tu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return nil, err
	}

	auth, err := tu.jwtService.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	t := entity.Todo{Title: createParams.Title, Description: createParams.Description, UserId: uint64(auth.Uid)}
	todo, err = tu.todoRepo.Create(&t)
	if err != nil {
		return nil, err
	}

	return
}

func (tu *todoUsecase) Show(ctx context.Context, todoId int) (todo *entity.Todo, err error) {
	accessToken, err := tu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return nil, err
	}

	todo, err = tu.todoRepo.Find(todoId)
	if err != nil {
		return nil, err
	}

	auth, err := tu.jwtService.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	if uint64(auth.Uid) != todo.UserId {
		return nil, errors.New("no your todo")
	}

	return
}

func (tu *todoUsecase) Edit(ctx context.Context, updateParams *model.UpdateTodo) (todo *entity.Todo, err error) {
	accessToken, err := tu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return nil, err
	}

	auth, err := tu.jwtService.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	todoId, _ := strconv.Atoi(updateParams.ID)

	todo, err = tu.todoRepo.Find(todoId)
	if err != nil {
		return nil, err
	}
	if uint64(auth.Uid) != todo.UserId {
		return nil, errors.New("no your todo")
	}

	t := entity.Todo{ID: uint64(todoId), Description: updateParams.Description}
	todo, err = tu.todoRepo.Update(&t)
	if err != nil {
		return nil, err
	}

	return
}

func (tu *todoUsecase) Delete(ctx context.Context, todoId int) error {
	accessToken, err := tu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return err
	}

	auth, err := tu.jwtService.ParseToken(accessToken)
	if err != nil {
		return err
	}

	todo, err := tu.todoRepo.Find(todoId)
	if err != nil {
		return err
	}
	if uint64(auth.Uid) != todo.UserId {
		return errors.New("no your todo")
	}

	if err := tu.todoRepo.Delete(&entity.Todo{ID: uint64(todoId)}); err != nil {
		return err
	}

	return nil
}
