// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testUserReceivedMessages(t *testing.T) {
	t.Parallel()

	query := UserReceivedMessages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserReceivedMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
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

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserReceivedMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserReceivedMessages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserReceivedMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserReceivedMessageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserReceivedMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserReceivedMessageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserReceivedMessage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserReceivedMessageExists to return true, but got false.")
	}
}

func testUserReceivedMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userReceivedMessageFound, err := FindUserReceivedMessage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userReceivedMessageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserReceivedMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserReceivedMessages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserReceivedMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserReceivedMessages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserReceivedMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userReceivedMessageOne := &UserReceivedMessage{}
	userReceivedMessageTwo := &UserReceivedMessage{}
	if err = randomize.Struct(seed, userReceivedMessageOne, userReceivedMessageDBTypes, false, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, userReceivedMessageTwo, userReceivedMessageDBTypes, false, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userReceivedMessageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userReceivedMessageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserReceivedMessages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserReceivedMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userReceivedMessageOne := &UserReceivedMessage{}
	userReceivedMessageTwo := &UserReceivedMessage{}
	if err = randomize.Struct(seed, userReceivedMessageOne, userReceivedMessageDBTypes, false, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, userReceivedMessageTwo, userReceivedMessageDBTypes, false, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userReceivedMessageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userReceivedMessageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userReceivedMessageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func userReceivedMessageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserReceivedMessage) error {
	*o = UserReceivedMessage{}
	return nil
}

func testUserReceivedMessagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserReceivedMessage{}
	o := &UserReceivedMessage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage object: %s", err)
	}

	AddUserReceivedMessageHook(boil.BeforeInsertHook, userReceivedMessageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageBeforeInsertHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.AfterInsertHook, userReceivedMessageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageAfterInsertHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.AfterSelectHook, userReceivedMessageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageAfterSelectHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.BeforeUpdateHook, userReceivedMessageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageBeforeUpdateHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.AfterUpdateHook, userReceivedMessageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageAfterUpdateHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.BeforeDeleteHook, userReceivedMessageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageBeforeDeleteHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.AfterDeleteHook, userReceivedMessageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageAfterDeleteHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.BeforeUpsertHook, userReceivedMessageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageBeforeUpsertHooks = []UserReceivedMessageHook{}

	AddUserReceivedMessageHook(boil.AfterUpsertHook, userReceivedMessageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userReceivedMessageAfterUpsertHooks = []UserReceivedMessageHook{}
}

func testUserReceivedMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserReceivedMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userReceivedMessageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserReceivedMessageToOneMessageUsingMessage(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserReceivedMessage
	var foreign Message

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userReceivedMessageDBTypes, false, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MessageID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Message().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserReceivedMessageSlice{&local}
	if err = local.L.LoadMessage(ctx, tx, false, (*[]*UserReceivedMessage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Message == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Message = nil
	if err = local.L.LoadMessage(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Message == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserReceivedMessageToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserReceivedMessage
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userReceivedMessageDBTypes, false, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
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

	slice := UserReceivedMessageSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserReceivedMessage)(&slice), nil); err != nil {
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

func testUserReceivedMessageToOneSetOpMessageUsingMessage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserReceivedMessage
	var b, c Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userReceivedMessageDBTypes, false, strmangle.SetComplement(userReceivedMessagePrimaryKeyColumns, userReceivedMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Message{&b, &c} {
		err = a.SetMessage(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Message != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserReceivedMessages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MessageID))
		reflect.Indirect(reflect.ValueOf(&a.MessageID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID, x.ID)
		}
	}
}
func testUserReceivedMessageToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserReceivedMessage
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userReceivedMessageDBTypes, false, strmangle.SetComplement(userReceivedMessagePrimaryKeyColumns, userReceivedMessageColumnsWithoutDefault)...); err != nil {
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

		if x.R.UserReceivedMessages[0] != &a {
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

func testUserReceivedMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
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

func testUserReceivedMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserReceivedMessageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserReceivedMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserReceivedMessages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userReceivedMessageDBTypes = map[string]string{`ID`: `integer`, `UserID`: `integer`, `MessageID`: `integer`, `ReceivedAt`: `timestamp without time zone`, `ReadAt`: `timestamp without time zone`}
	_                          = bytes.MinRead
)

func testUserReceivedMessagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userReceivedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userReceivedMessageAllColumns) == len(userReceivedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserReceivedMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userReceivedMessageAllColumns) == len(userReceivedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserReceivedMessage{}
	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userReceivedMessageDBTypes, true, userReceivedMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userReceivedMessageAllColumns, userReceivedMessagePrimaryKeyColumns) {
		fields = userReceivedMessageAllColumns
	} else {
		fields = strmangle.SetComplement(
			userReceivedMessageAllColumns,
			userReceivedMessagePrimaryKeyColumns,
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

	slice := UserReceivedMessageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserReceivedMessagesUpsert(t *testing.T) {
	t.Parallel()

	if len(userReceivedMessageAllColumns) == len(userReceivedMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserReceivedMessage{}
	if err = randomize.Struct(seed, &o, userReceivedMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserReceivedMessage: %s", err)
	}

	count, err := UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userReceivedMessageDBTypes, false, userReceivedMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserReceivedMessage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserReceivedMessage: %s", err)
	}

	count, err = UserReceivedMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}