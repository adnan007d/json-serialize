package onserialize

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// This is to avoid some string allocation inside function
// I have no idea if this impacts anything
const (
	QUOTE         = "\""
	ESCAPED_QUOTE = "\\\""
	TRUE          = "true"
	FALSE         = "false"
	NULL          = "null"
)

func serializeBool(sb *strings.Builder, s bool) {
	if s {
		sb.WriteString(TRUE)
	} else {
		sb.WriteString(FALSE)
	}
}

func serializeString(sb *strings.Builder, s string) {
	sb.WriteRune('"')
	// Probably not good?
	s = strings.ReplaceAll(s, QUOTE, ESCAPED_QUOTE)
	sb.WriteString(s)
	sb.WriteRune('"')
}
func serializeInts(sb *strings.Builder, n int64) {
	sb.WriteString(strconv.FormatInt(n, 10))
}

func serializeUInts(sb *strings.Builder, un uint64) {
	sb.WriteString(strconv.FormatUint(un, 10))
}

func serializeFloats(sb *strings.Builder, f float64) {
	sb.WriteString(strconv.FormatFloat(f, 'g', -1, 64))
}

func serializeSlices(sb *strings.Builder, s reflect.Value) {
	sb.WriteRune('[')
	for i := 0; i < s.Len(); i++ {
		serialize(sb, s.Index(i).Interface())
		if i != s.Len()-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune(']')
}

func serializeMap(sb *strings.Builder, s reflect.Value) {
	sb.WriteRune('{')
	keys := s.MapKeys()
	for i, key := range keys {
		if key.Kind() != reflect.String {
			panic(fmt.Sprintf("Key should be string not %v", key.Kind()))
		}
		serializeString(sb, key.String())
		sb.WriteRune(':')
		serialize(sb, s.MapIndex(key).Interface())

		if i != len(keys)-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune('}')
}

func serialize(sb *strings.Builder, s interface{}) {
	if s == nil {
		sb.WriteString(NULL)
		return
	}

	value := reflect.ValueOf(s)

	switch value.Kind() {
	case reflect.Bool:
		serializeBool(sb, value.Bool())
	case reflect.String:
		serializeString(sb, value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		serializeInts(sb, value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		serializeUInts(sb, value.Uint())
	case reflect.Float32, reflect.Float64:
		serializeFloats(sb, value.Float())
	case reflect.Slice, reflect.Array:
		serializeSlices(sb, value)
	case reflect.Map:
		serializeMap(sb, value)
	default:
		panic(fmt.Sprintf("Unknown type %v:%v", value.Kind(), value))
	}

}

func Serialize(s interface{}) string {
	var sb strings.Builder
	serialize(&sb, s)
	return sb.String()
}
