package circuitbreaker

import "github.com/clipperhouse/typewriter"

var common = &typewriter.Template{
	Name: "breaker",
	Text: `
	func New{{.Type}}Command(name string, timeout time.Duration, inner {{.Type}}) {{.Type}}Command {

  	metricExecTime := metrics.NewTimer()
  	metricErrors := metrics.NewMeter()
  	metricTimeouts := metrics.NewMeter()

    tags := []string{
  		"cmd:" + name,
  	}

  	if env := os.Getenv("ENVIRONMENT"); env != "" {
  		tags = append(tags, "environment:"+env)
  	}

  	if app := os.Getenv("APPLICATION_NAME"); app != "" {
  		tags = append(tags, "app:"+app)
  	}

  	tagsStr := strings.Join(tags, ",")

  	metrics.GetOrRegister(fmt.Sprintf("cmd.executionTime[%v]", tagsStr), metricExecTime)
  	metrics.GetOrRegister(fmt.Sprintf("cmd.errors[%v]", tagsStr), metricErrors)
  	metrics.GetOrRegister(fmt.Sprintf("cmd.timeouts[%v]", tagsStr), metricTimeouts)

  	return {{.Type}}Command{
  		name:           name,
  		timeout:        timeout,
  		inner:          inner,
  		metricExecTime: metricExecTime,
  		metricErrors:   metricErrors,
  		metricTimeouts: metricTimeouts,
			ErrTimeout: 	  errors.New("Command {{.Type}} timed-out"),
  	}
  }

  type {{.Type}}Command struct {
  	name           string
  	timeout        time.Duration
  	inner          {{.Type}}
  	metricExecTime metrics.Timer
  	metricErrors   metrics.Meter
  	metricTimeouts metrics.Meter
		ErrTimeout		 error
  }

  `,
	// TypeConstraint: typewriter.Constraint{Comparable: true},
}
