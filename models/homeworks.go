// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Homework is an object representing the database table.
type Homework struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	DueDate     time.Time   `boil:"due_date" json:"due_date" toml:"due_date" yaml:"due_date"`
	CourseID    int         `boil:"course_id" json:"course_id" toml:"course_id" yaml:"course_id"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *homeworkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L homeworkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HomeworkColumns = struct {
	ID          string
	Title       string
	Description string
	DueDate     string
	CourseID    string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Title:       "title",
	Description: "description",
	DueDate:     "due_date",
	CourseID:    "course_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var HomeworkTableColumns = struct {
	ID          string
	Title       string
	Description string
	DueDate     string
	CourseID    string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "homeworks.id",
	Title:       "homeworks.title",
	Description: "homeworks.description",
	DueDate:     "homeworks.due_date",
	CourseID:    "homeworks.course_id",
	CreatedAt:   "homeworks.created_at",
	UpdatedAt:   "homeworks.updated_at",
}

// Generated where

var HomeworkWhere = struct {
	ID          whereHelperint
	Title       whereHelperstring
	Description whereHelpernull_String
	DueDate     whereHelpertime_Time
	CourseID    whereHelperint
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint{field: "\"homeworks\".\"id\""},
	Title:       whereHelperstring{field: "\"homeworks\".\"title\""},
	Description: whereHelpernull_String{field: "\"homeworks\".\"description\""},
	DueDate:     whereHelpertime_Time{field: "\"homeworks\".\"due_date\""},
	CourseID:    whereHelperint{field: "\"homeworks\".\"course_id\""},
	CreatedAt:   whereHelpertime_Time{field: "\"homeworks\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"homeworks\".\"updated_at\""},
}

// HomeworkRels is where relationship names are stored.
var HomeworkRels = struct {
	Course              string
	HomeworkSubmissions string
}{
	Course:              "Course",
	HomeworkSubmissions: "HomeworkSubmissions",
}

// homeworkR is where relationships are stored.
type homeworkR struct {
	Course              *Course                 `boil:"Course" json:"Course" toml:"Course" yaml:"Course"`
	HomeworkSubmissions HomeworkSubmissionSlice `boil:"HomeworkSubmissions" json:"HomeworkSubmissions" toml:"HomeworkSubmissions" yaml:"HomeworkSubmissions"`
}

// NewStruct creates a new relationship struct
func (*homeworkR) NewStruct() *homeworkR {
	return &homeworkR{}
}

// homeworkL is where Load methods for each relationship are stored.
type homeworkL struct{}

var (
	homeworkAllColumns            = []string{"id", "title", "description", "due_date", "course_id", "created_at", "updated_at"}
	homeworkColumnsWithoutDefault = []string{"title", "description", "due_date", "course_id", "created_at", "updated_at"}
	homeworkColumnsWithDefault    = []string{"id"}
	homeworkPrimaryKeyColumns     = []string{"id"}
)

type (
	// HomeworkSlice is an alias for a slice of pointers to Homework.
	// This should almost always be used instead of []Homework.
	HomeworkSlice []*Homework
	// HomeworkHook is the signature for custom Homework hook methods
	HomeworkHook func(context.Context, boil.ContextExecutor, *Homework) error

	homeworkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	homeworkType                 = reflect.TypeOf(&Homework{})
	homeworkMapping              = queries.MakeStructMapping(homeworkType)
	homeworkPrimaryKeyMapping, _ = queries.BindMapping(homeworkType, homeworkMapping, homeworkPrimaryKeyColumns)
	homeworkInsertCacheMut       sync.RWMutex
	homeworkInsertCache          = make(map[string]insertCache)
	homeworkUpdateCacheMut       sync.RWMutex
	homeworkUpdateCache          = make(map[string]updateCache)
	homeworkUpsertCacheMut       sync.RWMutex
	homeworkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var homeworkBeforeInsertHooks []HomeworkHook
var homeworkBeforeUpdateHooks []HomeworkHook
var homeworkBeforeDeleteHooks []HomeworkHook
var homeworkBeforeUpsertHooks []HomeworkHook

var homeworkAfterInsertHooks []HomeworkHook
var homeworkAfterSelectHooks []HomeworkHook
var homeworkAfterUpdateHooks []HomeworkHook
var homeworkAfterDeleteHooks []HomeworkHook
var homeworkAfterUpsertHooks []HomeworkHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Homework) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Homework) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Homework) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Homework) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Homework) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Homework) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Homework) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Homework) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Homework) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeworkAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHomeworkHook registers your hook function for all future operations.
func AddHomeworkHook(hookPoint boil.HookPoint, homeworkHook HomeworkHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		homeworkBeforeInsertHooks = append(homeworkBeforeInsertHooks, homeworkHook)
	case boil.BeforeUpdateHook:
		homeworkBeforeUpdateHooks = append(homeworkBeforeUpdateHooks, homeworkHook)
	case boil.BeforeDeleteHook:
		homeworkBeforeDeleteHooks = append(homeworkBeforeDeleteHooks, homeworkHook)
	case boil.BeforeUpsertHook:
		homeworkBeforeUpsertHooks = append(homeworkBeforeUpsertHooks, homeworkHook)
	case boil.AfterInsertHook:
		homeworkAfterInsertHooks = append(homeworkAfterInsertHooks, homeworkHook)
	case boil.AfterSelectHook:
		homeworkAfterSelectHooks = append(homeworkAfterSelectHooks, homeworkHook)
	case boil.AfterUpdateHook:
		homeworkAfterUpdateHooks = append(homeworkAfterUpdateHooks, homeworkHook)
	case boil.AfterDeleteHook:
		homeworkAfterDeleteHooks = append(homeworkAfterDeleteHooks, homeworkHook)
	case boil.AfterUpsertHook:
		homeworkAfterUpsertHooks = append(homeworkAfterUpsertHooks, homeworkHook)
	}
}

// One returns a single homework record from the query.
func (q homeworkQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Homework, error) {
	o := &Homework{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for homeworks")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Homework records from the query.
func (q homeworkQuery) All(ctx context.Context, exec boil.ContextExecutor) (HomeworkSlice, error) {
	var o []*Homework

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Homework slice")
	}

	if len(homeworkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Homework records in the query.
func (q homeworkQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count homeworks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q homeworkQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if homeworks exists")
	}

	return count > 0, nil
}

// Course pointed to by the foreign key.
func (o *Homework) Course(mods ...qm.QueryMod) courseQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CourseID),
	}

	queryMods = append(queryMods, mods...)

	query := Courses(queryMods...)
	queries.SetFrom(query.Query, "\"courses\"")

	return query
}

// HomeworkSubmissions retrieves all the homework_submission's HomeworkSubmissions with an executor.
func (o *Homework) HomeworkSubmissions(mods ...qm.QueryMod) homeworkSubmissionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"homework_submissions\".\"homework_id\"=?", o.ID),
	)

	query := HomeworkSubmissions(queryMods...)
	queries.SetFrom(query.Query, "\"homework_submissions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"homework_submissions\".*"})
	}

	return query
}

// LoadCourse allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (homeworkL) LoadCourse(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHomework interface{}, mods queries.Applicator) error {
	var slice []*Homework
	var object *Homework

	if singular {
		object = maybeHomework.(*Homework)
	} else {
		slice = *maybeHomework.(*[]*Homework)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &homeworkR{}
		}
		args = append(args, object.CourseID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &homeworkR{}
			}

			for _, a := range args {
				if a == obj.CourseID {
					continue Outer
				}
			}

			args = append(args, obj.CourseID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`courses`),
		qm.WhereIn(`courses.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Course")
	}

	var resultSlice []*Course
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Course")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for courses")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for courses")
	}

	if len(homeworkAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Course = foreign
		if foreign.R == nil {
			foreign.R = &courseR{}
		}
		foreign.R.Homeworks = append(foreign.R.Homeworks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CourseID == foreign.ID {
				local.R.Course = foreign
				if foreign.R == nil {
					foreign.R = &courseR{}
				}
				foreign.R.Homeworks = append(foreign.R.Homeworks, local)
				break
			}
		}
	}

	return nil
}

// LoadHomeworkSubmissions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (homeworkL) LoadHomeworkSubmissions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHomework interface{}, mods queries.Applicator) error {
	var slice []*Homework
	var object *Homework

	if singular {
		object = maybeHomework.(*Homework)
	} else {
		slice = *maybeHomework.(*[]*Homework)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &homeworkR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &homeworkR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`homework_submissions`),
		qm.WhereIn(`homework_submissions.homework_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load homework_submissions")
	}

	var resultSlice []*HomeworkSubmission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice homework_submissions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on homework_submissions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for homework_submissions")
	}

	if len(homeworkSubmissionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.HomeworkSubmissions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &homeworkSubmissionR{}
			}
			foreign.R.Homework = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.HomeworkID {
				local.R.HomeworkSubmissions = append(local.R.HomeworkSubmissions, foreign)
				if foreign.R == nil {
					foreign.R = &homeworkSubmissionR{}
				}
				foreign.R.Homework = local
				break
			}
		}
	}

	return nil
}

// SetCourse of the homework to the related item.
// Sets o.R.Course to related.
// Adds o to related.R.Homeworks.
func (o *Homework) SetCourse(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Course) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"homeworks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"course_id"}),
		strmangle.WhereClause("\"", "\"", 2, homeworkPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CourseID = related.ID
	if o.R == nil {
		o.R = &homeworkR{
			Course: related,
		}
	} else {
		o.R.Course = related
	}

	if related.R == nil {
		related.R = &courseR{
			Homeworks: HomeworkSlice{o},
		}
	} else {
		related.R.Homeworks = append(related.R.Homeworks, o)
	}

	return nil
}

// AddHomeworkSubmissions adds the given related objects to the existing relationships
// of the homework, optionally inserting them as new records.
// Appends related to o.R.HomeworkSubmissions.
// Sets related.R.Homework appropriately.
func (o *Homework) AddHomeworkSubmissions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*HomeworkSubmission) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.HomeworkID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"homework_submissions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"homework_id"}),
				strmangle.WhereClause("\"", "\"", 2, homeworkSubmissionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.HomeworkID = o.ID
		}
	}

	if o.R == nil {
		o.R = &homeworkR{
			HomeworkSubmissions: related,
		}
	} else {
		o.R.HomeworkSubmissions = append(o.R.HomeworkSubmissions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &homeworkSubmissionR{
				Homework: o,
			}
		} else {
			rel.R.Homework = o
		}
	}
	return nil
}

// Homeworks retrieves all the records using an executor.
func Homeworks(mods ...qm.QueryMod) homeworkQuery {
	mods = append(mods, qm.From("\"homeworks\""))
	return homeworkQuery{NewQuery(mods...)}
}

// FindHomework retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHomework(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Homework, error) {
	homeworkObj := &Homework{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"homeworks\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, homeworkObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from homeworks")
	}

	if err = homeworkObj.doAfterSelectHooks(ctx, exec); err != nil {
		return homeworkObj, err
	}

	return homeworkObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Homework) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no homeworks provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homeworkColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	homeworkInsertCacheMut.RLock()
	cache, cached := homeworkInsertCache[key]
	homeworkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			homeworkAllColumns,
			homeworkColumnsWithDefault,
			homeworkColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(homeworkType, homeworkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(homeworkType, homeworkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"homeworks\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"homeworks\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into homeworks")
	}

	if !cached {
		homeworkInsertCacheMut.Lock()
		homeworkInsertCache[key] = cache
		homeworkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Homework.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Homework) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	homeworkUpdateCacheMut.RLock()
	cache, cached := homeworkUpdateCache[key]
	homeworkUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			homeworkAllColumns,
			homeworkPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update homeworks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"homeworks\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, homeworkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(homeworkType, homeworkMapping, append(wl, homeworkPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update homeworks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for homeworks")
	}

	if !cached {
		homeworkUpdateCacheMut.Lock()
		homeworkUpdateCache[key] = cache
		homeworkUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q homeworkQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for homeworks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for homeworks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HomeworkSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeworkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"homeworks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, homeworkPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in homework slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all homework")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Homework) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no homeworks provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homeworkColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	homeworkUpsertCacheMut.RLock()
	cache, cached := homeworkUpsertCache[key]
	homeworkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			homeworkAllColumns,
			homeworkColumnsWithDefault,
			homeworkColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			homeworkAllColumns,
			homeworkPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert homeworks, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(homeworkPrimaryKeyColumns))
			copy(conflict, homeworkPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"homeworks\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(homeworkType, homeworkMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(homeworkType, homeworkMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert homeworks")
	}

	if !cached {
		homeworkUpsertCacheMut.Lock()
		homeworkUpsertCache[key] = cache
		homeworkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Homework record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Homework) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Homework provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), homeworkPrimaryKeyMapping)
	sql := "DELETE FROM \"homeworks\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from homeworks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for homeworks")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q homeworkQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no homeworkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from homeworks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for homeworks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HomeworkSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(homeworkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeworkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"homeworks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, homeworkPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from homework slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for homeworks")
	}

	if len(homeworkAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Homework) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHomework(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HomeworkSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HomeworkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeworkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"homeworks\".* FROM \"homeworks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, homeworkPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HomeworkSlice")
	}

	*o = slice

	return nil
}

// HomeworkExists checks if the Homework row exists.
func HomeworkExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"homeworks\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if homeworks exists")
	}

	return exists, nil
}