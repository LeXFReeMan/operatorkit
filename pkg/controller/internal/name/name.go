package name

import (
	"fmt"
	"strings"

	"github.com/LeXFReeMan/operatorkit/v7/pkg/handler"
)

func Name(r handler.Interface) string {
	split := strings.Split(fmt.Sprintf("%T", r), ".")

	if len(split) < 2 {
		panic("unable to parse handler name")
	}

	return strings.Replace(split[0], "*", "", 1)
}
