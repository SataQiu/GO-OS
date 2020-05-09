org 0x7c00

LOAD_ADDR equ 0x9000

ENTRY:
    mov ax, 0
    mov ss, ax
    mov ds, ax
    mov es, ax
    mov si, ax

READ_FLOPPY:
    mov ch, 1            ; 1 号柱面
    mov dh, 0            ; 0 号磁头
    mov cl, 2            ; 2 号扇区
    mov bx, LOAD_ADDR    ; 内核写入内存起始地址 es:bx
    mov ah, 0x02         ; ah = 2 表示读盘操作
    mov al, 0x01         ; al = 1 表示读取一个扇区
    mov dl, 0            ; 驱动器编号，一个软盘就是 0
    int 0x13             ; 调用 BIOS 中断，读盘
    jc  END              ; 失败，跳到 fin
    jmp LOAD_ADDR        ; 成功，跳到内核所在内存地址，控制权交给内核

END:
    hlt
    jmp END
