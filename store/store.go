package store

import (
	"errors"

	"github.com/ishisora/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	// 動作確認用の仮実装のため、あえてexportしている
	LastID entity.TaskID
	Tasks  map[entity.Task]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

// All はソート済みのタスク一覧を返す
func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		// LastIDは1から1ずつ増えている。初期値0からAddで先にインクリメントされてTasksのキーとなるため。
		// スライスはインデックス0から指定する必要があるため-1している
		tasks[i-1] = t
	}
	return tasks
}
