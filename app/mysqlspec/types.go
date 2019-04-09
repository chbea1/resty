package mysqlspec

import (
	"bytes"
	"encoding/json"
	"strings"
)

type Type int

const (
	ERR Type = iota
	INT
	STRING
	TIME
)

func (t Type) String() string {
	return type2String[t]
}

func GetType(sqlType string) Type {
	if strings.Contains(sqlType, "int") {
		return INT
	} else if strings.Contains(sqlType, "char") {
		return STRING
	} else if strings.Contains(sqlType, "timestamp") {
		return TIME
	} else {
		return ERR
	}
}

var type2String = map[Type]string{ERR: "ERR", INT: "INT", STRING: "STRING", TIME: "TIME"}

func (k Type) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(type2String[k])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *Type) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = GetType(j)
	return nil
}
