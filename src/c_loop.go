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
 * go2work: src/c_loop.go
 * Mon Apr  4 17:28:38 CEST 2022
 * Joe
 *
 * The main loop
 */

package main

import(
	"time"
)

func main_loop(dest_t [3]byte, options Options) {
	var curr_t [3]byte
	ticker := time.NewTicker(INTERVAL * time.Millisecond)
	quit := make(chan struct{})
	for {
		select {
		case <- ticker.C:
			curr_t = get_time()
			print_time_left(curr_t, dest_t)
			if curr_t[HOURS] == dest_t[HOURS] &&
				curr_t[MINS] == dest_t[MINS] &&
				curr_t[SECS] == dest_t[SECS] {
				file_id := choose_file(options)
				var args []string
				if file_id >= 0 {
					args = append(
						options.Player_options,
						options.Files[file_id],
					)
				} else {
					args = append(
						options.Player_options,
						DEF_FILES()[0],
					)
				}
				has_rang := false
				for {
					exec_player(
						options.Fortune,
						has_rang,
						options.Media_player,
						args...,
					)
					has_rang = true
				}
			}
		case <- quit:
			ticker.Stop()
			return
		}
	}
}
