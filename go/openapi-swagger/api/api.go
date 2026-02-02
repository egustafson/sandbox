package api

import (
	"sync"

	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-server-codegen.yaml ../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-models-codegen.yaml ../openapi.yaml

type ItemStore struct {
	items map[string]Item
	lock  sync.Mutex
}

func NewItemStore() *ItemStore {
	return &ItemStore{
		items: make(map[string]Item),
	}
}

// ItemStore implements ServerInterface
var _ ServerInterface = (*ItemStore)(nil)

func (s *ItemStore) PostItems(c *gin.Context) {
	var newItem PostItemsJSONRequestBody
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.items[newItem.Id] = newItem

	c.JSON(201, newItem)
}

func (s *ItemStore) DeleteItemsItemID(c *gin.Context, itemID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exists := s.items[itemID]; !exists {
		c.JSON(404, NotFound{
			Id:      itemID,
			Message: "Item not found",
		})
		return
	}

	delete(s.items, itemID)
	c.Status(204)
}

func (s *ItemStore) GetItems(c *gin.Context) {
	s.lock.Lock()
	defer s.lock.Unlock()

	itemList := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		itemList = append(itemList, item)
	}

	c.JSON(200, itemList)
}

func (s *ItemStore) GetItemsItemID(c *gin.Context, itemID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	item, exists := s.items[itemID]
	if !exists {
		c.JSON(404, NotFound{
			Id:      itemID,
			Message: "Item not found",
		})
		return
	}

	c.JSON(200, item)
}
