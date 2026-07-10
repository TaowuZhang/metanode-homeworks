package controllers

import (
	"errors"
	"net/http"

	"metanode-go-backend-homeworks/internal/blog/middleware"
	"metanode-go-backend-homeworks/internal/blog/services"
	"metanode-go-backend-homeworks/internal/blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct{ svc *services.CommentService }

func NewCommentController(svc *services.CommentService) *CommentController {
	return &CommentController{svc: svc}
}

type commentRequest struct {
	Content string `json:"content"`
}

func (ctl *CommentController) List(c *gin.Context) {
	comments, err := ctl.svc.List(parseID(c, "id"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, 40404, "post not found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, 50001, "database error")
		return
	}
	utils.OK(c, comments)
}

func (ctl *CommentController) Create(c *gin.Context) {
	var req commentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, 40000, "invalid request")
		return
	}
	comment, err := ctl.svc.Create(middleware.CurrentUserID(c), parseID(c, "id"), req.Content)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, 40404, "post not found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusBadRequest, 40001, err.Error())
		return
	}
	utils.Created(c, comment)
}
