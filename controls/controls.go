package ctrls

import(
	"github.com/gorilla/mux"
	"net/http"
	"db_mysql/db_models"
	"strconv"
	"db_mysql/utils"
)

var NewClass modls.Class

func GetClasses (w http.ResponseWriter, r*http.Request){
	classes:= modls.GetClassesDb()
	//convert from db format to json-client format
	resp,_ := json.Marshal(classes)
	w.Header().Set("content-type","applicaton/json")
	w.WriteHeader(http.StatusOK)
	w.write(resp)
}

func GetClass(w http.ResponseWriter,r*http.Request){
	//gets the request for url
	param := mux.Vars(r)
	//specifies the request required from url is ID 
	ID:= param["ID"]
	//converts the string ID to typeint
	Id,err := strconv.ParseInt(ID,0,0)
	if err != nil{
		fmt.Println("Paring error")
	}

	class,_ := modls.GetClassDb(Id)
	resp,_ := json.Marshal(class)

	w.Header().Set("content-type","applicaton/json")
	w.WriteHeader(http.StatusOK)
	w.write(resp)
}

func CreateClass(w http.ResponseWriter,r*http.Request){
	createCls:= &mdls.Class{}
	// convert the json requst to db format or that of createcls
	unmarsh.ChangeFmt(r,createCls)
	class := createCls.CreateClassDb()
	resp,_ := json.Marshal(class)

	w.WriteHeader(http.StatusOK)
	w.write(resp)

}

func DeleteClass(w http.ResponseWriter,r*http.Request){
	param := mux.Vars(r)
	ID := param["ID"]
	Id,err := strconv.ParseInt(ID,0,0)
	if err != nil{
		fmt.Println("Parsing error")
	}
	class := modls.DeleteClassDb(Id)
	resp,_ := json.Marshal(class)

	w.Header().Set("content-type","applicaton/json")
	w.WriteHeader(http.StatusOK)
	w.write(resp)
}

func UpdateClass (w http.ResponseWriter, r*http.Request){
	params := mux.Vars(r)
	updateCls := &modls.Class{}
	unmarsh.ChangeFmt(r,updateCls)

	ID := params["ID"]
	Id,err := strconv.ParseInt(ID,0,0)
	if err != nil{
		fmt.Println("Parsing Error")
	}
	pastClass,db := modls.GetClassDb(Id)

	//if the fields were specified in the request then update
	
	if updateCls.Name != ""{
		pastClass.Name = updateCls.Name
	}
	if updateCls.lesson != ""{
		pastClass.lesson = updateCls.lesson
	}
	if updateCls.Attendance != ""{
		pastClasss.Attendance = updateCls.Attendance 
	}
	db.Save(&pastClass)
	resp,_ := json.Marshal(pastClass)

	w.Header().Set("content-type","applicaton/json")
	w.WriteHeader(http.StatusOK)
	w.write(resp)
}
