package parking_space

type ParkingSpace struct {
	ParkingSpaceId string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Address        string `json:"address"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Availability   bool   `json:"availability"`
}
