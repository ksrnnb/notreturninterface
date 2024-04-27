package a

type Hoge interface{}

type Foo struct{}

func f() Hoge { // want "function f must not return interface a.Hoge, but struct"
	return nil
}

//lint:ignore notreturninterface OK
func h() Hoge {
	return nil
}

func g() Foo {
	return Foo{}
}
