package mysqlspec

import (
	"bytes"
	"encoding/json"
)

type NULL int

const (
	NO  NULL = 0
	YES NULL = 1
)

func (n NULL) String() string {
	return null2String[n]
}

var null2String = map[NULL]string{YES: "YES", NO: "NO"}
var null2ID = map[string]NULL{"YES": YES, "NO": NO}

func (k NULL) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(null2String[k])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *NULL) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = null2ID[j]
	return nil
}

func GetNull(value string) NULL {
	return null2ID[value]
}
