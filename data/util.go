package data

import (
    "database/sql"
    "github.com/d3code/pkg/log"
)

func CloseRows(rows *sql.Rows) {
    func(rows *sql.Rows) {
        if rows == nil {
            return
        }

        err := rows.Close()
        if err != nil {
            log.Log.Error(err)
        }
    }(rows)
}
