package phonecode_test

import (
	"github.com/baorv/phonecode"
	"log"
)

func Example() {
	reader, err := phonecode.New()
	if err != nil {
		log.Panicf("Error when open reader with message: %s", err.Error())
	}
	record := reader.LookupPhoneCode("VN")
	log.Printf("%s as calling code is: %s", record.CountryCode, record.Code)
}
