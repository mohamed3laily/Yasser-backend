package search

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

type Client struct {
	client    meilisearch.ServiceManager
	index     meilisearch.IndexManager
	indexName string
}

func NewClient(host, apiKey, indexName string) *Client {
	client := meilisearch.New(host, meilisearch.WithAPIKey(apiKey))
	
	index := client.Index(indexName)

	c := &Client{
		client:    client,
		index:     index,
		indexName: indexName,
	}

	c.setupIndex()

	return c
}

func (c *Client) setupIndex() {
	searchableAttributes := []string{
		"nameEn",
		"nameAr", 
		"vendorName",
		"items",
	}

	filterableAttributes := []string{
		"type",
		"vendorId",
		"districtId",
		"categoryId",
		"isActive",
	}



	// Set index settings
	task, err := c.index.UpdateSettings(&meilisearch.Settings{
		SearchableAttributes: searchableAttributes,
		FilterableAttributes: filterableAttributes,
	})

	if err != nil {
		log.Printf("Failed to update Meilisearch settings: %v", err)
		return
	}

	_, err = c.client.WaitForTask(task.TaskUID , 30*time.Second)
	if err != nil {
		log.Printf("Failed to wait for settings update task: %v", err)
	}
}

func (c *Client) AddDocuments(documents []SearchDocument) error {

    primaryKey := "id"
    task, err := c.index.AddDocuments(documents, &primaryKey)
    if err != nil {
        return fmt.Errorf("failed to add documents: %w", err)
    }
    taskInfo, err := c.client.WaitForTask(task.TaskUID, 30*time.Second)
    if err != nil {
        return fmt.Errorf("failed to wait for indexing task: %w", err)
    }

    if taskInfo.Status == meilisearch.TaskStatusFailed {
        return fmt.Errorf("meilisearch task failed: %v", taskInfo.Error)
    }

    return nil
}

func (c *Client) Search(request SearchRequest) ([]SearchResponse, error) {
	searchReq := &meilisearch.SearchRequest{
		Limit:  int64(request.Limit),
		Offset: int64(request.Offset),
	}

	var filters []string
	filters = append(filters, "isActive = true")
	filters = append(filters, fmt.Sprintf("districtId = %d", request.DistrictID))
	if request.Type != nil && *request.Type != "" {
		filters = append(filters, fmt.Sprintf("type = '%s'", *request.Type))
	}
	
	if len(filters) > 0 {
		searchReq.Filter = strings.Join(filters, " AND ")
	}

	searchResult, err := c.index.Search(request.Query, searchReq)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	return c.mapHitsToResponse(searchResult.Hits, request.Lang), nil
}

func (c *Client) mapHitsToResponse(hits meilisearch.Hits, lang string) []SearchResponse {
	responses := make([]SearchResponse, 0, len(hits))

	for _, hitMap := range hits {
		name := c.getLocalizedField(hitMap, "nameEn", "nameAr", lang)
		description := c.getLocalizedField(hitMap, "descriptionEn", "descriptionAr", lang)
		vendorName := c.getLocalizedField(hitMap, "vendorNameEn", "vendorNameAr", lang)

		response := SearchResponse{
			ID:              c.getString(hitMap, "id"),
			Type:            c.getString(hitMap, "type"),
			Name:            name,
			Description:     description,
			VendorName:      vendorName,
			Picture:         c.getString(hitMap, "picture"),
			BasePrice:       c.getInt(hitMap, "basePrice"),
			DiscountPercent: c.getFloat(hitMap, "discountPercent"),
			DiscountedPrice: c.getInt(hitMap, "discountedPrice"),
			VendorID:        c.getUint(hitMap, "vendorId"),
			CategoryID:      c.getUint(hitMap, "categoryId"),
		}
		responses = append(responses, response)
	}

	return responses
}

func (c *Client) getLocalizedField(hitMap meilisearch.Hit, enField, arField, lang string) string {
	if lang == "ar" {
		if arValue := c.getString(hitMap, arField); arValue != "" {
			return arValue
		}
	}
	return c.getString(hitMap, enField)
}

func (c *Client) getFloat(hitMap meilisearch.Hit, key string) float64 {
	if raw, ok := hitMap[key]; ok {
		var n float64
		if err := json.Unmarshal(raw, &n); err == nil {
			return n
		}
	}
	return 0.0
}


func (c *Client) getString(hitMap meilisearch.Hit, key string) string {
	if raw, ok := hitMap[key]; ok {
		var s string
		if err := json.Unmarshal(raw, &s); err == nil {
			return s
		}
	}
	return ""
}

func (c *Client) getInt(hitMap meilisearch.Hit, key string) int {
	if raw, ok := hitMap[key]; ok {
		var n float64
		if err := json.Unmarshal(raw, &n); err == nil {
			return int(n)
		}
	}
	return 0
}

func (c *Client) getUint(hitMap meilisearch.Hit, key string) uint {
	if raw, ok := hitMap[key]; ok {
		var n float64
		if err := json.Unmarshal(raw, &n); err == nil {
			return uint(n)
		}
	}
	return 0
}