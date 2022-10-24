package controller

import (
	"github.com/lorenzobenvenuti/ifttt"
)

func SonosPlay() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "sonosPlay"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)
}

func SonosStop() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "sonosStop"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)
}

func SonosVolUp() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "sonosVolUp"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)
}

func SonosVolDown() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "sonosVolDown"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)
}

func SonosMute() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "sonosMute"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)
}

func SonosPlayRadio(r string) {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "sonosPlayRadio"

	switch r {
	case "SRF1":
		value := "http://stream.srg-ssr.ch/drs1/mp3_128.m3u"
		values := []string{value, "", ""}
		iftttClient.Trigger(event, values)
	case "Radio Vintage":
		value := "https://vintage80s.ice.infomaniak.ch/vintage80s-high.mp3"
		values := []string{value, "", ""}
		iftttClient.Trigger(event, values)
	case "Radio Argovia":
		value := "http://icecast.argovia.ch/argovia-rc-96-aac"
		values := []string{value, "", ""}
		iftttClient.Trigger(event, values)
	}
}
