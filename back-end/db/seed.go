package db

func SeedDB(db *DB) error {
	if err := seedJumpHeights(db); err != nil {
		return err
	}

	if err := seedRanks(db); err != nil {
		return err
	}

	return nil
}
