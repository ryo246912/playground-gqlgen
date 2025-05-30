package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/ryo246912/playground-gqlgen/graph/db"
	"github.com/uptrace/bun"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	// *dataloadgen.Loader[K, V]を設定
	// K = 検索条件となるkeyの型
	// V = 検索結果の型
	StoreLoader *dataloadgen.Loader[string, *db.Store]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(conn *bun.DB) *Loaders {
	// define the data loader
	// ur := &userReader{db: conn}
	sr := &storeReader{db: conn}
	return &Loaders{
		StoreLoader: dataloadgen.NewLoader(sr.getStores, dataloadgen.WithWait(time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(conn *bun.DB, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(conn)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
