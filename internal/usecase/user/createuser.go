package user

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rikkunn23/kokoiko-app-backend-bff/config"
	"github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/user"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/user/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordRequired = errors.New("パスワードは必須です")
	ErrPasswordInvalid  = errors.New("パスワードは8〜128文字で入力してください")
	ErrEmailRequired    = errors.New("メールアドレスは必須です")
	ErrEmailTooLong     = errors.New("メールアドレスは256文字以下で入力してください")
	ErrTelRequired      = errors.New("電話番号は必須です")
	ErrTelTooLong       = errors.New("電話番号は15文字以下で入力してください")
	ErrSeiNameRequired  = errors.New("姓は必須です")
	ErrSeiNameTooLong   = errors.New("姓は50文字以下で入力してください")
	ErrMeiNameRequired  = errors.New("名は必須です")
	ErrMeiNameTooLong   = errors.New("名は50文字以下で入力してください")
)



func (u *Usecase) CreateUser(r *http.Request, reqBody user.RequestUserCreate) (user.ResponseUserCreate, error) {
	ctx := r.Context()

	if reqBody.Password == "" {
		return user.ResponseUserCreate{}, ErrPasswordRequired
	}
	if len(reqBody.Password) < 8 || len(reqBody.Password) > 128 {
		return user.ResponseUserCreate{}, ErrPasswordInvalid
	}
	if reqBody.Email == "" {
		return user.ResponseUserCreate{}, ErrEmailRequired
	}
	if len(reqBody.Email) > 256 {
		return user.ResponseUserCreate{}, ErrEmailTooLong
	}
	if reqBody.Tel == "" {
		return user.ResponseUserCreate{}, ErrTelRequired
	}
	if len(reqBody.Tel) > 15 {
		return user.ResponseUserCreate{}, ErrTelTooLong
	}
	if reqBody.SeiName == "" {
		return user.ResponseUserCreate{}, ErrSeiNameRequired
	}
	if len(reqBody.SeiName) > 50 {
		return user.ResponseUserCreate{}, ErrSeiNameTooLong
	}
	if reqBody.MeiName == "" {
		return user.ResponseUserCreate{}, ErrMeiNameRequired
	}
	if len(reqBody.MeiName) > 50 {
		return user.ResponseUserCreate{}, ErrMeiNameTooLong
	}

	// パスワードをハッシュ化
	hashedPassword, err := hashPassword(reqBody.Password)
	if err != nil {
		return user.ResponseUserCreate{}, err
	}



	// TODO configから取れてない件修正
	userEntity := entity.CreateUser{
		Email: 	 reqBody.Email,
		Tel:     reqBody.Tel,
		Password: hashedPassword,
		SeiName: reqBody.SeiName,
		MeiName: reqBody.MeiName,
		RecApp:  config.Get().App.RecApp,
		RecAcnt: config.Get().App.RecAcnt,
		UpdApp:  config.Get().App.UpdApp,
		UpdAcnt: config.Get().App.UpdAcnt,
	}

	fmt.Println("userEntity:", userEntity)

	userNo, err := u.repo.CreateUser(ctx, userEntity)
	if err != nil {
		return  user.ResponseUserCreate{}, err
	}

	// createToken := "TestToken" // JWTトークンを生成するための変数
	createToken, err := GenerateJWT(userEntity.Email, userEntity.Tel, userNo)
	if err != nil {
			return user.ResponseUserCreate{}, err
		}

	createRefreshToken, expiresAt, err := GenerateRefreshToken()
	if err != nil {
		return user.ResponseUserCreate{}, err
	}

	// DBに保存（refreshToken, userID, 有効期限）
	err = u.repo.SaveRefreshToken(ctx, userNo, createRefreshToken, expiresAt)
	if err != nil {
		return user.ResponseUserCreate{}, err
	}

	tokenRes := user.ResponseUserCreate{
		AccessToken: createToken,
		ExpiresIn: int(time.Until(expiresAt).Seconds()),
		RefreshToken: createRefreshToken,
	}

	return tokenRes, nil
}

// ハッシュ化の関数
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
