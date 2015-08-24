package pyplot

import (
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	lines []string = []string{
		"import matplotlib.pyplot as plt",
		"import numpy as np",
	}
)

func InsertLine(s string) { lines = append(lines, s) }

func Reset() {
	lines = []string{
		"import matplotlib.pyplot as plt",
		"import numpy as np",
	}
}

func Show() {
	lines = append(lines, "plt.show()")
	Execute()
}

func Execute() {
	f, err := ioutil.TempFile(".", "go.pyplot")
	if err != nil { panic(err.Error()) }
	defer os.Remove(f.Name())

	_, err = f.Write([]byte(strings.Join(lines, "\n")))
	if err != nil { panic(err.Error()) }

	out, err := exec.Command("python", f.Name()).CombinedOutput()
	fmt.Printf("%s", out)
}
