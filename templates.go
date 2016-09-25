package circuitbreaker

import "github.com/clipperhouse/typewriter"

type model struct {
	Type            typewriter.Type
	OutType         string
	MethodInputType string
}

var templates = typewriter.TemplateSlice{
	common,
	Command,
}
