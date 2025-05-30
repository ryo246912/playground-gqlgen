package dataloader

import (
	"context"

	"github.com/ryo246912/playground-gqlgen/graph/db"
	"github.com/uptrace/bun"
)

type storeReader struct {
	db *bun.DB
}

func (s *storeReader) getStores(ctx context.Context, storeIDs []string) ([]*db.Store, []error) {
	stores := make([]*db.Store, 0, len(storeIDs))

	err := s.db.NewSelect().Model(&stores).
		Where("store_id IN (?)", bun.In(storeIDs)).
		Scan(ctx)

	if err != nil {
		return nil, []error{err}
	}

	return stores, nil
}

// returns single by id efficiently
func GetStore(ctx context.Context, storeID string) (*db.Store, error) {
	loaders := For(ctx)
	return loaders.StoreLoader.Load(ctx, storeID)
}

// returns many by ids efficiently
func GetStores(ctx context.Context, storeIDs []string) ([]*db.Store, error) {
	loaders := For(ctx)
	return loaders.StoreLoader.LoadAll(ctx, storeIDs)
}
