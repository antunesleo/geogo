package geo

import "math"

// Point represents a location on earth
type Point struct {
	Lat  float64 `validate:"min=-90,max=90"`
	Long float64 `validate:"min=-180,max=190"`
}

func degreesToRadians(d float64) float64 {
	const x = math.Pi / 180
	return d * x
}

func (fp Point) DistanceBetween(sp Point) float64 {
	firstLat := degreesToRadians(fp.Lat)
	firstLong := degreesToRadians(fp.Long)
	secondLat := degreesToRadians(sp.Lat)
	secondLon := degreesToRadians(sp.Long)

	result := math.Acos(
		(math.Sin(firstLat) * math.Sin(secondLat)) +
			math.Cos(firstLat)*math.Cos(secondLat)*
				math.Cos(secondLon-firstLong),
	)
	const earth_readius_in_km = 6378.137
	return earth_readius_in_km * result
}
