package controller

import (
	"fmt"
	"github.com/go-playground/colors"
	"github.com/lorenzobenvenuti/ifttt"
	"image/color"
	"strings"
)

// Lights Living Room Control
func LivingLightStripToggle() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "livingLightStripToggle"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)

	fmt.Println("Triggered ")
}

func LivingLightStripBrght(valueIn float64) float64 {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "livingBrght"
	valueOut := valueIn

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
	return valueOut
}

func LivingLightStripCol(valueIn float64) {
	valueOut := ""

	if valueIn > 50 {
		valueOut = "warmwhite"
	}
	if valueIn <= 50 {
		valueOut = "coldwhite"
	}

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "livingBrght"

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
}

func LivingBubBrght(valueIn float64) float64 {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "bubBrght"
	valueOut := valueIn

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
	return valueOut
}
func LivingBubToggle() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "livingBubToggle"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)

	fmt.Println("Triggered")
}

func LivingBubColor(valueIn color.Color) {

	colorHex := fmt.Sprintf("%v", valueIn)
	colorHex = strings.ReplaceAll(colorHex, " ", ",")
	colorHex = strings.ReplaceAll(colorHex, "{", "(")
	colorHex = strings.ReplaceAll(colorHex, "}", ")")
	rgb := "rgb"
	colorHex = rgb + colorHex
	colorHex = colorHex[:len(colorHex)-5]
	end := ")"
	colorHex += end

	hex, err := colors.ParseRGB(colorHex)

	fmt.Println(err)

	newHex := hex.ToHEX()

	fmt.Println(newHex)

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "livingBubColor"
	valueOut := newHex
	fmt.Println(hex)

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
}

// Lights Kitchen Control
func KitchenMainToggle() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenMainToggle"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)

	fmt.Println("Triggered")
}

func KitchenCabToggle() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenCabToggle"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)

	fmt.Println("Triggered")
}

func KitchenCabCol(valueIn float64) {
	valueOut := ""

	if valueIn > 50 {
		valueOut = "warmwhite"
	}
	if valueIn <= 50 {
		valueOut = "coldwhite"
	}

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenCabCol"

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
}

func KitchenCabBrght(valueIn float64) float64 {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenCabBrght"
	valueOut := valueIn

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
	return valueOut
}

func KitchenBarToggle() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenBarToggle"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)

	fmt.Println("Triggered")
}

// Bar Light Controll
func TableLightToggle() {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenCabToggle"

	values := []string{"", "", ""}
	iftttClient.Trigger(event, values)

	fmt.Println("Triggered")
}

func TableLightCol(valueIn float64) {
	valueOut := ""

	if valueIn > 50 {
		valueOut = "warmwhite"
	}
	if valueIn <= 50 {
		valueOut = "coldwhite"
	}

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenCabCol"

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
}

func TableLightBrght(valueIn float64) float64 {

	iftttClient := ifttt.NewIftttClient(ApiKey)
	event := "kitchenCabBrght"
	valueOut := valueIn

	values := []string{fmt.Sprintf("%v", valueOut), "", ""}
	iftttClient.Trigger(event, values)
	fmt.Println("Triggered with: ", valueOut)
	return valueOut
}
