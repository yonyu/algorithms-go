package adt

import (
	"fmt"
)

type CityWeatherData struct { 
	name string 
	temperature float64 
	humidity float64 
	Wind  // Anomynous field. It is called struct embedding
}

// When we do not want to use a method derived from an embedded struct, func (w *Wind) Set( direction, speed float64)
// a new method with the same name can be declared with the embedding struct as its receiver. Go will then use the
// outer method when called with the embedding struct as receiver, and the inner method when called with an embedded
// struct value as receiver. This provides a way to override derived methods.
func (d *CityWeatherData) Set(name string, temperature, humidity, direction, speed float64) {
	d.name = name
	d.temperature = temperature
	d.humidity = humidity
	d.direction = direction
	d.speed = speed
}

func DemonstrateOverride() {
	d := CityWeatherData {}
	d.Set("Phoenix", 102, 32, 220.1, 8.6)
	fmt.Printf("%#v\n", d)

	d.Wind.Set(-80.3, 15.3)
	fmt.Printf("%#v\n", d)
}