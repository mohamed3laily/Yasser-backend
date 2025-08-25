package search

import (
	"fmt"
)

type Service interface {
	Search(request SearchRequest) ([]SearchResponse, error)
	IndexData() error
}

type service struct {
	client *Client
	repo   Repository
}

func NewService(client *Client, repo Repository) Service {
	return &service{
		client: client,
		repo:   repo,
	}
}

func (s *service) Search(request SearchRequest) ([]SearchResponse, error) {
	// Set default limit if not provided
	if request.Limit == 0 {
		request.Limit = 20
	}

	// Set default language if not provided
	if request.Lang == "" {
		request.Lang = "en"
	}

	return s.client.Search(request)
}

func (s *service) IndexData() error {
	// Index items
	items, err := s.repo.GetAllItemsForIndexing()
	if err != nil {
		return fmt.Errorf("failed to get items for indexing: %w", err)
	}

	if len(items) > 0 {
		if err := s.client.AddDocuments(items ); err != nil {
			return fmt.Errorf("failed to index items: %w", err)
		}
	}

	// Index vendors
	vendors, err := s.repo.GetAllVendorsForIndexing()
	if err != nil {
		return fmt.Errorf("failed to get vendors for indexing: %w", err)
	}
	fmt.Println("Vendors length: ", len(vendors))
	if len(vendors) > 0 {
		if err := s.client.AddDocuments(vendors); err != nil {
			return fmt.Errorf("failed to index vendors: %w", err)
		}
	}

	return nil
}