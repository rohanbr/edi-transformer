package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jf-tech/omniparser"
	"github.com/jf-tech/omniparser/transformctx"
)

func main() {

	schemaFile := "./schema.json"
	schemaFileBaseName := filepath.Base(schemaFile)
	schemaFileReader, err := os.Open(schemaFile)
	defer schemaFileReader.Close()

	inputFile := "./sample.log"
	inputFileBaseName := filepath.Base(inputFile)
	inputFileReader, err := os.Open(inputFile)
	defer inputFileReader.Close()

	schem, err := omniparser.NewSchema(schemaFileBaseName, schemaFileReader)

	if err != nil {
		fmt.Println("this is dsfs")
		fmt.Println(err.Error())
		return
	}
	transform, err := schem.NewTransform(inputFileBaseName, inputFileReader, &transformctx.Ctx{})

	var records []string
	for {
		output, err := transform.Read()
		if err == io.EOF {
			break
		}
		records = append(records, string(output))
		// output contains a []byte of the ingested and transformed record.
		fmt.Println("output", string(output))
		// Also transform.RawRecord() gives you access to the raw record.
		a, err := transform.RawRecord()
		fmt.Println(a.Checksum())
	}
}
