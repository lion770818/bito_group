package model

import (
	"errors"

	"github.com/shopspring/decimal"
)

// domain 领域对象

var (
	DefaultUserIDValue     = "0"
	DefaultUsernameValue   = ""
	DefaultUserGenderValue = 0
	DefaultUserHeightValue = 0
	DefaultPasswordValue   = ""
	DefaultCurrencyValue   = "CNY"
	DefaultAmountValue     = decimal.NewFromFloat(0)
	DefaultFeeValue, _     = NewAmount(decimal.NewFromFloat(0))
)

var (
	ErrAmountNotEnough = errors.New("余额不足")

	ErrUserId   = errors.New("UserID錯誤")
	ErrUsername = errors.New("Username錯誤")
)

type UserCheck struct {
	UserId   int64
	Username string
}

func ParameterCheck(check UserCheck) error {

	if check.UserId < 0 {
		return ErrUserId
	}

	if len(check.Username) == 0 {
		return ErrUsername
	}

	return nil
}

type UserID struct {
	value    int64
	username string
}

func NewUserID(check UserCheck) (*UserID, error) {

	// 参数检查
	err := ParameterCheck(check)
	if err != nil {
		return nil, err
	}

	return &UserID{
		value:    check.UserId,
		username: check.Username,
	}, nil
}

func NewUserID2(userId int64) (*UserID, error) {

	// 参数检查
	if userId < 0 {
		return nil, ErrUserId

	}
	return &UserID{
		value: userId,
	}, nil
}

func (u *UserID) Value() int64 {
	if u == nil {
		return 0
	}

	return u.value
}

func (u *UserID) GetName() string {
	if u == nil {
		return ""
	}

	return u.username
}

type Username struct {
	value string
}

func NewUsername(username string) (*Username, error) {
	// 省略参数检查
	return &Username{
		value: username,
	}, nil
}

func (u *Username) Value() string {
	if u == nil {
		return DefaultUsernameValue
	}

	return u.value
}

type Gender struct {
	value int
}

func NewGender(gender int) (*Gender, error) {
	// 省略参数检查
	return &Gender{
		value: gender,
	}, nil
}

func (u *Gender) Value() int {
	if u == nil {
		return DefaultUserGenderValue
	}

	return u.value
}

type Height struct {
	value int
}

func NewHeight(height int) (*Height, error) {
	// 省略参数检查
	return &Height{
		value: height,
	}, nil
}

func (u *Height) Value() int {
	if u == nil {
		return DefaultUserGenderValue
	}

	return u.value
}

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	// 省略参数检查
	return &Password{
		value: password,
	}, nil
}

func (u *Password) Value() string {
	if u == nil {
		return DefaultPasswordValue
	}

	return u.value
}

type Currency struct {
	value string
}

func NewCurrency(currency string) (*Currency, error) {
	// 省略参数检查
	return &Currency{
		value: currency,
	}, nil
}

func (u *Currency) Value() string {
	if u == nil {
		return DefaultCurrencyValue
	}

	return u.value
}

type Amount struct {
	value decimal.Decimal
}

func NewAmount(amount decimal.Decimal) (*Amount, error) {
	// 省略参数检查
	return &Amount{
		value: amount,
	}, nil
}

func (m *Amount) Value() decimal.Decimal {
	if m == nil {
		return DefaultAmountValue
	}

	return m.value
}

func (m *Amount) Add(amount *Amount) *Amount {
	return &Amount{
		value: m.value.Add(amount.value),
	}
}

type User struct {
	ID       *UserID
	Username *Username
	Gender   *Gender // 性別
	Height   *Height // 身高

	Password *Password
	Currency *Currency
	Amount   *Amount
}

func (u *User) CalcFee(fromAmount *Amount) (*Amount, error) {
	return NewAmount(fromAmount.Value().Mul(DefaultFeeValue.Value()))
}

// 付款
func (u *User) Pay(amount *Amount) error {
	// 省略参数检查
	if u.Amount.Value().LessThan(amount.Value()) {
		return ErrAmountNotEnough
	}

	u.Amount.value = u.Amount.Value().Sub(amount.Value())

	return nil
}

// 收款
func (u *User) Receive(amount *Amount) error {
	// 省略参数检查

	u.Amount.value = u.Amount.value.Add(amount.value)

	return nil
}

func (u *User) ToLoginResp(token string) *S2C_Login {
	return &S2C_Login{
		UserID:   u.ID.Value(),
		Username: u.Username.Value(),
		Token:    token,
	}
}

func (u *User) ToUserInfo() *S2C_UserInfo {
	return &S2C_UserInfo{
		UserID:   u.ID.Value(),
		Username: u.Username.Value(),
		Gender:   u.Gender.Value(),
		Height:   u.Height.Value(),
		Amount:   u.Amount.Value().String(),
		Currency: u.Currency.Value(),
	}
}

func (u *User) ToPO() *UserPO {

	return &UserPO{
		ID:       u.ID.Value(),
		Username: u.Username.Value(),
		Gender:   u.Gender.Value(),
		Height:   u.Height.Value(),
		Password: u.Password.Value(),
		Currency: u.Currency.Value(),
		Amount:   u.Amount.Value(),
	}
}

type LoginParams struct {
	Username *Username
	Password *Password
}

type RegisterParams struct {
	Username *Username
	Password *Password
	Gender   *Gender // 性別
	Height   *Height // 身高
}

func (c *RegisterParams) ToDomain() *User {
	return &User{
		Username: c.Username,
		Password: c.Password,
		Gender:   c.Gender,
		Height:   c.Height,
	}
}

type Rate struct {
	rate decimal.Decimal
}

func NewRate(rate decimal.Decimal) (*Rate, error) {
	// 省略参数检查
	return &Rate{
		rate: rate,
	}, nil
}

func (r *Rate) Exchange(amount *Amount) (*Amount, error) {
	return NewAmount(amount.Value().Mul(r.rate))
}

// 用來批配 尋找最多 N 個可能匹配的單身人士 條件
type UserQueryCheck struct {
	Username  string // 用戶名
	Gender    int    // 性別
	Height    int    // 身高
	NeedCount int    // 需要人數
}
