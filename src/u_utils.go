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
 * go2work: src/u_utils.go
 * Wed Mar 30 01:25:20 CEST 2022
 * Joe
 */

package main

import (
	"time"
	"strconv"
	"strings"
)

func time_to_seconds(time [3]byte) uint {
	return (3600 * uint(time[HOURS])) +
		(60 * uint(time[MINS])) +
		uint(time[SECS])
}

func seconds_to_time(seconds uint) [3]byte {
	var time [3]byte
	var hours, mins uint
	hours = seconds / 3600
	time[HOURS] = byte(hours)
	mins = (seconds - ((hours) * 3600)) / 60
	time[MINS] = byte(mins)
	time[SECS] = byte((seconds - (hours * 3600)) - mins * 60)
	return time
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

func get_test_time() [3]byte {
	return seconds_to_time(time_to_seconds(get_time()) + 3)
}
