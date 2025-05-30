//go:generate go run github.com/99designs/gqlgen generate
// Goのビルドツールに「このファイルでgo generateコマンドを実行したときに、どんなコマンドを走らせるか」を教える特別なコメント
// package宣言の直前や直後、importの前など、Goファイルの先頭付近に書きます。

package graph

import (
	"github.com/ryo246912/playground-gqlgen/graph/model"
	"github.com/uptrace/bun"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// リゾルバ（Resolver構造体）に参照し続けたいインスタンスを保持する
type Resolver struct {
	todos []*model.Todo
	DB    *bun.DB
}
