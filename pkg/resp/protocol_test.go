package resp

import (
	"bytes"
	"errors"
	"testing"
)

func TestSimpleString(t *testing.T) {
	s := Type{
		descriptor: '+',
		strVal:     "OK",
	}

	exp := "+OK\r\n"
	val := s.GetValue()
	if len(val) != len(exp) {
		t.Fatalf("length not equal, %d", len(val))
	}
	if !bytes.Equal(val, []byte(exp)) {
		t.Fatalf("%s %s", val, exp)
	}
}

func TestError(t *testing.T) {
	s := Type{
		descriptor: '-',
		errVal:     errors.New("Err unknown command"),
	}

	exp := "-Err unknown command\r\n"
	val := s.GetValue()
	if len(val) != len(exp) {
		t.Fatalf("length not equal, %d", len(val))
	}
	if !bytes.Equal(val, []byte(exp)) {
		t.Fatalf("%s %s", val, exp)
	}
}

func TestInteger(t *testing.T) {
	s := Type{
		descriptor: ':',
		intVal:     0,
	}

	exp := ":0\r\n"
	val := s.GetValue()
	if len(val) != len(exp) {
		t.Fatalf("length not equal, %d", len(val))
	}
	if !bytes.Equal(val, []byte(exp)) {
		t.Fatalf("%s %s", val, exp)
	}
}

func TestBulkString(t *testing.T) {
	s := Type{
		descriptor: '$',
		strVal:     "hello",
	}

	exp := "$5\r\nhello\r\n"
	val := s.GetValue()
	if len(val) != len(exp) {
		t.Fatalf("length not equal, %d", len(val))
	}
	if !bytes.Equal(val, []byte(exp)) {
		t.Fatalf("%s %s", val, exp)
	}
}

func TestArray(t *testing.T) {
	s := Type{
		descriptor: '*',
		arrayVal: []Type{
			{
				descriptor: '$',
				strVal:     "hello",
			},
			{
				descriptor: ':',
				intVal:     2,
			},
		},
	}

	exp := "*2\r\n$5\r\nhello\r\n:2\r\n"
	val := s.GetValue()
	if len(val) != len(exp) {
		t.Fatalf("length not equal, %d", len(val))
	}
	if !bytes.Equal(val, []byte(exp)) {
		t.Fatalf("%s %s", val, exp)
	}
}

func TestNestedArray(t *testing.T) {
	s := Type{
		descriptor: '*',
		arrayVal: []Type{
			{
				descriptor: '*',
				arrayVal: []Type{
					{
						descriptor: ':',
						intVal:     1,
					},
					{
						descriptor: ':',
						intVal:     2,
					},
					{
						descriptor: ':',
						intVal:     3,
					},
				},
			},
			{
				descriptor: '*',
				arrayVal: []Type{
					{
						descriptor: '+',
						strVal:     "Hello",
					},
					{
						descriptor: '-',
						errVal:     errors.New("World"),
					},
				},
			},
		},
	}

	exp := "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n"
	val := s.GetValue()
	if len(val) != len(exp) {
		t.Fatalf("length not equal, %d", len(val))
	}
	if !bytes.Equal(val, []byte(exp)) {
		t.Fatalf("%s %s", val, exp)
	}
}
