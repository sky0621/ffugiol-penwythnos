package domain

import (
	"context"
	"fmt"
)

type ItemHolder interface {
	GetItemHolder(ctx context.Context, id string) (*QueryItemHolderModel, error)
	GetItemHolders(ctx context.Context) ([]*QueryItemHolderModel, error)
	GetItemHolderAllCount(ctx context.Context) (int, error)
	CreateItemHolder(ctx context.Context, input CommandItemHolderModel) (string, error)
}

type QueryItemHolderModel struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  *string
	HoldItems []QueryItemModel
}

func (i *QueryItemHolderModel) String() string {
	return fmt.Sprintf("[domain/QueryItemHolderModel]ID:%s, FirstName:%s, LastName:%s, Nickname:%v", i.ID, i.FirstName, i.LastName, i.Nickname)
}

type CommandItemHolderModel struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  *string
}

func (i *CommandItemHolderModel) String() string {
	return fmt.Sprintf("[domain/CommandItemHolderModel]ID:%s, FirstName:%s, LastName:%s, Nickname:%v", i.ID, i.FirstName, i.LastName, i.Nickname)
}

type SearchItemHolderConditionModel struct {
	Nickname *string
}
