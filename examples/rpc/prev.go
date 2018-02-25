package main

import (
	"github.com/ArsenyZorin/mpv"
	"os/exec"
	"strings"
)

func main() {
	path := "https://www.youtube.com/watch?v="
	ipcc := mpv.NewIPCClient("/tmp/mpvsocket")
	c := mpv.NewClient(ipcc)
	c.PlaylistPrevious()

	file,_ := c.Filename()
	name := file[1 : len(file) - 1]

	cmd := exec.Command("youtube-dl", path + name, "--get-title")
	out, _ := cmd.CombinedOutput()
	cmd.Run()

	str := strings.Replace(string(out), "\n", "", -1)

	cmd = exec.Command("dunstify", str, "-r", "3040", "-u", "urgency")
	cmd.Run()
}
