package data

type TableSchema struct {
	Fields *[]RowSchema
}

type RowSchema struct {
	Field   string
	Type    string
	Null    NULL
	Key     KEY
	Default *string
	Extra   string
}

type KEY int

const (
	NONEKEY KEY = 0
	PRI     KEY = 1
	UNI     KEY = 2
	MUL     KEY = 3
)

func (k KEY) String() string {
	if k == NONEKEY {
		return ""
	} else if k == PRI {
		return "PRI"
	} else if k == UNI {
		return "UNI"
	} else if k == MUL {
		return "MUL"
	} else {
		return ""
	}
}

type NULL int

const (
	NO  NULL = 0
	YES NULL = 1
)

func (n NULL) String() string {
	if n == YES {
		return "YES"
	} else {
		return "NO"
	}
}

type TYPE int

const (
	MEDIUMINT TYPE = 0
	SMALLINT  TYPE = 1
	CHAR      TYPE = 2
	INT       TYPE = 3
	TIMESTAMP TYPE = 4
)

func GetKey(value string) KEY {
	switch value {
	case "PRI":
		return PRI
	case "UNI":
		return UNI
	case "MUL":
		return MUL
	default:
		return NONEKEY
	}
}

func GetNull(value string) NULL {
	switch value {
	case "NO":
		return NO
	case "YES":
		return YES
	default:
		return YES
	}
}
