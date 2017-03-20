package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"service/protocol"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
inputLoop:
	for {
		fmt.Println("please choose a service")
		fmt.Println("a. Import a collection\nb. FindCities\nc. Paging\n d. execute query\ne. exit")
		fmt.Println("enter your choice in (a,b,c,d)")
		input, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = bytes.TrimSpace(input)
		switch string(input) {
		case "a":
			val, err := importInputHandler(r)
			if err != nil {
				fmt.Println(err)
				continue
			}
			requestService(val, "import")
		case "b":
			val, err := citiesInputHandler(r)
			if err != nil {
				fmt.Println(err)
				continue
			}
			requestService(val, "cities")
		case "c":
			val, err := pagingInputHandler(r)
			if err != nil {
				fmt.Println(err)
				continue
			}
			requestService(val, "places")
		case "d":
			val, err := queryInputHandler(r)
			if err != nil {
				fmt.Println(err)
				continue
			}
			requestService(val, "records")
		case "e":
			break inputLoop
		default:
			fmt.Println("Unrecognized choice")
		}
	}

}

func requestService(val interface{}, url string) {
	data, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:8080/nagp/"+url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func importInputHandler(r *bufio.Reader) (*protocol.ImportParam, error) {
	fmt.Println("Enter database Name")
	database, err := r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	database = bytes.TrimSpace(database)

	fmt.Println("Enter Collection Name")
	collection, err := r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	collection = bytes.TrimSpace(collection)

	fmt.Println("Enter File path")
	path, err := r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	path = bytes.TrimSpace(path)

	return &protocol.ImportParam{
		Database:   string(database),
		Collection: string(collection),
		File:       string(path),
	}, nil

}

func citiesInputHandler(r *bufio.Reader) (*protocol.CitiesParam, error) {
	fmt.Println("if sort is not required leave empty like: 74.4,70,10,,20")
	fmt.Println("Enter longitude,Latitude,radius,sort,limit")
	data, err := r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	data = bytes.TrimSpace(data)
	strData := string(data)
	fields := strings.Split(strData, ",")
	switch len(fields) {
	case 3:
		lon, err := strconv.ParseFloat(fields[0], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		lat, err := strconv.ParseFloat(fields[1], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		rad, err := strconv.ParseFloat(fields[2], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		return &protocol.CitiesParam{
			Latitude:  lat,
			Longitude: lon,
			Radius:    rad,
		}, nil
	case 4:
		lat, err := strconv.ParseFloat(fields[0], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		lon, err := strconv.ParseFloat(fields[1], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		rad, err := strconv.ParseFloat(fields[2], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}

		return &protocol.CitiesParam{
			Latitude:  lat,
			Longitude: lon,
			Radius:    rad,
			Sort:      fields[3],
		}, nil
	case 5:
		lat, err := strconv.ParseFloat(fields[0], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		lon, err := strconv.ParseFloat(fields[1], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		rad, err := strconv.ParseFloat(fields[2], 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}

		limit, err := strconv.ParseUint(fields[4], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
		}
		return &protocol.CitiesParam{
			Latitude:  lat,
			Longitude: lon,
			Radius:    rad,
			Sort:      fields[3],
			Limit:     uint32(limit),
		}, nil
	default:
		return nil, fmt.Errorf("Insufficient or inappropriate parameters")
	}
}

func pagingInputHandler(r *bufio.Reader) (*protocol.PagingParam, error) {
	fmt.Println("Enter PageSize,PageIndex")
	data, err := r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	data = bytes.TrimSpace(data)
	strData := string(data)
	fields := strings.Split(strData, ",")
	if len(fields) < 2 {
		return nil, fmt.Errorf("Not enough parameters\n")
	}
	pageSize, err := strconv.ParseUint(fields[0], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
	}
	pageIndex, err := strconv.ParseUint(fields[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
	}
	return &protocol.PagingParam{
		PageIndex: uint32(pageIndex),
		PageSize:  uint32(pageSize),
	}, nil

}

func queryInputHandler(r *bufio.Reader) (*protocol.QueryParam, error) {
	fmt.Println("EnterQuery")
	data, err := r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	data = bytes.TrimSpace(data)

	return &protocol.QueryParam{
		Query: string(data),
	}, nil
}
