module blue/be/be_mock

go 1.18

//replace blue/lib/common => ../lib/common
//require (
//        blue/lib/common v0.0.1
//        blue/lib/msg v0.0.1
//)

require github.com/go-sql-driver/mysql v1.6.0 // indirect
