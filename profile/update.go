package profile

import (
	"fmt"
	"net/http"
	"github.com/dharnnie/linktor/db"
	"github.com/dharnnie/linktor/sess"
	"html/template"
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
	Dur string
	Level string
}



func UpdateServlet(w http.ResponseWriter, r *http.Request) {
	
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		//var user interface{}
		//user = Who{n}		
		
		b:= new(Basic)
		s:= new(SecondaryUp)
		e:= new(EducationUp)

		b.Fname, b.Sname, b.Email, b.Phone, b.Sex =  r.FormValue("fname"), r.FormValue("sname"), r.FormValue("email"), r.FormValue("phone"), r.FormValue("sex")
		s.Bio, s.Rel, s.DOB, s.Joined = r.FormValue("bio"), r.FormValue("rel"), r.FormValue("dob"), r.FormValue("joined")
		e.Inst, e.Prog, e.Fac, e.Dept, e.Mat, e.Dur, e.Level =r.FormValue("inst"), r.FormValue("prog"), r.FormValue("fac"), r.FormValue("dept"), r.FormValue("mat"),r.FormValue("dur"), r.FormValue("lev")

		db.UpdateBasic(n, b.Fname, b.Sname, b.Email, b.Phone, b.Sex )
		db.UpdateSecondary(n, s.Bio, s.Rel, s.DOB, s.Joined)
		db.UpdateEducation(n,e.Inst, e.Prog, e.Fac, e.Dept, e.Mat, e.Dur, e.Level)

		//ProcessPicUpdate(w,r)

		person := Who{n}
		person.EditProfile(w,r)

	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		fmt.Println("Session Linktor does not exist")
	}
}

func EditProfileServlet(w http.ResponseWriter, r *http.Request) {
	//typically get the session value - nick
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		user := Who{n}		
		log.Println("SessionExists at ViewProfile ", n) // here....
		user.EditProfile(w,r)
	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		log.Println("Session Linktor does not exist")
	}
}


func (wh Who) EditProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("EditProfile has been called")
	p := new(Profile)
	p.Nick = wh.Nick

	img :=  GetPicPath(p.Nick)
	fmt.Println("The nick from session is: -- ", p.Nick)
	fmt.Println("The image path has name: -- ", img)

	p.BasicProf.Fname, p.BasicProf.Sname, p.BasicProf.Email, p.BasicProf.Phone, p.BasicProf.Sex = db.GetBasic(wh.Nick)
	p.BasicProf.ImagePath = "../imgs/" + img
	log.Println("Your ppic link is..", p.BasicProf.ImagePath)
	p.Secondary.Bio, p.Secondary.Rel, p.Secondary.DOB, p.Secondary.Joined = db.GetSecondary(wh.Nick)
	p.Guardian.Name, p.Guardian.Phone, p.Guardian.Email = db.GetGuardian(wh.Nick)
	p.Education.Inst, p.Education.Prog, p.Education.Fac, p.Education.Dept, p.Education.Mat,p.Education.Duration,p.Education.Level = db.GetEducation(wh.Nick)

	//log.Println(p)

	t, err := template.ParseFiles("templates/p/edit-profile.html")
	smplErr(err, "Could not parse edit-profile")
	log.Println("--From EditProfile")
	t.Execute(w, p)
	//t.Execute(os.Stdout, p)
}

func UpdatePic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		ProcessPicUpdate(w,r)
		http.Redirect(w,r , "/profile/edit", 301)
	}
}

func smplErr(e error, m string){
	if e != nil{
		log.Println(m, e)
	}
}