package main

import (
    "fmt"
    "time"
)

// table layout
// serialnum=0: net_level=1: Purification site
// serialnum=1; net_level=2: iot-net-1
// serialnum=2; net_level=2: control-net

type siteInfo struct {
    Serial_num string;
    Net_level string;
    Applied_slice string;
    Net_property string;
    Egress_name string;
    Traffic string;
    Alarm_high string;
    Alarm_mid string;
    Alarm_low string;
}

var siteNum = 3 
var sList []siteInfo

func siteListInit() {

    sList = make([]siteInfo, siteNum)

    sList[0].Serial_num = fmt.Sprintf("%d", 0)
    sList[0].Net_level = fmt.Sprintf("%d", 1)
    sList[0].Applied_slice = "purification site"
    sList[0].Net_property = ""  // not possible to describe level-1 net property
    sList[0].Egress_name = "hostnet-1"

    sList[1].Serial_num = fmt.Sprintf("%d", 1)
    sList[1].Net_level = fmt.Sprintf("%d", 2)
    sList[1].Applied_slice = "iot-net-1"
    sList[1].Net_property = "mIoT"
    sList[1].Egress_name = "hostnet-2"

    sList[2].Serial_num = fmt.Sprintf("%d", 2)
    sList[2].Net_level = fmt.Sprintf("%d", 2)
    sList[2].Applied_slice = "control-net"
    sList[2].Net_property = "URLC"
    sList[2].Egress_name = "hostnet-2"
}

func sitelist_data_write(sn int, count int) {
    
    sList[sn].Traffic = fmt.Sprintf("%d", count);
    sList[sn].Alarm_high = fmt.Sprintf("%d", count);
    sList[sn].Alarm_mid = fmt.Sprintf("%d", count);
    sList[sn].Alarm_low = fmt.Sprintf("%d", count);

    stmt := fmt.Sprintf("INSERT INTO dash_sitelist (serialnum, net_level, applied_slice, net_property, egress_name, traffic, alarm_high, alarm_mid, alarm_low) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE net_level='%s', applied_slice='%s', net_property='%s', egress_name='%s', traffic='%s', alarm_high='%s', alarm_mid='%s', alarm_low='%s';", sList[sn].Serial_num, sList[sn].Net_level, sList[sn].Applied_slice, sList[sn].Net_property, sList[sn].Egress_name, sList[sn].Traffic, sList[sn].Alarm_high, sList[sn].Alarm_mid, sList[sn].Alarm_low, sList[sn].Net_level, sList[sn].Applied_slice, sList[sn].Net_property, sList[sn].Egress_name, sList[sn].Traffic, sList[sn].Alarm_high, sList[sn].Alarm_mid, sList[sn].Alarm_low);  

    fmt.Println(stmt)
    _, err := fedb.Exec(stmt)
    if(err != nil) {
        fmt.Println("DB insert error", err)
    }
}

func sitelist_data(period int) {

    count := 0

    siteListInit()
    for {
        fmt.Println("sitelist data")
	for i:=0; i<siteNum; i++ {
	    sitelist_data_write(i, count)
	}
	count++
	time.Sleep(time.Duration(period) * time.Second)
    }
}
