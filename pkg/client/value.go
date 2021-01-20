package client

import (
	"encoding/json"
	"time"
)

type Value interface {
	Error() error
	Raw() string
	Scan(v interface{}) error
	String() (string, bool)
	StringOr(def string) string
	Int32() (int32, bool)
	Int32Or(def int32) int32
	Int64() (int64, bool)
	Int64Or(def int64) int64
	Bool() (bool, bool)
	BoolOr(def bool) bool
	Float64() (float64, bool)
	Float64Or(def float64) float64
	Duration() (time.Duration, bool)
	DurationOr(def time.Duration) time.Duration
}

type jsonValue struct {
	err error
	val string
}

var _ Value = (*jsonValue)(nil)

func (v *jsonValue) Error() error {
	return v.err
}

func (v *jsonValue) Raw() string {
	return v.val
}

func (v *jsonValue) Scan(val interface{}) error {
	return json.Unmarshal([]byte(v.val), val)
}

func (v *jsonValue) String() (string, bool) {
	var ret string
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return "", false
	}
	return ret, true
}

func (v *jsonValue) StringOr(def string) string {
	if val, ok := v.String(); ok {
		return val
	}
	return def
}

func (v *jsonValue) Int32() (int32, bool) {
	var ret int32
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return 0, false
	}
	return ret, true
}

func (v *jsonValue) Int32Or(def int32) int32 {
	if val, ok := v.Int32(); ok {
		return val
	}
	return def
}

func (v *jsonValue) Int64() (int64, bool) {
	var ret int64
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return 0, false
	}
	return ret, true
}

func (v *jsonValue) Int64Or(def int64) int64 {
	if val, ok := v.Int64(); ok {
		return val
	}
	return def
}

func (v *jsonValue) Bool() (bool, bool) {
	var ret bool
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return false, false
	}
	return ret, true
}

func (v *jsonValue) BoolOr(def bool) bool {
	if val, ok := v.Bool(); ok {
		return val
	}
	return def
}

func (v *jsonValue) Float64() (float64, bool) {
	var ret float64
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return 0, false
	}
	return ret, true
}

func (v *jsonValue) Float64Or(def float64) float64 {
	if val, ok := v.Float64(); ok {
		return val
	}
	return def
}

func (v *jsonValue) Duration() (time.Duration, bool) {
	if str, ok := v.String(); ok {
		dur, err := time.ParseDuration(str)
		if err != nil {
			return 0, false
		}

		return dur, true
	}

	return 0, false
}

func (v *jsonValue) DurationOr(def time.Duration) time.Duration {
	if val, ok := v.Duration(); ok {
		return val
	}
	return def
}
