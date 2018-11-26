package phonecode_test

import (
	"testing"
	"github.com/baorv/phonecode"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var (
		reader, _ = phonecode.New()
		_, err    = phonecode.New("fake.csv")
	)
	assert.IsType(t, &phonecode.Reader{}, reader)
	assert.Error(t, err)
}

func TestReader_LookupPhoneCode(t *testing.T) {
	var (
		reader, _ = phonecode.New()
		record    = reader.LookupPhoneCode("VN")
		record2   = reader.LookupPhoneCode("AWM")
	)
	assert.Equal(t, "84", record.Code)
	assert.Nil(t, record2)
}
func TestReader_GetRecords(t *testing.T) {
	reader, _ := phonecode.New()
	assert.IsType(t, map[string]*phonecode.Record{}, reader.GetRecords())
}
