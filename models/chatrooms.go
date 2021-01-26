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

// Chatroom is an object representing the database table.
type Chatroom struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *chatroomR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chatroomL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChatroomColumns = struct {
	ID        string
	Name      string
	CreatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ChatroomWhere = struct {
	ID        whereHelperint
	Name      whereHelpernull_String
	CreatedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"chatrooms\".\"id\""},
	Name:      whereHelpernull_String{field: "\"chatrooms\".\"name\""},
	CreatedAt: whereHelpernull_Time{field: "\"chatrooms\".\"created_at\""},
}

// ChatroomRels is where relationship names are stored.
var ChatroomRels = struct {
	ChatroomEvents string
	ChatroomUsers  string
	Messages       string
}{
	ChatroomEvents: "ChatroomEvents",
	ChatroomUsers:  "ChatroomUsers",
	Messages:       "Messages",
}

// chatroomR is where relationships are stored.
type chatroomR struct {
	ChatroomEvents ChatroomEventSlice `boil:"ChatroomEvents" json:"ChatroomEvents" toml:"ChatroomEvents" yaml:"ChatroomEvents"`
	ChatroomUsers  ChatroomUserSlice  `boil:"ChatroomUsers" json:"ChatroomUsers" toml:"ChatroomUsers" yaml:"ChatroomUsers"`
	Messages       MessageSlice       `boil:"Messages" json:"Messages" toml:"Messages" yaml:"Messages"`
}

// NewStruct creates a new relationship struct
func (*chatroomR) NewStruct() *chatroomR {
	return &chatroomR{}
}

// chatroomL is where Load methods for each relationship are stored.
type chatroomL struct{}

var (
	chatroomAllColumns            = []string{"id", "name", "created_at"}
	chatroomColumnsWithoutDefault = []string{"name", "created_at"}
	chatroomColumnsWithDefault    = []string{"id"}
	chatroomPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChatroomSlice is an alias for a slice of pointers to Chatroom.
	// This should generally be used opposed to []Chatroom.
	ChatroomSlice []*Chatroom
	// ChatroomHook is the signature for custom Chatroom hook methods
	ChatroomHook func(context.Context, boil.ContextExecutor, *Chatroom) error

	chatroomQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chatroomType                 = reflect.TypeOf(&Chatroom{})
	chatroomMapping              = queries.MakeStructMapping(chatroomType)
	chatroomPrimaryKeyMapping, _ = queries.BindMapping(chatroomType, chatroomMapping, chatroomPrimaryKeyColumns)
	chatroomInsertCacheMut       sync.RWMutex
	chatroomInsertCache          = make(map[string]insertCache)
	chatroomUpdateCacheMut       sync.RWMutex
	chatroomUpdateCache          = make(map[string]updateCache)
	chatroomUpsertCacheMut       sync.RWMutex
	chatroomUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var chatroomBeforeInsertHooks []ChatroomHook
var chatroomBeforeUpdateHooks []ChatroomHook
var chatroomBeforeDeleteHooks []ChatroomHook
var chatroomBeforeUpsertHooks []ChatroomHook

var chatroomAfterInsertHooks []ChatroomHook
var chatroomAfterSelectHooks []ChatroomHook
var chatroomAfterUpdateHooks []ChatroomHook
var chatroomAfterDeleteHooks []ChatroomHook
var chatroomAfterUpsertHooks []ChatroomHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Chatroom) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Chatroom) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Chatroom) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Chatroom) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Chatroom) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Chatroom) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Chatroom) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Chatroom) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Chatroom) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChatroomHook registers your hook function for all future operations.
func AddChatroomHook(hookPoint boil.HookPoint, chatroomHook ChatroomHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		chatroomBeforeInsertHooks = append(chatroomBeforeInsertHooks, chatroomHook)
	case boil.BeforeUpdateHook:
		chatroomBeforeUpdateHooks = append(chatroomBeforeUpdateHooks, chatroomHook)
	case boil.BeforeDeleteHook:
		chatroomBeforeDeleteHooks = append(chatroomBeforeDeleteHooks, chatroomHook)
	case boil.BeforeUpsertHook:
		chatroomBeforeUpsertHooks = append(chatroomBeforeUpsertHooks, chatroomHook)
	case boil.AfterInsertHook:
		chatroomAfterInsertHooks = append(chatroomAfterInsertHooks, chatroomHook)
	case boil.AfterSelectHook:
		chatroomAfterSelectHooks = append(chatroomAfterSelectHooks, chatroomHook)
	case boil.AfterUpdateHook:
		chatroomAfterUpdateHooks = append(chatroomAfterUpdateHooks, chatroomHook)
	case boil.AfterDeleteHook:
		chatroomAfterDeleteHooks = append(chatroomAfterDeleteHooks, chatroomHook)
	case boil.AfterUpsertHook:
		chatroomAfterUpsertHooks = append(chatroomAfterUpsertHooks, chatroomHook)
	}
}

// One returns a single chatroom record from the query.
func (q chatroomQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Chatroom, error) {
	o := &Chatroom{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chatrooms")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Chatroom records from the query.
func (q chatroomQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChatroomSlice, error) {
	var o []*Chatroom

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Chatroom slice")
	}

	if len(chatroomAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Chatroom records in the query.
func (q chatroomQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chatrooms rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q chatroomQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chatrooms exists")
	}

	return count > 0, nil
}

// ChatroomEvents retrieves all the chatroom_event's ChatroomEvents with an executor.
func (o *Chatroom) ChatroomEvents(mods ...qm.QueryMod) chatroomEventQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"chatroom_events\".\"chatroom_id\"=?", o.ID),
	)

	query := ChatroomEvents(queryMods...)
	queries.SetFrom(query.Query, "\"chatroom_events\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"chatroom_events\".*"})
	}

	return query
}

// ChatroomUsers retrieves all the chatroom_user's ChatroomUsers with an executor.
func (o *Chatroom) ChatroomUsers(mods ...qm.QueryMod) chatroomUserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"chatroom_users\".\"chatroom_id\"=?", o.ID),
	)

	query := ChatroomUsers(queryMods...)
	queries.SetFrom(query.Query, "\"chatroom_users\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"chatroom_users\".*"})
	}

	return query
}

// Messages retrieves all the message's Messages with an executor.
func (o *Chatroom) Messages(mods ...qm.QueryMod) messageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"messages\".\"chatroom_id\"=?", o.ID),
	)

	query := Messages(queryMods...)
	queries.SetFrom(query.Query, "\"messages\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"messages\".*"})
	}

	return query
}

// LoadChatroomEvents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (chatroomL) LoadChatroomEvents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatroom interface{}, mods queries.Applicator) error {
	var slice []*Chatroom
	var object *Chatroom

	if singular {
		object = maybeChatroom.(*Chatroom)
	} else {
		slice = *maybeChatroom.(*[]*Chatroom)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatroomR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatroomR{}
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
		qm.From(`chatroom_events`),
		qm.WhereIn(`chatroom_events.chatroom_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load chatroom_events")
	}

	var resultSlice []*ChatroomEvent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice chatroom_events")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on chatroom_events")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chatroom_events")
	}

	if len(chatroomEventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChatroomEvents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &chatroomEventR{}
			}
			foreign.R.Chatroom = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChatroomID {
				local.R.ChatroomEvents = append(local.R.ChatroomEvents, foreign)
				if foreign.R == nil {
					foreign.R = &chatroomEventR{}
				}
				foreign.R.Chatroom = local
				break
			}
		}
	}

	return nil
}

// LoadChatroomUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (chatroomL) LoadChatroomUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatroom interface{}, mods queries.Applicator) error {
	var slice []*Chatroom
	var object *Chatroom

	if singular {
		object = maybeChatroom.(*Chatroom)
	} else {
		slice = *maybeChatroom.(*[]*Chatroom)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatroomR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatroomR{}
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
		qm.From(`chatroom_users`),
		qm.WhereIn(`chatroom_users.chatroom_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load chatroom_users")
	}

	var resultSlice []*ChatroomUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice chatroom_users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on chatroom_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chatroom_users")
	}

	if len(chatroomUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChatroomUsers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &chatroomUserR{}
			}
			foreign.R.Chatroom = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChatroomID {
				local.R.ChatroomUsers = append(local.R.ChatroomUsers, foreign)
				if foreign.R == nil {
					foreign.R = &chatroomUserR{}
				}
				foreign.R.Chatroom = local
				break
			}
		}
	}

	return nil
}

// LoadMessages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (chatroomL) LoadMessages(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatroom interface{}, mods queries.Applicator) error {
	var slice []*Chatroom
	var object *Chatroom

	if singular {
		object = maybeChatroom.(*Chatroom)
	} else {
		slice = *maybeChatroom.(*[]*Chatroom)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatroomR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatroomR{}
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
		qm.From(`messages`),
		qm.WhereIn(`messages.chatroom_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load messages")
	}

	var resultSlice []*Message
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice messages")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on messages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for messages")
	}

	if len(messageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Messages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &messageR{}
			}
			foreign.R.Chatroom = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChatroomID {
				local.R.Messages = append(local.R.Messages, foreign)
				if foreign.R == nil {
					foreign.R = &messageR{}
				}
				foreign.R.Chatroom = local
				break
			}
		}
	}

	return nil
}

// AddChatroomEvents adds the given related objects to the existing relationships
// of the chatroom, optionally inserting them as new records.
// Appends related to o.R.ChatroomEvents.
// Sets related.R.Chatroom appropriately.
func (o *Chatroom) AddChatroomEvents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChatroomEvent) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChatroomID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"chatroom_events\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"chatroom_id"}),
				strmangle.WhereClause("\"", "\"", 2, chatroomEventPrimaryKeyColumns),
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

			rel.ChatroomID = o.ID
		}
	}

	if o.R == nil {
		o.R = &chatroomR{
			ChatroomEvents: related,
		}
	} else {
		o.R.ChatroomEvents = append(o.R.ChatroomEvents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &chatroomEventR{
				Chatroom: o,
			}
		} else {
			rel.R.Chatroom = o
		}
	}
	return nil
}

// AddChatroomUsers adds the given related objects to the existing relationships
// of the chatroom, optionally inserting them as new records.
// Appends related to o.R.ChatroomUsers.
// Sets related.R.Chatroom appropriately.
func (o *Chatroom) AddChatroomUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChatroomUser) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChatroomID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"chatroom_users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"chatroom_id"}),
				strmangle.WhereClause("\"", "\"", 2, chatroomUserPrimaryKeyColumns),
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

			rel.ChatroomID = o.ID
		}
	}

	if o.R == nil {
		o.R = &chatroomR{
			ChatroomUsers: related,
		}
	} else {
		o.R.ChatroomUsers = append(o.R.ChatroomUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &chatroomUserR{
				Chatroom: o,
			}
		} else {
			rel.R.Chatroom = o
		}
	}
	return nil
}

// AddMessages adds the given related objects to the existing relationships
// of the chatroom, optionally inserting them as new records.
// Appends related to o.R.Messages.
// Sets related.R.Chatroom appropriately.
func (o *Chatroom) AddMessages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Message) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChatroomID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"messages\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"chatroom_id"}),
				strmangle.WhereClause("\"", "\"", 2, messagePrimaryKeyColumns),
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

			rel.ChatroomID = o.ID
		}
	}

	if o.R == nil {
		o.R = &chatroomR{
			Messages: related,
		}
	} else {
		o.R.Messages = append(o.R.Messages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &messageR{
				Chatroom: o,
			}
		} else {
			rel.R.Chatroom = o
		}
	}
	return nil
}

// Chatrooms retrieves all the records using an executor.
func Chatrooms(mods ...qm.QueryMod) chatroomQuery {
	mods = append(mods, qm.From("\"chatrooms\""))
	return chatroomQuery{NewQuery(mods...)}
}

// FindChatroom retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChatroom(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Chatroom, error) {
	chatroomObj := &Chatroom{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"chatrooms\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, chatroomObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chatrooms")
	}

	return chatroomObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Chatroom) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chatrooms provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chatroomColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	chatroomInsertCacheMut.RLock()
	cache, cached := chatroomInsertCache[key]
	chatroomInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			chatroomAllColumns,
			chatroomColumnsWithDefault,
			chatroomColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(chatroomType, chatroomMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chatroomType, chatroomMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"chatrooms\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"chatrooms\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into chatrooms")
	}

	if !cached {
		chatroomInsertCacheMut.Lock()
		chatroomInsertCache[key] = cache
		chatroomInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Chatroom.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Chatroom) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	chatroomUpdateCacheMut.RLock()
	cache, cached := chatroomUpdateCache[key]
	chatroomUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			chatroomAllColumns,
			chatroomPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update chatrooms, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"chatrooms\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, chatroomPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chatroomType, chatroomMapping, append(wl, chatroomPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update chatrooms row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for chatrooms")
	}

	if !cached {
		chatroomUpdateCacheMut.Lock()
		chatroomUpdateCache[key] = cache
		chatroomUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q chatroomQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for chatrooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for chatrooms")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChatroomSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatroomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"chatrooms\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, chatroomPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in chatroom slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all chatroom")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Chatroom) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chatrooms provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chatroomColumnsWithDefault, o)

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

	chatroomUpsertCacheMut.RLock()
	cache, cached := chatroomUpsertCache[key]
	chatroomUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			chatroomAllColumns,
			chatroomColumnsWithDefault,
			chatroomColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			chatroomAllColumns,
			chatroomPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert chatrooms, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(chatroomPrimaryKeyColumns))
			copy(conflict, chatroomPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"chatrooms\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(chatroomType, chatroomMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chatroomType, chatroomMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert chatrooms")
	}

	if !cached {
		chatroomUpsertCacheMut.Lock()
		chatroomUpsertCache[key] = cache
		chatroomUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Chatroom record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Chatroom) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Chatroom provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chatroomPrimaryKeyMapping)
	sql := "DELETE FROM \"chatrooms\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from chatrooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for chatrooms")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q chatroomQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no chatroomQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chatrooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chatrooms")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChatroomSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(chatroomBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatroomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"chatrooms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chatroomPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chatroom slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chatrooms")
	}

	if len(chatroomAfterDeleteHooks) != 0 {
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
func (o *Chatroom) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChatroom(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChatroomSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChatroomSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatroomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"chatrooms\".* FROM \"chatrooms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chatroomPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChatroomSlice")
	}

	*o = slice

	return nil
}

// ChatroomExists checks if the Chatroom row exists.
func ChatroomExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"chatrooms\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chatrooms exists")
	}

	return exists, nil
}
