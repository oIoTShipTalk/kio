// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeAddressing Addressing information of a node for all address families
//
// +k8s:deepcopy-gen=true
//
// swagger:model NodeAddressing
type NodeAddressing struct {

	// ipv4
	IPV4 *NodeAddressingElement `json:"ipv4,omitempty"`

	// ipv6
	IPV6 *NodeAddressingElement `json:"ipv6,omitempty"`
}

// Validate validates this node addressing
func (m *NodeAddressing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeAddressing) validateIPV4(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4) { // not required
		return nil
	}

	if m.IPV4 != nil {
		if err := m.IPV4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4")
			}
			return err
		}
	}

	return nil
}

func (m *NodeAddressing) validateIPV6(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	if m.IPV6 != nil {
		if err := m.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeAddressing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeAddressing) UnmarshalBinary(b []byte) error {
	var res NodeAddressing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
