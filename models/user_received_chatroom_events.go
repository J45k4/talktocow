// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// UserReceivedChatroomEvent is an object representing the database table.
type UserReceivedChatroomEvent struct {
	ID              int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID          int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ChatroomEventID int       `boil:"chatroom_event_id" json:"chatroom_event_id" toml:"chatroom_event_id" yaml:"chatroom_event_id"`
	ReceivedAt      time.Time `boil:"received_at" json:"received_at" toml:"received_at" yaml:"received_at"`
	ServerSendAt    null.Time `boil:"server_send_at" json:"server_send_at,omitempty" toml:"server_send_at" yaml:"server_send_at,omitempty"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userReceivedChatroomEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userReceivedChatroomEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserReceivedChatroomEventColumns = struct {
	ID              string
	UserID          string
	ChatroomEventID string
	ReceivedAt      string
	ServerSendAt    string
	CreatedAt       string
}{
	ID:              "id",
	UserID:          "user_id",
	ChatroomEventID: "chatroom_event_id",
	ReceivedAt:      "received_at",
	ServerSendAt:    "server_send_at",
	CreatedAt:       "created_at",
}

// Generated where

var UserReceivedChatroomEventWhere = struct {
	ID              whereHelperint
	UserID          whereHelperint
	ChatroomEventID whereHelperint
	ReceivedAt      whereHelpertime_Time
	ServerSendAt    whereHelpernull_Time
	CreatedAt       whereHelpertime_Time
}{
	ID:              whereHelperint{field: "\"user_received_chatroom_events\".\"id\""},
	UserID:          whereHelperint{field: "\"user_received_chatroom_events\".\"user_id\""},
	ChatroomEventID: whereHelperint{field: "\"user_received_chatroom_events\".\"chatroom_event_id\""},
	ReceivedAt:      whereHelpertime_Time{field: "\"user_received_chatroom_events\".\"received_at\""},
	ServerSendAt:    whereHelpernull_Time{field: "\"user_received_chatroom_events\".\"server_send_at\""},
	CreatedAt:       whereHelpertime_Time{field: "\"user_received_chatroom_events\".\"created_at\""},
}

// UserReceivedChatroomEventRels is where relationship names are stored.
var UserReceivedChatroomEventRels = struct {
	ChatroomEvent string
	User          string
}{
	ChatroomEvent: "ChatroomEvent",
	User:          "User",
}

// userReceivedChatroomEventR is where relationships are stored.
type userReceivedChatroomEventR struct {
	ChatroomEvent *ChatroomEvent `boil:"ChatroomEvent" json:"ChatroomEvent" toml:"ChatroomEvent" yaml:"ChatroomEvent"`
	User          *User          `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userReceivedChatroomEventR) NewStruct() *userReceivedChatroomEventR {
	return &userReceivedChatroomEventR{}
}

// userReceivedChatroomEventL is where Load methods for each relationship are stored.
type userReceivedChatroomEventL struct{}

var (
	userReceivedChatroomEventAllColumns            = []string{"id", "user_id", "chatroom_event_id", "received_at", "server_send_at", "created_at"}
	userReceivedChatroomEventColumnsWithoutDefault = []string{"user_id", "chatroom_event_id", "received_at", "server_send_at", "created_at"}
	userReceivedChatroomEventColumnsWithDefault    = []string{"id"}
	userReceivedChatroomEventPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserReceivedChatroomEventSlice is an alias for a slice of pointers to UserReceivedChatroomEvent.
	// This should generally be used opposed to []UserReceivedChatroomEvent.
	UserReceivedChatroomEventSlice []*UserReceivedChatroomEvent
	// UserReceivedChatroomEventHook is the signature for custom UserReceivedChatroomEvent hook methods
	UserReceivedChatroomEventHook func(context.Context, boil.ContextExecutor, *UserReceivedChatroomEvent) error

	userReceivedChatroomEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userReceivedChatroomEventType                 = reflect.TypeOf(&UserReceivedChatroomEvent{})
	userReceivedChatroomEventMapping              = queries.MakeStructMapping(userReceivedChatroomEventType)
	userReceivedChatroomEventPrimaryKeyMapping, _ = queries.BindMapping(userReceivedChatroomEventType, userReceivedChatroomEventMapping, userReceivedChatroomEventPrimaryKeyColumns)
	userReceivedChatroomEventInsertCacheMut       sync.RWMutex
	userReceivedChatroomEventInsertCache          = make(map[string]insertCache)
	userReceivedChatroomEventUpdateCacheMut       sync.RWMutex
	userReceivedChatroomEventUpdateCache          = make(map[string]updateCache)
	userReceivedChatroomEventUpsertCacheMut       sync.RWMutex
	userReceivedChatroomEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userReceivedChatroomEventBeforeInsertHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventBeforeUpdateHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventBeforeDeleteHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventBeforeUpsertHooks []UserReceivedChatroomEventHook

var userReceivedChatroomEventAfterInsertHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventAfterSelectHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventAfterUpdateHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventAfterDeleteHooks []UserReceivedChatroomEventHook
var userReceivedChatroomEventAfterUpsertHooks []UserReceivedChatroomEventHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserReceivedChatroomEvent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserReceivedChatroomEvent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserReceivedChatroomEvent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserReceivedChatroomEvent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserReceivedChatroomEvent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserReceivedChatroomEvent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserReceivedChatroomEvent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserReceivedChatroomEvent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserReceivedChatroomEvent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userReceivedChatroomEventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserReceivedChatroomEventHook registers your hook function for all future operations.
func AddUserReceivedChatroomEventHook(hookPoint boil.HookPoint, userReceivedChatroomEventHook UserReceivedChatroomEventHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userReceivedChatroomEventBeforeInsertHooks = append(userReceivedChatroomEventBeforeInsertHooks, userReceivedChatroomEventHook)
	case boil.BeforeUpdateHook:
		userReceivedChatroomEventBeforeUpdateHooks = append(userReceivedChatroomEventBeforeUpdateHooks, userReceivedChatroomEventHook)
	case boil.BeforeDeleteHook:
		userReceivedChatroomEventBeforeDeleteHooks = append(userReceivedChatroomEventBeforeDeleteHooks, userReceivedChatroomEventHook)
	case boil.BeforeUpsertHook:
		userReceivedChatroomEventBeforeUpsertHooks = append(userReceivedChatroomEventBeforeUpsertHooks, userReceivedChatroomEventHook)
	case boil.AfterInsertHook:
		userReceivedChatroomEventAfterInsertHooks = append(userReceivedChatroomEventAfterInsertHooks, userReceivedChatroomEventHook)
	case boil.AfterSelectHook:
		userReceivedChatroomEventAfterSelectHooks = append(userReceivedChatroomEventAfterSelectHooks, userReceivedChatroomEventHook)
	case boil.AfterUpdateHook:
		userReceivedChatroomEventAfterUpdateHooks = append(userReceivedChatroomEventAfterUpdateHooks, userReceivedChatroomEventHook)
	case boil.AfterDeleteHook:
		userReceivedChatroomEventAfterDeleteHooks = append(userReceivedChatroomEventAfterDeleteHooks, userReceivedChatroomEventHook)
	case boil.AfterUpsertHook:
		userReceivedChatroomEventAfterUpsertHooks = append(userReceivedChatroomEventAfterUpsertHooks, userReceivedChatroomEventHook)
	}
}

// One returns a single userReceivedChatroomEvent record from the query.
func (q userReceivedChatroomEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserReceivedChatroomEvent, error) {
	o := &UserReceivedChatroomEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_received_chatroom_events")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserReceivedChatroomEvent records from the query.
func (q userReceivedChatroomEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserReceivedChatroomEventSlice, error) {
	var o []*UserReceivedChatroomEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserReceivedChatroomEvent slice")
	}

	if len(userReceivedChatroomEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserReceivedChatroomEvent records in the query.
func (q userReceivedChatroomEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_received_chatroom_events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userReceivedChatroomEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_received_chatroom_events exists")
	}

	return count > 0, nil
}

// ChatroomEvent pointed to by the foreign key.
func (o *UserReceivedChatroomEvent) ChatroomEvent(mods ...qm.QueryMod) chatroomEventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChatroomEventID),
	}

	queryMods = append(queryMods, mods...)

	query := ChatroomEvents(queryMods...)
	queries.SetFrom(query.Query, "\"chatroom_events\"")

	return query
}

// User pointed to by the foreign key.
func (o *UserReceivedChatroomEvent) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadChatroomEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userReceivedChatroomEventL) LoadChatroomEvent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserReceivedChatroomEvent interface{}, mods queries.Applicator) error {
	var slice []*UserReceivedChatroomEvent
	var object *UserReceivedChatroomEvent

	if singular {
		object = maybeUserReceivedChatroomEvent.(*UserReceivedChatroomEvent)
	} else {
		slice = *maybeUserReceivedChatroomEvent.(*[]*UserReceivedChatroomEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userReceivedChatroomEventR{}
		}
		args = append(args, object.ChatroomEventID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userReceivedChatroomEventR{}
			}

			for _, a := range args {
				if a == obj.ChatroomEventID {
					continue Outer
				}
			}

			args = append(args, obj.ChatroomEventID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`chatroom_events`),
		qm.WhereIn(`chatroom_events.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ChatroomEvent")
	}

	var resultSlice []*ChatroomEvent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ChatroomEvent")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for chatroom_events")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chatroom_events")
	}

	if len(userReceivedChatroomEventAfterSelectHooks) != 0 {
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
		object.R.ChatroomEvent = foreign
		if foreign.R == nil {
			foreign.R = &chatroomEventR{}
		}
		foreign.R.UserReceivedChatroomEvents = append(foreign.R.UserReceivedChatroomEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChatroomEventID == foreign.ID {
				local.R.ChatroomEvent = foreign
				if foreign.R == nil {
					foreign.R = &chatroomEventR{}
				}
				foreign.R.UserReceivedChatroomEvents = append(foreign.R.UserReceivedChatroomEvents, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userReceivedChatroomEventL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserReceivedChatroomEvent interface{}, mods queries.Applicator) error {
	var slice []*UserReceivedChatroomEvent
	var object *UserReceivedChatroomEvent

	if singular {
		object = maybeUserReceivedChatroomEvent.(*UserReceivedChatroomEvent)
	} else {
		slice = *maybeUserReceivedChatroomEvent.(*[]*UserReceivedChatroomEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userReceivedChatroomEventR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userReceivedChatroomEventR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userReceivedChatroomEventAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserReceivedChatroomEvents = append(foreign.R.UserReceivedChatroomEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserReceivedChatroomEvents = append(foreign.R.UserReceivedChatroomEvents, local)
				break
			}
		}
	}

	return nil
}

// SetChatroomEvent of the userReceivedChatroomEvent to the related item.
// Sets o.R.ChatroomEvent to related.
// Adds o to related.R.UserReceivedChatroomEvents.
func (o *UserReceivedChatroomEvent) SetChatroomEvent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ChatroomEvent) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_received_chatroom_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"chatroom_event_id"}),
		strmangle.WhereClause("\"", "\"", 2, userReceivedChatroomEventPrimaryKeyColumns),
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

	o.ChatroomEventID = related.ID
	if o.R == nil {
		o.R = &userReceivedChatroomEventR{
			ChatroomEvent: related,
		}
	} else {
		o.R.ChatroomEvent = related
	}

	if related.R == nil {
		related.R = &chatroomEventR{
			UserReceivedChatroomEvents: UserReceivedChatroomEventSlice{o},
		}
	} else {
		related.R.UserReceivedChatroomEvents = append(related.R.UserReceivedChatroomEvents, o)
	}

	return nil
}

// SetUser of the userReceivedChatroomEvent to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserReceivedChatroomEvents.
func (o *UserReceivedChatroomEvent) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_received_chatroom_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userReceivedChatroomEventPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userReceivedChatroomEventR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserReceivedChatroomEvents: UserReceivedChatroomEventSlice{o},
		}
	} else {
		related.R.UserReceivedChatroomEvents = append(related.R.UserReceivedChatroomEvents, o)
	}

	return nil
}

// UserReceivedChatroomEvents retrieves all the records using an executor.
func UserReceivedChatroomEvents(mods ...qm.QueryMod) userReceivedChatroomEventQuery {
	mods = append(mods, qm.From("\"user_received_chatroom_events\""))
	return userReceivedChatroomEventQuery{NewQuery(mods...)}
}

// FindUserReceivedChatroomEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserReceivedChatroomEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserReceivedChatroomEvent, error) {
	userReceivedChatroomEventObj := &UserReceivedChatroomEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_received_chatroom_events\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userReceivedChatroomEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_received_chatroom_events")
	}

	return userReceivedChatroomEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserReceivedChatroomEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_received_chatroom_events provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userReceivedChatroomEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userReceivedChatroomEventInsertCacheMut.RLock()
	cache, cached := userReceivedChatroomEventInsertCache[key]
	userReceivedChatroomEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userReceivedChatroomEventAllColumns,
			userReceivedChatroomEventColumnsWithDefault,
			userReceivedChatroomEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userReceivedChatroomEventType, userReceivedChatroomEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userReceivedChatroomEventType, userReceivedChatroomEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_received_chatroom_events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_received_chatroom_events\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_received_chatroom_events")
	}

	if !cached {
		userReceivedChatroomEventInsertCacheMut.Lock()
		userReceivedChatroomEventInsertCache[key] = cache
		userReceivedChatroomEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserReceivedChatroomEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserReceivedChatroomEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userReceivedChatroomEventUpdateCacheMut.RLock()
	cache, cached := userReceivedChatroomEventUpdateCache[key]
	userReceivedChatroomEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userReceivedChatroomEventAllColumns,
			userReceivedChatroomEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_received_chatroom_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_received_chatroom_events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userReceivedChatroomEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userReceivedChatroomEventType, userReceivedChatroomEventMapping, append(wl, userReceivedChatroomEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_received_chatroom_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_received_chatroom_events")
	}

	if !cached {
		userReceivedChatroomEventUpdateCacheMut.Lock()
		userReceivedChatroomEventUpdateCache[key] = cache
		userReceivedChatroomEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userReceivedChatroomEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_received_chatroom_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_received_chatroom_events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserReceivedChatroomEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userReceivedChatroomEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_received_chatroom_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userReceivedChatroomEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userReceivedChatroomEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userReceivedChatroomEvent")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserReceivedChatroomEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_received_chatroom_events provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userReceivedChatroomEventColumnsWithDefault, o)

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

	userReceivedChatroomEventUpsertCacheMut.RLock()
	cache, cached := userReceivedChatroomEventUpsertCache[key]
	userReceivedChatroomEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userReceivedChatroomEventAllColumns,
			userReceivedChatroomEventColumnsWithDefault,
			userReceivedChatroomEventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userReceivedChatroomEventAllColumns,
			userReceivedChatroomEventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_received_chatroom_events, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userReceivedChatroomEventPrimaryKeyColumns))
			copy(conflict, userReceivedChatroomEventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_received_chatroom_events\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userReceivedChatroomEventType, userReceivedChatroomEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userReceivedChatroomEventType, userReceivedChatroomEventMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_received_chatroom_events")
	}

	if !cached {
		userReceivedChatroomEventUpsertCacheMut.Lock()
		userReceivedChatroomEventUpsertCache[key] = cache
		userReceivedChatroomEventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserReceivedChatroomEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserReceivedChatroomEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserReceivedChatroomEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userReceivedChatroomEventPrimaryKeyMapping)
	sql := "DELETE FROM \"user_received_chatroom_events\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_received_chatroom_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_received_chatroom_events")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userReceivedChatroomEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userReceivedChatroomEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_received_chatroom_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_received_chatroom_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserReceivedChatroomEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userReceivedChatroomEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userReceivedChatroomEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_received_chatroom_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userReceivedChatroomEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userReceivedChatroomEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_received_chatroom_events")
	}

	if len(userReceivedChatroomEventAfterDeleteHooks) != 0 {
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
func (o *UserReceivedChatroomEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserReceivedChatroomEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserReceivedChatroomEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserReceivedChatroomEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userReceivedChatroomEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_received_chatroom_events\".* FROM \"user_received_chatroom_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userReceivedChatroomEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserReceivedChatroomEventSlice")
	}

	*o = slice

	return nil
}

// UserReceivedChatroomEventExists checks if the UserReceivedChatroomEvent row exists.
func UserReceivedChatroomEventExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_received_chatroom_events\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_received_chatroom_events exists")
	}

	return exists, nil
}
