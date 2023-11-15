package user

import (
	"bito_group/internal/user/model"
	"reflect"
	"testing"
)

func TestQuerySinglePeople(t *testing.T) {

	check := model.UserQueryCheck{
		Username:  "cat111",
		Gender:    model.Gender_Male,
		Height:    175,
		NeedCount: 0,
	}

	userRepo := NewMysqlUserRepo(nil)
	list, err := userRepo.QuerySinglePeople(check)
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
