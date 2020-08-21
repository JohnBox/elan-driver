package main

import (
	"log"
	"os/exec"
	"strings"
)

func findInString(str, substr string) bool {
	return strings.Contains(str, substr)
}

func getId(deviceInfo string) string {
	return strings.Split(strings.Split(deviceInfo, "\t")[1], "=")[1]
}

func main() {
	command := "xinput"
	out, err := exec.Command(command).Output()

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(out), "\n")
	device := "touchpad"
	var deviceId string

	for _, line := range lines {
		if findInString(line, device) {
			deviceId = getId(line)
			break
		}
	}

	setProps := map[string]string{
		"Tapping Enabled":           "1",
		"Natural Scrolling Enabled": "1",
		"Middle Emulation Enabled":  "1",
		"Accel Speed":               "0.4",
	}
	listPropsCmd, _ := exec.Command("xinput", "list-props", deviceId).Output()

	allProps := strings.Split(string(listPropsCmd), "\n")

	for setPropName, setPropValue := range setProps {
		for _, prop := range allProps {
			if findInString(prop, setPropName) {
				propId := prop[strings.IndexRune(prop, '(')+1 : strings.IndexRune(prop, ')')]
				_, err := exec.Command("xinput", "set-prop", deviceId, propId, setPropValue).Output()
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		}
	}

	log.Printf("%v", setProps)
}
