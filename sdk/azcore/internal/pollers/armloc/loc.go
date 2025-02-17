//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package armloc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pollers"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

// Kind is the identifier of this type in a resume token.
const Kind = "ARM-Location"

// Applicable returns true if the LRO is using Location.
func Applicable(resp *http.Response) bool {
	return resp.StatusCode == http.StatusAccepted && resp.Header.Get(shared.HeaderLocation) != ""
}

// Poller is an LRO poller that uses the Location pattern.
type Poller struct {
	// The poller's type, used for resume token processing.
	Type string `json:"type"`

	// The URL for polling.
	PollURL string `json:"pollURL"`

	// The LRO's current state.
	CurState string `json:"state"`
}

// New creates a new Poller from the provided initial response.
func New(resp *http.Response, pollerID string) (*Poller, error) {
	log.Write(log.EventLRO, "Using Location poller.")
	locURL := resp.Header.Get(shared.HeaderLocation)
	if locURL == "" {
		return nil, errors.New("response is missing Location header")
	}
	if !pollers.IsValidURL(locURL) {
		return nil, fmt.Errorf("invalid polling URL %s", locURL)
	}
	p := &Poller{
		Type:     pollers.MakeID(pollerID, Kind),
		PollURL:  locURL,
		CurState: pollers.StatusInProgress,
	}
	return p, nil
}

// URL returns the polling URL.
func (p *Poller) URL() string {
	return p.PollURL
}

// State returns the current state of the LRO.
func (p *Poller) State() pollers.OperationState {
	if pollers.Succeeded(p.CurState) {
		return pollers.OperationStateSucceeded
	} else if pollers.IsTerminalState(p.CurState) {
		return pollers.OperationStateFailed
	}
	return pollers.OperationStateInProgress
}

// Update updates the Poller from the polling response.
func (p *Poller) Update(resp *http.Response) error {
	// location polling can return an updated polling URL
	if h := resp.Header.Get(shared.HeaderLocation); h != "" {
		p.PollURL = h
	}
	if exported.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		// if a 200/201 returns a provisioning state, use that instead
		state, err := pollers.GetProvisioningState(resp)
		if err != nil && !errors.Is(err, pollers.ErrNoBody) {
			return err
		}
		if state != "" {
			p.CurState = state
		} else {
			// a 200/201 with no provisioning state indicates success
			p.CurState = pollers.StatusSucceeded
		}
	} else if resp.StatusCode == http.StatusNoContent {
		p.CurState = pollers.StatusSucceeded
	} else if resp.StatusCode > 399 && resp.StatusCode < 500 {
		p.CurState = pollers.StatusFailed
	}
	// a 202 falls through, means the LRO is still in progress and we don't check for provisioning state
	return nil
}

// FinalGetURL returns the empty string as no final GET is required for this poller type.
func (p *Poller) FinalGetURL() string {
	return ""
}
