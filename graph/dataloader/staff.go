package dataloader

import (
	"context"

	"github.com/ryo246912/playground-gqlgen/graph/db"
	"github.com/uptrace/bun"
)

type staffReader struct {
	db *bun.DB
}

func (s *staffReader) getStaffs(ctx context.Context, storeIDs []string) ([]*db.Staff, []error) {
	staffs := make([]*db.Staff, 0, len(storeIDs))

	err := s.db.NewSelect().Model(&staffs).
		Where("store_id IN (?)", bun.In(storeIDs)).
		Scan(ctx)

	if err != nil {
		return nil, []error{err}
	}

	return staffs, nil
}

func GetStaff(ctx context.Context, storeID string) (*db.Staff, error) {
	loaders := For(ctx)
	return loaders.StaffLoader.Load(ctx, storeID)
}

func GetStaffs(ctx context.Context, storeIDs []string) ([]*db.Staff, error) {
	loaders := For(ctx)
	return loaders.StaffLoader.LoadAll(ctx, storeIDs)
}
