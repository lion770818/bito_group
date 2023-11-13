package model

import "encoding/json"

type AuthKey struct {
	UserID int64 `json:"user_id"`
}

type AuthInfo struct {
	UserID int64 `json:"user_id"`
}

func (s *AuthInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *AuthInfo) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, s)
}

func NewAuth(userID int64) (*AuthInfo, error) {
	// 省略参数检查
	return &AuthInfo{
		UserID: userID,
	}, nil
}
