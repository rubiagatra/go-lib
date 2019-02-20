package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceByteToInt64(t *testing.T) {
	someInteger := 8
	bt, _ := json.Marshal(someInteger)
	result := InterfaceByteToInt64(interface{}(bt))
	assert.EqualValues(t, someInteger, result)
}
