package part1

const (
	SeedToSoil            = "seed-to-soil"
	SoilToFertilizer      = "soil-to-fertilizer"
	FertilizerToWater     = "fertilizer-to-water"
	WaterToLight          = "water-to-light"
	LightToTemperature    = "light-to-temperature"
	TemperatureToHumidity = "temperature-to-humidity"
	HumidityToLocation    = "humidity-to-location"
)

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

func (maps Maps) update(dest, source string, vals Vals) {}