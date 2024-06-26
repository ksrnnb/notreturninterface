package a

import "context"

type Hoge interface{}

type Foo struct{}

func f() Hoge { // want "function f must not return interface a.Hoge, but struct"
	return nil
}

func g() Foo {
	return Foo{}
}

//lint:ignore notreturninterface OK
func h() Hoge {
	return nil
}

func x() error {
	return nil
}

func y() {
	return
}

func z() any {
	return nil
}

func returnContext() context.Context {
	return context.Background()
}
