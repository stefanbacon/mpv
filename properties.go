package mpv

import (
	"errors"
	"fmt"
)

func (c *Client) ObserveFloat(property string, cb func(value float64)) {
	c.observeFloatCB[property] = cb
	_, _ = c.Exec("observe_property", 0, property)
}

func (c *Client) ObserveString(property string, cb func(value string)) {
	c.observeStringCB[property] = cb
	_, _ = c.Exec("observe_property", 0, property)
}

func (c *Client) ObserveBool(property string, cb func(value bool)) {
	c.observeBoolCB[property] = cb
	_, _ = c.Exec("observe_property", 0, property)
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
