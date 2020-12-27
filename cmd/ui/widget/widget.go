package widget

import "errors"

type WidgetType string

const (
	WidgetOnOff   WidgetType = "onoff"
	WidgetText    WidgetType = "text"
	WidgetGroup   WidgetType = "group"
	WidgetOptions WidgetType = "options"
)

var (
	ErrMissingTitle    = errors.New("missing title")
	ErrInvalidOptions  = errors.New("invalid options")
	ErrMissingChildren = errors.New("missing group children")
	ErrMissingOptions  = errors.New("missing options")
	ErrUnknownType     = errors.New("unknown type")
)

type Widget struct {
	Title       string      `json:"title" yaml:"title"`
	Type        WidgetType  `json:"type" yaml:"type"`
	Description string      `json:"description,omitempty" yaml:"description"`
	Value       interface{} `json:"value" yaml:"-"`

	// Text widget
	Placeholder string `json:"placeholder,omitempty" yaml:"placeholder"`
	Big         bool   `json:"big,omitempty" yaml:"big"`

	// Group widget
	Children []*Widget `json:"children,omitempty" yaml:"children"`

	// Options widget
	Options []string `json:"options,omitempty" yaml:"options"`
}

func (w *Widget) IsValid() (valid bool, reason error) {
	if w.Title == "" {
		return false, ErrMissingTitle
	}

	switch w.Type {
	case WidgetOnOff:
		if w.Placeholder != "" || w.Big || w.Children != nil || w.Options != nil {
			return false, ErrInvalidOptions
		}

	case WidgetText:
		if w.Children != nil || w.Options != nil {
			return false, ErrInvalidOptions
		}

	case WidgetGroup:
		if w.Options != nil || w.Placeholder != "" || w.Big {
			return false, ErrInvalidOptions
		}
		if w.Children == nil {
			return false, ErrMissingChildren
		}

	case WidgetOptions:
		if w.Placeholder != "" || w.Big || w.Children != nil {
			return false, ErrInvalidOptions
		}
		if w.Options == nil {
			return false, ErrMissingOptions
		}

	default:
		return false, ErrUnknownType
	}

	return true, nil
}
