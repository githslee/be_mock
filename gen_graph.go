package main

import (
    "fmt"
    "time"
    "encoding/json"
)

type graphInfo struct {
    Serial_num string
    Xmin string
    Xmax string
    Xstep string
    Ymin string
    Ymax string
    Ystep string
    Nsamples string
    Data string
}

var graph graphInfo
var nsamples = 100

func graphInit() {

    graph.Serial_num = fmt.Sprintf("%d", 0)
    graph.Xmin = "0.0"
    graph.Xmax = "1000.0"
    graph.Xstep = "1.0"
    graph.Ymin = "-10.0"
    graph.Ymax = "10.0"
    graph.Nsamples = fmt.Sprintf("%d", nsamples)
}

func graph_data_write(count int) {
    
    var plot_data = make([]float32, nsamples)

    for i:=0; i<nsamples; i++ {
        plot_data[i] = (float32)(((i+count) % 20) - 10);
    }
    json_plot_data, err := json.Marshal(plot_data)
    if(err != nil) {
        fmt.Println(err)
    }

    stmt := fmt.Sprintf("INSERT INTO dash_graph (serialnum, xmin, xmax, xstep, ymin, ymax, nsamples, data) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE xmin='%s', xmax='%s', xstep='%s', ymin='%s', ymax='%s', nsamples='%s', data='%s';", graph.Serial_num, graph.Xmin, graph.Xmax, graph.Xstep, graph.Ymin, graph.Ymax, graph.Nsamples, string(json_plot_data), graph.Xmin, graph.Xmax, graph.Xstep, graph.Ymin, graph.Ymax, graph.Nsamples, string(json_plot_data))

    fmt.Println(stmt)

    _, err = fedb.Exec(stmt)
    if(err != nil) {
        fmt.Println("DB insert error", err)
    }
}

func graph_data(period int) {

    count := 0

    graphInit()
    for {
        fmt.Println("graph data")
	graph_data_write(count)
	count++
	time.Sleep(time.Duration(period) * time.Second)
    }
}
