package item

type Service interface {
	GetItemByID(id uint, lang string) (*ItemDetailResponse, error)
	GetVendorItemsGrouped(vendorID uint, lang string) ([]CategoryWithItemsResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetItemByID(id uint, lang string) (*ItemDetailResponse, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return ToItemDetailResponse(item, lang), nil
}

func (s *service) GetVendorItemsGrouped(vendorID uint, lang string) ([]CategoryWithItemsResponse, error) {
	items, err := s.repo.GetByVendorID(vendorID)
	if err != nil {
		return nil, err
	}
	return GroupItemsByCategory(items, lang), nil
}
