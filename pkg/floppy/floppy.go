package floppy

import (
	"os"

	"github.com/mohae/deepcopy"
)

// MagneticHead 代表磁头编号类型
type MagneticHead int32

const (
	// MagneticHead0 表示 0 号磁头
	MagneticHead0 MagneticHead = 0
	// MagneticHead1 表示 1 号磁头
	MagneticHead1 MagneticHead = 1
)

// Floppy 代表软盘类型
type Floppy struct {
	CylinderCount   int                         // 柱面数量
	SectorCount     int                         // 扇区数量
	SectorSize      int                         // 扇区大小
	CurMagneticHead MagneticHead                // 当前磁头
	CurCyliner      int                         // 当前柱面
	CurSector       int                         // 当前扇区
	data            map[MagneticHead][][][]byte // 数据存储
}

// NewFloppy 创建一个标准软盘对象
func NewFloppy(cylinderCount, sectorCount, sectorSize int) *Floppy {
	disk0 := make([][][]byte, cylinderCount)
	for i := 0; i < cylinderCount; i++ {
		disk0[i] = make([][]byte, sectorCount)
		for j := 0; j < sectorCount; j++ {
			disk0[i][j] = make([]byte, sectorSize)
		}
	}

	disk1 := deepcopy.Copy(disk0).([][][]byte)

	return &Floppy{
		CylinderCount:   cylinderCount,
		SectorCount:     sectorCount,
		SectorSize:      sectorSize,
		CurMagneticHead: MagneticHead0,
		CurCyliner:      0,
		CurSector:       0,
		data:            map[MagneticHead][][][]byte{MagneticHead0: disk0, MagneticHead1: disk1},
	}
}

func (f *Floppy) setMagneticHead(head MagneticHead) {
	f.CurMagneticHead = head
}

func (f *Floppy) setCylinder(cylinder int) {
	if cylinder < 0 {
		cylinder = 0
	}
	if cylinder > 79 {
		cylinder = 79
	}
	f.CurCyliner = cylinder
}

func (f *Floppy) setSector(sector int) {
	if sector < 0 {
		sector = 0
	}
	if sector > 17 {
		sector = 17
	}
	f.CurSector = sector
}

func (f *Floppy) Write(head MagneticHead, cylinder, sector int, data []byte) {
	f.setMagneticHead(head)
	f.setCylinder(cylinder)
	f.setSector(sector)
	copy(f.data[head][cylinder][sector], data)
}

func (f *Floppy) Read(head MagneticHead, cylinder, sector int) []byte {
	f.setMagneticHead(head)
	f.setCylinder(cylinder)
	f.setSector(sector)
	return f.data[head][cylinder][sector]
}

// DumpToFile 将软盘数据持久化到文件
func (f *Floppy) DumpToFile(filename string) error {
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	for i := 0; i < f.CylinderCount; i++ {
		for j := 0; j <= int(MagneticHead1); j++ {
			for k := 0; k < f.SectorCount; k++ {
				data := f.Read(MagneticHead(j), i, k)
				if _, err := fd.Write(data); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
