package service

import (
	"context"
	"fmt"

	"github.com/ishisora/go_todo_app/auth"
	"github.com/ishisora/go_todo_app/entity"
	"github.com/ishisora/go_todo_app/store"
)

type GetTask struct {
	DB   store.Queryer
	Repo TaskGetter
}

func (g *GetTask) GetTask(ctx context.Context, id entity.TaskID) (*entity.Task, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	t, err := g.Repo.GetTask(ctx, g.DB, id, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return t, nil
}
