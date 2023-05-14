package data

import (
    "crypto/tls"
    "crypto/x509"
    "database/sql"
    "errors"
    "fmt"
    "github.com/d3code/pkg/cfg"
    "github.com/d3code/zlog"
    "github.com/go-sql-driver/mysql"
    "io/ioutil"
)

// connectTCPSocket initializes a TCP connection pool
func connectTCPSocket(databaseConfig cfg.DatabaseConfig) (*sql.DB, error) {
    var (
        user         = databaseConfig.User         // e.g. 'my-db-user'
        password     = databaseConfig.Password     // e.g. 'my-db-password'
        databaseName = databaseConfig.DatabaseName // e.g. 'my-database'
        port         = databaseConfig.Port         // e.g. '3306'
        host         = databaseConfig.Host         // e.g. '127.0.0.1'
    )

    // configureSSLCertificates if databaseConfig.RootCertPath is present
    connectionStringOptionSSL, err := configureSSLCertificates(databaseConfig)
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    // Format the connection string
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true%s",
        user, password, host, port, databaseName, connectionStringOptionSSL)

    // databaseConnection is the pool of database connections.
    databaseConnection, err := sql.Open("mysql", connectionString)
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    return databaseConnection, nil
}

// configureSSLCertificates Configure SSL certificates
// For deployments that connect directly to Cloud SQL instances without
// using the Cloud SQL Proxy, configuring SSL certificates will ensure the
// connection is encrypted.
func configureSSLCertificates(databaseConfig cfg.DatabaseConfig) (string, error) {
    rootCert := databaseConfig.RootCertPath // e.g., '/path/to/my/server-ca.pem'
    if rootCert != "" {
        var (
            certPath = databaseConfig.CertPath // e.g. '/path/to/my/client-cert.pem'
            keyPath  = databaseConfig.KeyPath  // e.g. '/path/to/my/client-key.pem'
        )

        pem, err := ioutil.ReadFile(rootCert)
        if err != nil {
            zlog.Log.Error(err)
            return "", err
        }

        certPool := x509.NewCertPool()
        if ok := certPool.AppendCertsFromPEM(pem); !ok {
            return "", errors.New("unable to append root cert to certPool")
        }

        cert, err := tls.LoadX509KeyPair(certPath, keyPath)
        if err != nil {
            zlog.Log.Error(err)
            return "", err
        }

        err = mysql.RegisterTLSConfig("cloudsql", &tls.Config{
            RootCAs:               certPool,
            Certificates:          []tls.Certificate{cert},
            InsecureSkipVerify:    true,
            VerifyPeerCertificate: verifyPeerCertFunc(certPool),
        })
        if err != nil {
            zlog.Log.Error(err)
            return "", err
        }

        return "&tls=cloudsql", nil
    }

    return "", nil
}

// verifyPeerCertFunc returns a function that verifies the peer certificate is in the cert pool.
func verifyPeerCertFunc(pool *x509.CertPool) func([][]byte, [][]*x509.Certificate) error {
    return func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
        if len(rawCerts) == 0 {
            return errors.New("no certificates available to verify")
        }

        cert, err := x509.ParseCertificate(rawCerts[0])
        if err != nil {
            return err
        }

        opts := x509.VerifyOptions{Roots: pool}
        if _, err = cert.Verify(opts); err != nil {
            return err
        }

        return nil
    }
}
