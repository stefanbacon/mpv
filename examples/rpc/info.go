package main

import (
	"math"
	"os/exec"
	"strings"
	"fmt"
	"github.com/ArsenyZorin/mpv"
	"mpvtime"
)



func main() {
	path := "https://www.youtube.com/watch?v="
	ipcc := mpv.NewIPCClient("/tmp/mpvsocket")
	c := mpv.NewClient(ipcc)
	duration,_ := c.Duration()
	perc,_ := c.Position()
	file,_ := c.Filename()
	name := file[1 : len(file) - 1]

	cmd := exec.Command("youtube-dl", path + name, "--get-title")
	out, _ := cmd.CombinedOutput()
	cmd.Run()

	str := strings.Replace(string(out), "\n", "", -1)
	passed := GetTime(perc)
	common := GetTime(duration)

	printed_filename := fmt.Sprintf("%s: ", str)
	printed_time := PrintTime(passed) + "/" + PrintTime(common)

	cmd = exec.Command("dunstify", printed_filename + printed_time, "-r", "3040", "-u", "urgency")
	cmd.Run()
}
