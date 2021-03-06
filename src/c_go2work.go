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
 * go2work: src/c_go2work.go
 * Tue Apr  5 11:11:13 CEST 2022
 * Joe
 *
 * The main.
 */

package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var curr_t [3]byte
	var dest_t [3]byte
	var tmp int
	var options Options
	log.SetPrefix(PROGNAME + ": ")
	log.SetFlags(0)
	if len(os.Args) == 1 {
		print_help()
		os.Exit(1)
		return
	}
	switch os.Args[1] {
	case "-h", "--help":
		print_help()
		return
	case "-H", "--real-help":
		print_real_help()
		return
	case "-v", "--version":
		print_version()
		return
	case "-t", "--test":
		dest_t = get_test_time()
	default:
		if strings.Contains(os.Args[1], ":") == false {
			log.Fatal(LOG_FORMAT)
		}
		str_dest_t := strings.Split(os.Args[1], ":")
		tmp, _ = strconv.Atoi(str_dest_t[HOURS])
		dest_t[HOURS] = byte(tmp)
		tmp, _ = strconv.Atoi(str_dest_t[MINS])
		dest_t[MINS] = byte(tmp)
		dest_t[SECS] = 0
	}
	options = parse_options()
	first_checks(dest_t, options)
	curr_t = get_time()
	print_time_left(curr_t, dest_t)
	main_loop(dest_t, options)
}
