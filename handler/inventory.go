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

	path := fmt.Sprintf("http://178.128.108.162:8080/images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	newItem, err := h.service.CreateItem(input, path)
	if err != nil {
		response := helper.APIResponse("Failed to create  item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create item", http.StatusOK, "success", inventory.FormatInventoryDetail(newItem))
	c.JSON(http.StatusOK, response)

}

func (h *inventoryHandler) GetInventory(c *gin.Context) {
	// userID, _ := strconv.Atoi(c.Query("user_id"))

	items, err := h.service.GetItems()

	if err != nil {
		response := helper.APIResponse("Error to get items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of items", http.StatusOK, "success", inventory.FormatInventory(items))
	c.JSON(http.StatusOK, response)
}

func (h *inventoryHandler) GetInventoryByID(c *gin.Context) {
	var input inventory.GetItemDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventoryDetails, err := h.service.GetItemByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Items", http.StatusOK, "success", inventory.FormatInventoryById(inventoryDetails))
	c.JSON(http.StatusOK, response)
}
