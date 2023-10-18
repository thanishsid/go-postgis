package postgis

import (
	"database/sql"
	"os"
	"reflect"
)

type Fatalistic interface {
	Fatal(args ...interface{})
}

func openTestConn(t Fatalistic) *sql.DB {
	dsn := os.Getenv("PGDSN")

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis")
	if err != nil {
		t.Fatal("PostGIS extension create failed.")
	}

	return conn
}

func compareGeometry(db *sql.DB, g1 Geometry, g2 Geometry) (bool, error) {
	if err := db.QueryRow("SELECT $1;", g1).Scan(g2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(g1, g2), nil
}
