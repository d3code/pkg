package data

import (
    "database/sql"
    "github.com/d3code/zlog"
)

func ClosePrepare(prepare *sql.Stmt) {
    prepareError := prepare.Close()
    if prepareError != nil {
        zlog.Log.Error(prepareError)
    }
}
