package pc

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc *Pc) Ping() {
	fmt.Println(myPc.Brand, "pong")
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPc Pc) String() string {
	return fmt.Sprintf("PC STRUCT\nRAM -> %dGB\nDISK SPACE -> %d\nBRAND -> %s\n", myPc.Ram, myPc.Disk, myPc.Brand)
}
