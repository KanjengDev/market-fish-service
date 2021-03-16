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

func FormatInventory(inventory []Inventory) []InventoryFormatter {
	if len(inventory) == 0 {
		return []InventoryFormatter{}
	}

	var campaignsFormatter []InventoryFormatter

	for _, campaign := range inventory {
		campaignFormatter := FormatInventoryDetail(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type InventoryFormatterDetails struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Price       int64     `json:"price"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	Stock       int       `json:"stock"`
	User        OwnerShop `json:"owner_shop"`
}

type OwnerShop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func FormatInventoryById(inventory Inventory) InventoryFormatterDetails {
	itemInventory := InventoryFormatterDetails{}
	itemInventory.ID = inventory.ID
	itemInventory.Name = inventory.Name
	itemInventory.Description = inventory.Description
	itemInventory.Price = inventory.Price
	itemInventory.Stock = inventory.Stock
	itemInventory.UserID = inventory.UserID
	itemInventory.ImageURL = inventory.FileName

	user := inventory.User
	ownerShop := OwnerShop{}

	ownerShop.Name = user.Name
	ownerShop.Address = user.Address

	itemInventory.User = ownerShop

	return itemInventory
}
