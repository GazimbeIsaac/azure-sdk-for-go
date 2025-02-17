//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package loc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pollers"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

// Kind is the identifier of this type in a resume token.
const Kind = "Location"

// Applicable returns true if the LRO is using Location.
func Applicable(resp *http.Response) bool {
	return resp.Header.Get(shared.HeaderLocation) != ""
}

// Poller is an LRO poller that uses the Location pattern.
type Poller struct {
	Type     string `json:"type"`
	PollURL  string `json:"pollURL"`
	FinalGET string `json:"finalGET"`
	CurState int    `json:"state"`
}

// New creates a new Poller from the provided initial response.
func New(resp *http.Response, finalState pollers.FinalStateVia, pollerID string) (*Poller, error) {
	log.Write(log.EventLRO, "Using Location poller.")
	locURL := resp.Header.Get(shared.HeaderLocation)
	if locURL == "" {
		return nil, errors.New("response is missing Location header")
	}
	if !pollers.IsValidURL(locURL) {
		return nil, fmt.Errorf("invalid polling URL %s", locURL)
	}
	// TODO: calculate the final GET URL.  can be empty
	finalGET := ""
	return &Poller{
		Type:     pollers.MakeID(pollerID, Kind),
		PollURL:  locURL,
		FinalGET: finalGET,
		CurState: resp.StatusCode,
	}, nil
}

func (p *Poller) URL() string {
	return p.PollURL
}

func (p *Poller) State() pollers.OperationState {
	if p.CurState == http.StatusAccepted {
		return pollers.OperationStateInProgress
	} else if p.CurState > 199 && p.CurState < 300 {
		// any 2xx other than a 202 indicates success
		return pollers.OperationStateSucceeded
	}
	return pollers.OperationStateFailed
}

func (p *Poller) Update(resp *http.Response) error {
	// if the endpoint returned a location header, update cached value
	if loc := resp.Header.Get(shared.HeaderLocation); loc != "" {
		p.PollURL = loc
	}
	p.CurState = resp.StatusCode
	return nil
}

func (p *Poller) FinalGetURL() string {
	return p.FinalGET
}
