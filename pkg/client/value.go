package hat

import "encoding/json"

type Value interface {
	Error() error
	Raw() string
	String() string
	StringOr(def string) string
	Int32() int32
	Int32Or(def int32) int32
	Bool() bool
	BoolOr(def bool) bool
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

func (v *jsonValue) String() string {
	return v.StringOr("")
}

func (v *jsonValue) StringOr(def string) string {
	var ret string
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return def
	}
	return ret
}

func (v *jsonValue) Int32() int32 {
	return v.Int32Or(0)
}

func (v *jsonValue) Int32Or(def int32) int32 {
	var ret int32
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return def
	}
	return ret
}

func (v *jsonValue) Bool() bool {
	return v.BoolOr(false)
}

func (v *jsonValue) BoolOr(def bool) bool {
	var ret bool
	if err := json.Unmarshal([]byte(v.val), &ret); err != nil {
		return def
	}
	return ret
}
