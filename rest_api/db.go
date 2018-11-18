package main

import (
	"github.com/gocql/gocql"
	"log"
)

var Session *gocql.Session

func cassandraInit() {
	var err error

	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "nehalem"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	log.Println("cassandra init done")
}
