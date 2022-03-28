package ddd

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

type AbstractId interface {
	Scalarable
	GetIdentity() any
}

// --------------------------------------------------------------------------------------------------------------------
// LOCAL ID
// --------------------------------------------------------------------------------------------------------------------

type LocalId struct {
	AbstractId
	identity int
}

func (id *LocalId) GetIdentity() any {
	return id.identity
}

func (id *LocalId) ToString() string {
	return strconv.Itoa(id.identity)
}

func (id *LocalId) ToScalar() any {
	return id.identity
}

func ParseLocalId(value any) (*LocalId, error) {
	var err error
	var identity int
	switch value.(type) {
	case int:
		identity = value.(int)
	case uint:
		identity = int(value.(uint))
	case int8:
		identity = int(value.(int8))
	case uint8:
		identity = int(value.(uint8))
	case int16:
		identity = int(value.(int16))
	case uint16:
		identity = int(value.(uint16))
	case int32:
		identity = int(value.(int32))
	case uint32:
		identity = int(value.(uint32))
		goto checkCastIssue
	case int64:
		identity = int(value.(int64))
		goto checkCastIssue
	case uint64:
		identity = int(value.(uint64))
		goto checkCastIssue
	case string:
		identity, err = strconv.Atoi(value.(string))
	default:
		err = errors.New("incorrect format")
	}
formLocalId:
	if err != nil {
		return nil, NewInvalidArgumentError("Invalid integer: %s.", err)
	}
	return &LocalId{identity: identity}, nil

checkCastIssue:
	if fmt.Sprintf("%d", value) != strconv.Itoa(identity) {
		err = errors.New(fmt.Sprintf("cannot convert %d to int", value))
	}
	goto formLocalId
}

func LocalIdFrom(value any) *LocalId {
	id, err := ParseLocalId(value)
	if err != nil {
		panic(err)
	}
	return id
}

// --------------------------------------------------------------------------------------------------------------------
// GLOBAL ID
// --------------------------------------------------------------------------------------------------------------------

type GlobalId struct {
	AbstractId
	identity uuid.UUID
}

func (id *GlobalId) GetIdentity() any {
	return id.identity
}

func (id *GlobalId) ToString() string {
	return id.identity.String()
}

func (id *GlobalId) ToScalar() any {
	return id.identity.String()
}

func NewGlobalId() *GlobalId {
	return &GlobalId{identity: uuid.New()}
}

func ParseGlobalId(value any) (*GlobalId, error) {
	var err error
	var identity uuid.UUID
	switch value.(type) {
	case string:
		identity, err = uuid.Parse(value.(string))
	case []byte:
		identity, err = uuid.FromBytes(value.([]byte))
	default:
		err = errors.New("incorrect format")
	}
	if err != nil {
		return nil, NewInvalidArgumentError("Invalid UUID: %s.", err.Error())
	}
	return &GlobalId{identity: identity}, nil
}

func GlobalIdFrom(value any) *GlobalId {
	id, err := ParseGlobalId(value)
	if err != nil {
		panic(err)
	}
	return id
}
