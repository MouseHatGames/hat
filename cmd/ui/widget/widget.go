package widget

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"gopkg.in/yaml.v3"
)

var (
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
	Title       string     `json:"title" yaml:"title"`
	Type        WidgetType `json:"type" yaml:"type"`
	Description string     `json:"description,omitempty" yaml:"description"`
	Colspan     int        `json:"colspan" yaml:"colspan"`

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
		return fmt.Errorf("invalid widget '%s': %w", w.Title, err)
	}

	return nil
}

func (w *Widget) isValid() (reason error) {
	if w.Title == "" {
		return ErrMissingTitle
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

func (w *Widget) UnmarshalValue(str string) (value interface{}, err error) {
	b := []byte(str)

	switch w.Type {
	case WidgetOnOff:
		var on bool
		err = json.Unmarshal(b, &on)
		value = on

	case WidgetText:
		var str string
		err = json.Unmarshal(b, &str)
		value = str

	case WidgetOptions:
		if w.StoreIndex {
			var idx int
			err = json.Unmarshal(b, &idx)
			value = idx
		} else {
			var sel string
			err = json.Unmarshal(b, &sel)
			value = sel
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
		if w.StoreIndex {
			if v, ok := val.(float64); ok {
				return json.Marshal(math.Floor(v))
			}
		} else {
			if v, ok := val.(string); ok {
				return json.Marshal(v)
			}
		}

	default:
		return nil, ErrInvalidType
	}

	return nil, ErrInvalidValueType
}
