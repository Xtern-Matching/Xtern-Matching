package tests

import (
	//"os"

	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"os"
)

func createStudent(ctx context.Context) (models.Student, error) {
	var student models.Student
	student.FirstName = "Darla"
	student.LastName = "leach"
	student.Email = "darlaleach@stockpost.com"
	student.University = "Rose-Hulman Institute of Technology"
	student.Major = "Computer Engineering"
	student.GradYear = "2017"
	student.WorkStatus = "US Citizen"
	student.HomeState = "West Virginia"
	student.Gender = "female"
	student.Skills = []models.Skill{{Name: "SQL", Category: "Database"}, {Name: "HTML", Category: "Frontend"}}
	student.Github = "https://github.com/xniccum"
	student.Linkin = ""
	student.PersonalSite = ""
	student.Interests = []string{"Product Management", "Software Engineer- Middle-tier Dev."}
	student.Grade = 5
	student.Status = "Stage 1 Approved"
	student.Resume = "public/data_mocks/sample.pdf"
	student.Active = true

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	key, err := datastore.Put(ctx, key, &student)
	if err != nil {
		return models.Student{}, err
	}
	// time.Sleep(time.Millisecond * 500)

	return student, nil
}

func TestAddMappedStudent(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err.Error())
	}
	defer done()

	studentJson := `{
		"firstName": "Alison",
			"lastName": "Hurley",
			"email": "alisonhurley@retrack.com",
			"university": "Rose-Hulman Institute of Technology",
			"major": "Design",
			"gradYear": "2017",
			"workStatus": "US Citizen",
			"gender": "female",
			"skills": [
	],
		"resume": "",
			"githubUrl": "https://github.com/Sp4rkfun",
			"linkedinUrl": null,
			"personalWebiteUrl": null,
			"interestedIn": [
		"Security",
			"Product Management",
			"Software Engineer- Middle-tier Dev."
		],
		"status": "Remaining",
			"interestedInEmail": "true",
			"homeState": "Delaware"
	}`
	//var studentI map[string]interface{}
	studentI := `{
	"firstName" : "firstName",
	 "lastName" : "lastName",
	 "email" : "email",
	 "university" : "university",
	 "major" : "major",
	}`
	var studentMapping map[string]string
	json.Unmarshal([]byte(studentI), &studentMapping)
	var studentJSON map[string]interface{}
	json.Unmarshal([]byte(studentJson), &studentJSON)
	_, err = services.AddMappedStudent(ctx, studentMapping, studentJSON)

	students, _, _ := services.GetStudents(ctx, nil)
	json.NewEncoder(os.Stdout).Encode(students)
	//{
	//	"name": "C#",
	//	"category": "General"
	//},
	//{
	//"name": ".Net",
	//"category": "Full-Stack"
	//}
}

//
//import (
//	//"os"
//	//"Xtern-Matching/handlers/services"
//	"Xtern-Matching/models"
//	//"reflect"
//	"testing"
//	//"time"
//	"google.golang.org/appengine/aetest"
//	"github.com/stretchr/testify/assert"
//	"Xtern-Matching/handlers/services"
//	"os"
//)
//
///*
//	Creates a fresh sample student to use in testing
//*/
//func GetStudent1() models.Student {
//	var student models.Student
//	student.FirstName = "Darla"
//	student.LastName = "leach"
//	student.Email = "darlaleach@stockpost.com"
//	//student.University = "Rose-Hulman Institute of Technology"
//	//student.Major = "Computer Engineering"
//	student.GradYear = "2017"
//	//student.WorkStatus = "US Citizen"
//	//student.HomeState = "West Virginia"
//	//student.Gender = "female"
//	student.Skills = []models.Skill{{Name: "SQL", Category: "Database"}, {Name: "HTML", Category: "Frontend"}}
//	//student.Github = "https://github.com/xniccum"
//	//student.Linkin = ""
//	//student.PersonalSite = ""
//	//student.Interests = []string{"Product Management", "Software Engineer- Middle-tier Dev."}
//	//student.EmailIntrest = "false"
//	//student.R1Grade = models.Grade{Text: "C", Value: "8"}
//	student.Status = "Stage 1 Approved"
//	student.Resume = "public/data_mocks/sample.pdf"
//	student.Active = true
//	return student
//}
//
///*
//	Tests the ability to create a Student
//*/
//func TestPost(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//
//	file, err := os.Open("resources/sample.pdf")
//	if err != nil {
//
//	}
//
//	student := GetStudent1()
//
//	// Basic input
//	_, err = services.NewStudent(ctx, student)
//	//if !assert.Nil(t, err, "Error creating student") {
//	//	t.Fatal(err)
//	//}
//	//
//	//// Check for bad input
//	//empty_student := models.Student{}
//	//status, err := services.NewStudent(ctx, empty_student)
//	//if !assert.Equal(t, status, http.StatusBadRequest, "Accepted empty student") {
//	//	t.Fatal(err)
//	//}
//}
//
///*
//	Tests both bulk retrieval and individual query.
//	Additionally ensures an invalid id can't be retrieved.
//*/
//func TestGet(t *testing.T) {
//	//ctx, done, err := aetest.NewContext()
//	//if !assert.Nil(t, err, "Error instantiating context") {
//	//	t.Fatal(err)
//	//}
//	//defer done()
//	//
//	//// Add student to get
//	//student := GetStudent1()
//	//
//	//_, err = services.NewStudent(ctx, student)
//	//if !assert.Nil(t, err, "Error creating student") {
//	//	t.Fatal(err)
//	//}
//	//
//	//student.FirstName = "Lee"
//	//student.LastName = "Robinson"
//	//student.Email = "leeson@stockpost.com"
//	//
//	//_, err = services.NewStudent(ctx, student)
//	//if !assert.Nil(t, err, "Error creating student") {
//	//	t.Fatal(err)
//	//}
//	//
//	//time.Sleep(time.Millisecond * 500)
//	//
//	//students, err := services.GetStudents(ctx)
//	//if !assert.Nil(t, err, "Error retrieving students") {
//	//	t.Fatal(err)
//	//}
//	//if !assert.Equal(t, 2, len(students), "Error in number of students retrieved") {
//	//	t.Fatal(students)
//	//}
//	//
//	//first := students[0].FirstName
//	//last := students[0].LastName
//	//student, err = services.GetStudent(ctx, id)
//	//if !assert.Nil(t, err, "Error retrieving student") {
//	//	t.Fatal(err)
//	//}
//	//if !assert.Equal(t, first, student.FirstName, "Mismatch in student first name") ||
//	//	!assert.Equal(t, last, student.LastName, "Mismatch in student last name") {
//	//	t.Fatal("Grabbed wrong student")
//	//}
//	//
//	//student, err = services.GetStudent(ctx, -4)
//	//if !reflect.DeepEqual(student, (models.Student{})) {
//	//	t.Fatal("student shouldn't exist")
//	//}
//
//}
//
//func TestResumePost(t *testing.T) {		 func TestResumePost(t *testing.T) {
//	// Due to current dependecy credentials, turning off this test until better alternative can be found		 	// Due to current dependecy credentials, turning off this test until better alternative can be found
//	//ctx, done, err := aetest.NewContext()		 	//ctx, done, err := aetest.NewContext()
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//	//defer done()		 	//defer done()
//	//		 	//
//	//student, err := createStudent(ctx)		 	//student, err := createStudent(ctx)
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//	//		 	//
//	//file, err := os.Open("resources/sample.pdf")		 	//file, err := os.Open("resources/sample.pdf")
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//	//defer file.Close()		 	//defer file.Close()
//	//		 	//
//	//err = services.UpdateResume(ctx, student.Id, file)		 	//err = services.UpdateResume(ctx, student.Id, file)
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//}
