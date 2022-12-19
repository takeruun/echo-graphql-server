package resolver_test

import (
	"app/entity"
	"app/graph/generated"
	"app/graph/model"
	"app/graph/resolver"
	"app/test_utils/mock_usecases"
	"encoding/json"
	"errors"
	"strconv"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func userSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase = mock_usecases.NewMockAuthUsecase(ctrl)

	resolvers = resolver.Resolver{
		AuthUsecase: mockAuthUsecase,
	}

	return func() {
		defer ctrl.Finish()
	}
}

func TestSignUp(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	var (
		ID       uint64 = 1
		name            = "test"
		email           = "test@example.com"
		password        = "password"
	)
	signInParams := model.SignUpInput{Name: name, Email: email, Password: password}
	user := entity.User{ID: ID, Name: name, Email: email}

	t.Run("success", func(t *testing.T) {
		mockAuthUsecase.EXPECT().SignUp(gomock.Any(), &signInParams).Return(&user, nil)

		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

		var resp struct {
			SignUp struct {
				ID, Name string
			}
		}

		query := `
			mutation {
				signUp(
					signUpInput: {name: "test", email: "test@example.com", password: "password"}
				) {
					id
					name
				}
			}
		`
		c.MustPost(query, &resp)
		assert.Equal(t, strconv.Itoa(int(user.ID)), resp.SignUp.ID)
		assert.Equal(t, user.Name, resp.SignUp.Name)
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("failed signup")
		mockAuthUsecase.EXPECT().SignUp(gomock.Any(), &signInParams).Return(nil, err)

		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

		type eRes struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		}
		var resp []eRes

		query := `
			mutation {
				signUp(
					signUpInput: {name: "test", email: "test@example.com", password: "password"}
				) {
					id
					name
				}
			}
		`

		res := c.Post(query, &resp)
		json.Unmarshal([]byte(res.Error()), &resp)
		assert.Equal(t, err.Error(), resp[0].Message)
	})
}

func TestSignIn(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	var (
		ID       uint64 = 1
		name            = "test"
		email           = "test@example.com"
		password        = "password"
	)
	signInParams := model.SignInInput{Email: email, Password: password}
	user := entity.User{ID: ID, Name: name, Email: email}

	t.Run("success", func(t *testing.T) {
		mockAuthUsecase.EXPECT().SignIn(gomock.Any(), &signInParams).Return(&user, nil)

		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

		var resp struct {
			SignIn struct {
				ID, Name string
			}
		}

		query := `
			mutation {
				signIn(
					signInInput: {email: "test@example.com", password: "password"}
				) {
					id
					name
				}
			}
		`
		c.MustPost(query, &resp)
		assert.Equal(t, strconv.Itoa(int(user.ID)), resp.SignIn.ID)
		assert.Equal(t, user.Name, resp.SignIn.Name)
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("failed signup")
		mockAuthUsecase.EXPECT().SignIn(gomock.Any(), &signInParams).Return(nil, err)

		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

		type eRes struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		}
		var resp []eRes

		query := `
			mutation {
				signIn(
					signInInput: {email: "test@example.com", password: "password"}
				) {
					id
					name
				}
			}
		`

		res := c.Post(query, &resp)
		json.Unmarshal([]byte(res.Error()), &resp)
		assert.Equal(t, err.Error(), resp[0].Message)
	})
}
