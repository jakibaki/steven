// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate protocol_builder $GOFILE Login serverbound

package protocol

// LoginStart is sent immeditately after switching into the login
// state. The passed username is used by the server to authenticate
// the player in online mode.
//
// Currently the packet id is: 0x00
type LoginStart struct {
	Username string
}

// EncryptionResponse is sent as a reply to EncryptionRequest. All
// packets following this one must be encrypted with AES/CFB8
// encryption.
//
// Currently the packet id is: 0x01
type EncryptionResponse struct {
	// The key for the AES/CFB8 cipher encrypted with the
	// public key
	SharedSecret []byte `length:"VarInt"`
	// The verify token from the request encrypted with the
	// public key
	VerifyToken []byte `length:"VarInt"`
}
