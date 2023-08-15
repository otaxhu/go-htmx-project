package wrappers

type Tx interface {
	Commit() error
	Rollback() error
}
