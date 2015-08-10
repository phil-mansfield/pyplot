package pyplot

import (
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	lines []string = []string{ "import matplotlib.pyplot as plt" }
)

func InsertLine(s string) { lines = append(lines, s) }

func Reset() { lines = []string{ "import matplotlib.pyplot as plt" } }

func Show() {
	lines = append(lines, "plt.show()")
	f, err := ioutil.TempFile(".", "go.pyplot")
	if err != nil { panic(err.Error()) }
	defer os.Remove(f.Name())

	_, err = f.Write([]byte(strings.Join(lines, "\n")))
	
	out, err := exec.Command("python", f.Name()).CombinedOutput()
	fmt.Printf("%s", out)
}
