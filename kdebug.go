package kame

import "fmt"

type kdebug bool

func (d kdebug) pf(format string, a ...interface{}) {
	if d {
		if len(a) > 0 {
			fmt.Printf(format, a...)
		} else {
			fmt.Printf(format)
		}
	}
}

func (d kdebug) pln(msg string) {
	if d {
		fmt.Println(msg)
	}
}
