package controllers

import (
	"net/http"

	"metanode-go-backend-homeworks/internal/blog/middleware"
	"metanode-go-backend-homeworks/internal/blog/services"
	"metanode-go-backend-homeworks/internal/blog/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{ svc *services.AuthService }

func NewAuthController(svc *services.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

type authRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ctl *AuthController) Register(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, 40000, "invalid request")
		return
	}
	user, err := ctl.svc.Register(req.Username, req.Email, req.Password)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, 40001, err.Error())
		return
	}
	utils.Created(c, gin.H{"id": user.ID, "username": user.Username, "email": user.Email})
}

func (ctl *AuthController) Login(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, 40000, "invalid request")
		return
	}
	token, user, err := ctl.svc.Login(req.Username, req.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, 40001, err.Error())
		return
	}
	utils.OK(c, gin.H{"token": token, "user": gin.H{"id": user.ID, "username": user.Username, "email": user.Email}})
}

func (ctl *AuthController) Profile(c *gin.Context) {
	user, err := ctl.svc.Profile(middleware.CurrentUserID(c))
	if err != nil {
		utils.Error(c, http.StatusNotFound, 40401, "user not found")
		return
	}
	utils.OK(c, gin.H{"id": user.ID, "username": user.Username, "email": user.Email})
}
