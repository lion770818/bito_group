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

	QuerySinglePeople(check model.UserQueryCheck) ([]model.S2C_MatchPeople, error)
}

var (
	ErrUserUsernameOrPassword = errors.New("用户名或者密码错误")
	ErrUserNotFound           = errors.New("用户不存在")
	ErrUserPointNil           = errors.New("用户指標錯誤")
	ErrUserCount              = errors.New("用户人數錯誤")
	ErrUserGender             = errors.New("用户性別")
	ErrUserHeight             = errors.New("用户身高錯誤")
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

	logs.Debugf("Save Username:%s, Gender:%d, Height:%d", user.Username, user.Gender, user.Height)
	r.UserMap[user.Username.Value()] = *user

	check := model.UserCheck{
		UserId:   int64(len(r.UserMap)),
		Username: user.Username.Value(),
	}

	user.ID, _ = model.NewUserID(check)

	return user, nil
}

// 約會配對核心計算
func (r *MysqlUserRepo) QuerySinglePeople(check model.UserQueryCheck) ([]model.S2C_MatchPeople, error) {

	//logs.Debugf("MysqlUserRepo::QuerySinglePeople")
	var matchlist []model.S2C_MatchPeople

	if len(check.Username) == 0 {
		return nil, ErrUserNotFound
	}
	if check.NeedCount == 0 {
		return nil, ErrUserCount
	}
	if check.Gender < model.Gender_Unknow || check.Gender > model.Gender_Max {
		return nil, ErrUserGender
	}
	if check.Height == 0 {
		return nil, ErrUserHeight
	}

	//logs.Debugf("QuerySinglePeople check%+v", check)
	/*
	   匹配規則如下：
	   - 單一人有四個輸入參數：姓名、身高、性別、人數 想要約會。

	   - 男孩只能配對身高較低的女孩。 相反，女孩與男孩相匹配
	   更高。
	   - 一旦女孩和男孩配對，他們都會用完一次約會。 當他們的約會次數, 變為零，則應將它們從匹配系統中刪除。
	*/
	needCountNow := 0
	isFind := false
	for _, value := range r.UserMap {

		if isFind {
			break
		}
		// relu 檢查人數
		if needCountNow >= check.NeedCount {
			break // 人數已滿
		}
		// relu 檢查是否是自己
		if value.Username.Value() == check.Username {
			continue // 無法自己一組
		}
		// relu 檢查性別
		if value.Gender.Value() == check.Gender {
			continue // 無法同性別一組
		}

		// 根據性別做出篩選規則
		switch check.Gender {
		case model.Gender_Male: // 男性
			if check.Height > value.Height.Value() {

				// 找到了
				isFind = true
			}
		case model.Gender_Woman: // 女性
			if check.Height <= value.Height.Value() {

				// 找到了
				isFind = true
			}
		default:
			// 目前不支援第三性
		}

		if isFind {
			// 組合配對資訊
			teamMemberA := model.C2S_Register_Base{
				Username: check.Username,
				Gender:   check.Gender,
				Height:   check.Height,
			}
			teamMemberB := model.C2S_Register_Base{
				Username: value.Username.Value(),
				Gender:   value.Gender.Value(),
				Height:   value.Height.Value(),
			}
			matchTeam := model.S2C_MatchPeople{
				TeamIndex:   needCountNow,
				TeamMemberA: teamMemberA,
				TeamMemberB: teamMemberB,
			}

			//logs.Debugf("matchTeam:%+v", matchTeam)

			matchlist = append(matchlist, matchTeam)

			// 計算約會次數

			// 設定離開資訊
			isFind = true
			needCountNow++
			break
		}

	}

	return matchlist, nil
}
