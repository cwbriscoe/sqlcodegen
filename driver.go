package main

import (
	"fmt"
	"os"

	"github.com/cwbriscoe/sqlcodegen/test"
	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func main() {
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

	test.RunTests(conn)
	/*
		test := &test.Sqlgentest{}
		test.in.parm1 = "0000118532"
		err = test.fetch1(conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(test.out) */
}
