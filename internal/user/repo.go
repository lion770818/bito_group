package user

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/user/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Get(*model.UserID) (*model.User, error)
	Del(*model.UserID) error
	GetUserByLoginParams(*model.LoginParams) (*model.User, error)
	GetUserByRegisterParams(*model.RegisterParams) (bool, error)
	Save(*model.User) (*model.User, error)

	QuerySinglePeople(check model.UserQueryCheck) ([]model.S2C_UserInfo, error)
}

var (
	ErrUserUsernameOrPassword = errors.New("用户名或者密码错误")
	ErrUserNotFound           = errors.New("用户不存在")
	ErrUserPointNil           = errors.New("用户指標錯誤")
	ErrUserOther              = errors.New("其他錯誤")
)

var _ UserRepo = &MysqlUserRepo{}

type MysqlUserRepo struct {
	db      *gorm.DB
	UserMap map[string]model.User // key = username
}

func NewMysqlUserRepo(db *gorm.DB) *MysqlUserRepo {
	return &MysqlUserRepo{
		db:      db,
		UserMap: make(map[string]model.User),
	}
}

func (r *MysqlUserRepo) GetUserByLoginParams(params *model.LoginParams) (*model.User, error) {
	var userPO model.UserPO
	var db = r.db
	var err error

	if params.Username.Value() != "" {
		err = db.Where("username = ? AND password = ?", params.Username.Value(), params.Password.Value()).First(&userPO).Error
	}
	// TODO: 支持其他参数查找

	if err != nil {
		return nil, ErrUserUsernameOrPassword
	}

	return userPO.ToDomain()
}

// 搜尋是否有註冊
func (r *MysqlUserRepo) GetUserByRegisterParams(params *model.RegisterParams) (bool, error) {

	_, ok := r.UserMap[params.Username.Value()]

	if !ok {
		return ok, ErrUserNotFound
	}

	return ok, nil
}

func (r *MysqlUserRepo) Get(id *model.UserID) (*model.User, error) {

	if id == nil {
		return nil, ErrUserPointNil
	}

	// 使用內存的記憶體結構
	user, ok := r.UserMap[id.GetName()]
	if !ok {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (r *MysqlUserRepo) Del(id *model.UserID) error {

	if id == nil {
		return ErrUserPointNil
	}

	// 使用內存的記憶體結構
	user, ok := r.UserMap[id.GetName()]
	if !ok {
		return ErrUserNotFound
	}

	// 資料錯誤
	if user.Username.Value() != id.GetName() {
		return ErrUserNotFound
	}

	// 刪除內存的用戶資料
	delete(r.UserMap, id.GetName())

	return nil
}

func (r *MysqlUserRepo) Save(user *model.User) (*model.User, error) {

	if user == nil || user.Username == nil {
		return nil, ErrUserPointNil
	}

	logs.Debugf("Save Username=%v", user.Username.Value())
	r.UserMap[user.Username.Value()] = *user

	check := model.UserCheck{
		UserId:   int64(len(r.UserMap)),
		Username: user.Username.Value(),
	}

	user.ID, _ = model.NewUserID(check)

	return user, nil
}

func (r *MysqlUserRepo) QuerySinglePeople(check model.UserQueryCheck) ([]model.S2C_UserInfo, error) {

	if len(check.Username) == 0 {
		return nil, ErrUserNotFound
	}

	//logs.Debugf("QuerySinglePeople Username=%v", check.Username)

	// TODO

	return nil, nil
}
