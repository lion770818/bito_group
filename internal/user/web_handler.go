package user

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/servers/web/response"
	"bito_group/internal/user/model"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "bito_group/docs"
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

// PingExample godoc
// @Summary 移除帳號
// @Description Remove a user from the matching system so that the user cannot be matched anymore
// @Schemes
// @Tags user
// @Accept json
// @Produce json
// @Param			message	body	model.UserCheck		true		"要檢查的帳號"
// @Success 	200 	{object} 	model.S2C_Login
// @Failure     500		{object}	response.HTTPError
// @Failure     400		{object}	response.HTTPError
// @Router /v1/RemoveSinglePerson [delete]
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
// @Summary 註冊帳號
// @Description Add a new user to the matching system and find any possible matches for the new user
// @Schemes
// @Tags user
// @Accept json
// @Produce json
// @Param			message	body	model.C2S_Register		true		"要註冊的帳號"
// @Success 	200 	{object} 	model.S2C_Login
// @Failure     500		{object}	response.HTTPError
// @Failure     400		{object}	response.HTTPError
// @Router /v1/AddSinglePersonAndMatch [post]
func (u *UserHandler) AddSinglePersonAndMatch(c *gin.Context) {
	// @Param name query string true "用户姓名"
	var err error
	req := &model.C2S_Register{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		//httputil.NewError(c, http.StatusBadRequest, err)
		response.ErrFromSwagger(c, http.StatusBadRequest, err.Error())
		return
	}

	logs.Debugf("AddSinglePersonAndMatch reg:%+v", req)

	// 转化为领域对象 + 参数验证
	registerParams, err := req.ToDomain()
	if err != nil {
		response.ErrFromSwagger(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err := u.UserApp.Register(registerParams)
	if err != nil {
		response.ErrFromSwagger(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}

// PingExample godoc
// @Summary 尋找最多 N 個可能匹配的單身人士
// @Description QuerySinglePeople : Find the most N possible matched single people
// @Schemes
// @Tags user
// @Accept json
// @Produce json
// @Param			message	body	model.UserQueryCheck		true		"要匹配的單身人士"
// @Success 	200 	{object} 	model.S2C_MatchPeople
// @Failure     500		{object}	response.HTTPError
// @Failure     400		{object}	response.HTTPError
// @Router /v1/QuerySinglePeople [post]
func (u *UserHandler) QuerySinglePeople(c *gin.Context) {

	req := &model.UserQueryCheck{}
	if err := c.ShouldBindJSON(req); err != nil {
		logs.Errorf("ShouldBindJSON failed, err: %+v", err)
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 尋找匹配的約會對象
	matchPeople, err := u.UserApp.QuerySinglePeople(*req)
	if err != nil {
		logs.Errorf("QuerySinglePeople failed, err: %+v", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回用户信息
	response.Ok(c, matchPeople)
}
