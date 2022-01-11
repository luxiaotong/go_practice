package pca

import "strings"

type TableType int32

const (
	TableType_UNKNOWN TableType = 0 // don't use
	TableType_TABLE   TableType = 1
	TableType_VIEW    TableType = 2
)

type TableFieldsInfo struct {
	Fields     []*FieldInfo `json:"fields,omitempty"`
	RecordSize uint32       `json:"record_size,omitempty"`
}

type FieldInfo struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	IsPrimary bool   `json:"is_primary,omitempty"`
}

type TableInfo struct {
	Type       TableType `json:"type,omitempty"`
	Name       string    `json:"name,omitempty"`
	FieldCount uint32    `json:"field_count,omitempty"`
}

type Schema string

func (sc Schema) Field() string {
	s := string(sc)
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}
