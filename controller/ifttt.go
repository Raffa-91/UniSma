package controller

import (
	"github.com/lorenzobenvenuti/ifttt"
)

func TestLight() {
	apiKey := "  "
	event := "testTrigger"

	iftttClient := ifttt.NewIftttClient(apiKey)
	values := []string{"firstValue", "secondValue"}
	iftttClient.Trigger(event, values)
}
