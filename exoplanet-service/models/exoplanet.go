package models

import (
	"errors"
	"fmt"
	"math/rand"
	//
	// "strconv"
)

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    float64       `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        float64       `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}

var Exoplanets = make(map[string]Exoplanet)

func AddExoplanet(exoplanet Exoplanet) (Exoplanet, error) {
	if exoplanet.Name == "" || exoplanet.Description == "" || exoplanet.Distance < 10 || exoplanet.Distance > 1000 ||
		exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return Exoplanet{}, errors.New("invalid exoplanet data")
	}

	if exoplanet.Type == Terrestrial && (exoplanet.Mass < 0.1 || exoplanet.Mass > 10) {
		return Exoplanet{}, errors.New("invalid exoplanet mass for Terrestrial type")
	}

	exoplanet.ID = fmt.Sprintf("%d", rand.Int())
	Exoplanets[exoplanet.ID] = exoplanet
	return exoplanet, nil
}
func UpdateExoplanet(id string, exoplanet Exoplanet) (Exoplanet, error) {
	if _, exists := Exoplanets[id]; !exists {
		return Exoplanet{}, errors.New("exoplanet not found")
	}

	exoplanet.ID = id
	Exoplanets[id] = exoplanet
	return exoplanet, nil
}

func GetGravity(exoplanet Exoplanet) float64 {
	if exoplanet.Type == GasGiant {
		return 0.5 / (exoplanet.Radius * exoplanet.Radius)
	}
	return exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
}

func CalculateFuel(distance float64, gravity float64, crewCapacity int) float64 {
	return distance / (gravity * gravity) * float64(crewCapacity)
}
