package data

import (
    "cloud.google.com/go/cloudsqlconn"
    "context"
    "database/sql"
    "fmt"
    "github.com/d3code/pkg/cfg"
    "github.com/go-sql-driver/mysql"
    "net"
)

// connectWithConnector initializes a SQL Cloud connector connection pool
func connectWithConnector(databaseConfig cfg.DatabaseConfig) (*sql.DB, error) {
    var (
        user                   = databaseConfig.User           // e.g. 'my-db-user'
        password               = databaseConfig.Password       // e.g. 'my-db-password'
        databaseName           = databaseConfig.DatabaseName   // e.g. 'my-database'
        instanceConnectionName = databaseConfig.ConnectionName // e.g. 'project:region:instance'
        private                = databaseConfig.Private
    )

    dialer, err := cloudsqlconn.NewDialer(context.Background())
    if err != nil {
        return nil, fmt.Errorf("cloudsqlconn.NewDialer: %v", err)
    }

    mysql.RegisterDialContext("cloudsqlconn",
        func(ctx context.Context, addr string) (net.Conn, error) {
            if private {
                return dialer.Dial(ctx, instanceConnectionName, cloudsqlconn.WithPrivateIP())
            }
            return dialer.Dial(ctx, instanceConnectionName)
        })

    connectionString := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true", user, password, databaseName)
    databaseConnection, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, fmt.Errorf("sql.Open: %v", err)
    }

    return databaseConnection, nil
}
