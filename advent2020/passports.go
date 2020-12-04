package advent2020

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear, IssueYear, ExpirationYear, Height string
	HairColor, EyeColor, PassportID, CountryID   string
}

func (p Passport) IsValid() bool {
	return p.hasValidBirthYear() && p.hasValidIssueYear() &&
		p.hasValidExpirationYear() && p.hasValidHeight() &&
		p.hasValidHairColor() && p.hasValidEyeColor() && p.hasValidID()
}

func (p Passport) hasValidBirthYear() bool {
	if p.BirthYear == "" {
		return false
	}

	year, err := strconv.Atoi(p.BirthYear)

	if err != nil {
		return false
	}

	return year >= 1920 && year < 2003
}

func (p Passport) hasValidIssueYear() bool {
	if p.IssueYear == "" {
		return false
	}

	year, err := strconv.Atoi(p.IssueYear)

	if err != nil {
		return false
	}

	return year >= 2010 && year < 2021
}

func (p Passport) hasValidExpirationYear() bool {
	if p.ExpirationYear == "" {
		return false
	}

	year, err := strconv.Atoi(p.ExpirationYear)

	if err != nil {
		return false
	}

	return year >= 2020 && year < 2031
}

var heightRegex = regexp.MustCompile(`^(\d+)(cm|in)$`)

func (p Passport) hasValidHeight() bool {
	matches := heightRegex.FindStringSubmatch(p.Height)

	if matches == nil {
		return false
	}

	units := matches[2]
	height, err := strconv.Atoi(matches[1])

	if err != nil {
		return false
	}

	switch units {
	case "cm":
		return height >= 150 && height < 194
	default:
		return height >= 59 && height < 77
	}
}

var hairColorRegex = regexp.MustCompile(`^#[a-f0-9]{6}$`)

func (p Passport) hasValidHairColor() bool {
	matches := hairColorRegex.FindStringSubmatch(p.HairColor)

	if matches == nil {
		return false
	}

	return true
}

func (p Passport) hasValidEyeColor() bool {
	validEyeColors := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}

	_, present := validEyeColors[p.EyeColor]

	return present
}

var idRegex = regexp.MustCompile(`^[0-9]{9}$`)

func (p Passport) hasValidID() bool {
	matches := idRegex.FindStringSubmatch(p.PassportID)

	if matches == nil {
		return false
	}

	return true
}

func ParsePassports(r io.Reader) []Passport {
	var passports []Passport

	scanner := bufio.NewScanner(r)

	passport := Passport{}

	availableFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid",
	}

	for scanner.Scan() {
		rawEntry := scanner.Text()

		if len(rawEntry) == 0 {
			passports = append(passports, passport)
			passport = Passport{}
			continue
		}

		fields := parseFields(rawEntry)

		for _, field := range availableFields {
			if value, present := fields[field]; present {
				switch field {
				case "byr":
					passport.BirthYear = value
				case "iyr":
					passport.IssueYear = value
				case "eyr":
					passport.ExpirationYear = value
				case "hgt":
					passport.Height = value
				case "hcl":
					passport.HairColor = value
				case "ecl":
					passport.EyeColor = value
				case "pid":
					passport.PassportID = value
				case "cid":
					passport.CountryID = value
				}
			}
		}
	}

	passports = append(passports, passport)

	return passports
}

func parseFields(rawEntry string) map[string]string {
	fields := map[string]string{}

	fieldPairs := strings.Split(rawEntry, " ")

	for _, fieldPair := range fieldPairs {
		keyValues := strings.Split(fieldPair, ":")
		fields[keyValues[0]] = keyValues[1]
	}

	return fields
}
