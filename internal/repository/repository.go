package repository

import (
	"context"
	"todo/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db     *sqlx.DB
	genSQL sq.StatementBuilderType
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db:     db,
		genSQL: sq.StatementBuilder.PlaceholderFormat(sq.Question),
	}
}

func (r *Repository) Update(ctx context.Context, todo domain.Todo) error {
	query, args, _ := r.genSQL.
		Update("todo").
		Set("body", todo.Body).
		Set("is_done", todo.IsDone).
		Set("priority", todo.Priority).
		Where(sq.Eq{"id": todo.ID}).
		ToSql()
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *Repository) List(ctx context.Context) ([]domain.Todo, error) {
	query, _, _ := r.genSQL.Select("*").From("todo").ToSql()
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return []domain.Todo{}, err
	}
	var todos []domain.Todo
	for rows.Next() {
		var t domain.Todo
		err = rows.StructScan(&t)
		if err != nil {
			return todos, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query, args, _ := r.genSQL.Delete("todo").Where(sq.Eq{"id": id}).ToSql()
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *Repository) Get(ctx context.Context, id int64) (domain.Todo, error) {
	query, args, _ := r.genSQL.
		Select("*").
		From("todo").
		Where(sq.Eq{"id": id}).
		ToSql()
	row := r.db.QueryRowxContext(ctx, query, args...)
	var t domain.Todo
	err := row.StructScan(&t)
	if err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}

func (r *Repository) Insert(ctx context.Context, todo *domain.Todo) (int64, error) {
	query, args, _ := r.genSQL.
		Insert("todo").
		Columns("body", "priority").
		Values(todo.Body, todo.Priority).
		ToSql()
	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
