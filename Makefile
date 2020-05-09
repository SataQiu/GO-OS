.PHONY: all
all: boot kernel
	@go run hack/main.go

.PHONY: clean
clean:
	@rm -rf _output/*

.PHONY: boot
boot: output
	@cd boot && nasm boot.asm -o boot && mv boot ../_output/boot.bin

.PHONY: kernel
kernel: output
	@cd kernel && nasm kernel.asm -o kernel && mv kernel ../_output/kernel.bin

.PHONY: output
	@mkdir _output
