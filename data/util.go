package data

import (
    "database/sql"
    "github.com/d3code/pkg/zlog"
)

func CloseRows(rows *sql.Rows) {
    func(rows *sql.Rows) {
        if rows == nil {
            return
        }

        err := rows.Close()
        if err != nil {
            zlog.Log.Error(err)
        }
    }(rows)
}
