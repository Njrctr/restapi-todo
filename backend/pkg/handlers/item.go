package handler

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/Njrctr/restapi-todo/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create todo item
// @Security ApiKeyAuth
// @Tags items
// @Description create todo item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param list_id path int true "List Id"
// @Param input body models.TodoItemCreateUpdate true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list_id param")
		return
	}

	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item_id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": item_id,
	})
}

type getAllItemsResponce struct {
	Data []models.TodoItem `json:"data"`
}

// @Summary Get All Items
// @Security ApiKeyAuth
// @Tags items
// @Description get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Param list_id path int true "List Id"
// @Success 200 {object} getAllItemsResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("list_id"))
	fmt.Print("listId: ", listId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list_id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponce{
		Data: items,
	})
}

// @Summary Get Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param item_id path int true "Item Id"
// @Success 200 {object} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{item_id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item_id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

// @Summary Update todo item
// @Security ApiKeyAuth
// @Tags items
// @Description update todo item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param item_id path int true "Item Id"
// @Param input body models.TodoItemCreateUpdate true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{item_id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid item_id param")
		return
	}

	var input models.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, itemId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{"ok"})
}

// @Summary Delete Item
// @Security ApiKeyAuth
// @Tags items
// @Description delete item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param item_id path int true "Item Id"
// @Success 200 {string} string ok
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{item_id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item_id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}
