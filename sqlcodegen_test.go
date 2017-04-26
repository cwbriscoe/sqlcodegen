package sqlcodegen

import (
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx"
)

//var conn *pgx.Conn

func TestMain(m *testing.M) {
	var err error

	dbcfg := &pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "website",
		User:     "dba",
		Password: "dba",
	}

	conn, err = pgx.Connect(*dbcfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
}

func TestSingleSelect(t *testing.T) {
	test := &sqlgentest{}
	test.in.parm1 = "0000118532"
	err := test.fetch1(conn)
	if err != nil {
		t.Error(err)
	}
}
