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

func testChatrooms(t *testing.T) {
	t.Parallel()

	query := Chatrooms()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testChatroomsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
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

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChatroomsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Chatrooms().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChatroomsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChatroomSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChatroomsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ChatroomExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Chatroom exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChatroomExists to return true, but got false.")
	}
}

func testChatroomsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	chatroomFound, err := FindChatroom(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if chatroomFound == nil {
		t.Error("want a record, got nil")
	}
}

func testChatroomsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Chatrooms().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testChatroomsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Chatrooms().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChatroomsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chatroomOne := &Chatroom{}
	chatroomTwo := &Chatroom{}
	if err = randomize.Struct(seed, chatroomOne, chatroomDBTypes, false, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}
	if err = randomize.Struct(seed, chatroomTwo, chatroomDBTypes, false, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = chatroomOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = chatroomTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Chatrooms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChatroomsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	chatroomOne := &Chatroom{}
	chatroomTwo := &Chatroom{}
	if err = randomize.Struct(seed, chatroomOne, chatroomDBTypes, false, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}
	if err = randomize.Struct(seed, chatroomTwo, chatroomDBTypes, false, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = chatroomOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = chatroomTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func chatroomBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func chatroomAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Chatroom) error {
	*o = Chatroom{}
	return nil
}

func testChatroomsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Chatroom{}
	o := &Chatroom{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, chatroomDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Chatroom object: %s", err)
	}

	AddChatroomHook(boil.BeforeInsertHook, chatroomBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	chatroomBeforeInsertHooks = []ChatroomHook{}

	AddChatroomHook(boil.AfterInsertHook, chatroomAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	chatroomAfterInsertHooks = []ChatroomHook{}

	AddChatroomHook(boil.AfterSelectHook, chatroomAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	chatroomAfterSelectHooks = []ChatroomHook{}

	AddChatroomHook(boil.BeforeUpdateHook, chatroomBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	chatroomBeforeUpdateHooks = []ChatroomHook{}

	AddChatroomHook(boil.AfterUpdateHook, chatroomAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	chatroomAfterUpdateHooks = []ChatroomHook{}

	AddChatroomHook(boil.BeforeDeleteHook, chatroomBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	chatroomBeforeDeleteHooks = []ChatroomHook{}

	AddChatroomHook(boil.AfterDeleteHook, chatroomAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	chatroomAfterDeleteHooks = []ChatroomHook{}

	AddChatroomHook(boil.BeforeUpsertHook, chatroomBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	chatroomBeforeUpsertHooks = []ChatroomHook{}

	AddChatroomHook(boil.AfterUpsertHook, chatroomAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	chatroomAfterUpsertHooks = []ChatroomHook{}
}

func testChatroomsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChatroomsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(chatroomColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChatroomToManyChatroomEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chatroom
	var b, c ChatroomEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, chatroomEventDBTypes, false, chatroomEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chatroomEventDBTypes, false, chatroomEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChatroomID = a.ID
	c.ChatroomID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ChatroomEvents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChatroomID == b.ChatroomID {
			bFound = true
		}
		if v.ChatroomID == c.ChatroomID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChatroomSlice{&a}
	if err = a.L.LoadChatroomEvents(ctx, tx, false, (*[]*Chatroom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChatroomEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChatroomEvents = nil
	if err = a.L.LoadChatroomEvents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChatroomEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChatroomToManyChatroomUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chatroom
	var b, c ChatroomUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, chatroomUserDBTypes, false, chatroomUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chatroomUserDBTypes, false, chatroomUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChatroomID = a.ID
	c.ChatroomID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ChatroomUsers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChatroomID == b.ChatroomID {
			bFound = true
		}
		if v.ChatroomID == c.ChatroomID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChatroomSlice{&a}
	if err = a.L.LoadChatroomUsers(ctx, tx, false, (*[]*Chatroom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChatroomUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChatroomUsers = nil
	if err = a.L.LoadChatroomUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChatroomUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChatroomToManyMessages(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chatroom
	var b, c Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChatroomID = a.ID
	c.ChatroomID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Messages().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChatroomID == b.ChatroomID {
			bFound = true
		}
		if v.ChatroomID == c.ChatroomID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChatroomSlice{&a}
	if err = a.L.LoadMessages(ctx, tx, false, (*[]*Chatroom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Messages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Messages = nil
	if err = a.L.LoadMessages(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Messages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChatroomToManyAddOpChatroomEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chatroom
	var b, c, d, e ChatroomEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chatroomDBTypes, false, strmangle.SetComplement(chatroomPrimaryKeyColumns, chatroomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ChatroomEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, chatroomEventDBTypes, false, strmangle.SetComplement(chatroomEventPrimaryKeyColumns, chatroomEventColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ChatroomEvent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChatroomEvents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChatroomID {
			t.Error("foreign key was wrong value", a.ID, first.ChatroomID)
		}
		if a.ID != second.ChatroomID {
			t.Error("foreign key was wrong value", a.ID, second.ChatroomID)
		}

		if first.R.Chatroom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Chatroom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChatroomEvents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChatroomEvents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChatroomEvents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testChatroomToManyAddOpChatroomUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chatroom
	var b, c, d, e ChatroomUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chatroomDBTypes, false, strmangle.SetComplement(chatroomPrimaryKeyColumns, chatroomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ChatroomUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, chatroomUserDBTypes, false, strmangle.SetComplement(chatroomUserPrimaryKeyColumns, chatroomUserColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ChatroomUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChatroomUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChatroomID {
			t.Error("foreign key was wrong value", a.ID, first.ChatroomID)
		}
		if a.ID != second.ChatroomID {
			t.Error("foreign key was wrong value", a.ID, second.ChatroomID)
		}

		if first.R.Chatroom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Chatroom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChatroomUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChatroomUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChatroomUsers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testChatroomToManyAddOpMessages(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chatroom
	var b, c, d, e Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chatroomDBTypes, false, strmangle.SetComplement(chatroomPrimaryKeyColumns, chatroomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Message{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Message{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMessages(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChatroomID {
			t.Error("foreign key was wrong value", a.ID, first.ChatroomID)
		}
		if a.ID != second.ChatroomID {
			t.Error("foreign key was wrong value", a.ID, second.ChatroomID)
		}

		if first.R.Chatroom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Chatroom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Messages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Messages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Messages().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testChatroomsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
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

func testChatroomsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChatroomSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChatroomsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Chatrooms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	chatroomDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `CreatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testChatroomsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(chatroomPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(chatroomAllColumns) == len(chatroomPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testChatroomsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(chatroomAllColumns) == len(chatroomPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Chatroom{}
	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, chatroomDBTypes, true, chatroomPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(chatroomAllColumns, chatroomPrimaryKeyColumns) {
		fields = chatroomAllColumns
	} else {
		fields = strmangle.SetComplement(
			chatroomAllColumns,
			chatroomPrimaryKeyColumns,
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

	slice := ChatroomSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testChatroomsUpsert(t *testing.T) {
	t.Parallel()

	if len(chatroomAllColumns) == len(chatroomPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Chatroom{}
	if err = randomize.Struct(seed, &o, chatroomDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Chatroom: %s", err)
	}

	count, err := Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, chatroomDBTypes, false, chatroomPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Chatroom: %s", err)
	}

	count, err = Chatrooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
