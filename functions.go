package pyplot

import (
	"fmt"
	"log"
	"strings"
)

func plotKwargStart(args []interface{}) int {
	return len(args)
}

func Plot(args ...interface{}) {
	str, ok := ConvertType("Meow = \"djhfs\"", String)
	fmt.Println(str, ok)
	log.Fatal(":3")
	stringArgs := []string{}
	kwargStart := plotKwargStart(args)
	for i := 0; i < kwargStart; i += 3 {
		panic("NYI")
	}

	for i := kwargStart; i < len(args); i += 2 {
		panic("NYI")
	}

	strings.Join(
		[]string{"plt.plot(", strings.Join(stringArgs, ","), ")"}, "",
	)
}
