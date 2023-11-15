package user

import (
	"bito_group/internal/user/model"
	"reflect"
	"testing"
)

// 測試 /v1/QuerySinglePeople
func TestQuerySinglePeople(t *testing.T) {

	userRepo := NewMysqlUserRepo(nil)
	// 新增男性
	name, _ := model.NewUsername("cat111")
	gender, _ := model.NewGender(1)
	height, _ := model.NewHeight(175)
	user1 := &model.User{
		Username: name,
		Gender:   gender,
		Height:   height,
	}
	newUser1, err := userRepo.Save(user1)
	if err != nil {
		t.Errorf("save err=%v", err)
		return
	}
	t.Log("save newUser1=", newUser1)

	name, _ = model.NewUsername("cat222")
	gender, _ = model.NewGender(1)
	height, _ = model.NewHeight(169)
	user2 := &model.User{
		Username: name,
		Gender:   gender,
		Height:   height,
	}
	newUser2, err := userRepo.Save(user2)
	if err != nil {
		t.Errorf("save err=%v", err)
		return
	}
	t.Log("save newUser2=", newUser2)

	// 新增女性
	name, _ = model.NewUsername("girl001")
	gender, _ = model.NewGender(2)
	height, _ = model.NewHeight(170)
	girl1 := &model.User{
		Username: name,
		Gender:   gender,
		Height:   height,
	}
	newUser3, err := userRepo.Save(girl1)
	if err != nil {
		t.Errorf("save err=%v", err)
		return
	}
	t.Log("save newUser3=", newUser3)

	// 測試配對
	queryCheck := model.UserQueryCheck{
		Username:  "cat111",
		Gender:    model.Gender_Male,
		Height:    175,
		NeedCount: 1,
	}

	list, err := userRepo.QuerySinglePeople(queryCheck)
	if err != nil {
		t.Errorf("querySinglePeople err=%v", err)
		return
	}

	t.Log("querySinglePeople succes list=", list)

}

func TestMysqlUserRepo_Get(t *testing.T) {
	type args struct {
		id *model.UserID
	}
	tests := []struct {
		name    string
		r       *MysqlUserRepo
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MysqlUserRepo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
