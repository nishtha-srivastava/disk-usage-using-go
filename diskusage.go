package main

import (
	"fmt"
        "strings"
        "strconv"
        "os/exec"
 )


func main() {

     fmt.Println("Hello, there")

     diskInfo := getDiskInfo()
     fmt.Println("In main -- ")
     fmt.Println("Total disk size in MBs: ", diskInfo.Total)
     fmt.Println("Disk used in MBs: ", diskInfo.Used)
     fmt.Println("Disk available in MBs: ", diskInfo.Available)
     fmt.Println("% of Disk used: ", diskInfo.Used_percent)

}

type DiskInfo struct {
     Total float64
     Used float64
     Available float64
     Used_percent float64
}


func getDiskInfo() DiskInfo {
          
     cmd := "/bin/bash"
     //args := []string{"-c", "cd /var/vcap/store && du -ch | grep total | awk '{print $1}'"}
     args := []string{"-c", "df -h | grep /var/vcap/store"}
     out, err := exec.Command(cmd, args...).Output()
        if err != nil {
            fmt.Println(err)
            panic(err)
         }
     outstr := string(out)
     outstr_fields := strings.Fields(outstr)
     total, used, avail, used_percent := outstr_fields[1], outstr_fields[2], outstr_fields[3], outstr_fields[4]

     // ** Total Disk Size (MBs)
     total_disk := total[0:len(total)-1]
     //string to float
     total_disk_float, err := strconv.ParseFloat(total_disk, 64)
     if err != nil {
        fmt.Println(err)
     }
     total_disk_unit := total[len(total)-1:len(total)]
     if total_disk_unit == "G" {
        total_disk_float = 1024 * total_disk_float
     }
     //fmt.Println(total_disk_float)

     // ** Used Disk (MBs) **
     used_disk := used[0:len(used)-1]
     // string to float
     used_disk_float, err := strconv.ParseFloat(used_disk, 64)
     if err != nil {
        fmt.Println(err)
     }
     used_disk_unit := used[len(used)-1:len(used)]
     if used_disk_unit == "G" {
        used_disk_float = 1024 * used_disk_float
     }
     //fmt.Println(used_disk_float)

     // ** Available Disk (MBs) **
     avail_disk := avail[0:len(avail)-1]
     // string to float
     avail_disk_float, err := strconv.ParseFloat(avail_disk, 64)
     if err != nil {
        fmt.Println(err)
     }
     avail_disk_unit := avail[len(avail)-1:len(avail)]
     if avail_disk_unit == "G" {
        avail_disk_float = 1024 * avail_disk_float
     }
     //fmt.Println(avail_disk_float)

     // ** Disk Used % **
     disk_used_percentage := used_percent[0:len(used_percent)-1]
     // string to float
     disk_used_percentage_float, err := strconv.ParseFloat(disk_used_percentage, 64)
     if err != nil {
        fmt.Println(err)
     }
     //fmt.Println(disk_used_percentage_float)
     
     //construct DiskInfo
     diskInfo := DiskInfo{Total: total_disk_float, Used: used_disk_float, Available: avail_disk_float, Used_percent: disk_used_percentage_float}

     return diskInfo
}



