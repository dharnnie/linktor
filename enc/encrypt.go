package enc

import(
	//"golang.org/x/crypto/bcrypt"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"log"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
func Encrypt(key, p  string ) string {
	block, err := aes.NewCipher([]byte(key))
	smplErr(err, "Could not create block - encrypt")
	plain := []byte(p)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plain))
	cfb.XORKeyStream(cipherText, plain)
	return encodeBase64(cipherText)
}

func Decrypt(key, p string)string {
	block, err := aes.NewCipher([]byte(key))
	smplErr(err, "Could not create block - encrypt")
	cipherText := decodeBase64(p)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plain := make([]byte, len(cipherText))
	cfb.XORKeyStream(plain, cipherText)
	return string(plain)
}

func encodeBase64(b []byte) string {                                                                                                                                                                        
    return base64.StdEncoding.EncodeToString(b)                                                                                                                                                             
} 

func decodeBase64(s string) []byte {                                                                                                                                                                        
    data, err := base64.StdEncoding.DecodeString(s)                                                                                                                                                         
    if err != nil { panic(err) }                                                                                                                                                                            
    return data                                                                                                                                                                                             
}    


// func Encrypt(p string)string{
// 	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
// 	smplErr(err, "Error occured during encryption")
// 	fmt.Printf("Encrypted value of %s is %s", p,hash)
// 	return hash
// }

// func Decrypt(h string)string {
	
// }




func smplErr(e error, m string){
	if e != nil{
		log.Println(m, e)
	}
}	