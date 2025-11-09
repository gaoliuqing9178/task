package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/internal/dto"
	"todolist/internal/service"
)

// UserHandler 负责处理用户相关的 HTTP 请求
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler 创建一个新的 UserHandler 实例
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

//	@Summary		User login
//	@Description	Authenticate a user and return a token
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		dto.LoginReq	true	"User credentials"
//	@Success		200			{object}	vo.LoginResponse
//	@Failure		400			{object}	vo.ErrorResponse
//	@Failure		401			{object}	vo.ErrorResponse
//	@Router			/user/login [post]
//
// Login 处理用户登录请求
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	userVO, token, err := h.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user":    userVO,
		"token":   token,
	})
}

//	@Summary		Register a new user
//	@Description	Create a new user account
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.RegisterReq	true	"User registration info"
//	@Success		200		{object}	vo.RegisterResponse
//	@Failure		400		{object}	vo.ErrorResponse
//	@Failure		500		{object}	vo.ErrorResponse
//	@Router			/user/register [post]
//
// Register 处理用户注册请求
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	userVO, err := h.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"user":    userVO,
	})
}

//	@Summary		Update user information
//	@Description	Update the authenticated user's information
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.UpdateUserInfoReq	true	"User info to update"
//	@Success		200		{object}	vo.UpdateUserResponse
//	@Failure		400		{object}	vo.ErrorResponse
//	@Failure		500		{object}	vo.ErrorResponse
//	@Router			/user/update [put]
//
// UpdateUserInfo 处理更新用户信息的请求
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	var req dto.UpdateUserInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未登录"})
		return
	}

	userVO, err := h.userService.UpdateUserInfo(&req, id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"user":    userVO,
	})
}
