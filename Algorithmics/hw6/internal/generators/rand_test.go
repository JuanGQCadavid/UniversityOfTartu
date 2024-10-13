package generators

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	upperLimit     int32 = 999999999
	maxValuesSizes int32 = 1000000
)

func TestDoesComputerSupportMaxVals(t *testing.T) {
	tenList := GenerateRanList(maxValuesSizes, upperLimit)

	for _, val := range tenList {
		assert.True(t, val > 0, "Nuber is less than 0")
		assert.True(t, val < upperLimit, "Nuber is bigger than upperLimit")
	}

	assert.True(t, len(tenList) > 0, "len is less than 0")
}
func TestListGen(t *testing.T) {
	tenList := GenerateRanList(10, upperLimit)

	for _, val := range tenList {
		log.Println(val)
		assert.True(t, val > 0, "Nuber is less than 0")
		assert.True(t, val < upperLimit, "Nuber is bigger than upperLimit")
	}
	assert.True(t, len(tenList) > 0, "len is less than 0")
}

func TestNumberGen(t *testing.T) {
	number := GenRanInt(upperLimit)
	log.Println(number)
	assert.True(t, number > 0, "Nuber is less than 0")
	assert.True(t, number < upperLimit, "Nuber is bigger than upperLimit")
}
