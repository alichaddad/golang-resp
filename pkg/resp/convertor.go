package resp

import (
	"bufio"
	"strconv"
)

func ParseMessage(rd *bufio.Reader) []string {
	var values []string
	str, _ := rd.ReadString('\n')
	if str[0] == byte(PrefixArray) {
		size, _ := strconv.ParseInt(str[1:len(str)-2], 10, 0)
		for i := 0; i < int(size); i++ {
			str, _ := rd.ReadString('\n')
			if str[0] == byte(PrefixBulkString) {
				str, _ := rd.ReadString('\n')
				values = append(values, str[:len(str)-2])
			}
		}
	}
	return values
}

func NewBulkMessage(value string) Type {
	return Type{
		descriptor: PrefixBulkString,
		strVal:     value,
	}
}

func NewError(value error) Type {
	return Type{
		descriptor: PrefixError,
		errVal:     value,
	}
}

func NewInt(value int) Type {
	return Type{
		descriptor: PrefixInteger,
		intVal:     value,
	}
}
