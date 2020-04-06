package usecase

import (
	"context"

	system "github.com/sky0621/fs-mng-backend"

	"github.com/sky0621/fs-mng-backend/domain"
)

func NewItem(itemDomain domain.Item, io system.IO) Item {
	return &item{
		itemDomain: itemDomain,
		io:         io,
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
	io         system.IO
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
