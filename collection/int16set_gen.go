// Code generated - DO NOT EDIT.

package collection

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

// Int16Set holds a set of int16 values.
type Int16Set map[int16]bool

// NewInt16Set creates a new set from its input values.
func NewInt16Set(values ...int16) Int16Set {
	s := Int16Set{}
	s.Add(values...)
	return s
}

// Empty returns true if there are no values in the set.
func (s Int16Set) Empty() bool {
	return len(s) == 0
}

// Clear the set.
func (s Int16Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Add values to the set.
func (s Int16Set) Add(values ...int16) {
	for _, v := range values {
		s[v] = true
	}
}

// Contains returns true if the value exists within the set.
func (s Int16Set) Contains(value int16) bool {
	_, ok := s[value]
	return ok
}

// Clone returns a copy of the set.
func (s Int16Set) Clone() Int16Set {
	if s == nil {
		return nil
	}
	clone := Int16Set{}
	for value := range s {
		clone[value] = true
	}
	return clone
}

// Values returns all values in the set.
func (s Int16Set) Values() []int16 {
	values := make([]int16, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

// MarshalJSON implements the json.Marshaler interface.
func (s Int16Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s Int16Set) UnmarshalJSON(data []byte) error {
	var values []int16
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	s.Clear()
	s.Add(values...)
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (s Int16Set) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(s.Values())
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (s Int16Set) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var values []int16
	if err := unmarshal(&values); err != nil {
		return err
	}
	s.Clear()
	s.Add(values...)
	return nil
}
