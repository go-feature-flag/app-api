package dbmodel_test

import (
	"testing"

	"github.com/go-feature-flag/app-api/dao/dbmodel"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var theTestMap = map[string]interface{}{
	"key":         "value",
	"another_key": "another_value",
	"nested_key": map[string]interface{}{
		"child_key": "child_value",
	},
}

var theTestString = `{"another_key":"another_value","key":"value","nested_key":{"child_key":"child_value"}}`

func TestJSONBMarshalling(t *testing.T) {
	d := theTestMap
	s, _ := dbmodel.JSONB(d).Value()
	if v, ok := s.([]byte); ok {
		assert.JSONEq(t, theTestString, string(v))
	} else {
		assert.False(t, true, "should be a byte slice")
	}
}

func TestJSONBUnmarshalling(t *testing.T) {
	var theTest dbmodel.JSONB
	err := theTest.Scan([]byte(theTestString))
	assert.NoError(t, err)
	assert.Equal(t, dbmodel.JSONB(theTestMap), theTest)
}

func TestJSONBUnmarshalling_wrongType(t *testing.T) {
	var theTest dbmodel.JSONB
	err := theTest.Scan([]byte("toto"))
	assert.Error(t, err)
}
