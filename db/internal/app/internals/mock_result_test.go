package internals_test

type MockResult struct{}

func (r MockResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (r MockResult) RowsAffected() (int64, error) {
	return 1, nil
}
