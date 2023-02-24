package errors

import "fmt"

func CastFail(function, from, to string) {
	panic(fmt.Sprintf("%s ::: cast from %s to %s failed", function, from, to))
}
func UndefinedBehaviour(function string) {
	panic(fmt.Sprintf("%s ::: we should never get here", function))
}

func Unimplemented(function string) {
	panic(fmt.Sprintf("%s ::: unimplemented", function))
}

var Debug bool
