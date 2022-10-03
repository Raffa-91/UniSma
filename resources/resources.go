package resources

import (
	"bufio"
	"embed"
	"fyne.io/fyne/v2"
	"io/ioutil"
	"log"
)

//go:embed assets
var Assets embed.FS

type Resource string

const (
	SonosLogo Resource = "assets/sonos.png"
	Light     Resource = "assets/light.jpg"
)

func GetResource(res Resource, name string) *fyne.StaticResource {
	f, err := Assets.Open(string(res))
	r := bufio.NewReader(f)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return fyne.NewStaticResource(name, b)
}
