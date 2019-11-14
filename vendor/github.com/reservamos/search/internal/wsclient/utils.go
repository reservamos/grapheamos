package wsclient

import "math"

// RouteCoordinate is a Lat,Lon value pair giving a location of a terminal
type RouteCoordinate struct {
	Lat float64
	Lon float64
}

type SphericalCoordinate float64

func (s SphericalCoordinate) Radians() float64 {
	return float64(s) * math.Pi / 180.0
}

func (p1 RouteCoordinate) Distance(p2 RouteCoordinate) float64 {
	R := 6371e3
	φ1 := SphericalCoordinate(p1.Lat).Radians()
	φ2 := SphericalCoordinate(p2.Lat).Radians()
	Δφ := SphericalCoordinate(p2.Lat - p1.Lat).Radians()
	Δλ := SphericalCoordinate(p2.Lon - p1.Lon).Radians()

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	d := R * c
	return d / 1e3
}
