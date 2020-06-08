package main

import (
	"fmt"
	"regexp"
)


func AreaCode(rawNumber string) (areaCode string, err error) {
	digits, err := Number(rawNumber)
	if err != nil {
	  return digits, err
	}
	return digits[0:3], nil
  }
  
  func Format(rawNumber string) (formattedNumber string, err error) {
	digits, err := Number(rawNumber)
	if err != nil {
	  return digits, err
	}
	return "(" + digits[0:3] + ") " + digits[3:6] + "-" + digits[6:10], nil
  }
  
  func Number(rawNumber string) (tenDigits string, err error) {
	letterR := regexp.MustCompile(`[a-z]`)
	if letterR.MatchString(strings.ToLower(rawNumber)) {
	  return rawNumber, errors.New("letters not permitted")
	}
	spaceR := regexp.MustCompile(`\s+`)
	wordcharR := regexp.MustCompile(`\w+`)
	symbols := wordcharR.ReplaceAllString(spaceR.ReplaceAllString(rawNumber, ""), "")
	punctuationsR := regexp.MustCompile(`[^+\-)(.]`)
	if punctuationsR.MatchString(symbols) {
	  return rawNumber, errors.New("punctuations not permitted")
	}
	nondigitR := regexp.MustCompile(`\D`)
	digits := nondigitR.ReplaceAllString(rawNumber, "");
	if len(digits) > 11 {
	  return digits, errors.New("more than 11 digits")
	}
	if len(digits) < 10 {
	  return digits, errors.New("incorrect number of digits")
	}
	if len(digits) == 11 {
	  if string(digits[0]) == "1" {
		digits = strings.Join(strings.Split(digits, "")[1:], "")
	  } else {
		return digits, errors.New("invalid when 11 digits does not start with a 1")
	  }
	}
	if string(digits[0]) == "0" {
	  return digits, errors.New("area code cannot start with zero")
	}
	if string(digits[0]) == "1" {
	  return digits, errors.New("area code cannot start with one")
	}
	if string(digits[3]) == "0" {
	  return digits, errors.New("exchange code cannot start with zero")
	}
	if string(digits[3]) == "1" {
	  return digits, errors.New("exchange code cannot start with one")
	}
	return digits, nil
  }