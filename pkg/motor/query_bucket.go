package motor

import (
	"context"
)

func (mtr *motorNodeImpl) QueryWhereIs(ctx context.Context, did string) error {
	err := mtr.Resources.GetWhereIs(ctx, did, mtr.Address)
	if err != nil {
		return err
	}

	return nil
}

func (mtr *motorNodeImpl) QueryWhereIsByCreator(ctx context.Context) error {
	err := mtr.Resources.GetWhereIsByCreator(ctx, mtr.Address)
	if err != nil {
		return err
	}

	return nil
}
