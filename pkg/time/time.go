package time

import "time"

func LoadLocation() (*time.Location, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	return loc, err
}
