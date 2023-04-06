package data

import (
    "database/sql"
    "github.com/d3code/pkg/cfg"
    "github.com/d3code/pkg/log"
)

var database = make(map[string]*sql.DB)

// GetDatabase lazily instantiates a database connection pool. Users of Cloud Run or
// Cloud Functions may wish to skip this lazy instantiation and connect as soon
// as the function is loaded. This is primarily to help testing.
func GetDatabase(databaseName string) *sql.DB {
    if database[databaseName] == nil {
        databaseConfig := cfg.GetDatabaseConfig(databaseName)
        database[databaseName] = mustConnect(databaseConfig)
    }

    return database[databaseName]
}

// mustConnect creates a connection to the database based on environment configuration
func mustConnect(databaseConfig cfg.DatabaseConfig) *sql.DB {
    var (
        db  *sql.DB
        err error
    )

    if databaseConfig.ConnectionType == "tcp" {
        db, err = connectTCPSocket(databaseConfig)
        if err != nil {
            log.Log.Fatalf("connectTCPSocket: unable to connect: %s", err)
        }
    } else if databaseConfig.ConnectionType == "unix" {
        db, err = connectUnixSocket(databaseConfig)
        if err != nil {
            log.Log.Fatalf("connectUnixSocket: unable to connect: %s", err)
        }
    } else if databaseConfig.ConnectionType == "connector" {
        db, err = connectWithConnector(databaseConfig)
        if err != nil {
            log.Log.Fatalf("connectConnector: unable to connect: %s", err)
        }
    } else {
        log.Log.Fatal("Missing database connection_type")
    }

    if db == nil {
        log.Log.Fatal("Database was not created")
    }

    pingError := db.Ping()
    if pingError != nil {
        log.Log.Fatal(pingError)
    }

    return db
}
