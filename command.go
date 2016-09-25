package circuitbreaker

import "github.com/clipperhouse/typewriter"

var Command = &typewriter.Template{
	Name: "Command",
	Text: `
  func (cmd {{.Type}}Command) Execute(input {{.MethodInputType}}) (*{{.OutType}}, error) {

  	c := make(chan {{.Type}}CommandResponse)

  	go func() {
  		t := time.Now()
  		defer cmd.metricExecTime.UpdateSince(t)

  		response, err := cmd.inner.Execute(input)
  		c <- {{.Type}}CommandResponse{response, err}
  	}()

  	select {
  	case <-time.After(cmd.timeout):
  		cmd.metricTimeouts.Mark(1)

  		return nil, cmd.ErrTimeout
  	case resp := <-c:
  		response := resp.response
  		err := resp.err

  		if err != nil {
  			cmd.metricErrors.Mark(1)

  			return nil, err
  		}

  		return response, nil
  	}

  }

  type {{.Type}}CommandResponse struct {
  	response *{{.OutType}}
  	err      error
  }`,
	TypeParameterConstraints: []typewriter.Constraint{
		{Comparable: true},
		{Comparable: true},
	},
}
