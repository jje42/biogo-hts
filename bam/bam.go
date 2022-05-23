// Copyright ©2012 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bam implements BAM file format reading, writing and indexing.
// The BAM format is described in the SAM specification.
//
// http://samtools.github.io/hts-specs/SAMv1.pdf
package bam

import (
	"errors"
	"fmt"
	"os"
)

func GetSampleName(fn string) (string, error) {
	r, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %s", err)
	}
	defer r.Close()

	br, err := NewReader(r, 1)
	if err != nil {
		return "", fmt.Errorf("unable to create reader: %s", err)
	}
	defer br.Close()

	rgs := br.Header().RGs()
	if len(rgs) == 0 {
		return "", errors.New("BAM has no read groups")
	}
	sm := ""
	for _, rg := range rgs {
		if sm == "" {
			sm = rg.SM()
		}
		if sm != rg.SM() {
			return "", errors.New("non-identical SM tags detected in read groups")
		}
	}
	return sm, nil
}
