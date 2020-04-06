package usecase

import (
	"context"

	"github.com/sky0621/fs-mng-backend/domain"
)

func NewItem(itemDomain domain.Item) Item {
	return &item{
		itemDomain: itemDomain,
	}
}

type Item interface {
	GetItem(ctx context.Context, id string, selectFields []string) (*domain.QueryItemModel, error)
	GetItems(ctx context.Context) ([]*domain.QueryItemModel, error)
	GetItemsByItemHolderID(ctx context.Context, itemHolderID string) ([]*domain.QueryItemModel, error)
	CreateItem(ctx context.Context, input domain.CommandItemModel) (string, error)
}

type item struct {
	itemDomain domain.Item
}

func (i *item) GetItem(ctx context.Context, id string, selectFields []string) (*domain.QueryItemModel, error) {
	return i.itemDomain.GetItem(ctx, id, selectFields)
}

func (i *item) GetItems(ctx context.Context) ([]*domain.QueryItemModel, error) {
	return i.itemDomain.GetItems(ctx)
}

func (i *item) GetItemsByItemHolderID(ctx context.Context, itemHolderID string) ([]*domain.QueryItemModel, error) {
	return i.itemDomain.GetItemsByItemHolderID(ctx, itemHolderID)
}

func (i *item) CreateItem(ctx context.Context, input domain.CommandItemModel) (string, error) {
	return i.itemDomain.CreateItem(ctx, input)
}
