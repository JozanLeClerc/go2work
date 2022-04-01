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
 * go2work: src/p_options.go
 * Fri Apr  1 21:13:50 CEST 2022
 * Joe
 *
 * Options parsing.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"github.com/BurntSushi/toml"
)

func parse_options() Options {
	options := init_def_options()
	options_file := find_options_file()
	if options_file == "" {
		return options
	}
	options = parse_toml_file(options_file, options)
	return options
}

func init_def_options() Options {
	options := Options{
		files:			DEF_FILES(),
		media_player:	DEF_MEDIA_PLAYER,
		player_options: DEF_PLAYER_OPTIONS(),
		random:			DEF_RANDOM,
		use_fortune:	DEF_USE_FORTUNE,
	}
	return options
}

func find_options_file() string {
	val, exists := os.LookupEnv("XDG_CONFIG_HOME")
	var file_path string
	if exists == true {
		file_path = val + "/" + OPTIONS_FILE
		if check_file_exists(file_path) == true {
			return file_path
		}
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file_path = dir + "/.config/" + OPTIONS_FILE
	if check_file_exists(file_path) == true {
		return file_path
	}
	return ""
}

func parse_toml_file(options_file string, def_options Options) Options {
	options := def_options
	fmt.Println(options_file)
	_, err := toml.DecodeFile(options_file, &options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OPTIONS")
	fmt.Println("=======")
	fmt.Println(options.files)
	fmt.Println(options.media_player)
	fmt.Println(options.player_options)
	fmt.Println(options.random)
	fmt.Println(options.use_fortune)
	fmt.Println("=======\n")
	return options
}
