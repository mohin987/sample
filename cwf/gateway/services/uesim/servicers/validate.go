/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package storage

import (
	"errors"
	"regexp"

	"magma/cwf/cloud/go/protos"
)

// validateUEData ensures that a UE data proto is not nil and that it contains
// a valid IMSI, key, and opc.
func validateUEData(ue *protos.UEConfig) error {
	if ue == nil {
		return errors.New("Invalid Argument: UE data cannot be nil")
	}
	errIMSI := validateUEIMSI(ue.GetImsi())
	if errIMSI != nil {
		return errIMSI
	}
	errkey := validateUEKey(ue.GetAuthKey())
	if errkey != nil {
		return errkey
	}
	erropc := validateUEOpc(ue.GetAuthOpc())
	if erropc != nil {
		return erropc
	}
	return nil
}

// validateUEIMSI ensures that a UE's IMSI can be stored.
func validateUEIMSI(imsi string) error {
	if len(imsi) < 5 || len(imsi) > 15 {
		return errors.New("Invalid Argument: IMSI must be between 5 and 15 digits long")
	}
	isOnlyDigits, err := regexp.MatchString(`^[0-9]*$`, imsi)
	if err != nil || !isOnlyDigits {
		return errors.New("Invalid Argument: IMSI must only be digits")
	}
	return nil
}

// validateUEKey ensures that a UE's key can be stored.
func validateUEKey(k []byte) error {
	if k == nil {
		return errors.New("Invalid Argument: key cannot be nil")
	}
	if len(k) != 32 {
		return errors.New("Invalid Argument: key must be 32 bytes")
	}
	return nil
}

// validateUEOpc ensures that a UE's opc can be stored.
func validateUEOpc(opc []byte) error {
	if opc == nil {
		return errors.New("Invalid Argument: opc cannot be nil")
	}
	if len(opc) != 32 {
		return errors.New("Invalid Argument: opc must be 32 bytes")
	}
	return nil
}
