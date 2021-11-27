package methods

import (
	"fmt"
	"go-study/ch12/params"
	"strings"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {

	Print(time.Hour)

	fmt.Println()

	Print(new(strings.Replacer))

	fmt.Println()
	Print(params.Test)
}
