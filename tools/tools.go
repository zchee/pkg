// SPDX-FileCopyrightText: Copyright 2021 The pkg Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
