package main

import (
	"log"
	"os"

	"github.com/bouke/hc-icy"
	"github.com/brutella/hc/hap"
)

func main() {
	thermostat, err := icy.NewThermostat(os.Getenv("ICY_USERNAME"), os.Getenv("ICY_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	t, err := hap.NewIPTransport(hap.Config{Pin: "12348765"}, thermostat.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hap.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
