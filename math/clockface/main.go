package main

import (
	"os"
	"time"

	"github.com/flpStrri/go-with-tests/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
