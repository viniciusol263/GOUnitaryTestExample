package soma

import (
	"reflect"
	"testing"
)

var containsTest Contains

func TestContainer(t *testing.T) {
	// Tests InsertInt method
	testDataInsertInt := []struct {
		value_t  Value
		response []Value
	}{
		{value_t: Value{id: 0, value: 42}, response: []Value{{id: 0, value: 42}}},
		{value_t: Value{id: 1, value: 231}, response: []Value{{id: 0, value: 42}, {id: 1, value: 231}}},
		{value_t: Value{id: 2, value: 3}, response: []Value{{id: 0, value: 42}, {id: 1, value: 231}, {id: 2, value: 3}}},
		{value_t: Value{id: 3, value: 1034}, response: []Value{{id: 0, value: 42}, {id: 1, value: 231}, {id: 2, value: 3}, {id: 3, value: 1034}}},
		{value_t: Value{id: 4, value: 598}, response: []Value{{id: 0, value: 42}, {id: 1, value: 231}, {id: 2, value: 3}, {id: 3, value: 1034}, {id: 4, value: 598}}},
		{value_t: Value{id: 5, value: 3423}, response: []Value{{id: 0, value: 42}, {id: 1, value: 231}, {id: 2, value: 3}, {id: 3, value: 1034}, {id: 4, value: 598}}},
	}
	{
		for _, testValue := range testDataInsertInt {
			tmpTest, err := containsTest.insertInt(testValue.value_t.id, testValue.value_t.value)
			if !reflect.DeepEqual(tmpTest, testValue.response) {
				t.Fatalf("Received return %v - Expected return %v, error %v\n", tmpTest, testValue.response, err)
			}
		}
	}

	//Tests SearchId method
	testDataSearchId := []struct {
		id_t     int
		response *Value
	}{
		{id_t: 0, response: &Value{id: 0, value: 42}},
		{id_t: 1, response: &Value{id: 1, value: 231}},
		{id_t: 2, response: &Value{id: 2, value: 3}},
		{id_t: 3, response: &Value{id: 3, value: 1034}},
		{id_t: 4, response: &Value{id: 4, value: 598}},
		{id_t: 5, response: nil},
	}
	{
		for _, testValue := range testDataSearchId {
			tmpTest := containsTest.searchId(testValue.id_t)
			if !reflect.DeepEqual(testValue.response, tmpTest) {
				t.Fatalf("Received return %v - Expected return %v\n", tmpTest, testValue.response)
			}
		}
	}
}
