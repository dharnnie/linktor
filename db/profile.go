package db

import (
	"log"
	"database/sql"
)

const(
	GET_BASIC = "SELECT fname, sname, email, phone, sex FROM users WHERE nick = ?"
	GET_SECONDARY = "SELECT bio, rel, dob FROM users WHERE nick = ?"
	GET_GUARDIAN = "SELECT guardian, g_phone, g_email FROM contact WHERE nick = ?"
	GET_EDUCATION = "SELECT inst, program, faculty, dept, mat, duration, level FROM education WHERE nick = ?"
	UPDATE_PIC = "UPDATE pic SET pic = ? WHERE nick = ?"
	GET_PIC = "SELECT pic FROM pic WHERE nick = ?"
)


func GetBasic(n string)(sql.NullString,sql.NullString,sql.NullString,sql.NullString,sql.NullString) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at GetProfile")
	defer db.Close()

	query, err := db.Query(GET_BASIC, n)
	HandleDBError(err, "Error occured while getting basic")

	var Fname,Sname,Email,Phone,Sex sql.NullString
	for query.Next(){
		err := query.Scan(&Fname, &Sname, &Email, &Phone, &Sex)
		HandleDBError(err, "Error > Query.Scan >GET_BASIC")
	}
	return Fname, Sname,Email,Phone,Sex
}

func GetSecondary(n string)(sql.NullString,sql.NullString,sql.NullString,string){
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at Get-Profile")
	defer db.Close()

	query, err := db.Query(GET_SECONDARY, n)
	HandleDBError(err, "Error occured while getting secondary")

	var Bio,Rel,DOB sql.NullString
	var Joined string
	for query.Next(){
		err := query.Scan(&Bio, &Rel, &DOB)
		HandleDBError(err, "Error occured at Query.Scan in GET_SECONDARY")
	}
	Joined = "Some Date"
	return Bio,Rel,DOB,Joined
}

func GetGuardian(n string) (sql.NullString,sql.NullString,sql.NullString) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at GetProfile")
	defer db.Close()

	query, err := db.Query(GET_GUARDIAN, n)
	HandleDBError(err, "Error occured while getting guardian")

	var Name,Phone,Email sql.NullString
	for query.Next(){
		err := query.Scan(&Name, &Phone, &Email)
		HandleDBError(err, "Error occured at Query.Scan in GET_GUARDIAN")
	}
	return Name,Phone,Email
}

func GetEducation(n string) (sql.NullString,sql.NullString,sql.NullString,sql.NullString,sql.NullString,sql.NullInt64,sql.NullInt64) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at GetEducatio")
	defer db.Close()

	query, err := db.Query(GET_EDUCATION, n)
	HandleDBError(err, "Error occured while getting education")

	var Inst, Prog, Fac, Dept, Mat sql.NullString
	var Duration, Level sql.NullInt64
	for query.Next(){
		err := query.Scan(&Inst, &Prog, &Fac, &Dept, &Mat, &Duration, &Level)
		HandleDBError(err, "Error occured at Query.Scan in GET_EDUCATION")
	}
	return Inst,Prog,Fac, Dept, Mat, Duration, Level
}
func GetImage(n string)string {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at SavePicName")
	defer db.Close()

	query, err := db.Query(GET_PIC, n)
	HandleDBError(err, "Could not Query GET_PIC")
	var pPath string

	for query.Next(){
		err := query.Scan(&pPath)
		HandleDBError(err, "Could not scan GET_PIC")
	}
	return pPath
}

func UpdatePic(n, p string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at UPDATE_PIC")
	defer db.Close()

	prep, err := db.Prepare(UPDATE_PIC)
	HandleDBError(err, "Could not prep db at UPDATE_PIC")

	res, err := prep.Exec(&p,&n)
	HandleDBError(err, "Could not Exec at UPDATE_PIC")

	log.Printf("Direct from DB, I have Updated %s pic with %s", n, p)

	lr, err := res.LastInsertId()
	HandleDBError(err, "Error occured while getting LastInsertId at UPDATE_PIC")
	log.Println("LastInsertId for Update is : ", lr)
}

