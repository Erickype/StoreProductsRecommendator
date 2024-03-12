package util

import "os"

type Collection int64

const (
	Categories Collection = iota
)

func (c Collection) String() string {
	switch c {
	case Categories:
		return os.Getenv("CATEGORIES_COLLECTION")
	}
	return "unknown"
}
