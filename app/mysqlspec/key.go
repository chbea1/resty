package mysqlspec

import (
	"bytes"
	"encoding/json"
)

type KEY int

const (
	NONE KEY = iota
	PRI
	UNI
	MUL
)

func (k KEY) String() string {
	return key2String[k]
}

var key2String = map[KEY]string{
	NONE: "",
	PRI:  "PRI",
	UNI:  "UNI",
	MUL:  "MUL",
}

var key2ID = map[string]KEY{
	"":    NONE,
	"PRI": PRI,
	"UNI": UNI,
	"MUL": MUL,
}

func (k KEY) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(key2String[k])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *KEY) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = key2ID[j]
	return nil
}

func GetKey(value string) KEY {
	return key2ID[value]
}
