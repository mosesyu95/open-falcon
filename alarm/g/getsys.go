package g
import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func GetSysName(ip string) (string, string) {
	var system_name string
        var manager_name string
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/om?loc=Local&parseTime=true")
    CheckErr(err)
    sql :="select system_name,manager from system,server where system.system_id = server.system_id and server.main_ip=\""+ip+"\";"
    rows, err := db.Query(sql)
    CheckErr(err)
	defer rows.Close()
    for rows.Next() {

        err = rows.Scan(&system_name, &manager_name)
        CheckErr(err)
//       fmt.Println(system_name)
//        fmt.Println(manager_name)
    }
    db.Close()
    return system_name,manager_name
}

func CheckErr(err error) {
    if err != nil {
        log.Println(err)
        panic(err)
	return
    }
}

