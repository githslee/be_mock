package main

import (
    "fmt"
    "time"
)

type utilizationInfo struct {
    Serial_num string
    Host_cpu string
    Host_mem string
    Switch_cpu string
    Switch_mem string
}

var util utilizationInfo

func utilizationInit() {

    util.Serial_num = fmt.Sprintf("%d", 0)
}

func utilization_data_write(count int) {
    
    util.Host_cpu = fmt.Sprintf("%d", count%100)
    util.Host_mem = fmt.Sprintf("%d", count%100)
    util.Switch_cpu = fmt.Sprintf("%d", count%100)
    util.Switch_mem = fmt.Sprintf("%d", count%100)

    stmt := fmt.Sprintf("INSERT INTO dash_utilization (serialnum, host_cpu, host_mem, switch_cpu, switch_mem) VALUES ('%s', '%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE host_cpu='%s', host_mem='%s', switch_cpu='%s', switch_mem='%s';", util.Serial_num, util.Host_cpu, util.Host_mem, util.Switch_cpu, util.Switch_mem, util.Host_cpu, util.Host_mem, util.Switch_cpu, util.Switch_mem); 

    fmt.Println(stmt)

    _, err := fedb.Exec(stmt)
    if(err != nil) {
        fmt.Println("DB insert error", err)
    }
}

func utilization_data(period int) {

    count := 0

    utilizationInit()
    for {
        fmt.Println("utilization data")
	utilization_data_write(count)
	count++
	time.Sleep(time.Duration(period) * time.Second)
    }
}
