package mpv

import (
	"strconv"
)

func (c *Client) LoadFile(path string) error {
	_, err := c.Exec("loadfile", path)
	return err
}

const (
	SeekModeRelative = "relative"
	SeekModeAbsolute = "absolute"
)

func (c *Client) Seek(n int, mode string) error {
	_, err := c.Exec("seek", strconv.Itoa(n), mode)
	return err
}

func (c *Client) Pause() (bool, error) {
	return c.GetBoolProperty("pause")
}

func (c *Client) SetPause(pause bool) error {
	return c.SetProperty("pause", pause)
}

func (c *Client) Mute() (bool, error) {
	return c.GetBoolProperty("mute")
}

// SetMute mutes or unmutes the player.
func (c *Client) SetMute(mute bool) error {
	return c.SetProperty("mute", mute)
}

