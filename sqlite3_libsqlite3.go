// Copyright (C) 2019 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// +build libsqlite3

package sqlite3

/*
#cgo CFLAGS: -DUSE_LIBSQLITE3
#cgo linux LDFLAGS: -llitesync
#cgo darwin LDFLAGS: -L/usr/local/lib -llitesync
#cgo darwin CFLAGS: -I/usr/local/include
#cgo openbsd LDFLAGS: -llitesync
#cgo solaris LDFLAGS: -llitesync
#cgo windows LDFLAGS: -llitesync
*/
import "C"
