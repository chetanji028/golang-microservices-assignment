package ports

// Port represents a port entity
type Port struct {
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Coordinates [2]float64 `json:"coordinates"`
}
