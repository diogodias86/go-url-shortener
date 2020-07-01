package urlgenerator

import "github.com/teris-io/shortid"

var (
	sid, _ = shortid.New(1, shortid.DefaultABC, 2342)
)

func Generate() string {
	id, _ := sid.Generate()
	return id
}
