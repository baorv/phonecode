package phonecode

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"bufio"
	"github.com/pkg/errors"
)

// New instance of Reader
func New(args ...string) (*Reader, error) {
	var (
		file   = dataFile
		reader = new(Reader)
		err    error
	)
	if len(args) > 0 {
		file = args[0]
	}
	reader.file = file
	reader.records = make(map[string]*Record)
	err = reader.load()
	return reader, err
}

// Define some constants in package
const (
	dataFile = "data/phonecode.csv"
)

type (
	// Record is a struct that represent a record in csv file
	Record struct {
		// Calling code of a country
		Code string `json:"code"`

		// Country code
		CountryCode string `json:"country_code"`
	}

	// Reader is a struct wraps methods to lookup calling code
	Reader struct {
		file    string
		records map[string]*Record
	}
)

// Load csv data

func (reader *Reader) load() error {
	var (
		csvFile, err = os.Open(reader.file)
		r            = csv.NewReader(bufio.NewReader(csvFile))
	)
	if err != nil {
		return errors.New("Error when loading reader with message: " + err.Error())
	}
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		reader.records[line[1]] = &Record{
			Code:        line[0],
			CountryCode: line[1],
		}
	}
	return nil
}

// LookupPhoneCode is a method to lookup country with given phone code
func (reader *Reader) LookupPhoneCode(countryCode string) *Record {
	return reader.records[countryCode]
}

// GetRecords is a method to get all records from reader
func (reader *Reader) GetRecords() map[string]*Record {
	return reader.records
}
