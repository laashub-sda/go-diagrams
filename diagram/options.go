package diagram

import (
	"strconv"

	"github.com/blushft/go-diagrams/font"
)

type Options struct {
	Name       string
	FileName   string
	OutFormat  string
	Direction  string
	CurveStyle string
	Show       bool
	Label      string
	Pad        float64
	Splines    string
	NodeSep    float64
	RankSep    float64
	Font       font.Options
	Attributes map[string]string
}

func (o Options) attrs() map[string]string {
	m := map[string]string{
		"pad":       strconv.FormatFloat(o.Pad, 'f', -1, 64),
		"label":     o.Label,
		"splines":   o.Splines,
		"nodesep":   strconv.FormatFloat(o.NodeSep, 'f', -1, 64),
		"rankdir":   o.Direction,
		"ranksep":   strconv.FormatFloat(o.RankSep, 'f', -1, 64),
		"fontname":  o.Font.Name,
		"fontsize":  strconv.FormatInt(int64(o.Font.Size), 10),
		"fontcolor": o.Font.Color,
	}

	for k, v := range o.Attributes {
		m[k] = v
	}

	return m
}

type Option func(*Options)

func DefaultOptions(opts ...Option) Options {
	options := Options{
		Name:       "go-diagrams",
		FileName:   "go-diagram",
		OutFormat:  "dot",
		Label:      "go-diagrams",
		Direction:  string(LeftToRight),
		CurveStyle: "ortho",
		Show:       true,
		Pad:        2.0,
		Splines:    "ortho",
		NodeSep:    0.60,
		RankSep:    0.75,
		Font:       font.DefaultOptions(),
		Attributes: make(map[string]string),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func groupOptions(name string, idx int, options Options, opts ...Option) Options {
	nopts := options
	opts = append(opts,
		PenColor("#AEB6BE"),
		Shape("box"),
		Style("rounded"),
		LabelJustify("l"),
		bgFromIndex(idx),
	)

	nopts.Font.Size = 12
	nopts.Name = name

	for _, o := range opts {
		o(&nopts)
	}

	return nopts
}

func Filename(f string) Option {
	return func(o *Options) {
		o.FileName = f
	}
}

func Label(l string) Option {
	return func(o *Options) {
		o.Label = l
	}
}

func Direction(d string) Option {
	return func(o *Options) {
		o.Direction = d
	}
}

func WithAttribute(name, value string) Option {
	return func(o *Options) {
		o.Attributes[name] = value
	}
}

func WithAttributes(attrs map[string]string) Option {
	return func(o *Options) {
		for k, v := range attrs {
			o.Attributes[k] = v
		}
	}
}

func PenColor(c string) Option {
	return func(o *Options) {
		o.Attributes["pencolor"] = c
	}
}

func Shape(s string) Option {
	return func(o *Options) {
		o.Attributes["shape"] = s
	}
}

func Style(s string) Option {
	return func(o *Options) {
		o.Attributes["style"] = s
	}
}

func LabelJustify(j string) Option {
	return func(o *Options) {
		o.Attributes["labeljust"] = j
	}
}

func BackgroundColor(c string) Option {
	return func(o *Options) {
		o.Attributes["bgcolor"] = c
	}
}

func bgFromIndex(idx int) Option {
	bgcs := []string{"#E5F5FD", "#EBF3E7", "#ECE8F6", "#FDF7E3"}
	if idx-1 > len(bgcs) {
		idx = 0
	}

	return BackgroundColor(bgcs[idx])

}
