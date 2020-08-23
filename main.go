package main

import (
	"log"
	"os/exec"
	"strings"
)


func setDeviceProp(deviceID, propID, propValue string) {
	bytes, err := exec.Command("xinput", "set-prop", deviceID, propID, propValue).Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))
}

func getDevicePropID(propLine string) string {
	start := strings.IndexRune(propLine, '(')

	if start < 0 {
		return ""
	}

	end := strings.IndexRune(propLine,')')

	if end < 0 {
		return ""
	}

	return propLine[start+1:end]
}

func getDeviceID(deviceInfo string) string {
	return strings.Split(strings.Split(deviceInfo, "\t")[1], "=")[1]
}


func findInString(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}



func main() {
	command := "xinput"
	out, err := exec.Command(command).Output()

	if err != nil {
		log.Fatal(err)
	}

	device := "touchpad"
	var deviceId string

	for _, line := range strings.Split(string(out), "\n") {
		if findInString(line, device) {
			deviceId = getDeviceID(line)
			break
		}
	}

	setProps := map[string]string{
		"Tapping Enabled":           "1",
		"Natural Scrolling Enabled": "1",
		"Middle Emulation Enabled":  "1",
		"Accel Speed":               "0.4",
	}
	devicePropsBytes, err := exec.Command("xinput", "list-props", deviceId).Output()

	if err != nil {
		log.Println(err)
	}

	deviceProps := strings.Split(string(devicePropsBytes), "\n")

	for setPropName, setPropValue := range setProps {
		for _, prop := range deviceProps {
			if findInString(prop, setPropName) {
				propId := getDevicePropID(prop)
				setDeviceProp(deviceId, propId, setPropValue)
				break
			}
		}
	}
}
