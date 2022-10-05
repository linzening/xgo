package xgo

import(
	"os/exec"
)

func Add(a int,b int) int {
	return a + b
}
// 获取系统的Uname
func Uname() string {
	cmd := exec.Command("uname")
	out,err := cmd.Output()
	if err != nil {
		return "cmd error."
	} else {
		return string(out)
	}
}