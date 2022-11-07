package main

import (
    "fmt"
    "time"
)

// example map layout
// serialnum=0: facility_name=Purification site, applied_slice=iot-net-1
// serialnum=1: facility_name=Bridge, applied_slice=iot-net-1

type mapInfo struct {
    Serial_num string
    Facility_name string
    Applied_slice string
    Cpe string
    Alarm_high string;
    Alarm_mid string;
    Alarm_low string;
}

var mapNum = 2 
var mList []mapInfo

func mapListInit() {

    mList = make([]mapInfo, mapNum)

    mList[0].Serial_num = fmt.Sprintf("%d", 0)
    mList[0].Facility_name = "Purification site"
    mList[0].Applied_slice = "2"
    mList[0].Cpe = "10"

    mList[1].Serial_num = fmt.Sprintf("%d", 1)
    mList[1].Facility_name = "Bridge"
    mList[1].Applied_slice = "3"
    mList[1].Cpe = "11"
}

func maplist_data_write(sn int, count int) {
    
    mList[sn].Alarm_high = fmt.Sprintf("%d", count);
    mList[sn].Alarm_mid = fmt.Sprintf("%d", count);
    mList[sn].Alarm_low = fmt.Sprintf("%d", count);

    stmt := fmt.Sprintf("INSERT INTO dash_map (serialnum, facility_name, applied_slice, cpe, alarm_high, alarm_mid, alarm_low) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE facility_name='%s', applied_slice='%s', cpe='%s', alarm_high='%s', alarm_mid='%s', alarm_low='%s';", mList[sn].Serial_num, mList[sn].Facility_name, mList[sn].Applied_slice, mList[sn].Cpe, mList[sn].Alarm_high, mList[sn].Alarm_mid, mList[sn].Alarm_low, mList[sn].Facility_name, mList[sn].Applied_slice, mList[sn].Cpe, mList[sn].Alarm_high, mList[sn].Alarm_mid, mList[sn].Alarm_low)

    fmt.Println(stmt)
    _, err := fedb.Exec(stmt)
    if(err != nil) {
        fmt.Println("DB insert error", err)
    }
}

func map_data(period int) {

    count := 0

    mapListInit()
    for {
        fmt.Println("maplist data")
	for i:=0; i<mapNum; i++ {
	    maplist_data_write(i, count)
	}
	count++
	time.Sleep(time.Duration(period) * time.Second)
    }
}
