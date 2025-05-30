package dataloader

import (
	"context"

	"github.com/ryo246912/playground-gqlgen/graph/db"
	"github.com/uptrace/bun"
)

type addressReader struct {
	db *bun.DB
}

func (s *addressReader) getAddress(ctx context.Context, addressIDs []string) ([]*db.Address, []error) {
	addresses := make([]*db.Address, 0, len(addressIDs))

	err := s.db.NewSelect().Model(&addresses).
		Where("address_id IN (?)", bun.In(addressIDs)).
		Scan(ctx)

	if err != nil {
		return nil, []error{err}
	}

	return addresses, nil
}

func GetAddress(ctx context.Context, addressID string) (*db.Address, error) {
	loaders := For(ctx)
	return loaders.AddressLoader.Load(ctx, addressID)
}

func GetAddresses(ctx context.Context, addressIDs []string) ([]*db.Address, error) {
	loaders := For(ctx)
	return loaders.AddressLoader.LoadAll(ctx, addressIDs)
}
