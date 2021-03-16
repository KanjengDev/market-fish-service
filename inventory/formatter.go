package inventory

type InventoryFormatter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	Stock       int    `json:"stock"`
}

func FormatInventoryDetail(inventory Inventory) InventoryFormatter {
	inventoryDetailFormatter := InventoryFormatter{}
	inventoryDetailFormatter.ID = inventory.ID
	inventoryDetailFormatter.Name = inventory.Name
	inventoryDetailFormatter.Description = inventory.Description
	inventoryDetailFormatter.Price = inventory.Price
	inventoryDetailFormatter.Stock = inventory.Stock
	inventoryDetailFormatter.UserID = inventory.UserID
	inventoryDetailFormatter.ImageURL = inventory.FileName

	return inventoryDetailFormatter
}
