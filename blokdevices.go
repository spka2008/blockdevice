package main
import (
    "fmt"
    "os/exec"
    "strings"
)

func main () {
    out, _ := exec.Command("blkid", "-o", "device").CombinedOutput()
    out1 := strings.TrimRight(string(out),"\n")
    outlist := strings.Split(out1,"\n")
    fmt.Println(outlist)
}
