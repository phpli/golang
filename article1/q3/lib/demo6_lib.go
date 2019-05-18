package lib

import (
	"os"
	in "phpli/article1/q3/lib/internal"
)

func Hello(name string)  {
	in.Hello(os.Stdout,name)
}