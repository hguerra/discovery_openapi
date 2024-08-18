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




// SpecialEventResponse - Information about a special event.
type SpecialEventResponse struct {

	// Identifier for a special event.
	EventId string `json:"eventId"`

	// Name of the special event
	Name string `json:"name"`

	// Location where the special event is held
	Location string `json:"location"`

	// Description of the special event
	EventDescription string `json:"eventDescription"`

	// List of planned dates for the special event
	Dates []string `json:"dates"`

	// Price of a ticket for the special event
	Price float32 `json:"price"`
}

// AssertSpecialEventResponseRequired checks if the required fields are not zero-ed
func AssertSpecialEventResponseRequired(obj SpecialEventResponse) error {
	elements := map[string]interface{}{
		"eventId": obj.EventId,
		"name": obj.Name,
		"location": obj.Location,
		"eventDescription": obj.EventDescription,
		"dates": obj.Dates,
		"price": obj.Price,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSpecialEventResponseConstraints checks if the values respects the defined constraints
func AssertSpecialEventResponseConstraints(obj SpecialEventResponse) error {
	return nil
}