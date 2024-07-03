// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

package tls

import "github.com/zirngibl/qscanner-tls/internal/boring/fipstls"

// needFIPS returns fipstls.Required(), which is not available without the
// boringcrypto build tag.
func needFIPS() bool {
	return fipstls.Required()
}
