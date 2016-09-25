package circuitbreaker

import (
	"errors"
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewBreakerWriter())
	if err != nil {
		panic(err)
	}
}

// BreakerWriter implements Typewriter to be used to generate
// a command like pattern
type BreakerWriter struct{}

// NewBreakerWriter is used to create a new BreakerWriter
func NewBreakerWriter() *BreakerWriter {
	return &BreakerWriter{}
}

func (cw *BreakerWriter) Name() string {
	return "breaker"
}

func (cw *BreakerWriter) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return []typewriter.ImportSpec{
		typewriter.ImportSpec{
			Name: "metrics",
			Path: "github.com/rcrowley/go-metrics",
		},
	}
}

func (cw *BreakerWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(cw)
	if !found {
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	m := model{
		Type: t,
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	for _, v := range tag.Values {

		if len(v.TypeParameters) != 2 {
			return errors.New("Only command is supported. Need all inputs")
		}

		outType := v.TypeParameters[0]
		methodInputType := v.TypeParameters[1]

		m := model{
			Type:            t,
			OutType:         outType.Name,
			MethodInputType: methodInputType.Name,
		}

		tmpl, err := templates.ByTagValue(t, v)

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	return nil
}
