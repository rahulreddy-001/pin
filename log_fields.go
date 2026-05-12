package pin

import "fmt"

type Fields []Field

type Field struct {
	Key   string
	Value any
}

func Any(key string, val any) Field {
	return Field{key, val}
}

func (fields Fields) Encode() string {
	encoded := "{"
	for i, field := range fields {
		encoded += fmt.Sprintf(`"%s":"%s"`, field.Key, field.Value)
		if len(fields)-1 != i {
			encoded += ","
		}
	}
	if len(encoded) == 1 {
		return ""
	}
	return encoded + "}"
}
