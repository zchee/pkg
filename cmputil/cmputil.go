// SPDX-FileCopyrightText: Copyright 2021 The pkg Authors
// SPDX-License-Identifier: BSD-3-Clause

// copy and edit from https://play.golang.org/p/PKGzR3fXs-P.

package cmputil

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// transformJSON transforms any Go string that looks like JSON into
// a generic data structure that represents that JSON input.
// We use an AcyclicTransformer to avoid having the transformer
// apply on outputs of itself (lest we get stuck in infinite recursion).
func TransformJSON() cmp.Option {
	return cmpopts.AcyclicTransformer("TransformJSON", func(s string) interface{} {
		var v interface{}
		if err := json.Unmarshal([]byte(s), &v); err != nil {
			return s // use unparseable input as the output
		}
		return v
	})
}

var iso8601Regex = regexp.MustCompile(`^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(\\.[0-9]+)?(Z)?$`)

// IgnoreISO8601 ignores any Go strings that matches the ISO8601 format.
func IgnoreISO8601() cmp.Option {
	return cmp.FilterValues(func(x, y string) bool {
		return iso8601Regex.MatchString(x) && iso8601Regex.MatchString(y)
	}, cmp.Ignore())

	if diff := cmp.Diff(want, got, transformJSON, ignoreISO8601); diff != "" {
		fmt.Printf("mismatch (-want +got):\n%v", diff)
	}
}
