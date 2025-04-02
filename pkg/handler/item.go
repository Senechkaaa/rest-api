package handler

import (
	"github.com/Senechkaaa/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItems(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(listId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) getItemsById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItems(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})

}

func (h *Handler) deleteItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		"ok",
	})
}
