package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// Tool struct
type Tool struct {
	Id       int
	Name     string
	Category string
	URL      string
	Rating   int
	Notes    string
}
