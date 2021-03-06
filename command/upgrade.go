package trainCommand

import (
	"fmt"
)

const CmdBinPath = "$GOPATH/bin/train"

func Upgrade() {
	bash("go get -u github.com/shaoshing/train")
	bash("go build -o " + CmdBinPath + " github.com/shaoshing/train/cmd")
	fmt.Println("Installed latest train command into " + CmdBinPath)
}
