// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Erc1155SwapPairs erc1155 swap pairs
//
// swagger:model Erc1155SwapPairs
type Erc1155SwapPairs struct {

	// pairs
	Pairs []*Erc1155SwapPair `json:"pairs"`

	// total
	Total int64 `json:"total"`
}

// Validate validates this erc1155 swap pairs
func (m *Erc1155SwapPairs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePairs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Erc1155SwapPairs) validatePairs(formats strfmt.Registry) error {
	if swag.IsZero(m.Pairs) { // not required
		return nil
	}

	for i := 0; i < len(m.Pairs); i++ {
		if swag.IsZero(m.Pairs[i]) { // not required
			continue
		}

		if m.Pairs[i] != nil {
			if err := m.Pairs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pairs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this erc1155 swap pairs based on the context it is used
func (m *Erc1155SwapPairs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePairs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Erc1155SwapPairs) contextValidatePairs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pairs); i++ {

		if m.Pairs[i] != nil {
			if err := m.Pairs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pairs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Erc1155SwapPairs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Erc1155SwapPairs) UnmarshalBinary(b []byte) error {
	var res Erc1155SwapPairs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}