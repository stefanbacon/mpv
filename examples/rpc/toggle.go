package main

import "github.com/ArsenyZorin/mpv"

func main() {
	ipcc := mpv.NewIPCClient("/tmp/mpvsocket")
	c := mpv.NewClient(ipcc)
	pause,_ := c.Pause()
	c.SetPause(!pause)
}
