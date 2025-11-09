package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/internal/dto"
	"todolist/internal/service"
)

type ToDoListHandler struct {
	toDoService *service.ToDoService
}

func NewToDoListHandler(toDoService *service.ToDoService) *ToDoListHandler {
	return &ToDoListHandler{
		toDoService: toDoService,
	}
}

//	@Summary		Create a new todo item
//	@Description	Add a new todo item to the list
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		dto.ToDoAddReq	true	"Todo item to add"
//	@Success		200		{object}	vo.SuccessResponse
//	@Failure		400		{object}	vo.ErrorResponse
//	@Failure		500		{object}	vo.ErrorResponse
//	@Router			/todos/add [post]
//
// Create 创建待办事项
func (h *ToDoListHandler) Create(c *gin.Context) {
	var req dto.ToDoAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未登录"})
		return
	}

	if err := h.toDoService.Create(req, userId.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

//	@Summary		Update a todo item by ID
//	@Description	Update a specific todo item by its ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Todo ID"
//	@Param			todo	body		dto.ToDoUpdateReq	true	"Todo item to update"
//	@Success		200		{object}	vo.SuccessResponse
//	@Failure		400		{object}	vo.ErrorResponse
//	@Failure		401		{object}	vo.ErrorResponse
//	@Failure		403		{object}	vo.ErrorResponse
//	@Failure		500		{object}	vo.ErrorResponse
//	@Router			/todos/update/{id} [put]
//
// UpdateByID 更新待办事项
func (h *ToDoListHandler) UpdateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.ToDoUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从认证信息中获取用户ID
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// 检查待办事项是否属于当前用户
	todo, err := h.toDoService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if todo.UserId != userId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	if err := h.toDoService.UpdateByID(req, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

//	@Summary		Delete a todo item by ID
//	@Description	Delete a specific todo item by its ID
//	@Tags			todos
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	vo.SuccessResponse
//	@Failure		400	{object}	vo.ErrorResponse
//	@Failure		401	{object}	vo.ErrorResponse
//	@Failure		403	{object}	vo.ErrorResponse
//	@Failure		500	{object}	vo.ErrorResponse
//	@Router			/todos/delete/{id} [delete]
//
// Delete 删除待办事项
func (h *ToDoListHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// 从认证信息中获取用户ID
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// 检查待办事项是否属于当前用户
	todo, err := h.toDoService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if todo.UserId != userId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	if err := h.toDoService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

//	@Summary		Get a todo item by ID
//	@Description	Get a specific todo item by its ID
//	@Tags			todos
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	vo.ToDoVO
//	@Failure		400	{object}	vo.ErrorResponse
//	@Failure		401	{object}	vo.ErrorResponse
//	@Failure		403	{object}	vo.ErrorResponse
//	@Failure		500	{object}	vo.ErrorResponse
//	@Router			/todos/get/{id} [get]
//
// GetByID 获取单个待办事项
func (h *ToDoListHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// 从认证信息中获取用户ID
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	todo, err := h.toDoService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 检查待办事项是否属于当前用户
	if todo.UserId != userId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

//	@Summary		List all todo items for a user
//	@Description	Get a list of all todo items for the authenticated user
//	@Tags			todos
//	@Produce		json
//	@Success		200	{array}		vo.ToDoVO
//	@Failure		400	{object}	vo.ErrorResponse
//	@Failure		500	{object}	vo.ErrorResponse
//	@Router			/todos/list [get]
//
// ListByUser 获取用户的待办事项列表
func (h *ToDoListHandler) ListByUser(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
	}

	todos, err := h.toDoService.ListByUser(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}
