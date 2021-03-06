package inventory

type Service interface {
	CreateItem(input ItemInput, fileLocation string) (Inventory, error)
	GetItems() ([]Inventory, error)
	GetItemByID(input GetItemDetailInput) (Inventory, error)
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

func (s *service) GetItems() ([]Inventory, error) {
	items, err := s.repository.FindAll()
	if err != nil {
		return items, err
	}

	return items, nil
}

func (s *service) GetItemByID(input GetItemDetailInput) (Inventory, error) {
	items, err := s.repository.FindByID(input.ID)
	if err != nil {
		return items, err
	}

	return items, nil
}
