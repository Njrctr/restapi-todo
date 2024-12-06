package repository

import (
	"fmt"

	todo "github.com/Njrctr/restapi-todo"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tr, err := r.db.Begin() //* Старт транзакции
	if err != nil {
		return 0, err
	}

	var todoListId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := r.db.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&todoListId); err != nil {
		tr.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tr.Exec(createUsersListQuery, userId, todoListId)
	if err != nil {
		tr.Rollback()
		return 0, err
	}

	return todoListId, tr.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id=$1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id=$1 AND ul.list_id=$2", todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
