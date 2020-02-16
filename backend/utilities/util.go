package utilities

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

type DynamicStruct struct {
	CType  interface{}
	Writer http.ResponseWriter
	Req    *http.Request
}

type DecoderMetadata struct {
	IgnoreUnknownKeys bool
	AliasTag          string
}

// Consider making this user configurable at a higher level?
var decoderOpts = &DecoderMetadata{IgnoreUnknownKeys: true, AliasTag: "json"}

func GenerateDecoder(opts *DecoderMetadata) *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(opts.IgnoreUnknownKeys)
	decoder.SetAliasTag(opts.AliasTag)
	return decoder
}

func PersistRequest(s *DynamicStruct) {
	r := s.Req
	rw := s.Writer
	data := s.CType

	err := r.ParseForm()
	if err != nil {
		log.Printf("HTTP %d - %s", 500, err.Error())
		http.Error(rw, err.Error(), 500)
	}

	decoder := GenerateDecoder(decoderOpts)
	err = decoder.Decode(data, r.Form)

	if err != nil {
		log.Printf("HTTP %d - %s", 500, err.Error())
		http.Error(rw, err.Error(), 500)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

}
