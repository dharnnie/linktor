package profile

import(
 	"log"
 	"net/http"
 	"os"
 	"github.com/dharnnie/linktor/sess"
 	"io"
 	"github.com/dharnnie/linktor/db"
 	"fmt"
 	"path/filepath"
)

func ProcessPicUpdate(w http.ResponseWriter, r *http.Request) {
 	r.ParseMultipartForm(32 << 20)

 	src, hdr, err := r.FormFile("pic")
 	smplErr2(w,err, "Error occured at r.FormFile()")
 	defer src.Close()

	//rename the file
	log.Println("Name of uploaded file is: ",hdr.Filename)
	var flName string
	var extension string
	extension = getExt(hdr.Filename)

	// plan to check for incompatible file types

 	if sess.SessionExists(w,r){
 		sess.InitSessionValues(w,r)
 		
 		flName = sess.GetSessionNick(w,r)
 		log.Println("The renamed file is: ", flName)
 		flName = flName + extension
 	}
	
 	dst, err := os.Create("assets/imgs/" + flName)
 	if err != nil {
         fmt.Println(err)
         return
     }
     defer dst.Close()
     io.Copy(dst, src)
     saveImagePath(sess.GetSessionNick(w,r), flName)
}

func saveImagePath(n string, flNm string) {
	db.UpdatePic(n, flNm)
}

func GetPicPath(n string)string {
	return db.GetImage(n)
}

func getExt(f string)string {
	ext := filepath.Ext(f)
	return ext
}

func smplErr2(w http.ResponseWriter, e error, m string){
 	if e != nil{
 		log.Println(e)
 		http.Error(w, "Error uploading picture", http.StatusInternalServerError)
 	}
}