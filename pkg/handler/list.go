package handler

import (
	"fmt"
	"github.com/Senechkaaa/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	fmt.Println(userId)

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	fmt.Println(userId)

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	// извлекает значение из нашего URL
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	// извлекает значение из нашего URL
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}
