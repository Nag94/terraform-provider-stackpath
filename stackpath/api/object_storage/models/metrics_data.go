// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsData The data points in a metrics collection
// swagger:model MetricsData
type MetricsData struct {

	// matrix
	Matrix *MetricsDataMatrix `json:"matrix,omitempty"`

	// vector
	Vector *MetricsDataVector `json:"vector,omitempty"`
}

// Validate validates this metrics data
func (m *MetricsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsData) validateMatrix(formats strfmt.Registry) error {

	if swag.IsZero(m.Matrix) { // not required
		return nil
	}

	if m.Matrix != nil {
		if err := m.Matrix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

func (m *MetricsData) validateVector(formats strfmt.Registry) error {

	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if m.Vector != nil {
		if err := m.Vector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsData) UnmarshalBinary(b []byte) error {
	var res MetricsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetricsDataMatrix A set of time series containing a range of data points over time for each time series
// swagger:model MetricsDataMatrix
type MetricsDataMatrix struct {

	// A data point's value
	Results []*MetricsDataMatrixResultsItems0 `json:"results"`
}

// Validate validates this metrics data matrix
func (m *MetricsDataMatrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsDataMatrix) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDataMatrix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDataMatrix) UnmarshalBinary(b []byte) error {
	var res MetricsDataMatrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetricsDataMatrixResultsItems0 Time series containing a range of data points over time for each time series
// swagger:model MetricsDataMatrixResultsItems0
type MetricsDataMatrixResultsItems0 struct {

	// The data points' labels
	Metric map[string]string `json:"metric,omitempty"`

	// Time series data point values
	Values []*MetricsDataMatrixResultsItems0ValuesItems0 `json:"values"`
}

// Validate validates this metrics data matrix results items0
func (m *MetricsDataMatrixResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsDataMatrixResultsItems0) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDataMatrixResultsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDataMatrixResultsItems0) UnmarshalBinary(b []byte) error {
	var res MetricsDataMatrixResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetricsDataMatrixResultsItems0ValuesItems0 An individual metric data point
// swagger:model MetricsDataMatrixResultsItems0ValuesItems0
type MetricsDataMatrixResultsItems0ValuesItems0 struct {

	// The time that a data point was recorded
	UnixTime string `json:"unixTime,omitempty"`

	// A data point's value
	Value string `json:"value,omitempty"`
}

// Validate validates this metrics data matrix results items0 values items0
func (m *MetricsDataMatrixResultsItems0ValuesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDataMatrixResultsItems0ValuesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDataMatrixResultsItems0ValuesItems0) UnmarshalBinary(b []byte) error {
	var res MetricsDataMatrixResultsItems0ValuesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetricsDataVector A set of time series containing a single sample for each time series, all sharing the same timestamp
// swagger:model MetricsDataVector
type MetricsDataVector struct {

	// A data point's value
	Results []*MetricsDataVectorResultsItems0 `json:"results"`
}

// Validate validates this metrics data vector
func (m *MetricsDataVector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsDataVector) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vector" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDataVector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDataVector) UnmarshalBinary(b []byte) error {
	var res MetricsDataVector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetricsDataVectorResultsItems0 Time series containing a single sample for each time series, all sharing the same timestamp
// swagger:model MetricsDataVectorResultsItems0
type MetricsDataVectorResultsItems0 struct {

	// The data points' labels
	Metric map[string]string `json:"metric,omitempty"`

	// value
	Value *MetricsDataVectorResultsItems0Value `json:"value,omitempty"`
}

// Validate validates this metrics data vector results items0
func (m *MetricsDataVectorResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsDataVectorResultsItems0) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDataVectorResultsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDataVectorResultsItems0) UnmarshalBinary(b []byte) error {
	var res MetricsDataVectorResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetricsDataVectorResultsItems0Value An individual metric data point
// swagger:model MetricsDataVectorResultsItems0Value
type MetricsDataVectorResultsItems0Value struct {

	// The time that a data point was recorded
	UnixTime string `json:"unixTime,omitempty"`

	// A data point's value
	Value string `json:"value,omitempty"`
}

// Validate validates this metrics data vector results items0 value
func (m *MetricsDataVectorResultsItems0Value) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDataVectorResultsItems0Value) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDataVectorResultsItems0Value) UnmarshalBinary(b []byte) error {
	var res MetricsDataVectorResultsItems0Value
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}