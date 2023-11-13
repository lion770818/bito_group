package user

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/user/model"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("用户已存在")
)

// 各種usecase
type UserAppInterface interface {
	Login(login *model.LoginParams) (*model.S2C_Login, error)
	Register(register *model.RegisterParams) (*model.S2C_Login, error)

	Get(userID *model.UserID) (*model.S2C_UserInfo, error)
	Del(userID *model.UserID) error
	QuerySinglePeople(model.UserQueryCheck) ([]model.S2C_UserInfo, error)
}

type UserApp struct {
	userRepo UserRepo
}

func NewUserApp(userRepo UserRepo) UserAppInterface {
	// 有跟用戶有關的其他 usercase 都可以丟進來
	return &UserApp{
		userRepo: userRepo,
	}
}

// Login
func (u *UserApp) Login(login *model.LoginParams) (*model.S2C_Login, error) {
	// 登录
	user, err := u.userRepo.GetUserByLoginParams(login)
	if err != nil {
		return nil, err
	}

	// // 生成 token
	// authInfo := &model.AuthInfo{
	// 	UserID: user.ID.Value(),
	// }
	// token, err := u.authRepo.Set(authInfo)
	// if err != nil {
	// 	return nil, err
	// }
	token := ""

	return user.ToLoginResp(token), nil
}

// Register 注册 + 自动登录
func (u *UserApp) Register(register *model.RegisterParams) (*model.S2C_Login, error) {
	// 检查是否已经注册
	isFind, err := u.userRepo.GetUserByRegisterParams(register)
	if isFind || err == nil {
		return nil, ErrUserAlreadyExists
	}

	// 注册
	user, err := u.userRepo.Save(register.ToDomain())
	if err != nil {
		return nil, err
	}

	// 生成 token
	token := ""

	return user.ToLoginResp(token), nil
}

// Get 获取用户信息
func (u *UserApp) Get(userID *model.UserID) (*model.S2C_UserInfo, error) {
	user, err := u.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	return user.ToUserInfo(), nil
}

func (u *UserApp) Del(userID *model.UserID) error {
	user, err := u.userRepo.Get(userID)
	if err != nil {
		return err
	}

	logs.Debugf("delete username=%v", user.Username)
	err = u.userRepo.Del(userID)

	return err
}

func (u *UserApp) QuerySinglePeople(check model.UserQueryCheck) ([]model.S2C_UserInfo, error) {

	return nil, nil
}
