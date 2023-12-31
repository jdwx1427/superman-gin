package app

import (
	"net/http"
	"superman-gin/app/common/request"
	"superman-gin/app/common/response"
	"superman-gin/app/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// @Summary 用户注册账号
// @Description 创建项目
// @Tags 项目管理
// @Accept  application/json
// @Product application/json
// @Param data body request.Register true "手机号，名称,密码"
// @Success 200 {object} response.Response{data=object} "{"code": 200, "data": [...]}"
// @Router /register [post]
// @Security Bearer
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// @Summary 用户登录账号
// @Description 创建项目
// @Tags 项目管理
// @Accept  application/json
// @Product application/json
// @Param data body request.Login true "手机号，密码"
// @Success 200 {object} response.Response{data=object} "{"code": 200, "data": [...]}"
// @Router /login [post]
// @Security Bearer
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(c, form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}

// @BasePath /api/auth

// 测试 godoc
// @Summary 测试一下
// @Schemes
// @Description do ping
// @Tags 测试
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
