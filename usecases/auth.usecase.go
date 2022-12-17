package usecases

import (
	"app/database"
	"app/entity"
	"app/graph/model"
	"app/services"
	"time"

	"context"
	"errors"
)

const (
	ACCESS_TOKEN_KEY = "_echo_graphql_access_token"
)

type AuthUsecase interface {
	SignIn(ctx context.Context, signInParams *model.SignInInput) (user *model.User, err error)
	SignUp(ctx context.Context, signInParams *model.SignUpInput) (user *model.User, err error)
	Show(ctx context.Context) (user *model.User, err error)
	Delete(ctx context.Context) error
}

type authUsecase struct {
	userRepo      database.UserRepository
	cookieService services.CookieService
	cyptoService  services.CyptoService
	jwtService    services.JwtService
}

func NewAuthUsecase(userRepo database.UserRepository, cookieService services.CookieService, cyptoService services.CyptoService, jwtService services.JwtService) AuthUsecase {
	return &authUsecase{
		userRepo:      userRepo,
		cookieService: cookieService,
		cyptoService:  cyptoService,
		jwtService:    jwtService,
	}
}

func (uu *authUsecase) SignIn(ctx context.Context, signInParams *model.SignInInput) (user *model.User, err error) {
	loginUser, err := uu.userRepo.FindByEmail(signInParams.Email)
	if err != nil {
		return nil, err
	}

	if !uu.cyptoService.ComparePasswords(loginUser.HashPassword, []byte(signInParams.Password)) {
		return nil, errors.New("Authentication Failure")
	}

	accessToken, err := uu.jwtService.GenerateToken(uint(loginUser.ID), time.Now())
	if err != nil {
		return nil, err
	}

	err = uu.cookieService.SetCookie(ctx, ACCESS_TOKEN_KEY, accessToken)
	if err != nil {
		return nil, err
	}

	user = entity.ToModelUser(loginUser)

	return
}

func (uu *authUsecase) SignUp(ctx context.Context, signInParams *model.SignUpInput) (user *model.User, err error) {
	hashPwd, err := uu.cyptoService.HashAndSalt([]byte(signInParams.Password))
	if err != nil {
		return nil, err
	}

	u := entity.User{Name: signInParams.Name, Email: signInParams.Email, HashPassword: hashPwd}
	loginUser, err := uu.userRepo.Create(&u)
	if err != nil {
		return nil, err
	}

	accessToken, err := uu.jwtService.GenerateToken(uint(loginUser.ID), time.Now())
	if err != nil {
		return nil, err
	}

	err = uu.cookieService.SetCookie(ctx, ACCESS_TOKEN_KEY, accessToken)
	if err != nil {
		return nil, err
	}

	user = entity.ToModelUser(loginUser)

	return
}

func (uu *authUsecase) Show(ctx context.Context) (user *model.User, err error) {
	accessToken, err := uu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return nil, err
	}

	auth, err := uu.jwtService.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	loginUser, err := uu.userRepo.Find(uint64(auth.Uid))
	if err != nil {
		return nil, err
	}

	user = entity.ToModelUser(loginUser)

	return
}

func (uu *authUsecase) Delete(ctx context.Context) error {
	accessToken, err := uu.cookieService.GetCookieValue(ctx, ACCESS_TOKEN_KEY)
	if err != nil {
		return err
	}

	auth, err := uu.jwtService.ParseToken(accessToken)
	if err != nil {
		return err
	}

	if err := uu.userRepo.Delete(uint64(auth.Uid)); err != nil {
		return nil
	}

	if err := uu.cookieService.DeleteCookie(ctx, ACCESS_TOKEN_KEY); err != nil {
		return err
	}

	return nil
}
