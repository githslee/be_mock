package main

import (
    "fmt"
    "os"
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var fedb *sql.DB
const (
    port = 3306
    user = "db_write"
    hostname = "fe-db.cqlofnlw6kug.ap-northeast-2.rds.amazonaws.com"
    dbname = "dashboard"
)

func check_err(err error, msg string) {
    if(err != nil) {
	fmt.Println(msg, err)
	os.Exit(1)
    }
}

func connect_db() {

    var err error
    var version string

    mysqlconn := fmt.Sprintf("%s@tcp(%s:%d)/%s", user, hostname, port, dbname)
    fedb, err = sql.Open("mysql", mysqlconn)
    check_err(err, "db_conection error!")

    fedb.QueryRow("SELECT VERSION()").Scan(&version)
    fmt.Println("Connected to:", hostname, version)
}

func create_table() {

    var err error

    // Dashboard sitelist  table
    fmt.Println("== create dashboard-sitelist-table")
    stmt := "CREATE TABLE if not exists dash_sitelist (serialnum varchar(8), net_level varchar(4), applied_slice varchar(64), net_property varchar(64), egress_name varchar(64), traffic varchar(8), alarm_high varchar(8), alarm_mid varchar(8), alarm_low varchar(8), PRIMARY KEY (serialnum));"
    //fmt.Println(stmt)
    _, err = fedb.Exec(stmt)
    check_err(err, "table creation error")

    // Dashboard map table
    fmt.Println("== create dashboard-map-table")
    stmt = "CREATE TABLE if not exists dash_map (serialnum varchar(8), facility_name varchar(64), alarm_high varchar(8), alarm_med varchar(8), alarm_low varchar(8), applied_slice varchar(64), cpe varchar(8), PRIMARY KEY (serialnum));"
    //fmt.Println(stmt)
    _, err = fedb.Exec(stmt)
    check_err(err, "table creation error")

    // Dashboard summary table
    fmt.Println("== create dashboard-summary-table")
    stmt = "CREATE TABLE if not exists dash_summary (serialnum varchar (8), site varchar(16), applied_slice varchar(16), sdnc varchar(16), switch varchar(16), alarm_high varchar(8), alarm_med varchar(8), alarm_low varchar(8), PRIMARY KEY (serialnum));"
    //fmt.Println(stmt)
    _, err = fedb.Exec(stmt)
    check_err(err, "table creation error")

    // Dashboard graph table
    // data separation rule 1.0,2.0,3.0 ....
    fmt.Println("== create dashboard-graph-table")
    stmt = "CREATE TABLE if not exists dash_graph (serialnum varchar (8), xmin varchar(16), xmax varchar(16), xstep varchar(16), ymin varchar(16), ymax varchar(16), nsamples varchar(16), data varchar(32000), PRIMARY KEY (serialnum));"
    //fmt.Println(stmt)
    _, err = fedb.Exec(stmt)
    check_err(err, "table creation error")

    // Dashboard resource utilization table
    fmt.Println("== create dashboard-resource-utilization-table")
    stmt = "CREATE TABLE if not exists dash_utilization (serialnum varchar (8), host_cpu varchar(16), host_mem varchar(16), switch_cpu varchar(16), switch_mem varchar(16), PRIMARY KEY (serialnum));"
    //fmt.Println(stmt)
    _, err = fedb.Exec(stmt)
    check_err(err, "table creation error")

    // Dashboard alarm table
    fmt.Println("== create dashboard-alarm-table")
    stmt = "CREATE TABLE if not exists dash_alarm (serialnum varchar (8), time varchar(64), site varchar(64), system varchar(64), type varchar(16), source varchar(16), event varchar(64), event_msg varchar(512), PRIMARY KEY (serialnum));"
    //fmt.Println(stmt)
    _, err = fedb.Exec(stmt)
    check_err(err, "table creation error")

}

func main() {

    // data generation period
    period := 5 // 5 sec

    // db init
    connect_db()
    create_table()

    // data generation
    go sitelist_data(period)
    go map_data(period);
    go summary_data(period);
    go graph_data(period);
    go utilization_data(period);
    go alarm_data(period);

    for {
        time.Sleep(1 * time.Second)
    }
}
