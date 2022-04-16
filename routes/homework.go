package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreateCourseBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func CreateCourse(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	body := CreateCourseBody{}

	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid body"))

		return
	}

	course := models.Course{
		Name:        body.Name,
		Description: null.StringFromPtr(body.Description),
	}

	err = course.Insert(ctx.Request.Context(), db, boil.Infer())
	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	ctx.JSON(200, course)
}

type UpdateCourseBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func UpdateCourse(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	courseID := ctx.Param("course_id")

	courseIDNum, err := strconv.Atoi(courseID)

	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid course id"))

		return
	}

	course, err := models.FindCourse(ctx.Request.Context(), db, courseIDNum)
	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	body := UpdateCourseBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	course.Name = body.Name
	course.Description = null.StringFromPtr(body.Description)

	_, err = course.Update(ctx.Request.Context(), db, boil.Infer())
	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	ctx.JSON(200, course)
}

func GetHomeworks(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	homeworks, err := models.Homeworks().All(ctx.Request.Context(), db)

	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	ctx.JSON(200, homeworks)
}

type CreateHomeworkBody struct {
	Title       string  `json:"title"`
	Description *string `json:"body"`
}

func CreateHomework(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	courseID, err := getNumParam(ctx, "courseId")

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid course id"))
	}

	body := CreateHomeworkBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid body"))

		return
	}

	homework := models.Homework{
		CourseID:    courseID,
		Title:       body.Title,
		Description: null.StringFromPtr(body.Description),
	}

	err = homework.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	ctx.JSON(200, homework)
}

type UpdateHomeworkBody struct {
	Title       string  `json:"title"`
	Description *string `json:"body"`
}

func UpdateHomework(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	homeworkID := ctx.Param("homework_id")

	homeworkIDNum, err := strconv.Atoi(homeworkID)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid homework id"))

		return
	}

	homework, err := models.FindHomework(ctx.Request.Context(), db, homeworkIDNum)
	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	body := UpdateHomeworkBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	homework.Title = body.Title
	homework.Description = null.StringFromPtr(body.Description)

	_, err = homework.Update(ctx.Request.Context(), db, boil.Infer())
	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	ctx.JSON(200, homework)
}

func GetHomework(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	courseID, err := getNumParam(ctx, "courseId")

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid course id"))
	}

	homeworkID, err := getNumParam(ctx, "homeworkId")

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid homework id"))
	}

	homework, err := models.Homeworks(
		qm.Where("course_id = ?", courseID),
		qm.Where("id = ?", homeworkID),
	).One(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	ctx.JSON(200, homework)
}

func GetCourse(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	courseID := ctx.Param("course_id")

	courseIDNum, err := strconv.Atoi(courseID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid course id"})
	}

	course, err := models.FindCourse(ctx.Request.Context(), db, courseIDNum)
	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	ctx.JSON(200, course)
}

func GetCourses(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	courses, err := models.Courses().All(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	ctx.JSON(200, courses)
}

type SubmitHomeworkBody struct {
	Submission string `json:"body"`
}

func SubmitHomework(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	homeworkID := ctx.Param("homework_id")

	homeworkIDNum, err := strconv.Atoi(homeworkID)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid homework id"))

		return
	}

	homework, err := models.FindHomework(ctx.Request.Context(), db, homeworkIDNum)
	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	body := SubmitHomeworkBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	submission := models.HomeworkSubmission{
		HomeworkID: homework.ID,
		Submission: body.Submission,
	}

	err = submission.Insert(ctx.Request.Context(), db, boil.Infer())
	if err != nil {
		log.Printf("Error: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))

		return
	}

	ctx.JSON(200, submission)
}

func GetHomeworkSubmissions(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	homeworkID := ctx.Param("homework_id")

	homeworkIDNum, err := strconv.Atoi(homeworkID)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid homework id"))

		return
	}

	homework, err := models.FindHomework(ctx.Request.Context(), db, homeworkIDNum)
	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	submissions, err := models.HomeworkSubmissions(models.HomeworkSubmissionWhere.HomeworkID.EQ(homework.ID)).All(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
	}

	ctx.JSON(200, submissions)
}

type SubmitCommentBody struct {
	Comment string `json:"comment"`
}

func SubmitComment(ctx *gin.Context) (*models.HomeworkSubmissionComment, error) {
	db := GetDBFromContext(ctx)

	homeworkID := ctx.Param("homework_id")

	homeworkIDNum, err := strconv.Atoi(homeworkID)

	if err != nil {
		return nil, err
	}

	homework, err := models.FindHomework(ctx.Request.Context(), db, homeworkIDNum)
	if err != nil {
		return nil, err
	}

	body := SubmitCommentBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		return nil, err
	}

	comment := models.HomeworkSubmissionComment{
		HomeworkSubmissionID: homework.ID,
		Comment:              body.Comment,
	}

	err = comment.Insert(ctx.Request.Context(), db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
