package interfaces

type IKwRepository interface {
	TxBegin() error
	TxCommit() error
	TxRollback() error
}
