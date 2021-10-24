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

func testEvents(t *testing.T) {
	t.Parallel()

	query := Events()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
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

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Events().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Event exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EventExists to return true, but got false.")
	}
}

func testEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	eventFound, err := FindEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if eventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Events().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Events().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	eventOne := &Event{}
	eventTwo := &Event{}
	if err = randomize.Struct(seed, eventOne, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}
	if err = randomize.Struct(seed, eventTwo, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Events().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	eventOne := &Event{}
	eventTwo := &Event{}
	if err = randomize.Struct(seed, eventOne, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}
	if err = randomize.Struct(seed, eventTwo, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func eventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Event) error {
	*o = Event{}
	return nil
}

func testEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Event{}
	o := &Event{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, eventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Event object: %s", err)
	}

	AddEventHook(boil.BeforeInsertHook, eventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	eventBeforeInsertHooks = []EventHook{}

	AddEventHook(boil.AfterInsertHook, eventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	eventAfterInsertHooks = []EventHook{}

	AddEventHook(boil.AfterSelectHook, eventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	eventAfterSelectHooks = []EventHook{}

	AddEventHook(boil.BeforeUpdateHook, eventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	eventBeforeUpdateHooks = []EventHook{}

	AddEventHook(boil.AfterUpdateHook, eventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	eventAfterUpdateHooks = []EventHook{}

	AddEventHook(boil.BeforeDeleteHook, eventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	eventBeforeDeleteHooks = []EventHook{}

	AddEventHook(boil.AfterDeleteHook, eventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	eventAfterDeleteHooks = []EventHook{}

	AddEventHook(boil.BeforeUpsertHook, eventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	eventBeforeUpsertHooks = []EventHook{}

	AddEventHook(boil.AfterUpsertHook, eventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	eventAfterUpsertHooks = []EventHook{}
}

func testEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(eventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
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

func testEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Events().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	eventDBTypes = map[string]string{`ID`: `integer`, `EventText`: `text`, `CreatedAt`: `timestamp without time zone`}
	_            = bytes.MinRead
)

func testEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(eventAllColumns) == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventDBTypes, true, eventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(eventAllColumns) == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Event{}
	if err = randomize.Struct(seed, o, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventDBTypes, true, eventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(eventAllColumns, eventPrimaryKeyColumns) {
		fields = eventAllColumns
	} else {
		fields = strmangle.SetComplement(
			eventAllColumns,
			eventPrimaryKeyColumns,
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

	slice := EventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(eventAllColumns) == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Event{}
	if err = randomize.Struct(seed, &o, eventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Event: %s", err)
	}

	count, err := Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, eventDBTypes, false, eventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Event: %s", err)
	}

	count, err = Events().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
