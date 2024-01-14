package keeper

// Migrator is a struct for handling in-place state migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns Migrator instance for the state migration.
func NewMigrator(k Keeper) Migrator {
	return Migrator{
		keeper: k,
	}
}

// Migrate1to2 migrates the module state from version 1 to version 2.
// func (m Migrator) Migrate1to2(ctx context.Context) error {
// 	return v2.Migrate(/* migrations arguments */)
// }
