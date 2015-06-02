package count

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	SELECT = "SELECT %s FROM %s WHERE id=%s"
	UPDATE = "UPDATE %s SET %s=%s%s WHERE id=%s"
)

func Add(udb *sql.DB, table, field string, id, num int) {
	var op string
	if num >= 0 {
		op = fmt.Sprintf("+%d", num)
	} else {
		op = fmt.Sprintf("-%d", -num)
	}

	s := fmt.Sprintf(UPDATE, table, field, field, op, "$1")
	_, err := udb.Exec(s, &id)
	if err != nil {
		log.Panicln(err)
	}
}

func Get(udb *sql.DB, table, field string, id int) int64 {
	var value sql.NullInt64
	s := fmt.Sprintf(SELECT, field, table, "$1")
	err := udb.QueryRow(s, &id).Scan(&value)
	if err != nil {
		log.Panicln(err)
	}
	return value.Int64
}
