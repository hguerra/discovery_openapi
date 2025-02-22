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




// UpdateSpecialEventRequest - Request payload for updating an existing special event. Only included fields are updated in the event.
type UpdateSpecialEventRequest struct {

	// Name of the special event
	Name string `json:"name,omitempty"`

	// Location where the special event is held
	Location string `json:"location,omitempty"`

	// Description of the special event
	EventDescription string `json:"eventDescription,omitempty"`

	// List of planned dates for the special event
	Dates []string `json:"dates,omitempty"`

	// Price of a ticket for the special event
	Price float32 `json:"price,omitempty"`
}

// AssertUpdateSpecialEventRequestRequired checks if the required fields are not zero-ed
func AssertUpdateSpecialEventRequestRequired(obj UpdateSpecialEventRequest) error {
	return nil
}

// AssertUpdateSpecialEventRequestConstraints checks if the values respects the defined constraints
func AssertUpdateSpecialEventRequestConstraints(obj UpdateSpecialEventRequest) error {
	return nil
}
