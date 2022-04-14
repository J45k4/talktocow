// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testHomeworks(t *testing.T) {
	t.Parallel()

	query := Homeworks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHomeworksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHomeworksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Homeworks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHomeworksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HomeworkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHomeworksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HomeworkExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Homework exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HomeworkExists to return true, but got false.")
	}
}

func testHomeworksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	homeworkFound, err := FindHomework(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if homeworkFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHomeworksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Homeworks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHomeworksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Homeworks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHomeworksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	homeworkOne := &Homework{}
	homeworkTwo := &Homework{}
	if err = randomize.Struct(seed, homeworkOne, homeworkDBTypes, false, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}
	if err = randomize.Struct(seed, homeworkTwo, homeworkDBTypes, false, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = homeworkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = homeworkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Homeworks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHomeworksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	homeworkOne := &Homework{}
	homeworkTwo := &Homework{}
	if err = randomize.Struct(seed, homeworkOne, homeworkDBTypes, false, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}
	if err = randomize.Struct(seed, homeworkTwo, homeworkDBTypes, false, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = homeworkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = homeworkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func homeworkBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func homeworkAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Homework) error {
	*o = Homework{}
	return nil
}

func testHomeworksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Homework{}
	o := &Homework{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, homeworkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Homework object: %s", err)
	}

	AddHomeworkHook(boil.BeforeInsertHook, homeworkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	homeworkBeforeInsertHooks = []HomeworkHook{}

	AddHomeworkHook(boil.AfterInsertHook, homeworkAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	homeworkAfterInsertHooks = []HomeworkHook{}

	AddHomeworkHook(boil.AfterSelectHook, homeworkAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	homeworkAfterSelectHooks = []HomeworkHook{}

	AddHomeworkHook(boil.BeforeUpdateHook, homeworkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	homeworkBeforeUpdateHooks = []HomeworkHook{}

	AddHomeworkHook(boil.AfterUpdateHook, homeworkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	homeworkAfterUpdateHooks = []HomeworkHook{}

	AddHomeworkHook(boil.BeforeDeleteHook, homeworkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	homeworkBeforeDeleteHooks = []HomeworkHook{}

	AddHomeworkHook(boil.AfterDeleteHook, homeworkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	homeworkAfterDeleteHooks = []HomeworkHook{}

	AddHomeworkHook(boil.BeforeUpsertHook, homeworkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	homeworkBeforeUpsertHooks = []HomeworkHook{}

	AddHomeworkHook(boil.AfterUpsertHook, homeworkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	homeworkAfterUpsertHooks = []HomeworkHook{}
}

func testHomeworksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHomeworksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(homeworkColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHomeworkToManyHomeworkSubmissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Homework
	var b, c HomeworkSubmission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, homeworkSubmissionDBTypes, false, homeworkSubmissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, homeworkSubmissionDBTypes, false, homeworkSubmissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.HomeworkID = a.ID
	c.HomeworkID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.HomeworkSubmissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.HomeworkID == b.HomeworkID {
			bFound = true
		}
		if v.HomeworkID == c.HomeworkID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HomeworkSlice{&a}
	if err = a.L.LoadHomeworkSubmissions(ctx, tx, false, (*[]*Homework)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.HomeworkSubmissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.HomeworkSubmissions = nil
	if err = a.L.LoadHomeworkSubmissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.HomeworkSubmissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHomeworkToManyAddOpHomeworkSubmissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Homework
	var b, c, d, e HomeworkSubmission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, homeworkDBTypes, false, strmangle.SetComplement(homeworkPrimaryKeyColumns, homeworkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*HomeworkSubmission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, homeworkSubmissionDBTypes, false, strmangle.SetComplement(homeworkSubmissionPrimaryKeyColumns, homeworkSubmissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*HomeworkSubmission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddHomeworkSubmissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.HomeworkID {
			t.Error("foreign key was wrong value", a.ID, first.HomeworkID)
		}
		if a.ID != second.HomeworkID {
			t.Error("foreign key was wrong value", a.ID, second.HomeworkID)
		}

		if first.R.Homework != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Homework != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.HomeworkSubmissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.HomeworkSubmissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.HomeworkSubmissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testHomeworkToOneCourseUsingCourse(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Homework
	var foreign Course

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, homeworkDBTypes, false, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, courseDBTypes, false, courseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Course struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CourseID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Course().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := HomeworkSlice{&local}
	if err = local.L.LoadCourse(ctx, tx, false, (*[]*Homework)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Course == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Course = nil
	if err = local.L.LoadCourse(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Course == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testHomeworkToOneSetOpCourseUsingCourse(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Homework
	var b, c Course

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, homeworkDBTypes, false, strmangle.SetComplement(homeworkPrimaryKeyColumns, homeworkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Course{&b, &c} {
		err = a.SetCourse(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Course != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Homeworks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CourseID != x.ID {
			t.Error("foreign key was wrong value", a.CourseID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CourseID))
		reflect.Indirect(reflect.ValueOf(&a.CourseID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CourseID != x.ID {
			t.Error("foreign key was wrong value", a.CourseID, x.ID)
		}
	}
}

func testHomeworksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHomeworksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HomeworkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHomeworksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Homeworks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	homeworkDBTypes = map[string]string{`ID`: `integer`, `Title`: `character varying`, `Description`: `text`, `DueDate`: `timestamp with time zone`, `CourseID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testHomeworksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(homeworkPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(homeworkAllColumns) == len(homeworkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHomeworksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(homeworkAllColumns) == len(homeworkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Homework{}
	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, homeworkDBTypes, true, homeworkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(homeworkAllColumns, homeworkPrimaryKeyColumns) {
		fields = homeworkAllColumns
	} else {
		fields = strmangle.SetComplement(
			homeworkAllColumns,
			homeworkPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := HomeworkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHomeworksUpsert(t *testing.T) {
	t.Parallel()

	if len(homeworkAllColumns) == len(homeworkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Homework{}
	if err = randomize.Struct(seed, &o, homeworkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Homework: %s", err)
	}

	count, err := Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, homeworkDBTypes, false, homeworkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Homework struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Homework: %s", err)
	}

	count, err = Homeworks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
