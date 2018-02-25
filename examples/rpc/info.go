package main

import (
	"math"
	"os/exec"
	"strings"
	"fmt"
	"github.com/ArsenyZorin/mpv"
)

type Time struct {
	hh float64
	mm float64
	ss float64
}

func GetTime(seconds float64) Time {
	time := Time{0, 0, seconds}
	if time.ss >= 60 {
		time.ss = math.Mod(time.ss, 60)
		time.mm = (seconds - time.ss) / 60
		if time.mm >= 60 {
			minutes := time.mm
			time.mm = math.Mod(time.mm, 60)
			time.hh = (minutes - time.mm) / 60
		//	time.hh = math.Mod(time.mm, 60) minutes
//			time.mm = time.mm - time.hh * 60
		}
	}
	return time
}

func PrintTime (time Time) {
	fmt.Printf("%02.0f:%02.0f:%02.0f", time.hh, time.mm, time.ss)
}

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
	fmt.Printf("%s: ", str)
	PrintTime(passed)
	fmt.Printf("/")
	PrintTime(common)
}
