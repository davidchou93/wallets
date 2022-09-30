package types

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/lib/pq"
)

var nullString = []byte("null")

// NullInt64 is a wrapper of sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// Int64 returns a NullInt64 instance
func Int64(n int64) NullInt64 {
	return NullInt64{
		sql.NullInt64{
			Int64: n,
			Valid: true,
		},
	}
}

// MarshalJSON serializes NullInt64 to JSON.
func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return nullString, nil
}

// UnmarshalJSON deserializes NullInt64 from JSON
func (v *NullInt64) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return v.Scan(nil)
	}
	var n int64
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	return v.Scan(n)
}

// UnmarshalParam deserializes NullInt64 from form,query for echo.Context.Bind
func (v *NullInt64) UnmarshalParam(src string) error {
	// scan for null
	if src == "null" {
		return v.Scan(nil)
	}
	n, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return err
	}
	return v.Scan(n)
}

// NullFloat64 is a wrapper of sql.NullFloat64
type NullFloat64 struct {
	sql.NullFloat64
}

// Float64 returns a NullFloat64 instance
func Float64(f float64) NullFloat64 {
	return NullFloat64{
		sql.NullFloat64{
			Float64: f,
			Valid:   true,
		},
	}
}

// MarshalJSON serializes NullFloat64 to JSON
func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	}
	return nullString, nil
}

// UnmarshalJSON deserializes NullFloat64 from JSON
func (v *NullFloat64) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return v.Scan(nil)
	}
	var d float64
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	return v.Scan(d)
}

// UnmarshalParam deserializes NullFloat64 from form,query for echo.Context.Bind
func (v *NullFloat64) UnmarshalParam(src string) error {
	// scan for null
	if src == "null" {
		return v.Scan(nil)
	}
	d, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return err
	}
	return v.Scan(d)
}

// NullString is a wrapper of sql.NullString
type NullString struct {
	sql.NullString
}

// String returns a NullString instance
func String(s string) NullString {
	return NullString{
		sql.NullString{
			String: s,
			Valid:  true,
		},
	}
}

// MarshalJSON serializes NullString to JSON
func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements json.Unmarshaler for NullString
func (v *NullString) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return v.Scan(nil)
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return v.Scan(s)
}

// UnmarshalParam deserializes NullString from form,query for echo.Context.Bind
func (v *NullString) UnmarshalParam(src string) error {
	// scan for null
	if src == "null" {
		return v.Scan(nil)
	}
	return v.Scan(src)
}

// NullTime is a wrapper of sql.NullTime
type NullTime struct {
	sql.NullTime
}

// Time returns a NullTime instance
func Time(t time.Time) NullTime {
	return NullTime{
		sql.NullTime{
			Time:  t,
			Valid: true,
		},
	}
}

// MarshalJSON implements json.Mashaler
func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON correctly deserializes a NullTime from JSON.
func (v *NullTime) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return v.Scan(nil)
	}
	// scan for JSON timestamp
	var t time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	return v.Scan(t)
}

// UnmarshalParam deserializes NullFloat64 from form,query for echo.Context.Bind
func (v *NullTime) UnmarshalParam(src string) error {
	// scan for null
	if src == "null" {
		return v.Scan(nil)
	}
	t, err := time.Parse(time.RFC3339, src)
	if err != nil {
		return err
	}
	return v.Scan(t)
}

// NullBool is a wrapper of sql.NullBool
type NullBool struct {
	sql.NullBool
}

// Bool returns a NullBool instance
func Bool(t bool) NullBool {
	return NullBool{
		sql.NullBool{
			Bool:  t,
			Valid: true,
		},
	}
}

// MarshalJSON correctly serializes a NullBool to JSON.
func (v NullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	}
	return nullString, nil
}

// UnmarshalJSON correctly deserializes a NullBool from JSON.
func (v *NullBool) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return v.Scan(s)
}

// UnmarshalParam deserializes NullFloat64 from form,query for echo.Context.Bind
func (v *NullBool) UnmarshalParam(src string) error {
	b, err := strconv.ParseBool(src)
	if err != nil {
		return err
	}
	return v.Scan(b)
}

// JSONB ...
type JSONB struct {
	json.RawMessage
}

// JSONBOf returns a new instance of JSONB
func JSONBOf(jt string) JSONB {
	return JSONB{
		RawMessage: []byte(jt),
	}
}

// JSONBFrom returns a new JSONB instance from src which implements MarshalJSON
func JSONBFrom(src interface{}) (JSONB, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return JSONB{}, err
	}
	return JSONB{RawMessage: b}, nil
}

// Scan ...
func (j *JSONB) Scan(v interface{}) error {
	switch v := v.(type) {
	case string:
		j.RawMessage = make(json.RawMessage, len([]byte(v)))
		copy(j.RawMessage, []byte(v))
	case []byte:
		j.RawMessage = make(json.RawMessage, len(v))
		copy(j.RawMessage, v)
	case json.RawMessage:
		j.RawMessage = make(json.RawMessage, len(v))
		copy(j.RawMessage, v)
	case nil:
		j.RawMessage = nil
	default:
		return fmt.Errorf("Incompatible type %T for JSONB", v)
	}
	return nil
}

// Value ...
func (j JSONB) Value() (driver.Value, error) {
	if j.RawMessage == nil {
		return nil, nil
	}
	return string(j.RawMessage), nil
}

// UnmarshalParam deserializes NullJSONText from form,query for echo.Context.Bind
func (j *JSONB) UnmarshalParam(src string) error {
	return j.Scan(src)
}

// StringArray ...
type StringArray pq.StringArray

// Scan
func (sa *StringArray) Scan(v interface{}) error {
	return (*pq.StringArray)(sa).Scan(v)
}

// Value
func (sa StringArray) Value() (driver.Value, error) {
	return (pq.StringArray)(sa).Value()
}

// StringSlice
func (sa StringArray) StringSlice() []string {
	return ([]string)(sa)
}

// Int64Array ...
type Int64Array pq.Int64Array

// Scan
func (ia *Int64Array) Scan(v interface{}) error {
	return (*pq.Int64Array)(ia).Scan(v)
}

// Value
func (ia Int64Array) Value() (driver.Value, error) {
	return (pq.Int64Array)(ia).Value()
}

// Int64Slice
func (ia Int64Array) Int64Slice() []int64 {
	return ([]int64)(ia)
}
