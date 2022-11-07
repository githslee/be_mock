package main

import (
    "fmt"
    "time"
)

type sumInfo struct {
    Serial_num string
    Site string
    Applied_slice string
    Sdnc string
    Switch string;
    Alarm_high string;
    Alarm_mid string;
    Alarm_low string;
}

var sumNum = 2 
var summary sumInfo

func sumListInit() {

    summary.Serial_num = fmt.Sprintf("%d", 0)
    summary.Site = "0/10"
    summary.Applied_slice = "0/10"
    summary.Sdnc = "0/10"
    summary.Switch = "0/10"
}

func summary_data_write(sn int, count int) {
    
    summary.Alarm_high = fmt.Sprintf("%d", count);
    summary.Alarm_mid = fmt.Sprintf("%d", count);
    summary.Alarm_low = fmt.Sprintf("%d", count);

    stmt := fmt.Sprintf("INSERT INTO dash_summary (serialnum, site, applied_slice, sdnc, switch, alarm_high, alarm_mid, alarm_low) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE site='%s', applied_slice='%s', sdnc='%s', switch='%s', alarm_high='%s', alarm_mid='%s', alarm_low='%s';", summary.Serial_num, summary.Site, summary.Applied_slice, summary.Sdnc, summary.Switch, summary.Alarm_high, summary.Alarm_mid, summary.Alarm_low, summary.Site, summary.Applied_slice, summary.Sdnc, summary.Switch, summary.Alarm_high, summary.Alarm_mid, summary.Alarm_low)

    fmt.Println(stmt)
    _, err := fedb.Exec(stmt)
    if(err != nil) {
        fmt.Println("DB insert error", err)
    }
}

func summary_data(period int) {

    count := 0

    sumListInit()
    for {
        fmt.Println("summary data")
	summary_data_write(0, count)
	count++
	time.Sleep(time.Duration(period) * time.Second)
    }
}
