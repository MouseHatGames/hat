package widget

import (
	"encoding/json"
	"errors"
	"math"

	"gopkg.in/yaml.v3"
)

var (
	ErrMissingType      = errors.New("missing type")
	ErrMissingTitle     = errors.New("missing title")
	ErrInvalidOptions   = errors.New("invalid options")
	ErrMissingChildren  = errors.New("missing group children")
	ErrMissingOptions   = errors.New("missing options")
	ErrUnknownType      = errors.New("unknown type")
	ErrInvalidType      = errors.New("invalid type")
	ErrInvalidValueType = errors.New("invalid value type")
)

type WidgetType string

const (
	WidgetOnOff   WidgetType = "onoff"
	WidgetText    WidgetType = "text"
	WidgetGroup   WidgetType = "group"
	WidgetOptions WidgetType = "options"
)

type Widget struct {
	// Used because map order is not preserved when marshalling to JSON
	Path  string `json:"path" yaml:"-"`
	Index int    `json:"-" yaml:"-"`

	Title       string     `json:"title" yaml:"title"`
	Type        WidgetType `json:"type" yaml:"type"`
	Description string     `json:"description,omitempty" yaml:"description"`
	Colspan     int        `json:"colspan" yaml:"colspan"`

	ParamNames []string `json:"params,omitempty" yaml:"-"`

	// Text widget
	Placeholder string `json:"placeholder,omitempty" yaml:"placeholder"`
	Big         bool   `json:"big,omitempty" yaml:"big"`

	// Group widget
	Children map[string]*Widget `json:"children,omitempty" yaml:"children"`

	// Options widget
	Options    []string `json:"options,omitempty" yaml:"options"`
	StoreIndex bool     `json:"-" yaml:"storeIndex"`
}

var _ yaml.Unmarshaler = (*Widget)(nil)

func (w *Widget) UnmarshalYAML(value *yaml.Node) error {
	type alias Widget
	widget := &alias{Colspan: 1}

	if err := value.Decode(widget); err != nil {
		return err
	}

	*w = Widget(*widget)

	if err := w.isValid(); err != nil {
		return err
	}

	return nil
}

func (w *Widget) guessType() WidgetType {
	if w.Placeholder != "" || w.Big != false {
		return WidgetText
	} else if w.Children != nil {
		return WidgetGroup
	} else if w.Options != nil {
		return WidgetOptions
	}

	return ""
}

func (w *Widget) isValid() (reason error) {
	if w.Title == "" {
		return ErrMissingTitle
	}

	if w.Type == "" {
		if guess := w.guessType(); guess != "" {
			w.Type = guess
		} else {
			return ErrMissingType
		}
	}

	switch w.Type {
	case WidgetOnOff:
		if w.Placeholder != "" || w.Big || w.Children != nil || w.Options != nil {
			return ErrInvalidOptions
		}

	case WidgetText:
		if w.Children != nil || w.Options != nil {
			return ErrInvalidOptions
		}

	case WidgetGroup:
		if w.Options != nil || w.Placeholder != "" || w.Big {
			return ErrInvalidOptions
		}
		if w.Children == nil {
			return ErrMissingChildren
		}

	case WidgetOptions:
		if w.Placeholder != "" || w.Big || w.Children != nil {
			return ErrInvalidOptions
		}
		if w.Options == nil {
			return ErrMissingOptions
		}

	default:
		return ErrUnknownType
	}

	return nil
}

func (w *Widget) UnmarshalValue(jsonval []byte) (value interface{}, err error) {
	switch w.Type {
	case WidgetOnOff:
		var on bool
		err = json.Unmarshal(jsonval, &on)
		value = on

	case WidgetText:
		var str string
		err = json.Unmarshal(jsonval, &str)
		value = str

	case WidgetOptions:
		if w.StoreIndex {
			var idx int
			err = json.Unmarshal(jsonval, &idx)
			value = idx
		} else {
			var sel string
			err = json.Unmarshal(jsonval, &sel)

			for i, opt := range w.Options {
				if opt == sel {
					value = i
					break
				}
			}
		}

	default:
		err = ErrInvalidType
	}

	return
}

func (w *Widget) MarshalValue(val interface{}) ([]byte, error) {
	switch w.Type {
	case WidgetOnOff:
		if v, ok := val.(bool); ok {
			return json.Marshal(v)
		}

	case WidgetText:
		if v, ok := val.(string); ok {
			return json.Marshal(v)
		}

	case WidgetOptions:
		var idx int32
		if v, ok := val.(float64); ok {
			idx = int32(math.Floor(v))

			if w.StoreIndex {
				return json.Marshal(idx)
			}

			return json.Marshal(w.Options[idx])
		}

	default:
		return nil, ErrInvalidType
	}

	return nil, ErrInvalidValueType
}
