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
 * go2work: src/u_checks.go
 * Tue Apr  5 11:13:23 CEST 2022
 * Joe
 *
 * Useful checks.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func check_media_player(media_player string) bool {
	cmd := exec.Command("command", "-v", media_player)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func check_fortune() bool {
	cmd := exec.Command("command", "-v", FORTUNE_BIN)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func check_time_format(time [3]byte) bool {
	if time[HOURS] > 23 {
		return false
	}
	if time[MINS] > 59 {
		return false
	}
	return true
}

func check_file_exists(file string) bool {
	if file == "" {
		fmt.Println("empty")
	}
	_, err := os.Stat(file)
	if os.IsNotExist(err) == true {
		return false
	}
	return true
}

func first_checks(dest_t [3]byte, options Options) {
	if check_time_format(dest_t) == false {
		log.Fatal(LOG_FORMAT)
		return
	}
	if check_media_player(options.Media_player) == false {
		log.Fatal("media player (" + options.Media_player + ") not found")
		return
	}
	if options.Fortune == true && check_fortune() == false {
		print_fortune_not_found()
	}
}
