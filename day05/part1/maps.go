package part1

import (
	"fmt"
	"log"
	"strings"
)

const (
	SeedToSoil            = "seed-to-soil"
	SoilToFertilizer      = "soil-to-fertilizer"
	FertilizerToWater     = "fertilizer-to-water"
	WaterToLight          = "water-to-light"
	LightToTemperature    = "light-to-temperature"
	TemperatureToHumidity = "temperature-to-humidity"
	HumidityToLocation    = "humidity-to-location"
)

type UpdateMap struct {
	Name   string
	Ranges []Ranges
}

func convertLinesToUpdateMap(lines []string) UpdateMap {
	var updateMap UpdateMap

	updateMap.Name = strings.Fields(lines[0])[0]
	for _, line := range lines[1:] {
		updateMap.Ranges = append(updateMap.Ranges, convertLineToRanges(line))
	}

	return updateMap
}

type Maps struct {
	SeedToSoil            map[int]int
	SoilToFertilizer      map[int]int
	FertilizerToWater     map[int]int
	WaterToLight          map[int]int
	LightToTemperature    map[int]int
	TemperatureToHumidity map[int]int
	HumidityToLocation    map[int]int
}

func NewMaps() Maps {
	return Maps{
		SeedToSoil:            map[int]int{},
		SoilToFertilizer:      map[int]int{},
		FertilizerToWater:     map[int]int{},
		WaterToLight:          map[int]int{},
		LightToTemperature:    map[int]int{},
		TemperatureToHumidity: map[int]int{},
		HumidityToLocation:    map[int]int{},
	}
}

func (maps Maps) findLocation(seed int) int {
	soil, ok := maps.SeedToSoil[seed]
	if !ok {
		soil = seed
	}
	fmt.Println("soil", "=>", soil)

	fertilizer, ok := maps.SoilToFertilizer[soil]
	if !ok {
		fertilizer = soil
	}
	fmt.Println("fertilizer", "=>", fertilizer)

	water, ok := maps.FertilizerToWater[fertilizer]
	if !ok {
		water = fertilizer
	}
	fmt.Println("water", "=>", water)

	light, ok := maps.WaterToLight[water]
	if !ok {
		light = water
	}
	fmt.Println("light", "=>", light)

	temperature, ok := maps.LightToTemperature[light]
	if !ok {
		temperature = light
	}
	fmt.Println("temperature", "=>", temperature)

	humidity, ok := maps.TemperatureToHumidity[temperature]
	if !ok {
		humidity = temperature
	}
	fmt.Println("humidity", "=>", humidity)

	location, ok := maps.HumidityToLocation[humidity]
	if !ok {
		location = humidity
	}
	fmt.Println("location", "=>", location)

	return location
}

func (maps Maps) update(u UpdateMap) {
	m := maps.getMap(u.Name)

	for _, r := range u.Ranges {
		for i, src := range r.SourceRange {
			m[src] = r.DestinationRange[i]
		}
	}

}

func (maps Maps) getMap(name string) map[int]int {
	switch name {
	case SeedToSoil:
		return maps.SeedToSoil
	case SoilToFertilizer:
		return maps.SoilToFertilizer
	case FertilizerToWater:
		return maps.FertilizerToWater
	case WaterToLight:
		return maps.WaterToLight
	case LightToTemperature:
		return maps.LightToTemperature
	case TemperatureToHumidity:
		return maps.TemperatureToHumidity
	case HumidityToLocation:
		return maps.HumidityToLocation
	default:
		log.Fatal("cannot find map", name)
	}

	return nil
}
