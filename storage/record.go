// Provides types and vars for handling data records

package storage

import (
	"errors"
	"time"
)

type RecordKeyType string
type RecordValueType string
type TimestampType int64

type Record struct {
	Key   RecordKeyType   `json:"key"`
	Value RecordValueType `json:"value"`
}

type RecordTimedValue struct {
	Timestamp TimestampType   `json:"timestamp"`
	Value     RecordValueType `json:"value"`
}

type RecordHistory struct {
	Key    RecordKeyType      `json:"key"`
	Values []RecordTimedValue `json:"values"`
}

type recordValues struct {
	Values   map[TimestampType]RecordValueType
	Sequence []TimestampType
}

type RecordsType map[RecordKeyType]recordValues

func (this RecordsType) GetLast(key RecordKeyType) (v RecordValueType, err error) {
	r, ok := this[key]
	if !ok || len(r.Sequence) == 0 {
		err = errors.New("No record found!")
	} else {
		v = r.Values[r.Sequence[len(r.Sequence)-1]]
	}
	return
}

func (this RecordsType) GetAll(key RecordKeyType) (v []RecordTimedValue, err error) {
	r, ok := this[key]
	if !ok || len(r.Sequence) == 0 {
		err = errors.New("No record found!")
	} else {
		for _, t := range r.Sequence {
			v = append(v, RecordTimedValue{Timestamp: t, Value: r.Values[t]})
		}
	}
	return
}

func (this RecordsType) Get(key RecordKeyType, timestamp TimestampType) (v RecordValueType, err error) {
	r, ok := this[key]
	if !ok || len(r.Sequence) == 0 {
		err = errors.New("No record found!")
	} else {
		if v, ok = r.Values[timestamp]; !ok {
			err = errors.New("No timestamp found!")
		}
	}
	return
}

func (this RecordsType) CheckExistence(key RecordKeyType) bool {
	r, ok := this[key]
	return ok && len(r.Sequence) != 0
}

func (this RecordsType) Add(key RecordKeyType, value RecordValueType) {
	t := TimestampType(time.Now().Unix())
	r, ok := this[key]
	if !ok || len(r.Sequence) == 0 {
		r.Values = make(map[TimestampType]RecordValueType)
	}
	r.Values[t] = value
	r.Sequence = append(r.Sequence, t)
	this[key] = r
}

func (this RecordsType) Delete(key RecordKeyType) (err error) {
	r, ok := this[key]
	if !ok || len(r.Sequence) == 0 {
		err = errors.New("No record found!")
	} else {
		delete(this, key)
	}
	return
}

var Records RecordsType

var RecordsFilePath = "records.gob"
