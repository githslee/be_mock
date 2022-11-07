package main

import (
    "fmt"
    "time"
)

type alarmInfo struct {
    Serial_num string
    Time string
    Site string
    System string
    Type string
    Source string
    Event string
    Event_msg string
}

var alarm alarmInfo

func alarm_data_write(count int) {
    
    alarm.Serial_num = fmt.Sprintf("%d", count)
    alarm.Time = time.Now().String()
    alarm.Site = fmt.Sprintf("Site-%d", count)
    alarm.System = "System"
    alarm.Type = "Spooky"
    alarm.Source = "Switch"
    alarm.Event = fmt.Sprintf("Event-%d", count)
    alarm.Event_msg = fmt.Sprintf("Event-msg-%d", count)

    stmt := fmt.Sprintf("INSERT INTO dash_alarm (serialnum, time, site, system, type, source, event, event_msg) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE time='%s', site='%s', system='%s', type='%s', source='%s', event='%s', event_msg='%s';", alarm.Serial_num, alarm.Time, alarm.Site, alarm.System, alarm.Type, alarm.Source, alarm.Event, alarm.Event_msg, alarm.Time, alarm.Site, alarm.System, alarm.Type, alarm.Source, alarm.Event, alarm.Event_msg)

    fmt.Println(stmt)

    _, err := fedb.Exec(stmt)
    if(err != nil) {
        fmt.Println("DB insert error", err)
    }
}

func alarm_data(period int) {

    count := 0

    for {
        fmt.Println("alarm data")
	alarm_data_write(count)
	count++
	time.Sleep(time.Duration(period) * time.Second)
    }
}
