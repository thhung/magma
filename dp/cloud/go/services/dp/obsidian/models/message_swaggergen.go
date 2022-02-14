// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Message message
// swagger:model message
type Message struct {

	// Body of message
	// Read Only: true
	Body string `json:"body,omitempty"`

	// Fcc Id of cbsd involved in message
	// Read Only: true
	FccID string `json:"fcc_id,omitempty"`

	// message origin
	// Read Only: true
	// Enum: [SAS DP CBSD]
	From string `json:"from,omitempty"`

	// Serial number of cbsd involved in message
	// Read Only: true
	SerialNumber string `json:"serial_number,omitempty"`

	// Datetime of message
	// Read Only: true
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// message destination
	// Read Only: true
	// Enum: [SAS DP CBSD]
	To string `json:"to,omitempty"`

	// Type of message
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var messageTypeFromPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SAS","DP","CBSD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageTypeFromPropEnum = append(messageTypeFromPropEnum, v)
	}
}

const (

	// MessageFromSAS captures enum value "SAS"
	MessageFromSAS string = "SAS"

	// MessageFromDP captures enum value "DP"
	MessageFromDP string = "DP"

	// MessageFromCBSD captures enum value "CBSD"
	MessageFromCBSD string = "CBSD"
)

// prop value enum
func (m *Message) validateFromEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, messageTypeFromPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Message) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	// value enum
	if err := m.validateFromEnum("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageTypeToPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SAS","DP","CBSD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageTypeToPropEnum = append(messageTypeToPropEnum, v)
	}
}

const (

	// MessageToSAS captures enum value "SAS"
	MessageToSAS string = "SAS"

	// MessageToDP captures enum value "DP"
	MessageToDP string = "DP"

	// MessageToCBSD captures enum value "CBSD"
	MessageToCBSD string = "CBSD"
)

// prop value enum
func (m *Message) validateToEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, messageTypeToPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Message) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	// value enum
	if err := m.validateToEnum("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
