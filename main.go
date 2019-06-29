package main

import (
	"github.com/cchantan/finalexam/nroute"

	_ "github.com/lib/pq"
)

func main() {
	r := nroute.Nroute()
	r.Run(":2019")
}
