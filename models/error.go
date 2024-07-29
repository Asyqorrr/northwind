package models

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrAccessForbidden = errors.New("access forbidden")
	ErrDataNotFound    = errors.New("data not found")
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidationError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = err.Error()
	return &e
}

func NewError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = err.Error()
	return &e
}

func Nullable[T any](row *T, err error) (*T, error) {
	if err == nil {
		return row, nil
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func NullableList[T any](rows []*T, err error) ([]*T, error) {
	if err == nil {
		return rows, nil
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func NullableID(row string, err error) (string, error) {
	if err == nil {
		return row, nil
	}

	if err == pgx.ErrNoRows {
		return "", nil
	}

	return "", err
}
