package lib

import (
	"github.com/phpli/golang/article1/q3/lib/internal"
	"os"
)

func Hello(name string)  {
	internal.Hello(os.Stdout,name)
}