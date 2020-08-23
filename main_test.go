package main

import (
	"os/exec"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTappingEnabled(t *testing.T) {
	assert := assert.New(t)
	setDeviceProp("11", "318", "1")
	expected := "libinput Natural Scrolling Enabled (318):\t1"
	bytes, err := exec.Command("xinput", "list-props", "11").Output()

	if err != nil {
		assert.Fail(err.Error())
	}

	actual := string(bytes)
	assert.Contains(actual, expected)

}

func TestGetPropId(t *testing.T) {
	assert := assert.New(t)
	data := "libinput Natural Scrolling Enabled (318):\t1"
	expected := "318"
	actual := getDevicePropID(data)
	assert.Equal(expected, actual)
}

func TestGetPropIdWithoutOpenBrace(t *testing.T) {
	assert := assert.New(t)
	data := "libinput Natural Scrolling Enabled 318):\t1"
	expected := ""
	actual := getDevicePropID(data)
	assert.Equal(expected, actual)
}

func TestGetPropIdWithoutCloseBrace(t *testing.T) {
	assert := assert.New(t)
	data := "libinput Natural Scrolling Enabled (318:\t1"
	expected := ""
	actual := getDevicePropID(data)
	assert.Equal(expected, actual)
}

func TestGetDeviceId(t *testing.T) {
	assert := assert.New(t)
	data := "⎜   ↳ Elan Touchpad                           \tid=11\t[slave  pointer  (2)]"
	expected := "11"
	actual := getDeviceID(data)
	assert.Equal(expected, actual)
}