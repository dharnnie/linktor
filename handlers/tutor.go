package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"html/template"
	"github.com/dharnnie/linktor/db"
	"github.com/dharnnie/linktor/sess"
)

type Request struct{
	Sender string // sender
	Course string
	Category string
	Institution string
	Description string
	RID string
}

type ThisUser struct{
	Nick string
	Done
}
type Done struct{
	Success string
}
type Become struct{
	Sender string
	Expertise string
	TBio string
}

func RequestTutorServlet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		if sess.SessionExists(w,r){
			sess.InitSessionValues(w,r)
			n := sess.GetSessionNick(w,r)
			fmt.Println("SessionExists at ViewProfile ", n) // here....
			gtd := Done{""}
			Person := ThisUser{n,gtd}
			Person.serveGetTutor(w)
			fmt.Println("serveGetTutor called!!")
		}else{
			t, err := template.ParseFiles("templates/homepage.html")
			smplErr(err, "Error at Index Servlet")
			t.Execute(w, nil)	
			fmt.Println("Session Linktor does not exist")
		}
	}else{
		if sess.SessionExists(w,r){
			sess.InitSessionValues(w,r)
			n := sess.GetSessionNick(w,r)
			r.ParseForm()
			sender := n
			request := Request{sender,r.FormValue("course"), r.FormValue("cats"), r.FormValue("inst"), r.FormValue("desc"), ""}
			request.GenerateRID() // sets the request ID - Pointer stuff
			request.SaveRequest(w,r)
			gtd := Done{"- Request Sent, expect a response soon."}
			Person := ThisUser{n,gtd}
			Person.serveGetTutor(w)// redirect to get tutor
		}else{
			t, err := template.ParseFiles("templates/homepage.html")
			smplErr(err, "Error at Index Servlet")
			t.Execute(w, nil)	
			fmt.Println("Session Linktor does not exist")
		}		
	}
}
			


func BecomeATutorServlet(w http.ResponseWriter, r *http.Request) {
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		fmt.Println("SessionExists at ViewProfile ", n)
		becomeT := Become{n, "",""} //  get danny from session
		this := ThisUser{n, Done{""}}
		if r.Method == "GET"{
			this.serveBecomeTutor(w)
		}else{
			becomeT.Expertise, becomeT.TBio = r.FormValue("expertise"), r.FormValue("tbio")
			becomeT.BecomeTutor()
			this = ThisUser{n,Done{"Recieved, expect a response"}}
			this.serveBecomeTutor(w)
		}
	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		fmt.Println("Session Linktor does not exist")
	}
}

func (b Become) BecomeTutor() {
	db.AddNewTutor(b.Sender,b.Expertise, b.TBio)
}

func (info ThisUser) serveBecomeTutor(w http.ResponseWriter){
	t, err := template.ParseFiles("templates/p/become-a-tutor.html")
	smplErr(err, "Could not parse become-a-tutor")
	t.Execute(w,info)
}

func  (info ThisUser) serveGetTutor(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/p/get-tutor.html")
	smplErr(err, "Could not parse get-tutor.html")
	t.Execute(w, info)
}


func (r *Request) GenerateRID() {
	f2 := string(r.Sender[0:2])
	l2 :=  string(r.Sender[len(r.Sender)-2:])
	rid:= f2 + string(rand.Int31n(99)) + l2
	fmt.Println("RID generated is: ", rid)
	r.RID = rid
}

func (re Request) SaveRequest(w http.ResponseWriter, r *http.Request) {
	db.PersistRequest(re.Sender, re.Course, re.Category, re.Institution, re.Description, re.RID)
}
