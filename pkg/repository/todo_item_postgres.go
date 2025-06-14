package repository

import (
	"fmt"
	"github.com/Senechkaaa/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2) RETURNING id", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, itemId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INNER JOIN %s li on li.item_id = tl.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INNER JOIN %s li on li.item_id = tl.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE tl.id = $1 AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s li, %s ul WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND tl.id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s li, %s ul WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND tl.id = $%d", todoItemsTable, setQuery, listsItemsTable, argId, argId+1)

	args = append(args, itemId, userId)
	_, err := r.db.Exec(query, args...)
	return err
}
