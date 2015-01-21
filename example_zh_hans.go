// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build zh_hans

package log

// abdeeffffffffffffffffff fmt.go:189
// const desc
const (
	// const desc
	item  = iota // item
	item2        // item2
	item3        // item3
)

// abdeeffffffffffffffffff fmt.go:199
// logger desc
type Logger struct {
	Flag int // flag
}

// abdeeffffffffffffffffff fmt.go:110
// new logger
func NewLogger(o int) error

// abdeeffffffffffffffffff fmt.go:120
// logger.close desc
func (l *Logger) Close() error
