# ========================
# =====    ===============
# ======  ================
# ======  ================
# ======  ====   ====   ==
# ======  ===     ==  =  =
# ======  ===  =  ==     =
# =  ===  ===  =  ==  ====
# =  ===  ===  =  ==  =  =
# ==     =====   ====   ==
# ========================
#
# go2work: Makefile
# Fri Apr  1 17:04:04 CEST 2022
# Joe
#
# GNU Makefile

default: build

SHELL		:= /bin/sh
OS			 = $(shell uname)

SRCS_DIR	 = src/
DATA_DIR	 = share/
TRGT_DIR	 = ./

SRCS_NAME	 = c_defs
SRCS_NAME	+= c_go2work
SRCS_NAME	+= c_player
SRCS_NAME	+= p_options
SRCS_NAME	+= u_checks
SRCS_NAME	+= u_prints
SRCS_NAME	+= u_utils

SRCS		 = $(addprefix ${SRCS_DIR}, $(addsuffix .go, ${SRCS_NAME}))

TRGT_NAME	 = go2work
TARGET		 = $(addprefix ${TRGT_DIR}, ${TRGT_NAME})

MKDIR		 = mkdir -p
RMDIR		 = rmdir
RM			 = rm -rf

build:
	go get github.com/BurntSushi/toml@latest
	go build -o ${TARGET} ${SRCS}

clean:
	${RM} ${TARGET}

run:
	go run ${SRCS}

.PHONY:	build clean run

# File prefixes info
# ------------------
# c_  -> core program
# p_  -> parsing
# u_  -> utils
