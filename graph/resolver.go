package graph

import "github.com/ryo246912/playground-gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// リゾルバ（Resolver構造体）に参照し続けたいインスタンスを保持する
type Resolver struct {
	todos []*model.Todo
}
