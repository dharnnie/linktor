package profile


import (
	"log"
	"net/http"
	"github.com/dharnnie/linktor/db"
	"html/template"
	"database/sql"
	//"os"
	"github.com/dharnnie/linktor/sess"
)
type BasicProf struct{
	Fname sql.NullString
	Sname sql.NullString
	Email sql.NullString
	Phone sql.NullString
	Sex sql.NullString
	ImagePath string
}

type Secondary struct{
	Bio sql.NullString
	Rel sql.NullString
	DOB sql.NullString
	Joined string
}

type Guardian struct{
	Name sql.NullString
	Phone sql.NullString
	Email sql.NullString
}

type Profile struct{
	Nick string
	BasicProf
	Secondary
	Guardian
	Education
}

type Education struct{
	Inst sql.NullString
	Prog sql.NullString
	Fac sql.NullString
	Dept sql.NullString
	Mat sql.NullString
	Duration sql.NullInt64
	Level sql.NullInt64
}

type Who struct{
	Nick string
}

func ViewProfileServlet(w http.ResponseWriter, r *http.Request){
	// typically get the session nick
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		user := Who{n}		
		log.Println("SessionExists at ViewProfile ", n) // here....
		user.ViewProfile(w,r)
	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		log.Println("Session Linktor does not exist")
	}
}

// func displays user's profile
func (wh Who) ViewProfile(w http.ResponseWriter, r *http.Request) {
	
	p := new(Profile)
	p.Nick = wh.Nick

	img := GetPicPath(p.Nick)
	// get the values and save in appropriate struct
	p.BasicProf.Fname, p.BasicProf.Sname, p.BasicProf.Email, p.BasicProf.Phone, p.BasicProf.Sex = db.GetBasic(wh.Nick)
	p.BasicProf.ImagePath = "../imgs/" + img
	p.Secondary.Bio, p.Secondary.Rel, p.Secondary.DOB, p.Secondary.Joined = db.GetSecondary(wh.Nick)
	p.Guardian.Name, p.Guardian.Phone, p.Guardian.Email = db.GetGuardian(wh.Nick)
	p.Education.Inst, p.Education.Prog, p.Education.Fac, p.Education.Dept, p.Education.Mat,p.Education.Duration,p.Education.Level = db.GetEducation(wh.Nick)

	log.Println(p)
	t, err := template.ParseFiles("templates/p/view-profile.html")
	smplErr(err, "Could not parse view-profile")
	t.Execute(w, p)
	//t.Execute(os.Stdout, p)
}




