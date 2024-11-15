package models

import (
	"errors"
)

// New makes a new foodItem, adds to foodManager and saves to file
func New(args []string) error {
	if len(args) < 3 {
		return errors.New("not enough arguments to make a record")
	}
	return nil

}

// Update updates a food record with new information
func Update(recordIndex int) error {
	return nil
}

// Delete deletes a food record
func Delete(recordIndex int) error {
	return nil
}

// List lists all food records
func List() {

}