package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shinz9474/InsightAps/InsightAPI/DBConnect"
	"github.com/Shinz9474/InsightAps/InsightAPI/Processors/SyncTestCaseProcessor"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Project struct {
	Project_ID     string //`xml:project_id`     //, `json:project_id`
	Project_name   string //`xml:project_name`   //, `json:project_name`
	Project_desc   string //`xml:project_desce`  //, `json:project_desce`
	Created_Date   string //`xml:created_date`   //, `json:created_date`
	Modified_Date  string //`xml:modified_date`  //, `json:modified_date`
	Project_status string //`xml:project_status` //, `json:project_status`
}

type Projects struct {
	Total_Projects int       `json:total_projects`
	Projects       []Project `json:projects`
}

type TestCase struct {
	TC_ID         string `json:"tc_id"`
	TC_NAME       string `json:"tc_name"`
	TC_DESC       string `json:"tc_desc"`
	COUNTRY       string `json:"country"`
	ENVIRONMENT   string `json:"environment"`
	CREATED_DATE  string `json:"tc_created_date"`
	MODIFIED_DATE string `json:"tc_modified_date"`
	TC_STATUS     string `json:"tc_status"`
}

type TestCases struct {
	Total_TCs int        `json:"total_TCs"`
	TestCases []TestCase `json:"testCases"`
}

type CreatedTestCase struct {
	Status           string   `json:status`
	TestCase_Created TestCase `json:testCase_Created`
}

type Messenger struct {
	Processing_message string `json:processing_message`
}

type Users struct {
	ID            int    `json:id`
	USER_ID       string `json:user_id`
	Username      string `json:username`
	User_Password string `json:password`
	User_Status   string `json:user_status`
}

type SyncTC struct {
	TestCases []string `json:testcases`
}

type SyncSpecificTC struct {
	TestCase string `json:testcase`
}

//Users
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	var message Messenger
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("'Access-Control-Allow-Origin", "*")
	username := r.FormValue("username")
	password := r.FormValue("password")
	query := "select User_Status from Users where Username='" + username + "' and User_Password='" + password + "';"
	result := DBConnect.Execute_Query("mylearning", query)
	if result.Next() {
		err := result.Scan(&user.User_Status)
		if err != nil {
			panic(err.Error())
		}
		log.Println("user status: " + user.User_Status)
		if user.User_Status == "Active" {
			message = Messenger{"Valid User"}
		} else {
			message = Messenger{"Inactive User"}
		}
	} else {
		message = Messenger{"Invalid User"}
	}
	json.NewEncoder(w).Encode(message)
}

//Projects
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var prj []Project
	var projects Projects
	//w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Content-Type", "application/json")
	query := "Select * from PROJECTS"
	result := DBConnect.Execute_Query("mylearning", query)
	var count int
	for result.Next() {
		var project Project
		err := result.Scan(&project.Project_ID, &project.Project_name, &project.Project_desc, &project.Created_Date,
			&project.Modified_Date, &project.Project_status)
		if err != nil {
			panic(err.Error)
		}
		count++
		prj = append(prj, Project{Project_ID: project.Project_ID, Project_name: project.Project_name, Project_desc: project.Project_desc,
			Created_Date: project.Created_Date, Modified_Date: project.Modified_Date, Project_status: project.Project_status})
	}
	projects = Projects{count, prj}
	//xml.NewEncoder(w).Encode(projects)
	json.NewEncoder(w).Encode(projects)
}

//Get all test cases
func GetTestCases(w http.ResponseWriter, r *http.Request) {
	var tc []TestCase
	var tcs TestCases
	//1. Set header, content-type as json
	w.Header().Set("Content-Type", "application/json")
	query := "select * from Test_cases"
	//2. Execute the query and get the result
	result := DBConnect.Execute_Query("mylearning", query)
	var count int
	//3. Read the result
	for result.Next() {
		var testCase TestCase
		//4. Scan copies the columns in the current row into the values pointed at by dest.
		//	 The number of values in dest must be the same as the number of columns in Rows.
		err := result.Scan(&testCase.TC_ID, &testCase.TC_NAME, &testCase.TC_DESC, &testCase.CREATED_DATE,
			&testCase.MODIFIED_DATE, &testCase.TC_STATUS, &testCase.COUNTRY, &testCase.ENVIRONMENT)
		if err != nil {
			panic(err.Error)
		}
		count++
		//5. Appending the value to TestCase slice
		tc = append(tc, TestCase{TC_ID: testCase.TC_ID, TC_NAME: testCase.TC_NAME,
			TC_DESC: testCase.TC_DESC, COUNTRY: testCase.COUNTRY, ENVIRONMENT: testCase.ENVIRONMENT,
			CREATED_DATE: testCase.CREATED_DATE, MODIFIED_DATE: testCase.MODIFIED_DATE, TC_STATUS: testCase.TC_STATUS})
	}
	//6. Create the TestCases type tcs for response
	tcs = TestCases{count, tc}
	//7. Encoding the tcs to the response writer
	json.NewEncoder(w).Encode(tcs)
}

//Get specific test case
func GetTestCase(w http.ResponseWriter, r *http.Request) {
	var tc []TestCase
	var tcs TestCases
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	log.Printf("Params: %v", params)
	query := "Select * from Test_Cases where tc_id =" + params["tc_id"] + ";"
	result := DBConnect.Execute_Query("mylearning", query)
	var count int
	for result.Next() {
		var testCase TestCase
		err := result.Scan(&testCase.TC_ID, &testCase.TC_NAME, &testCase.TC_DESC,
			&testCase.CREATED_DATE, &testCase.MODIFIED_DATE, &testCase.TC_STATUS, &testCase.COUNTRY, &testCase.ENVIRONMENT)
		if err != nil {
			panic(err.Error)
		}
		count++
		tc = append(tc, TestCase{TC_ID: testCase.TC_ID, TC_NAME: testCase.TC_NAME,
			TC_DESC: testCase.TC_DESC, COUNTRY: testCase.COUNTRY, ENVIRONMENT: testCase.ENVIRONMENT,
			CREATED_DATE: testCase.CREATED_DATE, MODIFIED_DATE: testCase.MODIFIED_DATE, TC_STATUS: testCase.TC_STATUS})
	}
	tcs = TestCases{count, tc}
	json.NewEncoder(w).Encode(tcs)
}

//Created new testcase
func CreateTestCase(w http.ResponseWriter, r *http.Request) {
	var tci TestCase
	var ctc CreatedTestCase
	w.Header().Set("Content-type", "application/json")
	var tc1 TestCase
	_ = json.NewDecoder(r.Body).Decode(&tc1)
	query := "insert into Test_cases(TC_NAME, TC_DESC, COUNTRY, ENVIRONMENT, TC_STATUS) values('" + tc1.TC_NAME + "', '" + tc1.TC_DESC + "', '" + tc1.COUNTRY + "', '" + tc1.ENVIRONMENT + "', '" + tc1.TC_STATUS + "');"
	DBConnect.Execute_Query("mylearning", query)
	query1 := "Select * from Test_Cases where tc_name = '" + tc1.TC_NAME + "'"
	result1 := DBConnect.Execute_Query("mylearning", query1)
	var testcase TestCase
	for result1.Next() {
		err := result1.Scan(&testcase.TC_ID, &testcase.TC_NAME, &testcase.TC_DESC,
			&testcase.CREATED_DATE, &testcase.MODIFIED_DATE, &testcase.TC_STATUS, &testcase.COUNTRY, &testcase.ENVIRONMENT)
		if err != nil {
			panic(err.Error())
		}
		tci = TestCase{testcase.TC_ID, testcase.TC_NAME, testcase.TC_DESC, testcase.COUNTRY, testcase.ENVIRONMENT, testcase.CREATED_DATE, testcase.MODIFIED_DATE, testcase.TC_STATUS}
	}
	ctc = CreatedTestCase{"Successfull", tci}
	json.NewEncoder(w).Encode(ctc)
}

//Update existing testcase
func UpdateTestCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var tc1 TestCase
	var pm Messenger
	_ = json.NewDecoder(r.Body).Decode(&tc1)
	test_query := "Select * from test_cases where tc_id = '" + params["tc_id"] + "';"
	test_result := DBConnect.Execute_Query("mylearning", test_query)
	if test_result.Next() != true {
		http.NotFound(w, r)
		pm = Messenger{"Invalid tc_id. " + params["tc_id"] + " does not exists"}
	} else {
		query := "update Test_cases set TC_NAME='" + tc1.TC_NAME + "', TC_DESC='" + tc1.TC_DESC + "', COUNTRY='" + tc1.COUNTRY + "', ENVIRONMENT='" + tc1.ENVIRONMENT + "', TC_STATUS='" + tc1.TC_STATUS + "'" +
			" where tc_id='" + params["tc_id"] + "';"
		DBConnect.Execute_Query("mylearning", query)
		pm = Messenger{params["tc_id"] + ", Updated successfully"}
	}
	json.NewEncoder(w).Encode(pm)
}

//Delete existing testcase
func DeleteTestCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var tc1 TestCase
	var pm Messenger
	_ = json.NewDecoder(r.Body).Decode(&tc1)
	test_query := "Select * from test_cases where tc_id = '" + params["tc_id"] + "'"
	test_result := DBConnect.Execute_Query("mylearning", test_query)
	if test_result.Next() != true {
		http.NotFound(w, r)
		pm = Messenger{"Invalid tc_id. " + params["tc_id"] + " does not exists"}
	} else {
		query := "Delete from Test_cases where tc_id='" + params["tc_id"] + "';"
		DBConnect.Execute_Query("mylearning", query)
		pm = Messenger{params["tc_id"] + ", Deleted successfully"}
	}
	json.NewEncoder(w).Encode(pm)
}

//Sync multiple testcases
func SyncTestCases(w http.ResponseWriter, r *http.Request) {
	var synctc SyncTC
	var message Messenger
	w.Header().Set("Content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&synctc)
	fmt.Printf("Test Cases to be synced: \n%v", synctc)
	for _, v := range synctc.TestCases {
		SyncTestCaseProcessor.Start_Processor(v)
	}
	message = Messenger{"Test Cases synced successfully"}
	json.NewEncoder(w).Encode(message)
}

//Sync specific testcase
func SyncSpecificTestCase(w http.ResponseWriter, r *http.Request) {
	var syncspecifictc SyncSpecificTC
	var message Messenger
	w.Header().Set("Content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&syncspecifictc)
	fmt.Printf("Test Cases to be synced: \n%v", syncspecifictc)
	SyncTestCaseProcessor.Start_Processor(syncspecifictc.TestCase)
	message = Messenger{"Test Cases synced successfully"}
	json.NewEncoder(w).Encode(message)
}

//GetAllTestSteps
//GetSpecificTestStep
//CreateTestStep
//UpdateTestStep
//DeleteTestStep

func main() {

	//Init router
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELTE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//Users route handlers
	r.HandleFunc("/user/authenticate", AuthenticateUser).Methods("GET")

	//Sync test cases
	r.HandleFunc("/api/utils/synctestcases", SyncTestCases).Methods("POST")
	r.HandleFunc("/api/utils/syncspecifictestcases", SyncSpecificTestCase).Methods("POST")

	//Project route handlers
	r.HandleFunc("/api/projects", GetAllProjects).Methods("GET")

	//TestCases Route Handlers
	r.HandleFunc("/api/testcases", GetTestCases).Methods("GET")
	r.HandleFunc("/api/testcases", CreateTestCase).Methods("POST")
	r.HandleFunc("/api/testcases/{tc_id}", GetTestCase).Methods("GET")
	r.HandleFunc("/api/testcases/{tc_id}", UpdateTestCase).Methods("PUT")
	r.HandleFunc("/api/testcases/{tc_id}", DeleteTestCase).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r)))

}
