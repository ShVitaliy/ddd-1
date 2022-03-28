package ddd

import (
	"github.com/google/uuid"
	"math/bits"
	"testing"
)

func TestLocalIdParsing(t *testing.T) {
	cases := []struct {
		value  any
		result int
		error  string
	}{
		{1, 1, ""},
		{int8(-3), -3, ""},
		{uint8(255), 255, ""},
		{int16(12), 12, ""},
		{uint16(32000), 32000, ""},
		{int32(-1), -1, ""},
		{uint32(100), 100, ""},
		{int64(-10), -10, ""},
		{uint64(5), 5, ""},
		{uint64(1<<64 - 1), 0, "Invalid integer: cannot convert 18446744073709551615 to int."},
		{"123", 123, ""},
		{"abc", 0, "Invalid integer: strconv.Atoi: parsing \"abc\": invalid syntax."},
	}

	if bits.UintSize == 32 {
		cases = append(cases, []struct {
			value  any
			result int
			error  string
		}{
			{uint32(1<<32 - 1), 0, "Invalid integer: cannot convert 4294967295 to int."},
			{int64(1<<63 - 1), 0, "Invalid integer: cannot convert 9223372036854775807 to int."},
		}...)
	}

	for _, c := range cases {
		id, err := ParseLocalId(c.value)
		if c.error == "" && err != nil {
			t.Errorf("Failed to parse \"%v\" (type %T) to local id: \"%s\"", c.value, c.value, err.Error())
		} else if c.error != "" {
			if err == nil {
				t.Errorf(
					"Parsing \"%v\" (type %T) must be failed with error \"%s\", but no error was given.",
					c.value,
					c.value,
					c.error,
				)
			} else if err.Error() != c.error {
				t.Errorf(
					"Parsing \"%v\" (type %T) must be failed with error \"%s\", but \"%s\" was given.",
					c.value,
					c.value,
					c.error,
					err.Error(),
				)
			}
		} else if id.identity != c.result {
			t.Errorf(
				"Parsing \"%v\" (type %T) with result %d does not match expected value %d.",
				c.value,
				c.value,
				id.GetIdentity(),
				c.result,
			)
		}
	}
}

func TestLocalIdFrom(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("LocalIdFrom: parsing must panic, but it does not.")
		}
	}()

	// valid parsing
	if LocalIdFrom("123").identity != 123 {
		t.Error("LocalIdFrom: invalid parsing.")
	}

	// invalid parsing
	LocalIdFrom("abc")
}

func TestGlobalIdParsing(t *testing.T) {
	cases := []struct {
		value  any
		result uuid.UUID
		error  string
	}{
		{
			"a0d0d3d7-c0e1-4880-8ea8-c7dd17b1ebba",
			uuid.MustParse("a0d0d3d7-c0e1-4880-8ea8-c7dd17b1ebba"),
			"",
		},
		{
			[]byte{43, 99, 88, 47, 99, 163, 65, 140, 150, 188, 247, 42, 178, 215, 22, 76},
			uuid.MustParse("2b63582f-63a3-418c-96bc-f72ab2d7164c"),
			"",
		},
		{
			"abcdefgh-abcd-abcd-abcd-abcdefghijkl",
			uuid.Nil,
			"Invalid UUID: invalid UUID format.",
		},
		{
			[]byte{43, 99, 88},
			uuid.Nil,
			"Invalid UUID: invalid UUID (got 3 bytes).",
		},
	}

	for _, c := range cases {
		id, err := ParseGlobalId(c.value)
		if c.error == "" && err != nil {
			t.Errorf("Failed to parse \"%v\" (type %T) to global id: \"%s\"", c.value, c.value, err.Error())
		} else if c.error != "" {
			if err == nil {
				t.Errorf(
					"Parsing \"%v\" (type %T) must be failed with error \"%s\", but no error was given.",
					c.value,
					c.value,
					c.error,
				)
			} else if err.Error() != c.error {
				t.Errorf(
					"Parsing \"%v\" (type %T) must be failed with error \"%s\", but \"%s\" was given.",
					c.value,
					c.value,
					c.error,
					err.Error(),
				)
			}
		} else if id.identity != c.result {
			t.Errorf(
				"Parsing \"%v\" (type %T) with result %s does not match expected value %s.",
				c.value,
				c.value,
				id.GetIdentity(),
				c.result,
			)
		}
	}
}

func TestGlobalIdFrom(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("GlobalIdFrom: parsing must panic, but it does not.")
		}
	}()

	// valid parsing
	if GlobalIdFrom("a0d0d3d7-c0e1-4880-8ea8-c7dd17b1ebba").identity !=
		uuid.MustParse("a0d0d3d7-c0e1-4880-8ea8-c7dd17b1ebba") {
		t.Error("GlobalIdFrom: invalid parsing.")
	}

	// invalid parsing
	LocalIdFrom("abc")
}

func TestNewGlobalId(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("NewGlobalId: %s", err)
		}
	}()

	NewGlobalId()
}
