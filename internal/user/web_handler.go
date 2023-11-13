package user

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/servers/web/response"
	"bito_group/internal/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserApp UserAppInterface
}

func NewUserHandler(userApp UserAppInterface) *UserHandler {
	return &UserHandler{
		UserApp: userApp,
	}
}

// AddSinglePersonAndMatch
func (u *UserHandler) Login(c *gin.Context) {
	var err error
	req := &model.C2S_Login{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	loginParams, err := req.ToDomain()
	if err != nil {
		logs.Debugf("錯誤的參數 err=%v", err.Error())
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err := u.UserApp.Login(loginParams)
	if err != nil {
		logs.Errorf("[Login] failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}

// UserInfo 获取用户信息
func (u *UserHandler) UserInfo(c *gin.Context) {
	var check model.UserCheck
	err := c.BindJSON(&check)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := model.NewUserID(check)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := u.UserApp.Get(userID)
	if err != nil {
		logs.Errorf("[UserInfo] failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回用户信息
	response.Ok(c, userInfo)
}

// UserInfo 刪除單一用戶
func (u *UserHandler) RemoveSinglePerson(c *gin.Context) {

	req := &model.UserCheck{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := model.NewUserID(*req)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := u.UserApp.Get(userID)
	if err != nil {
		logs.Errorf("[UserInfo] failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = u.UserApp.Del(userID)
	if err != nil {
		logs.Errorf("[UserInfo] failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回用户信息
	response.Ok(c, userInfo)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [post]
func (u *UserHandler) Register(c *gin.Context) {
	var err error
	req := &model.C2S_Register{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	registerParams, err := req.ToDomain()
	if err != nil {
		logs.Errorf("[Register] failed, err: %+v", err)
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err := u.UserApp.Register(registerParams)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}

func (u *UserHandler) QuerySinglePeople(c *gin.Context) {

	req := &model.UserCheck{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := model.NewUserID(*req)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := u.UserApp.Get(userID)
	if err != nil {
		logs.Errorf("[UserInfo] failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = u.UserApp.Del(userID)
	if err != nil {
		logs.Errorf("[UserInfo] failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回用户信息
	response.Ok(c, userInfo)
}
