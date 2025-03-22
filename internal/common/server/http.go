package server

import (
	"common"
	"common/client/user"
	"github.com/gin-gonic/gin"
	client "github.com/jiahuipaung/Codefolio_Backend/common/client/user"
	"github.com/jiahuipaung/Codefolio_Backend/user/app"
)

type HTTPServer struct {
	common.BaseResponse
	app app.Application
}

func RunHttpServer(serverName string, wrapper func(router *gin.Engine)) {

}

func PostAuthSignin(c *gin.Context) {
	var (
		req  client.EmailPasswordSignin
		err  error
		resp user.Response
	)
	if err = c.ShouldBindJSON(&req); err != nil {

	}
}

func PostAuthSignup(c *gin.Context) {
	var (
		req client.UserSignup
		err error
	)
}
