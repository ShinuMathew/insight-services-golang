package SyncTestCaseProcessor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Shinz9474/InsightAps/InsightAPI/Logger"
	"github.com/Shinz9474/InsightAps/InsightAPI/Plugins/CSVHandler"
	"github.com/Shinz9474/InsightAps/InsightAPI/Plugins/DBConnector"
)

type TestCaseSteps_DB struct {
	StepID          string
	Step_ID         string
	Test_Step_Desc  string
	Keyword         string
	LocatorType     string
	Target          string
	Value_s         string
	Comment         string
	Tc_name         string
	Country         string
	Environment     string
	Created_date    string
	Modified_date   string
	TCDesign_status string
}

type TestCaseSteps_CSV struct {
	StepID       string
	Test_step    string
	Keyword      string
	Locator_type string
	Target       string
	Value        string
	Comments     string
	Status       string
}

type CSV_To_MySQL struct {
	Processor_name string
	Source         struct {
		Source_type     string
		Source_location string
	}
	Processor struct {
		Processor_name string
	}
	Target struct {
		Target_type   string
		Database_name string
		Table_name    string
	}
}

type TC_Info struct {
	TC_name     string
	TC_ID       string
	Country     string
	Environment string
	Project     string
	TC_Type     string
	TC_desc     string
}

func Start_Processor(tc_name string) {
	var tc_info TC_Info
	var csv_to_mysql CSV_To_MySQL
	var tc_design []TestCaseSteps_CSV
	DBConnector.ReadDBConfig("Plugins/DBConnector/MySQL_dbconfig.json")
	//Read processor config
	fileName := "Processors/SyncTestCaseProcessor/SyncTestCaseProcessor.json"
	processorConfig, err := os.Open(fileName)
	Logger.CheckError(err, "Following error occured when trying to parse the file: "+fileName)
	defer processorConfig.Close()
	//Parse processor config to local struct
	jsonFileParser := json.NewDecoder(processorConfig)
	err = jsonFileParser.Decode(&csv_to_mysql)
	Logger.CheckError(err, "Following error occured when trying to parse the file: "+fileName)
	//Query DB and check existing records for selected test case
	query := "select * from TC_Design where tc_name='" + tc_name + "';"
	db := DBConnector.Open_DBConnection()
	result := DBConnector.Execute_query(db, query)
	DBConnector.Close_DBConnection(db)
	if !result.Next() {
		log.Printf("There is no record existing for the selected Test Case: '%s'. \nProcessing fresh data transfer... \n", tc_name)
	} else {
		log.Printf("Updating DB with steps for the selected Test Case: '%s'.\n", tc_name)
		//Delete existing rows if already exists
		query = "Delete from TC_Design where tc_name='" + tc_name + "';"
		db = DBConnector.Open_DBConnection()
		result = DBConnector.Execute_query(db, query)
		DBConnector.Close_DBConnection(db)
	}
	//Split test case name to get the TC params
	TCName_Split := strings.Split(tc_name, "_")
	tc_info = TC_Info{TC_name: tc_name, Country: TCName_Split[0], Environment: TCName_Split[1], Project: TCName_Split[2],
		TC_Type: TCName_Split[3], TC_desc: TCName_Split[4]}
	//Generate TC CSV file location string
	fileLocation := csv_to_mysql.Source.Source_location + "/" + tc_info.Country + "/" + tc_info.Environment + "/" + tc_info.TC_name + "/" + tc_info.TC_name + ".csv"
	//Read CSV
	reader := CSVHandler.ReadCSV(fileLocation)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			Logger.CheckError(err, "Error occured when reading from CSV")
		}
		tc_design = append(tc_design, TestCaseSteps_CSV{StepID: line[0], Test_step: line[1], Keyword: line[2], Locator_type: line[3],
			Target: line[4], Value: line[5], Comments: line[6], Status: line[7]})
	}
	db = DBConnector.Open_DBConnection()
	//Insert test steps to DB
	for _, v := range tc_design {
		fmt.Printf("\nInserting test step: %v\n", v)
		v.Test_step = strings.Replace(v.Test_step, "'", "''", -1)
		v.Target = strings.Replace(v.Target, "'", "''", -1)
		v.Comments = strings.Replace(v.Comments, "'", "''", -1)
		v.Keyword = strings.Replace(v.Keyword, "'", "''", -1)
		v.Locator_type = strings.Replace(v.Locator_type, "'", "''", -1)
		if v.StepID != "StepID" {
			query = "Insert into TC_Design(TC_ID, STEP_ID, TEST_STEP_DESC, KEYWORD, LOCATOR_TYPE, TARGET, VALUE_S, COMMENTS, TC_NAME, COUNTRY, ENVIRONMENT, STEP_STATUS)" +
				"values((select tc_id from Test_cases where tc_name='" + tc_name + "'), " + v.StepID + ",  '" + v.Test_step + "', '" + v.Keyword + "', '" + v.Locator_type + "'" +
				" ,'" + v.Target + "', '" + v.Value + "', '" + v.Comments + "', '" + tc_name + "', '" + tc_info.Country + "', '" + tc_info.Environment + "', '" + v.Status + "');"
			insert_result := DBConnector.Execute_query(db, query)
			if insert_result.Err() != nil {
				log.Fatalf("Error occured when inserting the Step: '%v' record in DB.\n\n Insert Query: \n%v", v.Test_step, query)
			} else {
				log.Printf("Step: '%v' for TestCase: '%v' inserted in DB", v.Test_step, tc_name)
			}
		}
	}
	db.Close()
	DBConnector.Close_DBConnection(db)
	defer DBConnector.Close_DBConnection(db)
}
