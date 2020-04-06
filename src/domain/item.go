package domain

import (
	"context"
	"fmt"
)

type Item interface {
	GetItem(ctx context.Context, id string, selectFields []string) (*QueryItemModel, error)
	GetItems(ctx context.Context) ([]*QueryItemModel, error)
	GetItemsByItemHolderID(ctx context.Context, itemHolderID string) ([]*QueryItemModel, error)
	CreateItem(ctx context.Context, input CommandItemModel) (string, error)
}

type QueryItemModel struct {
	ID           string
	Name         string
	Price        int
	ItemHolderID string
}

func (i *QueryItemModel) String() string {
	return fmt.Sprintf("[domain/QueryItemModel]ID:%s, Name:%s, Price:%d, ItemHolderID:%s", i.ID, i.Name, i.Price, i.ItemHolderID)
}

type CommandItemModel struct {
	ID           string
	Name         string
	Price        int
	ItemHolderID string
}

func (i *CommandItemModel) String() string {
	return fmt.Sprintf("[domain/CommandItemModel]ID:%s, Name:%s, Price:%d, ItemHolderID:%s", i.ID, i.Name, i.Price, i.ItemHolderID)
}
