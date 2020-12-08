package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passportsRaw, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	passportsSplit := strings.Split(string(passportsRaw), "\r\n\r\n")
	validPassportsCount := 0
	for _, passportEntryRaw := range passportsSplit {
		passportEntryAsString := strings.ReplaceAll(passportEntryRaw, "\r\n", " ")
		passportEntry := make(map[string]string)
		for _, passportField := range strings.Split(passportEntryAsString, " ") {
			passportFieldParsed := strings.Split(passportField, ":")
			passportEntry[passportFieldParsed[0]] = passportFieldParsed[1]
		}
		valid := isValid(passportEntry)
		if valid {
			validPassportsCount++
		}
	}
	fmt.Println("The number of valid passports is:", validPassportsCount)
}

func isValid(passportEntry map[string]string) (isValid bool) {
	requiredKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, k := range requiredKeys {
		if _, key := passportEntry[k]; !key {
			return false
		}
	}
	for _, k := range requiredKeys {
		switch k {
		case "byr":
			if byr, err := strconv.Atoi(passportEntry[k]); err != nil || (byr < 1920 || byr > 2002) {
				return false
			}
		case "iyr":
			if iyr, err := strconv.Atoi(passportEntry[k]); err != nil || (iyr < 2010 || iyr > 2020) {
				return false
			}
		case "eyr":
			if eyr, err := strconv.Atoi(passportEntry[k]); err != nil || (eyr < 2020 || eyr > 2030) {
				return false
			}
		case "hgt":
			r := regexp.MustCompile("^([0-9]+)(cm|in)$")
			matched := r.FindStringSubmatch(passportEntry[k])
			if len(matched) == 0 {
				return false
			} else if matched[2] == "cm" {
				height, _ := strconv.Atoi(matched[1])
				if height < 150 || height > 193 {
					return false
				}
			} else if matched[2] == "in" {
				height, _ := strconv.Atoi(matched[1])
				if height < 59 || height > 76 {
					return false
				}
			}
		case "hcl":
			if ok, err := regexp.MatchString("^#[0-9a-f]{6}$", passportEntry[k]); err != nil || !ok {
				return false
			}
		case "ecl":
			if ok, err := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", passportEntry[k]); err != nil || !ok {
				return false
			}
		case "pid":
			if ok, err := regexp.MatchString("^[0-9]{9}$", passportEntry[k]); err != nil || !ok {
				return false
			}
		}
	}
	return true
}
