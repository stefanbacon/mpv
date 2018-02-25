package main

import "github.com/blang/mpv"

func main() {
	ipcc := mpv.NewIPCClient("/tmp/mpvsocket")
	c := mpv.NewClient(ipcc)
	pause,_ = c.Pause()
	c.SetPause(!pause)
}
