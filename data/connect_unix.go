package data

import (
    "database/sql"
    "fmt"
    "github.com/d3code/pkg/configuration"
    "github.com/d3code/pkg/log"
    _ "github.com/go-sql-driver/mysql"
)

// connectUnixSocket initializes a Unix socket connection pool for a Cloud SQL instance of MySQL
func connectUnixSocket(databaseConfig configuration.DatabaseConfig) (*sql.DB, error) {
    var (
        user           = databaseConfig.User
        password       = databaseConfig.Password
        databaseName   = databaseConfig.DatabaseName
        unixSocketPath = databaseConfig.ConnectionName // /cloudsql/project:region:instance
    )

    connectionString := fmt.Sprintf("%s:%s@unix(/%s)/%s?parseTime=true", user, password, unixSocketPath, databaseName)

    // databaseConnection is the pool of databaseConnection connections.
    databaseConnection, err := sql.Open("mysql", connectionString)
    if err != nil {
        log.Log.Error(err)
        return nil, err
    }

    return databaseConnection, nil
}
