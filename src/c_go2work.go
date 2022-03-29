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
 * Tue Mar 29 22:32:07 CEST 2022
 * Joe
 *
 * The main.
 */

package main

import (
	"log"
	"os"
	"time"
	"strconv"
	"strings"
)

const (
	PROGNAME	= "go2work"
	VERSION		= "0.1.0"
	HOURS		= 0
	MINS		= 1
	SECS		= 2
)

func main() {
	var dest_t [3]byte
	var tmp int
	log.SetPrefix(PROGNAME + ": ")
	log.SetFlags(0)
	if len(os.Args[0:]) == 1 {
		log.Fatal("No arguments")
		return
	}
	switch os.Args[1] {
	case "-h":
		print_help()
		return
	case "-H":
		print_real_help()
		return
	case "-v":
		print_version()
		return
	}
	curr_t := get_time()
	str_dest_t := strings.Split(os.Args[1], ":")
	tmp, _ = strconv.Atoi(str_dest_t[HOURS])
	dest_t[HOURS] = byte(tmp)
	tmp, _ = strconv.Atoi(str_dest_t[MINS])
	dest_t[MINS] = byte(tmp)
	dest_t[SECS] = 0
	// dest_t = [2]int{0, 0}
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <- ticker.C:
			curr_t = get_time()
			// print_time(curr_t)
			print_time_left(curr_t, dest_t)
			if curr_t[HOURS] == dest_t[HOURS] && curr_t[MINS] == dest_t[MINS] {
				exec_player(
					true,
					"mpv",
					"--no-video",
					"/usr/home/jozan/mu/progressive/progressive_black_metal/deathspell_omega/2010_paracletus/02_wings_of_predation.flac",
				)
				return
			}
		case <- quit:
			ticker.Stop()
			return
		}
	}
}

func get_time() [3]byte {
	var curr_t [3]byte
	var tmp int
	now := time.Now()
	t := strings.Split(now.Format("15:04:05"), ":")
	tmp, _ = strconv.Atoi(t[HOURS])
	curr_t[HOURS] = byte(tmp)
	tmp, _ = strconv.Atoi(t[MINS])
	curr_t[MINS] = byte(tmp)
	tmp, _ = strconv.Atoi(t[SECS])
	curr_t[SECS] = byte(tmp)
	return curr_t
}
