package g
import (
    "database/sql"
//    "time"
    "fmt"
    "io"
    _ "github.com/go-sql-driver/mysql"
    "crypto/md5"
)

    func GetMap(endpoint string,counter string) (string) {
    str :=endpoint + counter
    w := md5.New()
    io.WriteString(w, str)
    ck := fmt.Sprintf("%x", w.Sum(nil))
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dashboard?loc=Local&parseTime=true")
    CheckErr(err)
    _,err1 := db.Query("insert ignore into tmp_graph (endpoints, counters, ck) values(\"" + endpoint + "\", \"" + counter+ "\", \""+ ck + "\") ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id);")
    CheckErr(err1)
        sql :="select id from tmp_graph where endpoints=\""+endpoint+"\" and counters=\""+counter + "\";"
    rows, err2 := db.Query(sql)
    CheckErr(err2)    
    var id_ string
    defer rows.Close()
    for rows.Next() {
    err = rows.Scan(&id_)
    CheckErr(err)
    }
    db.Close()
    return id_
}
func GetC(counter string,arg string)  (string){
    if arg == ""{
        return counter
    }else {
        return counter+"/"+arg
    }
}
