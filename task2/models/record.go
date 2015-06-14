// task2 project record.go
// Provides types and vars for handling data records

package models

type RecordKeyType string
type RecordValueType string

type Record struct {
	Key   RecordKeyType   `json:"key"`
	Value RecordValueType `json:"value"`
}

type RecordsMapType map[RecordKeyType]RecordValueType

var RecordsMap RecordsMapType

var RecordsFilePath = "records.gob"
