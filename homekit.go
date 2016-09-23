package main

import (
	"log"
	"os"

	"github.com/bouke/hc-icy"
	"github.com/brutella/hc"
)

func main() {
	thermostat, err := icy.NewThermostat(os.Getenv("ICY_USERNAME"), os.Getenv("ICY_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	t, err := hc.NewIPTransport(hc.Config{Pin: "12344321"}, thermostat.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
