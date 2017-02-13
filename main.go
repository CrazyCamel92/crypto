package main

import ("fmt"
"html"
"log"
	"net/http"
	"crypto/sha256"
	"os"
	"crypto/rand"
	"crypto/rsa"
)


func main (){
	fmt.Print("Starting Server...");
	Pvk,Prk := generateKeys();
	startServer(Prk,Pvk);


}
func startServer(prk rsa.PrivateKey,pvk rsa.PublicKey){
http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*");
	fmt.Printf("listening on port 8080");
	plain := r.URL.RawPath;
	fmt.Printf("encrypting "+plain)
	data := encrypt(plain,pvk);
	fmt.Fprintf(w, "plain: "+plain+ " encrypted value: " + data, html.EscapeString(r.URL.Path));
});
log.Fatal(http.ListenAndServe(":8080",nil));
}

func encrypt(messageText string,key rsa.PublicKey) string {
	message := []byte(messageText);
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, &key, message, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return  string(ciphertext);
}
func generateKeys () (rsa.PublicKey,rsa.PrivateKey){
	jimenaPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	jimenaPublicKey := &jimenaPrivateKey.PublicKey
	return *jimenaPublicKey,*jimenaPrivateKey;
}