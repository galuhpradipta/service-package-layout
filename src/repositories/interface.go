package repositories

import "github.com/galuhpradipta/service-package-layout/src/interfaces"

type KwRepository struct {
	interfaces.IDbHandler
}

func (r *KwRepository) TxBegin() (err error) {
	if r.Begin(); err != nil {
		return err
	}
	return nil
}

func (r *KwRepository) TxCommit() (err error) {
	if r.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *KwRepository) TxRollback() (err error) {
	if r.Rollback(); err != nil {
		return err
	}
	return nil
}
