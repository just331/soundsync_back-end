package model

// CRUD model for Role

import (
	"github.com/globalsign/mgo/bson"
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

// SaveRoles handles creating and updating user roles
func SaveRoles(myRoles *Roles) error {
	err := connection.Collection("roles").Save(myRoles)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		log.Fatal(vErr.Errors)
	} else {
		log.Fatal(err.Error())
	}
	return err
}

// FindRoleByID handles IDing roles of user by their ID
func FindRoleByID(id string) (*Roles, error) {
	role := &Roles{}
	err := connection.Collection("role").FindById(bson.ObjectIdHex(id), role)
	return role, err
}

// DeleteRoles deletes a users roles
func DeleteRoles(myRoles *Roles) error {
	err := connection.Collection("role").DeleteDocument(myRoles)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
