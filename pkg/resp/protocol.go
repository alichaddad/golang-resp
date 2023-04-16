package resp

import (
	"bytes"
	"strconv"
)

const (
	PrefixSimpleString = '+'
	PrefixError        = '-'
	PrefixInteger      = ':'
	PrefixBulkString   = '$'
	PrefixArray        = '*'
)

type Type struct {
	descriptor rune
	intVal     int
	strVal     string
	errVal     error
	arrayVal   []Type
	isNull     bool
}

func (t Type) GetValue() []byte {
	return t.encodeValue()
}

func (t Type) encodeValue() []byte {
	switch t.descriptor {
	case PrefixSimpleString:
		return genericEncode(t, []byte(t.strVal))
	case PrefixError:
		return genericEncode(t, []byte(t.errVal.Error()))
	case PrefixInteger:
		return genericEncode(t, []byte(strconv.Itoa(t.intVal)))
	case PrefixBulkString:
		return encodeBulkString(t)
	case PrefixArray:
		return encodeArray(t)
	}
	return nil
}

func encodeBulkString(t Type) []byte {
	if t.isNull {
		return []byte("$-1\r\n")
	}
	var buffer bytes.Buffer
	prefixLen := []byte(strconv.Itoa(len(t.strVal)))
	buffer.WriteByte(byte(t.descriptor))
	buffer.Write(prefixLen)
	buffer.WriteByte('\r')
	buffer.WriteByte('\n')
	buffer.Write([]byte(t.strVal))
	buffer.WriteByte('\r')
	buffer.WriteByte('\n')
	return buffer.Bytes()
}

func encodeArray(t Type) []byte {
	if t.isNull {
		return []byte("*-1\r\n")
	}
	var buffer bytes.Buffer
	prefixLen := []byte(strconv.Itoa(len(t.arrayVal)))
	buffer.WriteByte(byte(t.descriptor))
	buffer.Write(prefixLen)
	buffer.WriteByte('\r')
	buffer.WriteByte('\n')

	for _, elem := range t.arrayVal {
		elemBytes := elem.encodeValue()
		buffer.Write(elemBytes)
	}
	return buffer.Bytes()
}

func genericEncode(t Type, value []byte) []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(t.descriptor))
	buffer.Write(value)
	buffer.WriteByte('\r')
	buffer.WriteByte('\n')
	return buffer.Bytes()
}
