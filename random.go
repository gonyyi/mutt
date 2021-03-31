// (C) 2021 GON Y YI.
// https://gonyyi.com/copyright.txt

package mutt

import (
	"math/rand"
	"sync"
	"time"
)

const (
	ALPHA              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMERIC            = "0123456789"
	ALPHANUMERIC       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ALPHANUMERIC_LOWER = "abcdefghijklmnopqrstuvwxyz0123456789"
	ALPHANUMERIC_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	HEX                = "0123456789ABCDEF"
)

var std = NewRandom()

// ==================================================================== NEW RANDOM
func NewRandom() *random {
	DefaultStr := []byte(ALPHANUMERIC)
	r := &random{
		randBStr:     DefaultStr,
		randBStrSize: len(DefaultStr),
	}
	r.UseMutex(true)
	return r
}

type random struct {
	rand         *rand.Rand
	randBStr     []byte
	randBStrSize int
	isMutex      bool
}

// SetString will set default string to be used for random
func (r *random) SetString(s string) {
	r.randBStr = []byte(s)
	r.randBStrSize = len(r.randBStr)
}

// UseMutex will take boolean value; true will make mutex available. This will reduce performance, however
// this will make job run safely.
func (r *random) UseMutex(tf bool) {
	if tf == true {
		r.rand = rand.New(&randLockSrc{
			src: rand.NewSource(time.Now().UnixNano()),
		})
		r.isMutex = true
	} else {
		r.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		r.isMutex = false
	}
}

// UpdateSeed will update the seed, however be caucious when run in goroutine.
// Use UseMutex(true) to avoid possible panic.
func (r *random) UpdateSeed() (ok bool) {
	if r.isMutex == true {
		r.rand.Seed(time.Now().UnixNano())
		return true
	}
	return false
}

// RandomInt retures int between min, max.
func (r *random) RandInt(min, max int) int {
	// if equals, returns input
	if min == max {
		return min
	} else if min > max {
		t := max
		max = min
		min = t
	} else {
		max = max + 1
	}
	return r.rand.Intn(max-min) + min
}

// RandomInt64 returns int64 between min, max.
func (r *random) RandInt64(min, max int64) int64 {
	// if equals, return whatever it is
	if min == max {
		return min
	} else if min > max {
		t := max
		max = min
		min = t
	} else {
		max = max + 1
	}
	return r.rand.Int63n(max-min) + min
}

// RandomInt64 returns int64
func (r *random) Int64(n int64) int64 {
	return r.rand.Int63n(n)
}

// RandomString will take string and length and returns the randomized value within the string.
// This has better performance than current mathd.NewRand ...
func (r *random) RandStr(b []byte, length int) string {
	var ls = len(b)
	if ls == 0 || length < 1 {
		return ""
	}
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		out[i] = b[r.rand.Intn(ls)]
	}
	return string(out[:])
}

// Random creates alphanumeric random string with length
func (r *random) Rand(length int) string {
	if length > 0 {
		out := make([]byte, length)
		for i := 0; i < length; i++ {
			out[i] = r.randBStr[r.rand.Intn(r.randBStrSize)] // len of randBytesAlphanumeric = 62
		}
		return string(out[:])
	}
	return ""
}

// ==================================================================== RAND LOCK
type randLockSrc struct {
	lock sync.Mutex // there is no benefit of using RMMutex as all values are read only once.. (random support to return diff stuff each time)
	src  rand.Source
}

func (r *randLockSrc) Int63() int64 { // to satisfy rand.Source interface
	r.lock.Lock()
	t := r.src.Int63() // this is faster than using defer
	r.lock.Unlock()
	return t
}
func (r *randLockSrc) Seed(seed int64) { // to satisfy rand.Source interface
	r.lock.Lock()
	r.src.Seed(seed)
	r.lock.Unlock()
}

// ==================================================================== STANDARD

// UpdateSeed will update the seed, however be caucious when run in goroutine.
// Use UseMutex(true) to avoid possible panic.
func RandomUpdateSeed() bool {
	return std.UpdateSeed()
}

// RandomInt will return int between min, max
func RandomInt(min, max int) int {
	return std.RandInt(min, max)
}

func RandomInt64(min, max int64) int64 {
	return std.RandInt64(min, max)
}

// RandomString will take string and length and returns the randomized value within the string.
// This has better performance than current mathd.NewRand ...
func RandomString(b []byte, length int) string {
	return std.RandStr(b, length)
}

// Random creates alphanumeric random string with length
func Random(length int) string {
	return std.Rand(length)
}

