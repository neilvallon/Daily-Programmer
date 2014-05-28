; nasm -f macho hello.asm && ld -macosx_version_min 10.9 -lSystem -o hello hello.o && ./hello

global _main

section .text
_syscall:
        int     0x80
        ret

_exit:
        mov     ebx,        0
        mov     eax,        1
        call    _syscall

_hello:
        push    dword       msg.len
        push    dword       msg
        push    dword       1
        mov     eax,        4

        call    _syscall
        add     esp,        12
        ret

_main:
        mov     ecx,        10          ; Counter

    loop:
        call    _hello                  ; Print hello world
        dec     ecx                     ; Decrement counter
        cmp     ecx,        0           ; Check counter > 0
        jg      loop

        call    _exit


section .data
msg:    db      "Hello, world!", 10
.len:   equ     $ - msg
