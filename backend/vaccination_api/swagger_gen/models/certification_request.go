// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificationRequest certification request
//
// swagger:model CertificationRequest
type CertificationRequest struct {

	// comorbidities
	Comorbidities []string `json:"comorbidities"`

	// enrollment type
	// Min Length: 1
	EnrollmentType string `json:"enrollmentType,omitempty"`

	// facility
	// Required: true
	Facility *CertificationRequestFacility `json:"facility"`

	// meta
	Meta interface{} `json:"meta,omitempty"`

	// pre enrollment code
	// Required: true
	PreEnrollmentCode *string `json:"preEnrollmentCode"`

	// program Id
	// Min Length: 1
	ProgramID string `json:"programId,omitempty"`

	// recipient
	// Required: true
	Recipient *CertificationRequestRecipient `json:"recipient"`

	// vaccination
	// Required: true
	Vaccination *CertificationRequestVaccination `json:"vaccination"`

	// vaccinator
	// Required: true
	Vaccinator *CertificationRequestVaccinator `json:"vaccinator"`
}

// Validate validates this certification request
func (m *CertificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrollmentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreEnrollmentCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccinator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequest) validateEnrollmentType(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrollmentType) { // not required
		return nil
	}

	if err := validate.MinLength("enrollmentType", "body", string(m.EnrollmentType), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequest) validateFacility(formats strfmt.Registry) error {

	if err := validate.Required("facility", "body", m.Facility); err != nil {
		return err
	}

	if m.Facility != nil {
		if err := m.Facility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequest) validatePreEnrollmentCode(formats strfmt.Registry) error {

	if err := validate.Required("preEnrollmentCode", "body", m.PreEnrollmentCode); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequest) validateProgramID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgramID) { // not required
		return nil
	}

	if err := validate.MinLength("programId", "body", string(m.ProgramID), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequest) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequest) validateVaccination(formats strfmt.Registry) error {

	if err := validate.Required("vaccination", "body", m.Vaccination); err != nil {
		return err
	}

	if m.Vaccination != nil {
		if err := m.Vaccination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vaccination")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequest) validateVaccinator(formats strfmt.Registry) error {

	if err := validate.Required("vaccinator", "body", m.Vaccinator); err != nil {
		return err
	}

	if m.Vaccinator != nil {
		if err := m.Vaccinator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vaccinator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequest) UnmarshalBinary(b []byte) error {
	var res CertificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestFacility certification request facility
//
// swagger:model CertificationRequestFacility
type CertificationRequestFacility struct {

	// address
	// Required: true
	Address *CertificationRequestFacilityAddress `json:"address"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this certification request facility
func (m *CertificationRequestFacility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestFacility) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestFacility) validateName(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestFacility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestFacility) UnmarshalBinary(b []byte) error {
	var res CertificationRequestFacility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestFacilityAddress certification request facility address
//
// swagger:model CertificationRequestFacilityAddress
type CertificationRequestFacilityAddress struct {

	// address line1
	// Required: true
	// Min Length: 1
	AddressLine1 *string `json:"addressLine1"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	// Required: true
	// Min Length: 1
	District *string `json:"district"`

	// pincode
	// Required: true
	// Min Length: 1
	Pincode *string `json:"pincode"`

	// state
	// Required: true
	// Min Length: 1
	State *string `json:"state"`
}

// Validate validates this certification request facility address
func (m *CertificationRequestFacilityAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistrict(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePincode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestFacilityAddress) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"addressLine1", "body", string(*m.AddressLine1), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestFacilityAddress) validateDistrict(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"district", "body", m.District); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"district", "body", string(*m.District), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestFacilityAddress) validatePincode(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"pincode", "body", m.Pincode); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"pincode", "body", string(*m.Pincode), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestFacilityAddress) validateState(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"state", "body", m.State); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"state", "body", string(*m.State), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestFacilityAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestFacilityAddress) UnmarshalBinary(b []byte) error {
	var res CertificationRequestFacilityAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestRecipient certification request recipient
//
// swagger:model CertificationRequestRecipient
type CertificationRequestRecipient struct {

	// address
	// Required: true
	Address *CertificationRequestRecipientAddress `json:"address"`

	// age
	// Min Length: 1
	Age string `json:"age,omitempty"`

	// contact
	// Required: true
	Contact []string `json:"contact"`

	// dob
	// Format: date
	Dob *strfmt.Date `json:"dob,omitempty"`

	// gender
	// Required: true
	// Min Length: 1
	Gender *string `json:"gender"`

	// identity
	// Required: true
	// Min Length: 1
	Identity *string `json:"identity"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// nationality
	// Required: true
	// Min Length: 1
	Nationality *string `json:"nationality"`
}

// Validate validates this certification request recipient
func (m *CertificationRequestRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationality(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestRecipient) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestRecipient) validateAge(formats strfmt.Registry) error {

	if swag.IsZero(m.Age) { // not required
		return nil
	}

	if err := validate.MinLength("recipient"+"."+"age", "body", string(m.Age), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipient) validateContact(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"contact", "body", m.Contact); err != nil {
		return err
	}

	for i := 0; i < len(m.Contact); i++ {

		if err := validate.MinLength("recipient"+"."+"contact"+"."+strconv.Itoa(i), "body", string(m.Contact[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *CertificationRequestRecipient) validateDob(formats strfmt.Registry) error {

	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := validate.FormatOf("recipient"+"."+"dob", "body", "date", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipient) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"gender", "body", m.Gender); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"gender", "body", string(*m.Gender), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipient) validateIdentity(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"identity", "body", m.Identity); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"identity", "body", string(*m.Identity), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipient) validateName(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipient) validateNationality(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"nationality", "body", m.Nationality); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"nationality", "body", string(*m.Nationality), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestRecipient) UnmarshalBinary(b []byte) error {
	var res CertificationRequestRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestRecipientAddress certification request recipient address
//
// swagger:model CertificationRequestRecipientAddress
type CertificationRequestRecipientAddress struct {

	// address line1
	// Required: true
	// Min Length: 1
	AddressLine1 *string `json:"addressLine1"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	// Required: true
	// Min Length: 1
	District *string `json:"district"`

	// pincode
	// Required: true
	// Min Length: 1
	Pincode *string `json:"pincode"`

	// state
	// Required: true
	// Min Length: 1
	State *string `json:"state"`
}

// Validate validates this certification request recipient address
func (m *CertificationRequestRecipientAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistrict(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePincode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestRecipientAddress) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"address"+"."+"addressLine1", "body", string(*m.AddressLine1), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipientAddress) validateDistrict(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"district", "body", m.District); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"address"+"."+"district", "body", string(*m.District), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipientAddress) validatePincode(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"pincode", "body", m.Pincode); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"address"+"."+"pincode", "body", string(*m.Pincode), 1); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestRecipientAddress) validateState(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"state", "body", m.State); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"address"+"."+"state", "body", string(*m.State), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestRecipientAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestRecipientAddress) UnmarshalBinary(b []byte) error {
	var res CertificationRequestRecipientAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestVaccination certification request vaccination
//
// swagger:model CertificationRequestVaccination
type CertificationRequestVaccination struct {

	// batch
	Batch string `json:"batch,omitempty"`

	// date
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// Dose number for example 1 for first dose of 2 doses
	// Required: true
	// Minimum: 1
	Dose *float64 `json:"dose"`

	// effective start
	// Required: true
	// Format: date
	EffectiveStart *strfmt.Date `json:"effectiveStart"`

	// effective until
	// Required: true
	// Format: date
	EffectiveUntil *strfmt.Date `json:"effectiveUntil"`

	// manufacturer
	// Required: true
	Manufacturer *string `json:"manufacturer"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Total number of doses required for this vaccination.
	// Required: true
	// Minimum: 1
	TotalDoses *float64 `json:"totalDoses"`
}

// Validate validates this certification request vaccination
func (m *CertificationRequestVaccination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveUntil(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDoses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestVaccination) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("vaccination"+"."+"date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateDose(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"dose", "body", m.Dose); err != nil {
		return err
	}

	if err := validate.Minimum("vaccination"+"."+"dose", "body", float64(*m.Dose), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateEffectiveStart(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"effectiveStart", "body", m.EffectiveStart); err != nil {
		return err
	}

	if err := validate.FormatOf("vaccination"+"."+"effectiveStart", "body", "date", m.EffectiveStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateEffectiveUntil(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"effectiveUntil", "body", m.EffectiveUntil); err != nil {
		return err
	}

	if err := validate.FormatOf("vaccination"+"."+"effectiveUntil", "body", "date", m.EffectiveUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateTotalDoses(formats strfmt.Registry) error {

	if err := validate.Required("vaccination"+"."+"totalDoses", "body", m.TotalDoses); err != nil {
		return err
	}

	if err := validate.Minimum("vaccination"+"."+"totalDoses", "body", float64(*m.TotalDoses), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestVaccination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestVaccination) UnmarshalBinary(b []byte) error {
	var res CertificationRequestVaccination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestVaccinator certification request vaccinator
//
// swagger:model CertificationRequestVaccinator
type CertificationRequestVaccinator struct {

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this certification request vaccinator
func (m *CertificationRequestVaccinator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestVaccinator) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vaccinator"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("vaccinator"+"."+"name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestVaccinator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestVaccinator) UnmarshalBinary(b []byte) error {
	var res CertificationRequestVaccinator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
