// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testMessages(t *testing.T) {
	t.Parallel()

	query := Messages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
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

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Messages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MessageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MessageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Message exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MessageExists to return true, but got false.")
	}
}

func testMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	messageFound, err := FindMessage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if messageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Messages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Messages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	messageOne := &Message{}
	messageTwo := &Message{}
	if err = randomize.Struct(seed, messageOne, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err = randomize.Struct(seed, messageTwo, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = messageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = messageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Messages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	messageOne := &Message{}
	messageTwo := &Message{}
	if err = randomize.Struct(seed, messageOne, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err = randomize.Struct(seed, messageTwo, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = messageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = messageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func messageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func testMessagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Message{}
	o := &Message{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, messageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Message object: %s", err)
	}

	AddMessageHook(boil.BeforeInsertHook, messageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	messageBeforeInsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterInsertHook, messageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	messageAfterInsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterSelectHook, messageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	messageAfterSelectHooks = []MessageHook{}

	AddMessageHook(boil.BeforeUpdateHook, messageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	messageBeforeUpdateHooks = []MessageHook{}

	AddMessageHook(boil.AfterUpdateHook, messageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	messageAfterUpdateHooks = []MessageHook{}

	AddMessageHook(boil.BeforeDeleteHook, messageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	messageBeforeDeleteHooks = []MessageHook{}

	AddMessageHook(boil.AfterDeleteHook, messageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	messageAfterDeleteHooks = []MessageHook{}

	AddMessageHook(boil.BeforeUpsertHook, messageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	messageBeforeUpsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterUpsertHook, messageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	messageAfterUpsertHooks = []MessageHook{}
}

func testMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(messageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMessageToManyChatroomEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c ChatroomEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
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

	queries.Assign(&b.MessageID, a.ID)
	queries.Assign(&c.MessageID, a.ID)
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
		if queries.Equal(v.MessageID, b.MessageID) {
			bFound = true
		}
		if queries.Equal(v.MessageID, c.MessageID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MessageSlice{&a}
	if err = a.L.LoadChatroomEvents(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
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

func testMessageToManyAddOpChatroomEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c, d, e ChatroomEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
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

		if !queries.Equal(a.ID, first.MessageID) {
			t.Error("foreign key was wrong value", a.ID, first.MessageID)
		}
		if !queries.Equal(a.ID, second.MessageID) {
			t.Error("foreign key was wrong value", a.ID, second.MessageID)
		}

		if first.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Message != &a {
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

func testMessageToManySetOpChatroomEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c, d, e ChatroomEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ChatroomEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, chatroomEventDBTypes, false, strmangle.SetComplement(chatroomEventPrimaryKeyColumns, chatroomEventColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetChatroomEvents(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ChatroomEvents().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetChatroomEvents(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ChatroomEvents().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MessageID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MessageID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.MessageID) {
		t.Error("foreign key was wrong value", a.ID, d.MessageID)
	}
	if !queries.Equal(a.ID, e.MessageID) {
		t.Error("foreign key was wrong value", a.ID, e.MessageID)
	}

	if b.R.Message != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Message != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Message != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Message != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ChatroomEvents[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ChatroomEvents[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMessageToManyRemoveOpChatroomEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c, d, e ChatroomEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
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

	err = a.AddChatroomEvents(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ChatroomEvents().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveChatroomEvents(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ChatroomEvents().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MessageID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MessageID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Message != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Message != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Message != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Message != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ChatroomEvents) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ChatroomEvents[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ChatroomEvents[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMessageToOneChatroomUsingChatroom(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Message
	var foreign Chatroom

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, chatroomDBTypes, false, chatroomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chatroom struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChatroomID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Chatroom().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MessageSlice{&local}
	if err = local.L.LoadChatroom(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Chatroom == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Chatroom = nil
	if err = local.L.LoadChatroom(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Chatroom == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMessageToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Message
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MessageSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMessageToOneSetOpChatroomUsingChatroom(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c Chatroom

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, chatroomDBTypes, false, strmangle.SetComplement(chatroomPrimaryKeyColumns, chatroomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chatroomDBTypes, false, strmangle.SetComplement(chatroomPrimaryKeyColumns, chatroomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Chatroom{&b, &c} {
		err = a.SetChatroom(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Chatroom != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Messages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChatroomID != x.ID {
			t.Error("foreign key was wrong value", a.ChatroomID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChatroomID))
		reflect.Indirect(reflect.ValueOf(&a.ChatroomID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChatroomID != x.ID {
			t.Error("foreign key was wrong value", a.ChatroomID, x.ID)
		}
	}
}
func testMessageToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Messages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
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

func testMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MessageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Messages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	messageDBTypes = map[string]string{`ID`: `integer`, `MessageText`: `text`, `WrittenAt`: `timestamp without time zone`, `TransmitedAt`: `timestamp without time zone`, `ServerReceivedAt`: `timestamp without time zone`, `UserID`: `integer`, `ChatroomID`: `integer`, `Platform`: `character varying`, `CreatedAt`: `timestamp without time zone`}
	_              = bytes.MinRead
)

func testMessagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(messageAllColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, messageDBTypes, true, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(messageAllColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, messageDBTypes, true, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(messageAllColumns, messagePrimaryKeyColumns) {
		fields = messageAllColumns
	} else {
		fields = strmangle.SetComplement(
			messageAllColumns,
			messagePrimaryKeyColumns,
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

	slice := MessageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMessagesUpsert(t *testing.T) {
	t.Parallel()

	if len(messageAllColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Message{}
	if err = randomize.Struct(seed, &o, messageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Message: %s", err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, messageDBTypes, false, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Message: %s", err)
	}

	count, err = Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}