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

// Course is an object representing the database table.
type Course struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *courseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L courseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CourseColumns = struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var CourseTableColumns = struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "courses.id",
	Name:        "courses.name",
	Description: "courses.description",
	CreatedAt:   "courses.created_at",
	UpdatedAt:   "courses.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var CourseWhere = struct {
	ID          whereHelperint
	Name        whereHelperstring
	Description whereHelpernull_String
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint{field: "\"courses\".\"id\""},
	Name:        whereHelperstring{field: "\"courses\".\"name\""},
	Description: whereHelpernull_String{field: "\"courses\".\"description\""},
	CreatedAt:   whereHelpertime_Time{field: "\"courses\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"courses\".\"updated_at\""},
}

// CourseRels is where relationship names are stored.
var CourseRels = struct {
	CourseUsers string
	Homeworks   string
}{
	CourseUsers: "CourseUsers",
	Homeworks:   "Homeworks",
}

// courseR is where relationships are stored.
type courseR struct {
	CourseUsers CourseUserSlice `boil:"CourseUsers" json:"CourseUsers" toml:"CourseUsers" yaml:"CourseUsers"`
	Homeworks   HomeworkSlice   `boil:"Homeworks" json:"Homeworks" toml:"Homeworks" yaml:"Homeworks"`
}

// NewStruct creates a new relationship struct
func (*courseR) NewStruct() *courseR {
	return &courseR{}
}

// courseL is where Load methods for each relationship are stored.
type courseL struct{}

var (
	courseAllColumns            = []string{"id", "name", "description", "created_at", "updated_at"}
	courseColumnsWithoutDefault = []string{"name", "description", "created_at", "updated_at"}
	courseColumnsWithDefault    = []string{"id"}
	coursePrimaryKeyColumns     = []string{"id"}
)

type (
	// CourseSlice is an alias for a slice of pointers to Course.
	// This should almost always be used instead of []Course.
	CourseSlice []*Course
	// CourseHook is the signature for custom Course hook methods
	CourseHook func(context.Context, boil.ContextExecutor, *Course) error

	courseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	courseType                 = reflect.TypeOf(&Course{})
	courseMapping              = queries.MakeStructMapping(courseType)
	coursePrimaryKeyMapping, _ = queries.BindMapping(courseType, courseMapping, coursePrimaryKeyColumns)
	courseInsertCacheMut       sync.RWMutex
	courseInsertCache          = make(map[string]insertCache)
	courseUpdateCacheMut       sync.RWMutex
	courseUpdateCache          = make(map[string]updateCache)
	courseUpsertCacheMut       sync.RWMutex
	courseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var courseBeforeInsertHooks []CourseHook
var courseBeforeUpdateHooks []CourseHook
var courseBeforeDeleteHooks []CourseHook
var courseBeforeUpsertHooks []CourseHook

var courseAfterInsertHooks []CourseHook
var courseAfterSelectHooks []CourseHook
var courseAfterUpdateHooks []CourseHook
var courseAfterDeleteHooks []CourseHook
var courseAfterUpsertHooks []CourseHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Course) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Course) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Course) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Course) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Course) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Course) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Course) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Course) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Course) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCourseHook registers your hook function for all future operations.
func AddCourseHook(hookPoint boil.HookPoint, courseHook CourseHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		courseBeforeInsertHooks = append(courseBeforeInsertHooks, courseHook)
	case boil.BeforeUpdateHook:
		courseBeforeUpdateHooks = append(courseBeforeUpdateHooks, courseHook)
	case boil.BeforeDeleteHook:
		courseBeforeDeleteHooks = append(courseBeforeDeleteHooks, courseHook)
	case boil.BeforeUpsertHook:
		courseBeforeUpsertHooks = append(courseBeforeUpsertHooks, courseHook)
	case boil.AfterInsertHook:
		courseAfterInsertHooks = append(courseAfterInsertHooks, courseHook)
	case boil.AfterSelectHook:
		courseAfterSelectHooks = append(courseAfterSelectHooks, courseHook)
	case boil.AfterUpdateHook:
		courseAfterUpdateHooks = append(courseAfterUpdateHooks, courseHook)
	case boil.AfterDeleteHook:
		courseAfterDeleteHooks = append(courseAfterDeleteHooks, courseHook)
	case boil.AfterUpsertHook:
		courseAfterUpsertHooks = append(courseAfterUpsertHooks, courseHook)
	}
}

// One returns a single course record from the query.
func (q courseQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Course, error) {
	o := &Course{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for courses")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Course records from the query.
func (q courseQuery) All(ctx context.Context, exec boil.ContextExecutor) (CourseSlice, error) {
	var o []*Course

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Course slice")
	}

	if len(courseAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Course records in the query.
func (q courseQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count courses rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q courseQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if courses exists")
	}

	return count > 0, nil
}

// CourseUsers retrieves all the course_user's CourseUsers with an executor.
func (o *Course) CourseUsers(mods ...qm.QueryMod) courseUserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"course_users\".\"course_id\"=?", o.ID),
	)

	query := CourseUsers(queryMods...)
	queries.SetFrom(query.Query, "\"course_users\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"course_users\".*"})
	}

	return query
}

// Homeworks retrieves all the homework's Homeworks with an executor.
func (o *Course) Homeworks(mods ...qm.QueryMod) homeworkQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"homeworks\".\"course_id\"=?", o.ID),
	)

	query := Homeworks(queryMods...)
	queries.SetFrom(query.Query, "\"homeworks\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"homeworks\".*"})
	}

	return query
}

// LoadCourseUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (courseL) LoadCourseUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCourse interface{}, mods queries.Applicator) error {
	var slice []*Course
	var object *Course

	if singular {
		object = maybeCourse.(*Course)
	} else {
		slice = *maybeCourse.(*[]*Course)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &courseR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &courseR{}
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
		qm.From(`course_users`),
		qm.WhereIn(`course_users.course_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load course_users")
	}

	var resultSlice []*CourseUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice course_users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on course_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for course_users")
	}

	if len(courseUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CourseUsers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &courseUserR{}
			}
			foreign.R.Course = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CourseID {
				local.R.CourseUsers = append(local.R.CourseUsers, foreign)
				if foreign.R == nil {
					foreign.R = &courseUserR{}
				}
				foreign.R.Course = local
				break
			}
		}
	}

	return nil
}

// LoadHomeworks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (courseL) LoadHomeworks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCourse interface{}, mods queries.Applicator) error {
	var slice []*Course
	var object *Course

	if singular {
		object = maybeCourse.(*Course)
	} else {
		slice = *maybeCourse.(*[]*Course)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &courseR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &courseR{}
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
		qm.From(`homeworks`),
		qm.WhereIn(`homeworks.course_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load homeworks")
	}

	var resultSlice []*Homework
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice homeworks")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on homeworks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for homeworks")
	}

	if len(homeworkAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Homeworks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &homeworkR{}
			}
			foreign.R.Course = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CourseID {
				local.R.Homeworks = append(local.R.Homeworks, foreign)
				if foreign.R == nil {
					foreign.R = &homeworkR{}
				}
				foreign.R.Course = local
				break
			}
		}
	}

	return nil
}

// AddCourseUsers adds the given related objects to the existing relationships
// of the course, optionally inserting them as new records.
// Appends related to o.R.CourseUsers.
// Sets related.R.Course appropriately.
func (o *Course) AddCourseUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CourseUser) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CourseID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"course_users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"course_id"}),
				strmangle.WhereClause("\"", "\"", 2, courseUserPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.CourseID, rel.UserID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CourseID = o.ID
		}
	}

	if o.R == nil {
		o.R = &courseR{
			CourseUsers: related,
		}
	} else {
		o.R.CourseUsers = append(o.R.CourseUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &courseUserR{
				Course: o,
			}
		} else {
			rel.R.Course = o
		}
	}
	return nil
}

// AddHomeworks adds the given related objects to the existing relationships
// of the course, optionally inserting them as new records.
// Appends related to o.R.Homeworks.
// Sets related.R.Course appropriately.
func (o *Course) AddHomeworks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Homework) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CourseID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"homeworks\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"course_id"}),
				strmangle.WhereClause("\"", "\"", 2, homeworkPrimaryKeyColumns),
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

			rel.CourseID = o.ID
		}
	}

	if o.R == nil {
		o.R = &courseR{
			Homeworks: related,
		}
	} else {
		o.R.Homeworks = append(o.R.Homeworks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &homeworkR{
				Course: o,
			}
		} else {
			rel.R.Course = o
		}
	}
	return nil
}

// Courses retrieves all the records using an executor.
func Courses(mods ...qm.QueryMod) courseQuery {
	mods = append(mods, qm.From("\"courses\""))
	return courseQuery{NewQuery(mods...)}
}

// FindCourse retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCourse(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Course, error) {
	courseObj := &Course{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"courses\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, courseObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from courses")
	}

	if err = courseObj.doAfterSelectHooks(ctx, exec); err != nil {
		return courseObj, err
	}

	return courseObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Course) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no courses provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(courseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	courseInsertCacheMut.RLock()
	cache, cached := courseInsertCache[key]
	courseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			courseAllColumns,
			courseColumnsWithDefault,
			courseColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(courseType, courseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(courseType, courseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"courses\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"courses\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into courses")
	}

	if !cached {
		courseInsertCacheMut.Lock()
		courseInsertCache[key] = cache
		courseInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Course.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Course) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	courseUpdateCacheMut.RLock()
	cache, cached := courseUpdateCache[key]
	courseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			courseAllColumns,
			coursePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update courses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"courses\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, coursePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(courseType, courseMapping, append(wl, coursePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update courses row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for courses")
	}

	if !cached {
		courseUpdateCacheMut.Lock()
		courseUpdateCache[key] = cache
		courseUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q courseQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for courses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for courses")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CourseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), coursePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"courses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, coursePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in course slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all course")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Course) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no courses provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(courseColumnsWithDefault, o)

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

	courseUpsertCacheMut.RLock()
	cache, cached := courseUpsertCache[key]
	courseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			courseAllColumns,
			courseColumnsWithDefault,
			courseColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			courseAllColumns,
			coursePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert courses, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(coursePrimaryKeyColumns))
			copy(conflict, coursePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"courses\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(courseType, courseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(courseType, courseMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert courses")
	}

	if !cached {
		courseUpsertCacheMut.Lock()
		courseUpsertCache[key] = cache
		courseUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Course record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Course) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Course provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), coursePrimaryKeyMapping)
	sql := "DELETE FROM \"courses\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from courses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for courses")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q courseQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no courseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from courses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for courses")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CourseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(courseBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), coursePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"courses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, coursePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from course slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for courses")
	}

	if len(courseAfterDeleteHooks) != 0 {
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
func (o *Course) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCourse(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CourseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CourseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), coursePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"courses\".* FROM \"courses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, coursePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CourseSlice")
	}

	*o = slice

	return nil
}

// CourseExists checks if the Course row exists.
func CourseExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"courses\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if courses exists")
	}

	return exists, nil
}
