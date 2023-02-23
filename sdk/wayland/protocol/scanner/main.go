package main

import (
	"encoding/xml"
	"flag"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

var (
	protocolPath  string
	outputPath    string
	outputPackage string
)

func init() {
	flag.StringVar(&protocolPath, "protocol", "", "Path to protocol file")
	flag.StringVar(&outputPath, "out", "", "File to write generating code to")
	flag.StringVar(&outputPackage, "package", "", "Package to put generated code in")
	flag.Parse()

	if outputPath == "" {
		log.Fatal().
			Msg("--out is a required argument")
	}

	if outputPackage == "" {
		log.Fatal().
			Msg("--package is a required argument")
	}

	if protocolPath == "" {
		log.Fatal().
			Msg("--protocol is a required argument")
	}

	if _, err := os.Stat(protocolPath); os.IsNotExist(err) {
		log.Fatal().
			Str("path", protocolPath).
			Msg("file does not exist or is not accessible")
	}
}

func main() {
	protocolFile, err := os.Open(protocolPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Opening protocol file")
	}
	defer protocolFile.Close()

	protocolBytes, err := io.ReadAll(protocolFile)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Reading protocol file")
	}

	var protocol XMLProtocol
	if err := xml.Unmarshal(protocolBytes, &protocol); err != nil {
		log.Fatal().
			Err(err).
			Msg("Unmarshalling protocol file")
	}

	generatedCode := genProtocol(protocol, outputPackage)
	//generatedCode.NoFormat = true
	if err := generatedCode.Save(outputPath); err != nil {
		log.Warn().
			Err(err).
			Msg("Error saving file with formatting, trying without")

		generatedCode.NoFormat = true
		if err := generatedCode.Save(outputPath); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error saving file")

		}
	}
}
