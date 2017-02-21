package sess


import(
	//"github.com/gorilla/securecookie"
	"net/http"
	"github.com/gorilla/sessions"
	"fmt"
)

var store = sessions.NewCookieStore([]byte("elyoninternationalchristiancentreisdope"))
var activeUser string

func SaveSession(w http.ResponseWriter, r *http.Request, nick string) {
	session, err := store.Get(r, "Linktor")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["id"] = nick
	session.Values["LoggedIN"] = true
	session.Save(r,w)
	fmt.Println("Session has been saved as:\n", )
	fmt.Sprint(session.Values["id"])

}

func InitSessionValues(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "Linktor")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	activeUser = session.Values["id"].(string)
}

func SessionExists(w http.ResponseWriter, r *http.Request)bool{
	session, err := store.Get(r, "Linktor")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	if len(session.Values) > 0{
		return true
	}else{
		return false
	}
}

func GetSessionNick(w http.ResponseWriter, r *http.Request)string{
		return activeUser
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "Linktor")
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(r,w)
}

// cookie, err := r.Cookie("Linktor")
// 	if err != nil{
// 		fmt.Println("Logout issue from DeleteSession")
// 	}
// 	cookie.MaxAge = -1
// 	http.Redirect(w,r, "/", http.StatusFound)





// var hashKey = []byte("elyoninternationalchristiancentreisdope")
// var blockKey = []byte("trustmechristsavesandthatisallthatmatters")
// var s = securecookie.New(hashKey, blockKey)

// var uNick string

// func SetNewCookie(w http.ResponseWriter, r *http.Request, n string) {
// 	value := map[string]string{
// 		"nick" : n,
// 	}
// 	if encoded, err := s.Encode("Linktor", value); err == nil{
// 		cookie := &http.Cookie{
// 			Name: "Linktor",
// 			Value: encoded,
// 			Path: "/",
// 		}
// 		http.SetCookie(w, cookie)
// 	}
// }
// func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
// 	if cookie, err := r.Cookie("Linktor"); err == nil{
// 		value := make(map[string]string)
// 		if err = s.Decode("Linktor", cookie.Value, &value); err == nil{
// 			uNick = value["nick"]
// 		}
// 	}
// }

// func GetSessionNick(w http.ResponseWriter, r *http.Request)string {
// 	ReadCookieHandler(w,r)
// 	return uNick
// }
