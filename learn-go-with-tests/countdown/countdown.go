package countdown

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
