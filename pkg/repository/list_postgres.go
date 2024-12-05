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
