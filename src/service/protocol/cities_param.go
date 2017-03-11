package protocol

type CitiesParam struct {
	Latitude  float64 `json:"Lat"`
	Longitude float64 `json:"long"`
	Radius    float64 `json:"radius"`
	Sort      string  `json:"sort,omitempty"`
	Limit     uint32  `json:"limit,omitempty"`
}

/*
Lat: <latitude>
long: <longitude>
radius : <distance in m or Km>

sort : key - optional
limit : number of results - optional*/
