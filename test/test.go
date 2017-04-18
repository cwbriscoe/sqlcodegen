package test

import (
	"fmt"

	"github.com/jackc/pgx"
)

func RunTests(conn *pgx.Conn) {
	var err error

	test := &sqlgentest{}
	test.in.parm1 = "0000118532"
	err = test.fetch1(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test.out)
}
