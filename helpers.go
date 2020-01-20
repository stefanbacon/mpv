package mpv

import (
	"errors"
	"fmt"
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

// GetProperty reads a property by name and returns the data as a string.
func (c *Client) GetProperty(name string) (string, error) {
	res, err := c.Exec("get_property", name)
	if res == nil {
		return "", err
	}
	return fmt.Sprintf("%#v", res.Data), err
}

// SetProperty sets the value of a property.
func (c *Client) SetProperty(name string, value interface{}) error {
	_, err := c.Exec("set_property", name, value)
	return err
}

// GetFloatProperty reads a float property and returns the data as a float64.
func (c *Client) GetFloatProperty(name string) (float64, error) {
	res, err := c.Exec("get_property", name)
	if res == nil {
		return 0, err
	}
	if val, found := res.Data.(float64); found {
		return val, err
	}
	return 0, invalidType(name)
}

// GetBoolProperty reads a bool property and returns the data as a boolean.
func (c *Client) GetBoolProperty(name string) (bool, error) {
	res, err := c.Exec("get_property", name)
	if res == nil {
		return false, err
	}
	if val, found := res.Data.(bool); found {
		return val, err
	}
	return false, invalidType(name)
}

func invalidType(property string) error {
	return errors.New(fmt.Sprintf("invalid type for property %q", property))
}
