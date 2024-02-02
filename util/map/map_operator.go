package maps

import (
	"bytes"
	"encoding/gob"
)

func Copy(source interface{}, destination interface{}) {
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(source)
	gob.NewDecoder(buf).Decode(destination)
}
