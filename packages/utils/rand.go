/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package utils

import (
	"math/rand"
	"time"

	"github.com/IBAX-io/go-ibax/packages/crypto"
)

type Rand struct {
	src *rand.Rand
}

func (r *Rand) BytesSeed(b []byte) *rand.Rand {
	seed, _ := crypto.CalcChecksum(b)
	r.src.Seed(int64(seed))
	return r.src
}

func NewRand(seed int64) *Rand {
	return &Rand{
		src: rand.New(rand.NewSource(seed)),
	}
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

const (
	KC_RAND_KIND_NUM   = 0 // number
	KC_RAND_KIND_LOWER = 1 //
	KC_RAND_KIND_UPPER = 2 //
	KC_RAND_KIND_ALL   = 3 //
)

//
func Krand(size int64, kind int) []byte {
	return result
}

//
func RandNumber(size int64) string {
	result := Krand(size, KC_RAND_KIND_ALL)
	return string(result[:])
}