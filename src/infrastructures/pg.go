package infrastructures

import (
	"fmt"
	"log"

	"github.com/galuhpradipta/service-package-layout/src/interfaces"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	PgHandler struct {
		Conn *sqlx.DB
		Tx   *sqlx.Tx
	}

	PgRow struct {
		Rows *sqlx.Rows
	}
)

func (h *PgHandler) Begin() (err error) {
	h.Tx, err = h.Conn.Beginx()
	if err != nil {
		return err
	}
	return nil
}

func (h *PgHandler) Commit() (err error) {
	err = h.Tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (h *PgHandler) Rollback() (err error) {
	err = h.Tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (h *PgHandler) Execute(statement string) (interfaces.IRow, error) {
	rows, err := h.Conn.Queryx(statement)
	if err != nil {
		log.Println(err)
		return new(PgRow), err
	}

	row := new(PgRow)
	row.Rows = rows

	return row, nil
}

func (h *PgHandler) Query(statement string) (interfaces.IRow, error) {
	rows, err := h.Tx.Queryx(statement)

	if err != nil {
		fmt.Println(err)
		return new(PgRow), err
	}
	row := new(PgRow)
	row.Rows = rows

	return row, nil
}

func (r PgRow) Scan(dest ...interface{}) (err error) {
	err = r.Rows.Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func (r PgRow) Next() bool {
	return r.Rows.Next()
}

func (r PgRow) StructScan(dest interface{}) (err error) {
	err = r.Rows.StructScan(dest)
	if err != nil {
		return err
	}
	return nil
}
