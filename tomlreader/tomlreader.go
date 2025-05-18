package tomlreader

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Method for to read a toml file details
func ReadTomlFile(pFilename string) interface{} {
	var lFileDetails interface{}

	_, lErr := toml.DecodeFile(pFilename, &lFileDetails)
	if lErr != nil {
		log.Println("Error (TRRTF01) ", lErr.Error())
	}

	return lFileDetails
}
