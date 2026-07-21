package controllers

import (
	"errors"
	"net/http"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/middleware"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/services"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/utils"

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
	postID, ok := parseID(c, "id")
	if !ok {
		return
	}
	comments, err := ctl.svc.List(postID)
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
	postID, ok := parseID(c, "id")
	if !ok {
		return
	}
	var req commentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, 40000, "invalid request")
		return
	}
	comment, err := ctl.svc.Create(middleware.CurrentUserID(c), postID, req.Content)
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
