package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePassports(t *testing.T) {
	rawPassports := strings.NewReader(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)
	expected := []Passport{
		{
			EyeColor:       "gry",
			PassportID:     "860033327",
			ExpirationYear: "2020",
			HairColor:      "#fffffd",
			BirthYear:      "1937",
			IssueYear:      "2017",
			CountryID:      "147",
			Height:         "183cm",
		},
		{
			IssueYear:      "2013",
			EyeColor:       "amb",
			CountryID:      "350",
			ExpirationYear: "2023",
			PassportID:     "028048884",
			HairColor:      "#cfa07d",
			BirthYear:      "1929",
		},
		{
			HairColor:      "#ae17e1",
			IssueYear:      "2013",
			ExpirationYear: "2024",
			EyeColor:       "brn",
			PassportID:     "760753108",
			BirthYear:      "1931",
			Height:         "179cm",
		},
		{
			HairColor:      "#cfa07d",
			ExpirationYear: "2025",
			PassportID:     "166559648",
			IssueYear:      "2011",
			EyeColor:       "brn",
			Height:         "59in",
		},
	}

	actual := ParsePassports(rawPassports)
	assert.Equal(t, expected, actual)
}

func TestPassportValidity(t *testing.T) {
	cases := []struct {
		given    Passport
		expected bool
	}{
		{
			Passport{
				EyeColor:       "gry",
				PassportID:     "860033327",
				ExpirationYear: "2020",
				HairColor:      "#fffffd",
				BirthYear:      "1985",
				IssueYear:      "2017",
				CountryID:      "147",
				Height:         "183cm",
			},
			true,
		},
		{
			Passport{
				IssueYear:      "2013",
				EyeColor:       "amb",
				CountryID:      "350",
				ExpirationYear: "2023",
				PassportID:     "028048884",
				HairColor:      "#cfa07d",
				BirthYear:      "1929",
			},
			false,
		},
		{
			Passport{
				HairColor:      "#ae17e1",
				IssueYear:      "2013",
				ExpirationYear: "2024",
				EyeColor:       "brn",
				PassportID:     "760753108",
				BirthYear:      "1931",
				Height:         "179cm",
			},
			true,
		},
		{
			Passport{
				HairColor:      "#cfa07d",
				ExpirationYear: "2025",
				PassportID:     "166559648",
				IssueYear:      "2011",
				EyeColor:       "brn",
				Height:         "59in",
			},
			false,
		},
	}

	for _, tc := range cases {
		actual := tc.given.IsValid()
		assert.Equal(t, tc.expected, actual)
	}

}
