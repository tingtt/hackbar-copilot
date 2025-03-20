package graph

import (
	"hackbar-copilot/internal/domain/user"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type userAdapter user.User

func (c userAdapter) apply() *model.User {
	m := model.User{
		Email: string(c.Email),
		Name:  c.Name,
	}
	return &m
}
