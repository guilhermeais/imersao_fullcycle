package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	Id        string     `json:"id"`
	ClientId  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	Id        string    `json:"routeId"`
	ClientId  string    `json:"clientId"`
	Positions []float64 `json:"position"`
	Finished  bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.Id == "" {
		return errors.New("route id not informed")
	}

	file, err := os.Open("destinations/" + r.Id + ".txt")

	if err != nil {
		return err
	}

	defer file.Close() // fecha o arquivo após a execução da função

	scanner := bufio.NewScanner(file) // lê o arquivo linha a linha usando bugio scanner
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")     // separa linha em duas partes, latitude e longitude, usando a virgula como separador
		lat, errLat := strconv.ParseFloat(data[0], 64) // converte latitude de string para float64

		if errLat != nil {
			return errLat
		}

		long, errLong := strconv.ParseFloat(data[1], 64) // converte longitude de string para float64

		if errLong != nil {
			return errLong
		}

		r.Positions = append(r.Positions, Position{Lat: lat, Long: long}) // adiciona a posição na lista de posições

	}
	return nil
}

func (r *Route) ExportJSONPositionsAsString() ([]string, error) {
	var route PartialRoutePosition
	var result []string

	total := len(r.Positions)

	for index, value := range r.Positions {
		route.Id = r.Id
		route.ClientId = r.ClientId
		route.Positions = []float64{value.Lat, value.Long}
		route.Finished = index == total-1

		json, err := json.Marshal(route)

		if err != nil {
			return nil, err
		}

		result = append(result, string(json))
	}

	return result, nil
}
