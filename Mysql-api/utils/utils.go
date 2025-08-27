package unmarsh

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//changes from json-client format to struct/db format

func ChangeFmt(r*http.Request,t interface{}){
	body,err := ioutil.ReadAll(r.Body)
	if err == nil{
		if err := json.Unmashal([]byte(body),t);err != nil{
			return
		}
	}
}