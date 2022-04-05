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
# Mon Apr  4 22:51:40 CEST 2022
# Joe
#
# GNU Makefile

default: build

SHELL		:= /bin/sh
OS			:= $(shell uname)

DESTDIR		:= /usr/local/
SRCS_DIR	:= src/
DATA_DIR	:= share/
MAN_DIR		:= man/

SRCS_NAME	:= c_defs
SRCS_NAME	+= c_go2work
SRCS_NAME	+= c_loop
SRCS_NAME	+= c_player
SRCS_NAME	+= p_options
SRCS_NAME	+= u_checks
SRCS_NAME	+= u_prints
SRCS_NAME	+= u_utils

SRCS		:= $(addprefix ${SRCS_DIR}, $(addsuffix .go, ${SRCS_NAME}))

TARGET		:= go2work
MAN			:= $(addsuffix .1, ${TARGET})

MKDIR		:= mkdir -p
RMDIR		:= rmdir
RM			:= rm -f
RMDIR		:= rmdir
GZIP		:= gzip
GUNZIP		:= gunzip
INSTALL		:= install

deps:
	go get github.com/BurntSushi/toml@latest

build: ${TARGET}

${TARGET}: deps
	go build -o ${TARGET} ${SRCS}

install-bin: ${TARGET}
	${MKDIR} ${DESTDIR}bin
	${INSTALL} -m0555 ${TARGET} ${DESTDIR}bin

install-data:
	${MKDIR} ${DESTDIR}share/${TARGET}
	${INSTALL} -m0644 ${DATA_DIR}ring01.wav ${DESTDIR}share/${TARGET}
	${INSTALL} -m0644 ${DATA_DIR}ring02.wav ${DESTDIR}share/${TARGET}
	${INSTALL} -m0644 LICENSE ${DESTDIR}share/${TARGET}
	${INSTALL} -m0644 README ${DESTDIR}share/${TARGET}
	${MKDIR} ${DESTDIR}etc/${TARGET}
	${INSTALL} -m0644 ${DATA_DIR}${TARGET}.toml ${DESTDIR}etc/${TARGET}

install-doc:
	${MKDIR} ${DESTDIR}man/man1
	${GZIP} ${MAN_DIR}${MAN}
	${INSTALL} -m0444 ${MAN_DIR}${MAN}.gz ${DESTDIR}man/man1
	${GUNZIP} ${MAN_DIR}${MAN}

install: install-bin install-data install-doc

uninstall:
	${RM} ${DESTDIR}bin/${TARGET}
	${RM} ${DESTDIR}share/${TARGET}/ring01.wav
	${RM} ${DESTDIR}share/${TARGET}/ring02.wav
	${RM} ${DESTDIR}share/${TARGET}/LICENSE
	${RM} ${DESTDIR}share/${TARGET}/README
	${RMDIR} ${DESTDIR}share/${TARGET}
	${RM} ${DESTDIR}etc/${TARGET}/${TARGET}.toml
	${RMDIR} ${DESTDIR}etc/${TARGET}
	${RM} ${DESTDIR}man/man1/${MAN}.gz

clean:
	go clean
	${RM} ${TARGET}

.PHONY:	deps build install-bin install-data install-doc install uninstall clean

# File prefixes info
# ------------------
# c_  -> core program
# p_  -> parsing
# u_  -> utils
