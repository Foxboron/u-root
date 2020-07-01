// Copyright 2015-2020 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// TODO: make templates able to include other templates.
// e.g. "all" below should just say "core" and "boot". Use it to replace
// the old 'systemboot' template.
// Or just call it a day, now that we have the new directory structure, and dump the templates
// completely; that may be our best bet.
var templates = map[string][]string{
	"all": {
		"github.com/u-root/u-root/cmds/core/*",
		"github.com/u-root/u-root/cmds/boot/*boot*",
	},
	"boot": {
		"github.com/u-root/u-root/cmds/boot/*boot*",
	},
	// Core should be things you don't want to live without.
	"core": {
		"github.com/u-root/u-root/cmds/core/*",
	},
	// Minimal should be things you can't live without.
	"minimal": {
		"github.com/u-root/u-root/cmds/core/cat",
		"github.com/u-root/u-root/cmds/core/chmod",
		"github.com/u-root/u-root/cmds/core/cmp",
		"github.com/u-root/u-root/cmds/core/cp",
		"github.com/u-root/u-root/cmds/core/date",
		"github.com/u-root/u-root/cmds/core/dd",
		"github.com/u-root/u-root/cmds/core/df",
		"github.com/u-root/u-root/cmds/core/dhclient",
		"github.com/u-root/u-root/cmds/core/dmesg",
		"github.com/u-root/u-root/cmds/core/echo",
		"github.com/u-root/u-root/cmds/core/elvish",
		"github.com/u-root/u-root/cmds/core/find",
		"github.com/u-root/u-root/cmds/core/free",
		"github.com/u-root/u-root/cmds/core/gpgv",
		"github.com/u-root/u-root/cmds/core/grep",
		"github.com/u-root/u-root/cmds/core/gzip",
		"github.com/u-root/u-root/cmds/core/hostname",
		"github.com/u-root/u-root/cmds/core/id",
		"github.com/u-root/u-root/cmds/core/init",
		"github.com/u-root/u-root/cmds/core/insmod",
		"github.com/u-root/u-root/cmds/core/installcommand",
		"github.com/u-root/u-root/cmds/core/io",
		"github.com/u-root/u-root/cmds/core/ip",
		"github.com/u-root/u-root/cmds/core/kexec",
		"github.com/u-root/u-root/cmds/core/kill",
		"github.com/u-root/u-root/cmds/core/ln",
		"github.com/u-root/u-root/cmds/core/losetup",
		"github.com/u-root/u-root/cmds/core/ls",
		"github.com/u-root/u-root/cmds/core/lsmod",
		"github.com/u-root/u-root/cmds/core/mkdir",
		"github.com/u-root/u-root/cmds/core/mknod",
		"github.com/u-root/u-root/cmds/core/mount",
		"github.com/u-root/u-root/cmds/core/msr",
		"github.com/u-root/u-root/cmds/core/mv",
		"github.com/u-root/u-root/cmds/core/pci",
		"github.com/u-root/u-root/cmds/core/ping",
		"github.com/u-root/u-root/cmds/core/printenv",
		"github.com/u-root/u-root/cmds/core/ps",
		"github.com/u-root/u-root/cmds/core/pwd",
		"github.com/u-root/u-root/cmds/core/readlink",
		"github.com/u-root/u-root/cmds/core/rm",
		"github.com/u-root/u-root/cmds/core/rmmod",
		"github.com/u-root/u-root/cmds/core/seq",
		"github.com/u-root/u-root/cmds/core/shutdown",
		"github.com/u-root/u-root/cmds/core/sleep",
		"github.com/u-root/u-root/cmds/core/sync",
		"github.com/u-root/u-root/cmds/core/tail",
		"github.com/u-root/u-root/cmds/core/tee",
		"github.com/u-root/u-root/cmds/core/truncate",
		"github.com/u-root/u-root/cmds/core/umount",
		"github.com/u-root/u-root/cmds/core/uname",
		"github.com/u-root/u-root/cmds/core/unshare",
		"github.com/u-root/u-root/cmds/core/wc",
		"github.com/u-root/u-root/cmds/core/wget",
		"github.com/u-root/u-root/cmds/core/which",
	},
	// coreboot-app minimal environment
	"coreboot-app": {
		"github.com/u-root/u-root/cmds/core/cat",
		"github.com/u-root/u-root/cmds/exp/cbmem",
		"github.com/u-root/u-root/cmds/core/chroot",
		"github.com/u-root/u-root/cmds/core/cp",
		"github.com/u-root/u-root/cmds/core/dd",
		"github.com/u-root/u-root/cmds/core/dhclient",
		"github.com/u-root/u-root/cmds/core/dmesg",
		"github.com/u-root/u-root/cmds/core/elvish",
		"github.com/u-root/u-root/cmds/core/find",
		"github.com/u-root/u-root/cmds/core/grep",
		"github.com/u-root/u-root/cmds/core/id",
		"github.com/u-root/u-root/cmds/core/init",
		"github.com/u-root/u-root/cmds/core/insmod",
		"github.com/u-root/u-root/cmds/core/ip",
		"github.com/u-root/u-root/cmds/core/kill",
		"github.com/u-root/u-root/cmds/core/ls",
		"github.com/u-root/u-root/cmds/core/mount",
		"github.com/u-root/u-root/cmds/core/pci",
		"github.com/u-root/u-root/cmds/core/ping",
		"github.com/u-root/u-root/cmds/core/ps",
		"github.com/u-root/u-root/cmds/core/pwd",
		"github.com/u-root/u-root/cmds/core/rm",
		"github.com/u-root/u-root/cmds/core/rmmod",
		"github.com/u-root/u-root/cmds/core/shutdown",
		"github.com/u-root/u-root/cmds/core/sshd",
		"github.com/u-root/u-root/cmds/core/switch_root",
		"github.com/u-root/u-root/cmds/core/tail",
		"github.com/u-root/u-root/cmds/core/tee",
		"github.com/u-root/u-root/cmds/core/uname",
		"github.com/u-root/u-root/cmds/core/wget",
	},
	"plan9": {
		"github.com/u-root/u-root/cmds/core/init",
		"github.com/u-root/u-root/cmds/exp/rush",
	},
}
