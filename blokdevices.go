package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

func getattr (name string, attr string) (string){
    out,_ := exec.Command("lsblk","-lnpo", attr, name).CombinedOutput()
    out1 := strings.TrimRight(string(out),"\n")
    return out1
}
func main () {
    var rez, col, a, s string
    out, _ := exec.Command("blkid", "-o", "device").CombinedOutput()
    out1 := strings.TrimRight(string(out),"\n")
    outlist := strings.Split(out1,"\n")
    sort.Strings(outlist)
    for _, devices := range outlist{
        a=getattr(devices,"MOUNTPOINT")
        if strings.Contains(a,"boot") || strings.Contains(a,"SWAP") || strings.Contains(a,"var") { continue }
        s = "<span color='grey'>"+strings.TrimSpace(getattr(devices,"SIZE"))+"</span>"
        if a == "" {
            col = "gray"
            a = getattr(devices, "LABEL")
            if a == "" {
                a = "NO LABEL"
            }
        }else {
            col = "blue"
            al := strings.SplitAfter(a,"/")
            a = "/"+al[len(al)-1]
            s = "<span color='green'>"+strings.TrimSpace(getattr(devices, "FSAVAIL"))+"</span>/"+s
        }
        rez+=" |<span color='"+col+"'><b>["+devices[5:]+"]</b></span> "
        rez+=a+": " 
        rez+="<i>"+s+"</i>"
    }
    fmt.Println(rez)
    fmt.Println(rez)
}
