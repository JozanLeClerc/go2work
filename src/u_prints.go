/*
 * ========================
 * =====    ===============
 * ======  ================
 * ======  ================
 * ======  ====   ====   ==
 * ======  ===     ==  =  =
 * ======  ===  =  ==     =
 * =  ===  ===  =  ==  ====
 * =  ===  ===  =  ==  =  =
 * ==     =====   ====   ==
 * ========================
 *
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 2022 Joe
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the organization nor the
 *    names of its contributors may be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY JOE ''AS IS'' AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL JOE BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * go2work: src/u_prints.go
 * Mon Apr  4 19:27:33 CEST 2022
 * Joe
 *
 * Stuff to print
 */

package main

import (
	"fmt"
	"log"
)

func print_time(t [3]byte) {
	fmt.Print("\rTime is: ", t[HOURS], ":", t[MINS], ":", t[SECS])
}

func print_time_left(curr_t [3]byte, dest_t [3]byte) {
	var left_secs uint
	curr_t = get_time()
	curr_secs := time_to_seconds(curr_t)
	dest_secs := time_to_seconds(dest_t)
	if curr_secs <= dest_secs {
		left_secs = dest_secs - curr_secs
	} else {
		left_secs = (dest_secs + (24 * 3600)) - curr_secs
	}
	left_t := seconds_to_time(left_secs)
	fmt.Print("\r                         ")
	if left_secs < 60 {
		fmt.Print(
			"\r",
			left_t[SECS], "s",
			" left to sleep",
		)
	} else if left_secs < 3600 {
		fmt.Print(
			"\r",
			left_t[MINS], "m ",
			left_t[SECS], "s",
			" left to sleep",
		)
	} else {
		fmt.Print(
			"\r",
			left_t[HOURS], "h ",
			left_t[MINS],  "m ",
			left_t[SECS],  "s",
			" left to sleep",
		)
	}
}

func print_help() {
	fmt.Println(`Usage:
  go2work [options] time

Options:
  -h, --help         show this help menu
  -H, --real-help    show the real help menu
  -t, --test         run a test
  -v, --version      show version of go2work`)
}

func print_real_help() {
	fmt.Println("send help")
}

func print_version() {
	fmt.Println(PROGNAME, VERSION)
}

func print_fortune_not_found() {
	log.Println("beware, fortune is set on but was not found")
}

func print_no_files() {
	fmt.Print("\n\n")
	log.Println("beware, no files selected! Using default ringtone")
}
func print_file_not_found(file string) {
	fmt.Print("\n\n")
	log.Println("beware, file '" +
		file +
		"' not found! Using default ringtone",
	)
}
