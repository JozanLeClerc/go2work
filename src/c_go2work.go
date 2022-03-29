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
 * Tue Mar 29 20:10:23 CEST 2022
 * Joe
 *
 * The main.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"
)

const (
	PROGNAME = "go2work"
)

func main() {
	log.SetPrefix(PROGNAME + ": ")
	log.SetFlags(0)
	if len(os.Args[0:]) == 1 {
		log.Fatal("No arguments")
		return
	}
	curr_h, curr_m := get_time()
	dest_t := os.Args[1]
	dest_h := strings.Split(dest_t, ":")
	fmt.Println("dest_splitted: ", dest_h)
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	fmt.Println("Time is: " + curr_h + ":" + curr_m)
	for {
		select {
		case <- ticker.C:
			curr_h, curr_m = get_time()
			exec_player(false, "mpv", "--no-video", "/home/jozan/mu/rock/grunge/nirvana/1993_in_utero/04_rape_me.flac")
		case <- quit:
			ticker.Stop()
			return
		}
	}
}

func get_time() (string, string) {
	var h  string
	var m  string
	t := time.Now()
	h  = t.Format("15")
	m  = t.Format("04")
	return h, m
}
