package wrappers

//go:generate mockery --name Tx
type Tx interface {
	Commit() error
	Rollback() error
}
