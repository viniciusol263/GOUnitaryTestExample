package soma

import (
	"errors"
	"log"
)

const listSize = 5

var (
	ErrListFull      = errors.New("list is full")
	ErrValueNotFound = errors.New("Value not found")
)

type Value struct {
	id    int
	value int
}
type Contains struct {
	values []Value
}

func (c *Contains) insertInt(id int, value int) ([]Value, error) {
	if len(c.values) == listSize {
		return c.values, ErrListFull
	}
	c.values = append(c.values, Value{id, value})
	log.Printf("value %d was added into the list\n", value)
	return c.values, nil
}

func (c *Contains) searchId(id int) (*Value, error) {
	for _, value := range c.values {
		if value.id == id {
			log.Printf("Found Value %+v\n", value)
			return &value, nil
		}
	}
	return nil, ErrValueNotFound
}
