global _hello

section .text
_syscall:
        int     0x80
        ret

hello:
        push    dword       msg.len
        push    dword       msg
        push    dword       1
        mov     eax,        4

        call    _syscall
        add     esp,        12
        ret

_hello:
        mov     ecx,        [esp+4]    ; Counter

    loop:
        call    hello                  ; Print hello world
        dec     ecx                    ; Decrement counter
        cmp     ecx,        0          ; Check counter > 0
        jg      loop

        ret

section .data
msg:    db      "Hello, world!", 10
.len:   equ     $ - msg
