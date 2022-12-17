package resolver

import (
	"app/config"
	"app/database"
	"app/services"
	"app/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthUsecase usecases.AuthUsecase
	TodoUsecase usecases.TodoUsecase
	UserUsecase usecases.UserUsecase
}

func NewResolver(db *config.DB) *Resolver {
	cookieService := services.NewSessionService()
	cyptoService := services.NewCyptoService()
	jwtService := services.NewJwtService()

	userRepository := database.NewUserRepository(db)
	todoRepository := database.NewTodoRepository(db)

	authUsecase := usecases.NewAuthUsecase(userRepository, cookieService, cyptoService, jwtService)
	todoUsecase := usecases.NewTodoUsecase(todoRepository, cookieService, jwtService)
	userUsecase := usecases.NewUserUsecase(userRepository)

	return &Resolver{
		AuthUsecase: authUsecase,
		TodoUsecase: todoUsecase,
		UserUsecase: userUsecase,
	}
}
