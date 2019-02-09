package model

// CRUD model for Role

import (
	"log"
	"os"

	"github.com/go-bongo/bongo"
	/*
		"github.com/globalsign/mgo/bson"
		"github.com/go-bongo/bongo"
	*/)

var (
	config = &bongo.Config{
		ConnectionString: "",
		Database:         "soundsync",
	}

	connection, connErr = bongo.Connect(config)
)

func init() {
	config.ConnectionString = os.Getenv("DefaultConnection")

	connection, connErr = bongo.Connect(config)
	if connErr != nil {
		log.Fatal(connErr)
	}
}
