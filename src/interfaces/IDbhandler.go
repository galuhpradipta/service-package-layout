package interfaces

type IDbHandler interface {
	Begin() error
	Commit() error
	Rollback() error
	Execute(statement string) (IRow, error)
	Query(statement string) (IRow, error)
}

type IRow interface {
	StructScan(dest interface{}) error
	Scan(dest ...interface{}) error
	Next() bool
}
