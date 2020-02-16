package utilities

import "github.com/gorilla/schema"

func GenerateDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.SetAliasTag("json")
	return decoder
}
