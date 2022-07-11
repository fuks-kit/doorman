package simple

import (
	"encoding/json"
	"log"
)

func PrettifyJsonBytes(byt []byte) (pretty []byte) {
	var obj interface{}

	err := json.Unmarshal(byt, &obj)
	if err != nil {
		log.Fatalln(err)
	}

	pretty, _ = json.MarshalIndent(obj, "", "  ")

	return
}

func PrettifyMarshaler(obj json.Marshaler) (pretty []byte) {
	byt, err := obj.MarshalJSON()
	if err != nil {
		log.Fatalln(err)
	}

	return PrettifyJsonBytes(byt)
}

func PrettifyAny(obj interface{}) (pretty []byte) {
	byt, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	return byt
}
