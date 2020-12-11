package sample

import pb "gitlab.com/GRPC/pcbook/pb/proto"

//New keyboard return a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

//NewCPU return a sample CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 3.5)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

//NewGPU return a sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: 0,
		MaxGhz: 0,
		Memory: nil,
	}
	return gpu
}
