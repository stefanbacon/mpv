package main

import "github.com/ArsenyZorin/mpv"

func main() {
	ipcc := mpv.NewIPCClient("/tmp/mpvsocket")
	c := mpv.NewClient(ipcc)
	c.Seek(10, mpv.SeekModeRelative)
}
