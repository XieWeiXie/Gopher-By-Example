package vote

import (
	"gopkg.in/go-playground/validator.v9"
)

type CreateVoteParams struct {
	Title       string         `json:"title" validate:"required"`
	Description string         `json:"description"`
	Options     []OptionParams `json:"options" validate:"min=2"`
	Deadline    string         `json:"deadline"`
	Class       int            `json:"class" validate:"required"`
}

type OptionParams struct {
	Name string `json:"name" validate:"name"`
}

func (C CreateVoteParams) Valid() error {
	return validator.New().Struct(C)
}
