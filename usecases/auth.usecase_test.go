package usecases_test

import (
	"app/entity"
	"app/graph/model"
	"app/test_utils/mock_database"
	"app/test_utils/mock_services"
	"app/usecases"
	"errors"

	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func authSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository = mock_database.NewMockUserRepository(ctrl)
	mockCyptoService = mock_services.NewMockCyptoService(ctrl)
	mockJwtService = mock_services.NewMockJwtService(ctrl)
	mockCookieService = mock_services.NewMockCookieService(ctrl)

	return func() {}
}

func TestSignIn(t *testing.T) {
	authSetup := authSetUp(t)
	defer authSetup()

	var (
		email        = "test@example.com"
		expectedUser = entity.User{ID: 1, Email: email, Name: "test", HashPassword: "$2a$10$dipu0jX8IYvy.r.XsecEt.gF4XYsYmBIheZotlxLekYC9UKQWHKe2"}
		password     = "password"
	)

	t.Run("success", func(t *testing.T) {
		mockUserRepository.EXPECT().FindByEmail(email).Return(&expectedUser, nil)
		mockCyptoService.EXPECT().ComparePasswords(expectedUser.HashPassword, []byte(password)).Return(true)
		mockJwtService.EXPECT().GenerateToken(uint(expectedUser.ID), gomock.Any()).Return(JwtToken, nil)
		mockCookieService.EXPECT().SetCookie(gomock.Any(), usecases.ACCESS_TOKEN_KEY, JwtToken).Return(nil)

		testAuthUsecase = usecases.NewAuthUsecase(
			mockUserRepository,
			mockCookieService,
			mockCyptoService,
			mockJwtService,
		)

		loginUser, err := testAuthUsecase.SignIn(context.TODO(), &model.SignInInput{Email: email, Password: password})

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, loginUser.ID)
		assert.Equal(t, expectedUser.Email, loginUser.Email)
	})

	t.Run("If the authentication fails", func(t *testing.T) {
		mockUserRepository.EXPECT().FindByEmail(email).Return(&expectedUser, nil)
		mockCyptoService.EXPECT().ComparePasswords(expectedUser.HashPassword, []byte(password)).Return(false)

		testAuthUsecase = usecases.NewAuthUsecase(
			mockUserRepository,
			mockCookieService,
			mockCyptoService,
			mockJwtService,
		)

		_, err := testAuthUsecase.SignIn(context.TODO(), &model.SignInInput{Email: email, Password: password})

		assert.Error(t, err)
		assert.EqualError(t, err, "Authentication Failure")
	})
}

func TestSignUp(t *testing.T) {
	authSetup := authSetUp(t)
	defer authSetup()

	var (
		email        = "test@example.com"
		password     = "password"
		name         = "test"
		hashPassword = "$2a$10$dipu0jX8IYvy.r.XsecEt.gF4XYsYmBIheZotlxLekYC9UKQWHKe2"
		expectedUser = entity.User{ID: 1, Email: email, Name: name, HashPassword: hashPassword}
	)

	t.Run("success", func(t *testing.T) {
		mockCyptoService.EXPECT().HashAndSalt([]byte(password)).Return(hashPassword, nil)
		mockUserRepository.EXPECT().Create(&entity.User{Email: email, HashPassword: hashPassword, Name: name}).Return(&expectedUser, nil)
		mockJwtService.EXPECT().GenerateToken(uint(expectedUser.ID), gomock.Any()).Return(JwtToken, nil)
		mockCookieService.EXPECT().SetCookie(gomock.Any(), usecases.ACCESS_TOKEN_KEY, JwtToken).Return(nil)

		testAuthUsecase = usecases.NewAuthUsecase(
			mockUserRepository,
			mockCookieService,
			mockCyptoService,
			mockJwtService,
		)

		loginUser, err := testAuthUsecase.SignUp(context.TODO(), &model.SignUpInput{Email: email, Password: password, Name: name})

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, loginUser.ID)
		assert.Equal(t, expectedUser.Email, loginUser.Email)
	})

	t.Run("If hashing fails", func(t *testing.T) {
		err := errors.New("")

		mockCyptoService.EXPECT().HashAndSalt([]byte(password)).Return(hashPassword, err)
		mockUserRepository.EXPECT().Create(&entity.User{Email: email, HashPassword: hashPassword, Name: name}).Return(&expectedUser, nil)

		testAuthUsecase = usecases.NewAuthUsecase(
			mockUserRepository,
			mockCookieService,
			mockCyptoService,
			mockJwtService,
		)

		_, err = testAuthUsecase.SignUp(context.TODO(), &model.SignUpInput{Email: email, Password: password, Name: name})

		assert.Error(t, err)
	})
}
