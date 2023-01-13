package internal

import (
	"encoding/json"
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

func TestTime(t *testing.T) {
	is := is.New(t)

	t.Run("unmarshal", func(t *testing.T) {
		payload := `["2021-01-01", "1980-07-30"]`

		var times []Time
		err := json.Unmarshal([]byte(payload), &times)
		is.NoErr(err) // unmarshal
		is.Equal(times[1].String(), "1980-07-30")
	})

}
