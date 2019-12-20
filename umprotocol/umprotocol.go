// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2019 Renesas Inc.
// Copyright 2019 EPAM Systems Inc.
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

package umprotocol

/*******************************************************************************
 * Consts
 ******************************************************************************/

// Message types
const (
	RevertType  = "systemRevert"
	StatusType  = "status"
	UpgradeType = "systemUpgrade"
)

// Operation status
const (
	SuccessStatus    = "success"
	FailedStatus     = "failed"
	InProgressStatus = "inprogress"
)

/*******************************************************************************
 * Types
 ******************************************************************************/

// MessageHeader UM message header
type MessageHeader struct {
	Type  string `json:"type"`
	Error string `json:"error,omitempty"`
}

// UpgradeFileInfo upgrade file info
type UpgradeFileInfo struct {
	Target string `json:"target"`
	URL    string `json:"url"`
	Sha256 []byte `json:"sha256"`
	Sha512 []byte `json:"sha512"`
	Size   uint64 `json:"size"`
}

// UpgradeReq system upgrade request
type UpgradeReq struct {
	MessageHeader
	ImageVersion uint64            `json:"imageVersion"`
	Files        []UpgradeFileInfo `json:"files"`
}

// RevertReq system revert request
type RevertReq struct {
	MessageHeader
	ImageVersion uint64 `json:"imageVersion"`
}

// GetStatusReq get system status request
type GetStatusReq struct {
	MessageHeader
}

// StatusMessage status message
type StatusMessage struct {
	MessageHeader
	Operation        string `json:"operation"` // upgrade, revert
	Status           string `json:"status"`    // success, failed, inprogress
	OperationVersion uint64 `json:"operationVersion"`
	ImageVersion     uint64 `json:"imageVersion"`
}
