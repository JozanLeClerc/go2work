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
 * Tue Mar 29 22:15:02 CEST 2022
 * Joe
 *
 * Stuff to print
 */

package main

import (
	"fmt"
	"strconv"
)

func print_time(t []string) {
	fmt.Print("\rTime is: ", t[HOURS], ":", t[MINS], ":", t[SECS])
}

func print_time_left(curr_t []int, dest_t []int) {
	var left_t [2]int
	left_t = int{0, 0}
	h, _ := strconv.Atoi(curr)
	m
	s, _ := strconv.Atoi(curr_t[SECS])
	s = 60 - s
	fmt.Print(
		"\rTime left to sleep: ",
		left_t[HOURS], "h ",
		left_t[MINS], "m ",
		s, "s",
	)
}

func print_help() {
	fmt.Println("help")
}

func print_real_help() {
	fmt.Println("send help")
}

func print_version() {
	fmt.Println(PROGNAME, VERSION)
}
