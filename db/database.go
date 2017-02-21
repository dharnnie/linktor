package db

import(
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"	
)

type Basic struct{
	Fname string
	Sname string
	Email string
	Phone string
	Sex string
}

type SecondaryUp struct{
	Bio string
	Rel string
	DOB string
	Joined string
}

type EducationUp struct{
	Inst string
	Prog string
	Fac string
	Dept string
	Mat string
	Dur int64
	Level int64
}

const(
	db_type = "mysql"
	db_path = "root:mysqlrootpassword@/lts"
	NICK_EXISTS = "SELECT COUNT(*) FROM users WHERE nick = ?"
	SIGN_UP = "INSERT INTO users (nick, fname, sname, email,phone,bio,rel,dob,sex,joined,password) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	INIT_CONTACT = "INSERT INTO contact (nick, email, guardian, g_phone, g_email) VALUES(?,?,?,?,?)"
	INIT_LOCATION = "INSERT INTO location (nick, street, state, lga, college) VALUES(?,?,?,?,?)"
	INIT_EDUCATION = "INSERT INTO education (nick, inst,program,faculty, dept, mat, duration,level) VALUES(?,?,?,?,?,?,?,?)"
	INIT_PIC = "INSERT INTO pic (nick, pic) VALUES(?,?)"
	GET_PASSWORD = "SELECT password FROM users WHERE nick = ?"
	UPDATE_BASIC = "UPDATE users SET fname = ?, sname = ?, email = ?, phone = ?, sex = ? WHERE nick = ?"
	UP_BIO = "UPDATE users SET bio = ? WHERE nick = ?"
	UP_OTHERS = "UPDATE users SET rel = ?, dob = ?, joined = ? WHERE nick = ?"
	UP_EDUC = "UPDATE education SET inst = ?, program = ?, faculty = ?,dept = ?, mat = ?, duration = ?, level = ? WHERE nick = ?"
)

func SignUpAuth(n string)string{
	if exists(n){
		return "bad"
	}else{
		return "ok"
	}
}

func SignUp(n, f string , s,e string, b, p string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not open at SignUp")
	defer db.Close()

	prep, err := db.Prepare(SIGN_UP)
	HandleDBError(err, "Could not Prepare @ SignUp")

	res, err := prep.Exec(n,f,s,e,"Not set",b,"Not set","Not set","Not set","Not set",p)
	HandleDBError(err, "Could not Exec at SignUp")

	lr, err := res.LastInsertId()
	HandleDBError(err, "Error occured getting the LastInsertId")
	log.Println("LastInsertId is: ", lr)
}

func InitSignUp(n string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open @ InitSignUp")
	defer db.Close()

	prep, err := db.Prepare(INIT_LOCATION)
	HandleDBError(err, "Could not Prepare @ prep1 in InitSignUp")

	res, err := prep.Exec(n, "Not set","Not set","Not set","Not set")
	HandleDBError(err, "Could not Exec at InitSignUp")

	lr1, err := res.LastInsertId()
	HandleDBError(err, "Could not get LastInsertId lr1")
	log.Println("LastInsertId for location", lr1)

	prep2, err := db.Prepare(INIT_CONTACT)
	HandleDBError(err, "Could not Prepare @ prep2 in InitSignUp")

	res2, err := prep2.Exec(n,"Not set","Not set","Not set","Not set")
	HandleDBError(err, "Could not Exec at res2 in InitSignUp")

	lr2, err := res2.LastInsertId()
	HandleDBError(err, "Could not get LastInsertId lr2")
	log.Println("LastInsertId for location", lr2)

	prep3, err := db.Prepare(INIT_EDUCATION)
	HandleDBError(err, "Could not Prepare @ prep3 in InitSignUp")

	res3, err := prep3.Exec(n,"Not set","Not set","Not set","Not set","Not set",4,0)
	HandleDBError(err, "Could not Exec at res3 in InitSignUp")

	lr3, err := res3.LastInsertId()
	HandleDBError(err, "Could not get LastInsertId lr2")
	log.Println("LastInsertId for location", lr3)

	prep4, err := db.Prepare(INIT_PIC)
	HandleDBError(err, "Could not Prepare @ prep3 in INIT_PIC")

	res4, err := prep4.Exec(n,"linktorsample.jpg")
	HandleDBError(err, "Could not Exec at res3 in InitSignUp")

	lr4, err := res4.LastInsertId()
	HandleDBError(err, "Could not get LastInsertId lr2")
	log.Println("LastInsertId for PIC is", lr4)
}

func exists(n string)bool{
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db @ exists")
	defer db.Close()

	stmt, err := db.Query(NICK_EXISTS, n)
	HandleDBError(err, "Could not Query @ exists")
	noOfRows := checkCount(stmt)
	if noOfRows == 0{
		return false
	}else{
		return true
	}
}

func LoginAuth(n,p string)int {
	var val int
	exists := exists(n)
	if exists{
		pass := GetPassword(n)
		if pass == p{
			val = 20
		}else{
			val = 21
		}
	}else{
		val = 11
	}
	return val
}

func GetPassword(n string) string{
	var pass string
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open db at LoginDetMatch")
	defer db.Close()

	rows, err := db.Query(GET_PASSWORD, n)
	HandleDBError(err, "Could not Query at Stmt in LoginDetMatch")

	for rows.Next(){
		err := rows.Scan(&pass)
		if err != nil{
			log.Println("Error occured at rows.Next() in GetPassword")
		}
	}
	defer rows.Close()
	return pass

}

func UpdateBasic(n,fn string, sn,em string, ph,sx string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "COuld not open db at basic Update")
	defer db.Close()

	prep, err := db.Prepare(UPDATE_BASIC)
	HandleDBError(err, "Could not Prepare UPDATE_BASIC")

	res, err := prep.Exec(&fn,&sn,&em,&ph,&sx, &n)
	HandleDBError(err, "Could not Execute UPDATE_BASIC")

	l, _ := res.LastInsertId()
	log.Println("Last update was on : - ", l)
}

func UpdateSecondary(n,bi string, re, dObirth string, joi string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open at sql")
	defer db.Close()

	prepBio, err := db.Prepare(UP_BIO)
	HandleDBError(err, "Could not prepare UP_BIO")	

	resBio, err := prepBio.Exec(&bi, &n)
	HandleDBError(err, "Could not Execute UP_BIO")

	l1, _ := resBio.LastInsertId()
	log.Println("LastInsertId at resBio update is : - ", l1) 

	prepOthers, err := db.Prepare(UP_OTHERS)
	HandleDBError(err, "Could not prepare UP_OTHERS")

	resOthers, err := prepOthers.Exec(&re, &dObirth, &joi, &n)
	HandleDBError(err, "Could not Exec UP_OTHERS")

	l2, _ := resOthers.LastInsertId()
	log.Println("LastInsertId at resOthers update is :- ", l2)
}

func UpdateEducation(n, in string, prog, fa string, de, ma string, du, le string) {
	db, err := sql.Open(db_type, db_path)
	HandleDBError(err, "Could not Open at sql")
	defer db.Close()

	prepEduc, err := db.Prepare(UP_EDUC)
	HandleDBError(err, "Could not prepare UP_EDUC")

	res, err := prepEduc.Exec(&in, &prog, &fa, &de, &ma, &du,&le, &n)
	HandleDBError(err, "Could not Exec UP_EDUC")

	l, _ := res.LastInsertId()
	log.Println("LastInsertId at prepEduc update is  - ", l)
}

func checkCount(rows *sql.Rows)(count int){
	for rows.Next(){
		err := rows.Scan(&count)
		if err != nil{
			panic(err)
		}
	}
	return count
}
func HandleDBError(e error, info string){
	if e != nil{
		log.Println(info, e)
	}
}