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
# Thu Mar 24 15:00:55 CET 2022
# Joe
#
# GNU Makefile

default: run

SHELL		:= /bin/sh
OS			 = $(shell uname)

SRCS_DIR	 = src/
DATA_DIR	 = share/
TRGT_DIR	 = ./

SRCS_NAME	 = c_go2work
SRCS_NAME	+= c_player

SRCS		 = $(addprefix ${SRCS_DIR}, $(addsuffix .go, ${SRCS_NAME}))

TRGT_NAME	 = go2work
TARGET		 = $(addprefix ${TRGT_DIR}, ${TRGT_NAME})

MKDIR		 = mkdir -p
RMDIR		 = rmdir
RM			 = rm -rf

build:
	go build -o ${TARGET} ${SRCS}

clean:
	${RM} ${TARGET}

run:
	go run ${SRCS}

tools:
	${MAKE} --no-print-directory -C ${TOOLS_DIR} all

.PHONY:	build run

# File prefixes info
# ------------------
# c_  -> core program related
