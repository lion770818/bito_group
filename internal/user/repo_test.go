package user

import (
	"bito_group/internal/user/model"
	"testing"
)

func TestQuerySinglePeople(t *testing.T) {

	check := model.UserQueryCheck{
		Username: "cat111",
		Gender:   "男",
		Height:   175,
	}

	userRepo := NewMysqlUserRepo(nil)
	list, err := userRepo.QuerySinglePeople(check)
	if err != nil {
		t.Errorf("querySinglePeople err=%v", err)
		return
	}

	t.Log("querySinglePeople succes list=", list)

}
