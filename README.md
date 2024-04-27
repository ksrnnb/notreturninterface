# notreturninterface

`notreturninterface` is a static code analyzer which detects function returns interface.

## Install

You can install `notreturninterface` by `go install` command.

```bash
$ go install github.com/ksrnnb/notreturninterface/cmd/notreturninterface@latest
```

## How to use

You can run `notreturninterface` by `go vet` command.

```bash
$ go vet -vettool=$(which notreturninterface) ./...
```

## Example

```go
type Hoge interface{}

type Foo struct{}

func f() Hoge { // NG: return value is interface
	return nil
}

func g() Foo { // OK: return value is struct
	return Foo{}
}

func x() error { // OK: error is allowed
	return nil
}
```

You can ignore `notreturninterface` check by comment like`//lint:ignore notreturninterface reason`. Note that `reason` is required.

```go
//lint:ignore notreturninterface hoge is special
func h() Hoge { // OK: you can ignore
	return nil
}
```