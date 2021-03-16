package handler

import (
	"fmt"
	"market-fish-service/helper"
	"market-fish-service/inventory"
	"market-fish-service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type inventoryHandler struct {
	service inventory.Service
}

func NewInventoryHandler(service inventory.Service) *inventoryHandler {
	return &inventoryHandler{service}
}

func (h *inventoryHandler) CreateCampaign(c *gin.Context) {
	var input inventory.ItemInput
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create item", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload item image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	newItem, err := h.service.CreateItem(input, path)
	if err != nil {
		response := helper.APIResponse("Failed to create  item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create item", http.StatusOK, "success", inventory.FormatInventoryDetail(newItem))
	c.JSON(http.StatusOK, response)

}
