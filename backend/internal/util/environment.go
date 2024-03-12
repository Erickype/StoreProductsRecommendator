package util

import "os"

type Environment int64

const (
	MongodbUri Environment = iota
)

func (e Environment) String() string {
	switch e {
	case MongodbUri:
		return os.Getenv("MONGODB_URI")
	}
	return "unknown"
}
