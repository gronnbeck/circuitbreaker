# circuitbreaker

circuitbreaker is a tool for protecting your app against third-party services.
It's inspired by Netlfix' Hystrix.

It's not a library that you can pull using ``go get`` but a tool for
generating circuitbreakers for third-party method calls. It uses ``gen``
for generating these wrappers.

## Installation

```sh
go get github.com/clipperhouse/gen
go get github.com/gronnbeck/circuitbreaker
```

## Usage
First, you need to setup ``gen`` in the folder you want to generate the code.
Use ``gen add github.com/gronnbeck/circuitbreaker`` to setup ``gen``.

``circuitbreaker`` uses a ``go`` structure to generate the circuitbreaker
wrappers.

```go
// +gen breaker:"Command[Response, Input]"
type Structure struct {}
type Response struct {}
type Input struct {}

func (s Structure) Execute(input InputType) (*Response, error)
```

The example above will generate a ``StructureCommand`` with the
exactly same signature as ``Structure``s ``Execute`` function.

Notice that these structures expects to implement a function ``Execute``,
a response type (e.g. ``Response``) of any kind, and a input type expected to
Execute.

## Dependency needed in your project

You will need to add rcrowley's [go-metrics](https://github.com/rcrowley/go-metrics)
library for the generated code to work in your application.
I'll maybe look into loosely coupled the metrics functionality from
``circuitbreaker``. But that is future me's problem.

## Why roll my own circuit breaker tooling
There is a lot of good alternatives for implementing circuit breakers in Go.
Basically, you can do it by using ``go channels`` and ``go routines``.
Or you could use afex' [hystrix](https://github.com/afex/hystrix-go) library
which is quite popular.

I decided to build this because I wanted it to be integrated with datadog.
And I didn't like the async-like method signatures of afex' library.
The latter part can be to wrap that library and generate generics like this
tool does.
