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
 * go2work: src/c_defs.go
 * Fri Apr  1 17:11:48 CEST 2022
 * Joe
 *
 * Definitions.
 */

package main

const (
	PROGNAME			= "go2work"
	VERSION				= "0.1.0"
	HOURS				= 0
	MINS				= 1
	SECS				= 2
	INTERVAL			= 500
	LOG_FORMAT			= "bad time format"
	DEF_MEDIA_PLAYER	= "mpv"
	DEF_RANDOM			= true
	DEF_USE_FORTUNE		= true
)

func DEF_FILES() []string {
	return []string{
		"/usr/home/jozan/mu/progressive/progressive_black_metal/deathspell_omega/2010_paracletus/02_wings_of_predation.flac",
		"/usr/home/jozan/docs/dr2_art/sounds/sfx/broking_cursed_object.wav",
	}
}

func DEF_PLAYER_OPTIONS() []string {
	return []string{
		"--no-video",
		"--loop",
	}
}

type Options struct {
	files			[]string
	media_player	string
	player_options	[]string
	random			bool
	use_fortune		bool
}
