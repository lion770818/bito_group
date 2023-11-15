package user

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/servers/web/response"
	"bito_group/internal/user/model"
	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/swaggo/swag/example/celler/model"
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
// @Description Register
// @Summary 註冊帳號
// @Schemes
// @Tags example
// @Accept json
// @Produce json
//
//	@Param			message	body		model.C2S_Register	true	"要註冊的帳號"
//
// @Success 	200 	{string} 	register 	   account
// @Failure		500	    {object}	string "fail"
// @Router /v1/AddSinglePersonAndMatch [post]
func (u *UserHandler) AddSinglePersonAndMatch(c *gin.Context) {
	// @Param name query string true "用户姓名"
	var err error
	req := &model.C2S_Register{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		//response.ErrFromSwagger(c, http.StatusBadRequest, err.Error())
		//httputil.NewError(c, http.StatusBadRequest, errors.New("test"))
		return
	}

	// 转化为领域对象 + 参数验证
	registerParams, err := req.ToDomain()
	if err != nil {
		logs.Errorf("[AddSinglePersonAndMatch] failed, err: %+v", err)
		response.Err(c, http.StatusBadRequest, err.Error())
		//response.ErrFromSwagger(c, http.StatusBadRequest, err.Error())
		//httputil.NewError(c, http.StatusBadRequest, errors.New(fmt.Sprintf(err.Error())))
		return
	}

	// 调用应用层
	user, err := u.UserApp.Register(registerParams)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		//response.ErrFromSwagger(c, http.StatusInternalServerError, err.Error())
		//httputil.NewError(c, http.StatusInternalServerError, errors.New(fmt.Sprintf(err.Error())))
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
