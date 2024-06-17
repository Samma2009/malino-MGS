/*
    core.c - core syscalls like sync, reboot, etc...

    This code is part of libmsb (the Malino Syscall Bridge).

    Copyleft 2024 malino. This code is licensed under GPL2.
*/

#include <unistd.h>
#include <string.h>
#include "core.h"

void msb_sync() {
    syscall(0xa2);
}

long msb_reboot(unsigned int cmd) {
    return syscall(0xa9, 0xfee1dead, 0x28121969, cmd);
}

long msb_write(unsigned int fd, const char *buf) {
    return syscall(0x01, fd, buf, strlen(buf));
}

long msb_read(unsigned int fd, char *buf, unsigned long count) {
    return syscall(0, fd, buf, count);
}