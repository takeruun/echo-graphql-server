package usecases_test

import (
	"app/entity"
	"app/graph/model"
	"app/services"
	"app/test_utils/mock_database"
	"app/test_utils/mock_services"
	"app/usecases"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func todoSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepository = mock_database.NewMockTodoRepository(ctrl)
	mockCookieService = mock_services.NewMockCookieService(ctrl)
	mockJwtService = mock_services.NewMockJwtService(ctrl)

	return func() {}
}

func TestCreate(t *testing.T) {
	todoSetUp := todoSetUp(t)
	defer todoSetUp()

	var (
		auth         = services.Auth{Uid: uint(1)}
		title        = "test-test"
		description  = "description-test"
		expectedTodo = entity.Todo{Title: title, Description: description, UserId: uint64(auth.Uid), UpdatedAt: time.Now(), CreatedAt: time.Now()}
	)

	t.Run("success", func(t *testing.T) {
		mockCookieService.EXPECT().GetCookieValue(gomock.Any(), usecases.ACCESS_TOKEN_KEY).Return(JwtToken, nil)
		mockJwtService.EXPECT().ParseToken(JwtToken).Return(&auth, nil)
		mockTodoRepository.EXPECT().Create(&entity.Todo{Title: title, Description: description, UserId: 1}).Return(&expectedTodo, nil)

		todoUsecase := usecases.NewTodoUsecase(mockTodoRepository, mockCookieService, mockJwtService)

		result, err := todoUsecase.Create(context.TODO(), &model.CreateTodo{Title: title, Description: description})

		assert.NoError(t, err)
		assert.Equal(t, expectedTodo.Title, result.Title)
		assert.Equal(t, expectedTodo.Description, result.Description)
	})

	t.Run("Error in todo creation", func(t *testing.T) {
		var expectErr error = errors.New("error")
		mockCookieService.EXPECT().GetCookieValue(gomock.Any(), usecases.ACCESS_TOKEN_KEY).Return(JwtToken, nil)
		mockJwtService.EXPECT().ParseToken(JwtToken).Return(&auth, nil)
		mockTodoRepository.EXPECT().Create(&entity.Todo{Title: title, Description: description, UserId: 1}).Return(&expectedTodo, expectErr)
		todoUsecase := usecases.NewTodoUsecase(mockTodoRepository, mockCookieService, mockJwtService)

		_, err := todoUsecase.Create(context.TODO(), &model.CreateTodo{Title: title, Description: description})

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectErr)
	})
}
