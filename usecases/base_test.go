package usecases_test

import (
	"app/test_utils/mock_database"
	"app/test_utils/mock_services"
	"app/usecases"
)

const (
	JwtToken = "jwt-token"
)

var mockTodoRepository *mock_database.MockTodoRepository
var mockUserRepository *mock_database.MockUserRepository
var mockCookieService *mock_services.MockCookieService
var mockJwtService *mock_services.MockJwtService
var mockCyptoService *mock_services.MockCyptoService
var testAuthUsecase usecases.AuthUsecase
