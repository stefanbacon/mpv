package main

import "github.com/ArsenyZorin/mpv"

func main() {
	ipcc := mpv.NewIPCClient("/tmp/mpvsocket")
	c := mpv.NewClient(ipcc)
	mute,_ := c.Mute()
	c.SetMute(!mute)
}
