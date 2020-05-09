package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SataQiu/GO-OS/pkg/floppy"
)

func main() {
	fd := floppy.NewFloppy(80, 18, 512)

	// 写入 boot 引导数据（第一个扇区）
	b, err := ioutil.ReadFile("_output/boot.bin")
	if err != nil {
		fmt.Printf("Unable to read data from boot.bin, %v\n", err)
		os.Exit(-1)
	}

	bootData := make([]byte, 512)
	copy(bootData, b)
	bootData[510] = 0x55
	bootData[511] = 0xaa

	fd.Write(floppy.MagneticHead0, 0, 0, bootData)

	// 写入内核数据到 1 号柱面， 2 号扇区
	b, err = ioutil.ReadFile("_output/kernel.bin")
	if err != nil {
		fmt.Printf("Unable to read data from kernel.bin, %v\n", err)
		os.Exit(-1)
	}
	fd.Write(floppy.MagneticHead0, 1, 1, b)

	// 输出镜像文件
	if err := fd.DumpToFile("_output/system.img"); err != nil {
		fmt.Printf("Unable to dump disk to system.img, %v\n", err)
		os.Exit(-1)
	}
}
