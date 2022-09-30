package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandler struct {
}

func NewAdminHandler() *AdminHandler {
	handler := new(AdminHandler)
	return handler
}
func (h AdminHandler) Echo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
