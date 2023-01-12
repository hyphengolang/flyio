package harrypotter_test

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/google/uuid"
	intern "github.com/hyphengolang/flyio/internal/harry-potter"
	"github.com/hyphengolang/prelude/testing/is"
)

func TestCharacterTyp(t *testing.T) {
	is := is.New(t)

	birth, err := time.Parse("2006-01-02", "1980-07-31")
	is.NoErr(err) // parse birth day

	img, err := url.Parse("https://i.pinimg.com/originals/58/39/61/5839613bc887946211e72778d01da05f.jpg")
	is.NoErr(err) // parse image url

	c := intern.Character{
		ID:      uuid.New(),
		Name:    "Harry Potter",
		Blood:   intern.BloodMuggle,
		Species: intern.SpeciesHuman,
		Born:    &birth,
		Quote:   "I don't go looking for trouble. Trouble usually finds me.",
		ImgURL:  img,
	}

	p, err := json.Marshal(c)
	is.NoErr(err) // marshal character

	var c2 intern.Character
	err = json.Unmarshal(p, &c2)
	is.NoErr(err) // unmarshal character

	is.Equal(c, c2) // should be equal
}

func TestParseBloodTyp(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		name    string
		s       string
		want    intern.BloodTyp
		wantErr bool
	}{
		{
			name:    "muggle-born",
			s:       "muggle-born",
			want:    intern.BloodMuggle,
			wantErr: false,
		},
		{
			name:    "pure-blood",
			s:       "pure-blood",
			want:    intern.BloodPure,
			wantErr: false,
		},
		{
			name:    "half-blood",
			s:       "Half-Blood",
			want:    intern.BloodHalf,
			wantErr: false,
		},
		{
			name:    "invalid",
			s:       "invalid",
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := intern.ParseBloodTyp(tt.s)
			if err != nil {
				is.Equal(err != nil, tt.wantErr) // should throw error
				return
			}

			is.Equal(got, tt.want) // should be equal
		})
	}
}

func TestParseSpeciesTyp(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		name    string
		s       string
		want    intern.SpeciesTyp
		wantErr bool
	}{
		{
			name:    "human",
			s:       "human",
			want:    intern.SpeciesHuman,
			wantErr: false,
		},
		{
			name:    "half-giant",
			s:       "half-giant",
			want:    intern.SpeciesHalfGiant,
			wantErr: false,
		},
		{
			name:    "werewolf",
			s:       "werewolf",
			want:    intern.SpeciesWerewolf,
			wantErr: false,
		},
		{
			name:    "invalid",
			s:       "invalid",
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := intern.ParseSpeciesTyp(tt.s)
			if err != nil {
				is.Equal(err != nil, tt.wantErr) // should throw error
				return
			}

			is.Equal(got, tt.want) // should be equal
		})
	}
}
