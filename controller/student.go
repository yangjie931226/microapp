package controller

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"web/model"
	"web/proto/grades"
	"web/utils"
)

var rootTemplate *template.Template

func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"../../views/students.html",
		"../../views/student.html")

	if err != nil {
		return err
	}

	return nil
}

type StudentsHandler struct{}

func (sh StudentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2: // /students
		sh.renderStudents(w, r)
	case 3: // /students/{:id}
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderStudent(w, r, id)
	case 4: // /students/{:id}/grades
		id, err := strconv.Atoi(pathSegments[2])
		id = id
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.ToLower(pathSegments[3]) != "grades" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderGrades(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (sh StudentsHandler) renderStudents(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error retrieving students: ", err)
		}
	}()

	client := utils.GetMicroClient()

	gradeservice := grades.NewGradesService("go.micro.gradeservice", client)
	resp, err := gradeservice.GetAllStudents(context.Background(), &grades.GetAllRequest{})

	if err != nil {
		return
	}

	var s model.Students
	for _, stu := range resp.Data {
		dataGrades := []model.Grade{}
		for _, stuGrades := range stu.Grades {
			dataGrade := model.Grade{
				Title: stuGrades.Title,
				Type:  model.GradeType(stuGrades.Type),
				Score: stuGrades.Score,
			}
			dataGrades = append(dataGrades, dataGrade)
		}
		dataStu := model.Student{
			ID:        int(stu.Id),
			FirstName: stu.FirstName,
			LastName:  stu.LastName,
			Grades:    dataGrades,
		}
		s = append(s, dataStu)
	}

	fmt.Println(s)

	rootTemplate.Lookup("students.html").Execute(w, s)
}

func (sh StudentsHandler) renderStudent(w http.ResponseWriter, r *http.Request, id int) {

	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error retrieving students: ", err)
			return
		}
	}()

	client := utils.GetMicroClient()
	gradeservice := grades.NewGradesService("go.micro.gradeservice", client)

	resp, err := gradeservice.GetOneStudent(context.Background(), &grades.IdRequest{Id: int32(id)})

	if err != nil {
		return
	}

	var s model.Student
	dataGrades := []model.Grade{}
	for _, stuGrades := range resp.Data.Grades {
		dataGrade := model.Grade{
			Title: stuGrades.Title,
			Type:  model.GradeType(stuGrades.Type),
			Score: stuGrades.Score,
		}
		dataGrades = append(dataGrades, dataGrade)
	}
	s = model.Student{
		ID:        int(resp.Data.Id),
		FirstName: resp.Data.FirstName,
		LastName:  resp.Data.LastName,
		Grades:    dataGrades,
	}

	rootTemplate.Lookup("student.html").Execute(w, s)
}

func (sh StudentsHandler) renderGrades(w http.ResponseWriter, r *http.Request, id int) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		w.Header().Add("location", fmt.Sprintf("/students/%v", id))
		w.WriteHeader(http.StatusTemporaryRedirect)
	}()
	title := r.FormValue("Title")
	gradeType := r.FormValue("Type")
	score, err := strconv.ParseFloat(r.FormValue("Score"), 32)
	if err != nil {
		log.Println("Failed to parse score: ", err)
		return
	}
	g := &grades.Grade{
		Title: title,
		Type:  gradeType,
		Score: float32(score),
	}

	client := utils.GetMicroClient()

	gradeservice := grades.NewGradesService("go.micro.gradeservice", client)
	_, err = gradeservice.AddGrade(context.Background(), &grades.GradeRequest{
		Id: int32(id),
		Grade: g,
	})


	if err != nil {
		log.Println("Failed to save grade to Grading Service", err)
		return
	}

}
