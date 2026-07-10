package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/middleware"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/services"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct{ svc *services.PostService }

func NewPostController(svc *services.PostService) *PostController {
	return &PostController{svc: svc}
}

type postRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (ctl *PostController) List(c *gin.Context) {
	posts, err := ctl.svc.List()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, 50001, "database error")
		return
	}
	utils.OK(c, posts)
}

func (ctl *PostController) Get(c *gin.Context) {
	post, err := ctl.svc.Get(parseID(c, "id"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, 40404, "post not found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, 50001, "database error")
		return
	}
	utils.OK(c, post)
}

func (ctl *PostController) Create(c *gin.Context) {
	var req postRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, 40000, "invalid request")
		return
	}
	post, err := ctl.svc.Create(middleware.CurrentUserID(c), req.Title, req.Content)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, 40001, err.Error())
		return
	}
	utils.Created(c, post)
}

func (ctl *PostController) Update(c *gin.Context) {
	var req postRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, 40000, "invalid request")
		return
	}
	post, err := ctl.svc.Update(parseID(c, "id"), middleware.CurrentUserID(c), req.Title, req.Content)
	if errors.Is(err, services.ErrForbidden) {
		utils.Error(c, http.StatusForbidden, 40301, "only author can update post")
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, 40404, "post not found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, 50001, "database error")
		return
	}
	utils.OK(c, post)
}

func (ctl *PostController) Delete(c *gin.Context) {
	err := ctl.svc.Delete(parseID(c, "id"), middleware.CurrentUserID(c))
	if errors.Is(err, services.ErrForbidden) {
		utils.Error(c, http.StatusForbidden, 40301, "only author can delete post")
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, 40404, "post not found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, 50001, "database error")
		return
	}
	utils.OK(c, gin.H{"deleted": true})
}

func parseID(c *gin.Context, name string) uint {
	id, _ := strconv.ParseUint(c.Param(name), 10, 64)
	return uint(id)
}
