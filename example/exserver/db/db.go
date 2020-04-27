package db

import (
	"database/sql"
	"flag"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luno/reflex/example/internal/db"
	"github.com/luno/reflex/rsql"
)

var (
	db_example_uri = flag.String("db_example_uri", getDefaultURI(), "URI of reflex example server DB")

	Events1 = rsql.NewEventsTable("server_events1")
	Events2 = rsql.NewEventsTable("server_events2")
	Cursors = rsql.NewCursorsTable("server_cursors")
)

func Connect() (*sql.DB, error) {
	return db.Connect(*db_example_uri)
}

func ConnectForTesting(t *testing.T) (*sql.DB, error) {
	return db.ConnectForTesting(t, *db_example_uri, "schema.sql")
}

func getDefaultURI() string {
	uri := os.Getenv("DB_EXAMPLE_URI")
	if uri != "" {
		return uri
	}

	return "root:@tcp(localhost:3306)/test?parseTime=true"
}
