package user

import (
	"bito_group/internal/user/model"
	"log"
	"testing"
)

func TestQuerySinglePeople(t *testing.T) {

	check := model.UserQueryCheck{}

	userRepo := NewMysqlUserRepo(nil)
	list, err := userRepo.QuerySinglePeople(check)
	if err != nil {
		log.Printf("querySinglePeople err=%v", err)
		return
	}

	log.Printf("querySinglePeople list=%v", list)

}
