package util

import "os"

type Database int64

const (
	Products Database = iota
)

func (d Database) String() string {
	switch d {
	case Products:
		return os.Getenv("PRODUCTS_DATABASE")
	}
	return "unknown"
}
