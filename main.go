package main

import(
	"gopkg.in/couchbase/gocb.v1"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var bucket *gocb.Bucket
// defining a user profile with bare minimum info

type UserProfile struct {
	Username string `json:"username"`
        Firstname string `json:"firstname"`
        Lastname  string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main(){
	
	f, err := os.Open("test.json")
	defer f.Close()
	if err!=nil {
		log.Fatal(err)
	}
	// f is a File* which is a Reader type
	var datamap map[string]interface{}
	decoder := json.NewDecoder(f)
	err  = decoder.Decode(&datamap)
	if err!=nil{
		 log.Fatal(err)
	}
	userprofile := UserProfile{ Username : datamap["username"].(string),
                                    Firstname : datamap["firstname"].(string),
				    Lastname : datamap["lastname"].(string),
				    Email : datamap["email"].(string),
	    			    Password : datamap["password"].(string)}
	myCluster, _ := gocb.Connect("couchbase://localhost")
        myCluster.Authenticate(gocb.PasswordAuthenticator{
            Username: "SuperRootUser",
            Password: "dummypwd",
        })
	bucket, _ = myCluster.OpenBucket("travel-data", "")
	_, err = bucket.Upsert(datamap["username"].(string), userprofile, 0)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("exiting main")
}
