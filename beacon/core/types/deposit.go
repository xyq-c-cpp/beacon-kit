// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package types

import "strconv"

// TODO: move deposit off of protobuf once the staking prs are merged.

// Deposit into the consensus layer from the deposit contract in the execution
// layer.
//

type Deposit struct {
	// Public key of the validator, which is compatible to the implementations
	// of the PubKey interface in Cosmos SDK. 32-byte ed25519 public key is
	// preferred.
	ValidatorPubkey []byte `json:"validatorPubkey" ssz-max:"48"`
	// A staking credentials with
	// 1 byte prefix + 11 bytes padding + 20 bytes address = 32 bytes.
	Credentials []byte `json:"credentials"                  ssz-size:"32"`
	// Deposit amount in gwei.
	Amount uint64 `json:"amount"`
	// Signature of the deposit data.
	Signature []byte `json:"signature"       ssz-max:"96"`
}

// String returns a string representation of the Deposit.
func (d *Deposit) String() string {
	return "Deposit{" +
		"ValidatorPubkey: " + string(d.ValidatorPubkey) +
		", Credentials: " + string(d.Credentials) +
		", Amount: " + strconv.FormatUint(d.Amount, 10) +
		", Signature: " + string(d.Signature) +
		"}"
}