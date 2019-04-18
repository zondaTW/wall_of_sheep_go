package baseLib

import (
	"fmt"
)

func ErrorIf(err error, output string) {
	if err != nil {
		fmt.Println(output)
	}
}
