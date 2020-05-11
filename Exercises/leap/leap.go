// Package leap provides functions for work with leap years.
package leap

// IsLeapYear returns if the given year is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}