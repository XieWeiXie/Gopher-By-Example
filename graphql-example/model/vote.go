package model

import (
	"Gopher-By-Example/graphql-example/pkg/database"
	"time"
)

type base struct {
	Id        int64      `xorm:"pk autoincr notnull" json:"id"`
	CreatedAt time.Time  `xorm:"created" json:"created_at"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_at"`
	DeletedAt *time.Time `xorm:"deleted" json:"deleted_at"`
}

const (
	SINGLE = iota
	MULTIPLE
)

var ClassMap = map[int]string{}

func init() {
	ClassMap = make(map[int]string)
	ClassMap[SINGLE] = "SINGLE"
	ClassMap[MULTIPLE] = "MULTIPLE"
}

type Vote struct {
	base        `xorm:"extends"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OptionIds   []int64   `json:"option_ids"`
	Deadline    time.Time `json:"deadline"`
	Class       int       `json:"class"`
}

type VoteSerializer struct {
	Id          int64              `json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Options     []OptionSerializer `json:"options"`
	Deadline    time.Time          `json:"deadline"`
	Class       int                `json:"class"`
	ClassString string             `json:"class_string"`
}

func (V Vote) TableName() string {
	return "votes"
}

func (V Vote) Serializer() VoteSerializer {
	var optionSerializer []OptionSerializer
	var options []Option
	database.Engine.Where("id in (?)", V.OptionIds).Find(&options)
	for _, i := range options {
		optionSerializer = append(optionSerializer, i.Serializer())
	}
	return VoteSerializer{
		Id:          V.Id,
		CreatedAt:   V.CreatedAt.Truncate(time.Second),
		UpdatedAt:   V.UpdatedAt.Truncate(time.Second),
		Title:       V.Title,
		Description: V.Description,
		Options:     optionSerializer,
		Deadline:    V.Deadline,
		Class:       V.Class,
		ClassString: ClassMap[V.Class],
	}
}

type Option struct {
	base `xorm:"extends"`
	Name string `json:"name"`
}

type OptionSerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (O Option) TableName() string {
	return "options"
}

func (O Option) Serializer() OptionSerializer {
	return OptionSerializer{
		Id:        O.Id,
		CreatedAt: O.CreatedAt.Truncate(time.Second),
		UpdatedAt: O.UpdatedAt.Truncate(time.Second),
		Name:      O.Name,
	}
}
