package db

import(
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)


const(
	SAVE_REQUEST = "INSERT INTO requests (nick,course,category,school,description, rid) VALUES(?,?,?,?,?,?)"
	ADD_TUTOR = "INSERT INTO t_expertise (nick,expertise,t_bio) VALUES(?,?,?)"
)

func PersistRequest(n, c string, cat, sch string , desc, rid string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open at PersistRequest")
	defer db.Close()

	prep, err := db.Prepare(SAVE_REQUEST)
	HandleDBError(err, "Could not Prepare SAVE_REQUEST")

	res, err := prep.Exec(&n,&c,&cat,&sch,&desc,&rid)
	HandleDBError(err, "Could not Execute at PersistRequest")

	lr, err := res.LastInsertId()
	HandleDBError(err, "Error getting LastInsertId")
	fmt.Println("Last record is on: ", lr)
}

func AddNewTutor(n,ex string, tbio string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db  at AddNewTutor")
	defer db.Close()

	prep, err := db.Prepare(ADD_TUTOR)
	HandleDBError(err, "Could not prepare ADD_TUTOR")

	res, err := prep.Exec(&n, &ex, &tbio)
	HandleDBError(err, "Could not Exec ADD_TUTOR")

	lr, err := res.LastInsertId()
	HandleDBError(err, "Could not get LastInsertId")
	fmt.Println("Last row is: ", lr)
}






