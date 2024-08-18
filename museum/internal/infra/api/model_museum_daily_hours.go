// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redocly Museum API
 *
 * An imaginary, but delightful Museum API for interacting with museum services and information. Built with love by Redocly.
 *
 * API version: 1.0.0
 * Contact: team@redocly.com
 */

package api




// MuseumDailyHours - Daily operating hours for the museum.
type MuseumDailyHours struct {

	Date string `json:"date"`

	// Time the museum opens on a specific date. Uses 24 hour time format (`HH:mm`).
	TimeOpen string `json:"timeOpen" validate:"regexp=^([01]\\\\d|2[0-3]):?([0-5]\\\\d)$"`

	// Time the museum closes on a specific date. Uses 24 hour time format (`HH:mm`).
	TimeClose string `json:"timeClose" validate:"regexp=^([01]\\\\d|2[0-3]):?([0-5]\\\\d)$"`
}

// AssertMuseumDailyHoursRequired checks if the required fields are not zero-ed
func AssertMuseumDailyHoursRequired(obj MuseumDailyHours) error {
	elements := map[string]interface{}{
		"date": obj.Date,
		"timeOpen": obj.TimeOpen,
		"timeClose": obj.TimeClose,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertMuseumDailyHoursConstraints checks if the values respects the defined constraints
func AssertMuseumDailyHoursConstraints(obj MuseumDailyHours) error {
	return nil
}
