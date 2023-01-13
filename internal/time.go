package internal

import (
	"strings"
	"time"
)

// This is a custom time type that uses the YYYY-MM-DD format
//
// TODO - allow YYYY-MM and YYYY formats
type Time time.Time

func (c *Time) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		// NOTE - should return an error
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = Time(t) //set result using the pointer
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02") + `"`), nil
}

func (t Time) String() string { return time.Time(t).Format("2006-01-02") }

func ParseTime(s string) (Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Time{}, err
	}
	return Time(t), nil
}
