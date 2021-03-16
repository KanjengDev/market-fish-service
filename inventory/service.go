package inventory

type Service interface {
	CreateItem(input ItemInput, fileLocation string) (Inventory, error)
	GetItems(userID uint) ([]Inventory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateItem(input ItemInput, fileLocation string) (Inventory, error) {
	item := Inventory{}
	item.Name = input.Name
	item.Price = input.Price
	item.Description = input.Description
	item.Stock = input.Stock
	item.FileName = fileLocation
	item.UserID = input.User.ID

	newItem, err := s.repository.Save(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) GetItems(userID uint) ([]Inventory, error) {
	// if userID != 0 {
	// 	items, err := s.repository.FindByUserID(userID)
	// 	if err != nil {
	// 		return items, err
	// 	}

	// 	return items, nil
	// }
	items, err := s.repository.FindAll()
	if err != nil {
		return items, err
	}

	return items, nil
}
