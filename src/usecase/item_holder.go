package usecase

import (
	"context"

	"github.com/sky0621/fs-mng-backend/domain"
)

func NewItemHolder(itemHolderDomain domain.ItemHolder) ItemHolder {
	return &itemHolder{itemHolderDomain: itemHolderDomain}
}

type ItemHolder interface {
	GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error)
	GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error)
	CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error)
}

type itemHolder struct {
	itemHolderDomain domain.ItemHolder
}

func (i *itemHolder) GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error) {
	return i.itemHolderDomain.GetItemHolder(ctx, id)
}

func (i *itemHolder) GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error) {
	return i.itemHolderDomain.GetItemHolders(ctx)
}

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	return i.itemHolderDomain.CreateItemHolder(ctx, input)
}
