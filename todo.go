package todo

import "github.com/pkg/errors"

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	ID     int
	UserId int
	ListId int
}

type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UserItem struct {
	ID     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil &&  i.Description == nil {
		return errors.New("update values has no values")
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool    `json:"done" db:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil &&  i.Description == nil && i.Done == nil {
		return errors.New("update values has no values")
	}
	return nil
}
