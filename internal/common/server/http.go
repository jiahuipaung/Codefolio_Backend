package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jiahuipaung/Codefolio_Backend/internal/common"
	"github.com/jiahuipaung/Codefolio_Backend/internal/common/client/user"
	"github.com/jiahuipaung/Codefolio_Backend/internal/user/app"
)

type HTTPServer struct {
	common.BaseResponse
	app app.Application
}

func RunHttpServer(serverName string, wrapper func(router *gin.Engine)) {
	// TODO: Implement server initialization
}

func PostAuthSignin(c *gin.Context) {
	var (
		req  user.EmailPasswordSignin
		err  error
		resp user.Response
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		// TODO: Handle error
	}
}

func PostAuthSignup(c *gin.Context) {
	var (
		req user.UserSignup
		err error
	)
	// TODO: Implement signup
}
