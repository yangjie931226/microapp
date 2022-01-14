package handler

import (
	"context"
	"grades/model"
	pb "grades/proto"
)

type Grades struct{}

func (*Grades) GetAllStudents(ctx context.Context, req *pb.GetAllRequest, resp *pb.StudentsReply) error {
	data := []*pb.Student{}
	for _, stu := range model.AllStudents {
		dataGrades := []*pb.Grade{}
		for _, stuGrades := range stu.Grades {
			dataGrade := &pb.Grade{
				Title: stuGrades.Title,
				Type:  string(stuGrades.Type),
				Score: stuGrades.Score,
			}
			dataGrades = append(dataGrades, dataGrade)
		}
		dataStu := &pb.Student{
			Id:        int32(stu.ID),
			FirstName: stu.FirstName,
			LastName:  stu.LastName,
			Grades:    dataGrades,
		}
		data = append(data, dataStu)
	}
	resp.Code = 200
	resp.Msg = ""
	resp.Data = data

	return nil

}

func (*Grades) GetOneStudent(ctx context.Context, req *pb.IdRequest, resp *pb.StudentReply) error {
	id := int(req.Id)
	stu, err := model.AllStudents.GetByID(id)
	if err != nil {
		return err
	}
	dataGrades := []*pb.Grade{}
	for _, stuGrades := range stu.Grades {
		dataGrade := &pb.Grade{
			Title: stuGrades.Title,
			Type:  string(stuGrades.Type),
			Score: stuGrades.Score,
		}
		dataGrades = append(dataGrades, dataGrade)
	}
	data := &pb.Student{
		Id:        int32(stu.ID),
		FirstName: stu.FirstName,
		LastName:  stu.LastName,
		Grades:    dataGrades,
	}
	resp.Code = 200
	resp.Msg = ""
	resp.Data = data

	return nil
}

func (*Grades) AddGrade(ctx context.Context, req *pb.GradeRequest, resp *pb.GradesReply) error {
	id := int(req.Id)
	stu, err := model.AllStudents.GetByID(id)
	if err != nil {
		return err
	}

	g := model.Grade{
		Title: req.Grade.Title,
		Type:  model.GradeType(req.Grade.Type),
		Score: req.Grade.Score,
	}
	gResp := &pb.Grade{
		Title: req.Grade.Title,
		Type:  req.Grade.Type,
		Score: req.Grade.Score,
	}
	stu.Grades = append(stu.Grades, g)

	resp.Code = 200
	resp.Msg = ""
	resp.Data = gResp

	return nil
}
