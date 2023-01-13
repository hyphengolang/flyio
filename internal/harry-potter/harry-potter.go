package harrypotter

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
	intern "github.com/hyphengolang/flyio/internal"
)

var ErrInvalidTyp = errors.New("harry-potter: invalid type")

type Character struct {
	ID      uuid.UUID    `json:"id"`
	Name    string       `json:"name"`
	Blood   BloodTyp     `json:"blood"`
	Species SpeciesTyp   `json:"species"`
	Born    *intern.Time `json:"born"`
	Quote   string       `json:"quote"`
	ImgURL  *url.URL     `json:"imgUrl"`
}

func (c Character) String() string {
	return fmt.Sprintf("%s (%s, %s)", c.Name, c.Blood, c.Species)
}

type BloodTyp int

const (
	BloodMuggle BloodTyp = iota
	BloodPure
	BloodHalf
)

var bts = [...]string{
	"muggle-born",
	"pure-blood",
	"half-blood",
}

var btm = map[string]BloodTyp{
	"muggle-born": BloodMuggle,
	"pure-blood":  BloodPure,
	"half-blood":  BloodHalf,
}

func ParseBloodTyp(s string) (BloodTyp, error) {
	if b, ok := btm[strings.ToLower(s)]; ok {
		return b, nil
	}
	return -1, fmt.Errorf("%w: blood type of %s not allowed", ErrInvalidTyp, s)
}

func (t *BloodTyp) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if pt, err := ParseBloodTyp(s); err != nil {
		return err
	} else {
		*t = pt
		return nil
	}
}

func (b BloodTyp) String() string { return bts[b] }

type SpeciesTyp int

const (
	SpeciesHuman SpeciesTyp = iota
	SpeciesHalfGiant
	SpeciesWerewolf
)

var sta = [...]string{
	"human",
	"half-giant",
	"werewolf",
}

var stm = map[string]SpeciesTyp{
	"human":      SpeciesHuman,
	"half-giant": SpeciesHalfGiant,
	"werewolf":   SpeciesWerewolf,
}

func ParseSpeciesTyp(s string) (SpeciesTyp, error) {
	if s, ok := stm[strings.ToLower(s)]; ok {
		return s, nil
	}
	return -1, fmt.Errorf("%w: species type of %s not allowed", ErrInvalidTyp, s)
}

func (t *SpeciesTyp) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if pt, err := ParseSpeciesTyp(s); err != nil {
		return err
	} else {
		*t = pt
		return nil
	}
}

func (s SpeciesTyp) String() string { return sta[s] }
