// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package atom

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

// CloseRes closes the response (or if it's nil just no-ops)
func CloseRes(ctx context.Context, res *http.Response) {
	if res == nil {
		return
	}

	_, _ = io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()
}
