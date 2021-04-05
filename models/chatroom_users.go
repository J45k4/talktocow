// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// ChatroomUser is an object representing the database table.
type ChatroomUser struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ChatroomID int       `boil:"chatroom_id" json:"chatroom_id" toml:"chatroom_id" yaml:"chatroom_id"`
	CreatedAt  null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *chatroomUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chatroomUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChatroomUserColumns = struct {
	ID         string
	UserID     string
	ChatroomID string
	CreatedAt  string
}{
	ID:         "id",
	UserID:     "user_id",
	ChatroomID: "chatroom_id",
	CreatedAt:  "created_at",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ChatroomUserWhere = struct {
	ID         whereHelperint
	UserID     whereHelperint
	ChatroomID whereHelperint
	CreatedAt  whereHelpernull_Time
}{
	ID:         whereHelperint{field: "\"chatroom_users\".\"id\""},
	UserID:     whereHelperint{field: "\"chatroom_users\".\"user_id\""},
	ChatroomID: whereHelperint{field: "\"chatroom_users\".\"chatroom_id\""},
	CreatedAt:  whereHelpernull_Time{field: "\"chatroom_users\".\"created_at\""},
}

// ChatroomUserRels is where relationship names are stored.
var ChatroomUserRels = struct {
	Chatroom string
	User     string
}{
	Chatroom: "Chatroom",
	User:     "User",
}

// chatroomUserR is where relationships are stored.
type chatroomUserR struct {
	Chatroom *Chatroom `boil:"Chatroom" json:"Chatroom" toml:"Chatroom" yaml:"Chatroom"`
	User     *User     `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*chatroomUserR) NewStruct() *chatroomUserR {
	return &chatroomUserR{}
}

// chatroomUserL is where Load methods for each relationship are stored.
type chatroomUserL struct{}

var (
	chatroomUserAllColumns            = []string{"id", "user_id", "chatroom_id", "created_at"}
	chatroomUserColumnsWithoutDefault = []string{"user_id", "chatroom_id", "created_at"}
	chatroomUserColumnsWithDefault    = []string{"id"}
	chatroomUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChatroomUserSlice is an alias for a slice of pointers to ChatroomUser.
	// This should generally be used opposed to []ChatroomUser.
	ChatroomUserSlice []*ChatroomUser
	// ChatroomUserHook is the signature for custom ChatroomUser hook methods
	ChatroomUserHook func(context.Context, boil.ContextExecutor, *ChatroomUser) error

	chatroomUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chatroomUserType                 = reflect.TypeOf(&ChatroomUser{})
	chatroomUserMapping              = queries.MakeStructMapping(chatroomUserType)
	chatroomUserPrimaryKeyMapping, _ = queries.BindMapping(chatroomUserType, chatroomUserMapping, chatroomUserPrimaryKeyColumns)
	chatroomUserInsertCacheMut       sync.RWMutex
	chatroomUserInsertCache          = make(map[string]insertCache)
	chatroomUserUpdateCacheMut       sync.RWMutex
	chatroomUserUpdateCache          = make(map[string]updateCache)
	chatroomUserUpsertCacheMut       sync.RWMutex
	chatroomUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var chatroomUserBeforeInsertHooks []ChatroomUserHook
var chatroomUserBeforeUpdateHooks []ChatroomUserHook
var chatroomUserBeforeDeleteHooks []ChatroomUserHook
var chatroomUserBeforeUpsertHooks []ChatroomUserHook

var chatroomUserAfterInsertHooks []ChatroomUserHook
var chatroomUserAfterSelectHooks []ChatroomUserHook
var chatroomUserAfterUpdateHooks []ChatroomUserHook
var chatroomUserAfterDeleteHooks []ChatroomUserHook
var chatroomUserAfterUpsertHooks []ChatroomUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChatroomUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChatroomUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChatroomUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChatroomUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChatroomUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChatroomUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChatroomUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChatroomUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChatroomUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatroomUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChatroomUserHook registers your hook function for all future operations.
func AddChatroomUserHook(hookPoint boil.HookPoint, chatroomUserHook ChatroomUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		chatroomUserBeforeInsertHooks = append(chatroomUserBeforeInsertHooks, chatroomUserHook)
	case boil.BeforeUpdateHook:
		chatroomUserBeforeUpdateHooks = append(chatroomUserBeforeUpdateHooks, chatroomUserHook)
	case boil.BeforeDeleteHook:
		chatroomUserBeforeDeleteHooks = append(chatroomUserBeforeDeleteHooks, chatroomUserHook)
	case boil.BeforeUpsertHook:
		chatroomUserBeforeUpsertHooks = append(chatroomUserBeforeUpsertHooks, chatroomUserHook)
	case boil.AfterInsertHook:
		chatroomUserAfterInsertHooks = append(chatroomUserAfterInsertHooks, chatroomUserHook)
	case boil.AfterSelectHook:
		chatroomUserAfterSelectHooks = append(chatroomUserAfterSelectHooks, chatroomUserHook)
	case boil.AfterUpdateHook:
		chatroomUserAfterUpdateHooks = append(chatroomUserAfterUpdateHooks, chatroomUserHook)
	case boil.AfterDeleteHook:
		chatroomUserAfterDeleteHooks = append(chatroomUserAfterDeleteHooks, chatroomUserHook)
	case boil.AfterUpsertHook:
		chatroomUserAfterUpsertHooks = append(chatroomUserAfterUpsertHooks, chatroomUserHook)
	}
}

// One returns a single chatroomUser record from the query.
func (q chatroomUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChatroomUser, error) {
	o := &ChatroomUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chatroom_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChatroomUser records from the query.
func (q chatroomUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChatroomUserSlice, error) {
	var o []*ChatroomUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChatroomUser slice")
	}

	if len(chatroomUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChatroomUser records in the query.
func (q chatroomUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chatroom_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q chatroomUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chatroom_users exists")
	}

	return count > 0, nil
}

// Chatroom pointed to by the foreign key.
func (o *ChatroomUser) Chatroom(mods ...qm.QueryMod) chatroomQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChatroomID),
	}

	queryMods = append(queryMods, mods...)

	query := Chatrooms(queryMods...)
	queries.SetFrom(query.Query, "\"chatrooms\"")

	return query
}

// User pointed to by the foreign key.
func (o *ChatroomUser) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadChatroom allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (chatroomUserL) LoadChatroom(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatroomUser interface{}, mods queries.Applicator) error {
	var slice []*ChatroomUser
	var object *ChatroomUser

	if singular {
		object = maybeChatroomUser.(*ChatroomUser)
	} else {
		slice = *maybeChatroomUser.(*[]*ChatroomUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatroomUserR{}
		}
		args = append(args, object.ChatroomID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatroomUserR{}
			}

			for _, a := range args {
				if a == obj.ChatroomID {
					continue Outer
				}
			}

			args = append(args, obj.ChatroomID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`chatrooms`),
		qm.WhereIn(`chatrooms.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Chatroom")
	}

	var resultSlice []*Chatroom
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Chatroom")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for chatrooms")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chatrooms")
	}

	if len(chatroomUserAfterSelectHooks) != 0 {
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
		object.R.Chatroom = foreign
		if foreign.R == nil {
			foreign.R = &chatroomR{}
		}
		foreign.R.ChatroomUsers = append(foreign.R.ChatroomUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChatroomID == foreign.ID {
				local.R.Chatroom = foreign
				if foreign.R == nil {
					foreign.R = &chatroomR{}
				}
				foreign.R.ChatroomUsers = append(foreign.R.ChatroomUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (chatroomUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatroomUser interface{}, mods queries.Applicator) error {
	var slice []*ChatroomUser
	var object *ChatroomUser

	if singular {
		object = maybeChatroomUser.(*ChatroomUser)
	} else {
		slice = *maybeChatroomUser.(*[]*ChatroomUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatroomUserR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatroomUserR{}
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

	if len(chatroomUserAfterSelectHooks) != 0 {
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
		foreign.R.ChatroomUsers = append(foreign.R.ChatroomUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.ChatroomUsers = append(foreign.R.ChatroomUsers, local)
				break
			}
		}
	}

	return nil
}

// SetChatroom of the chatroomUser to the related item.
// Sets o.R.Chatroom to related.
// Adds o to related.R.ChatroomUsers.
func (o *ChatroomUser) SetChatroom(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Chatroom) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"chatroom_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"chatroom_id"}),
		strmangle.WhereClause("\"", "\"", 2, chatroomUserPrimaryKeyColumns),
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

	o.ChatroomID = related.ID
	if o.R == nil {
		o.R = &chatroomUserR{
			Chatroom: related,
		}
	} else {
		o.R.Chatroom = related
	}

	if related.R == nil {
		related.R = &chatroomR{
			ChatroomUsers: ChatroomUserSlice{o},
		}
	} else {
		related.R.ChatroomUsers = append(related.R.ChatroomUsers, o)
	}

	return nil
}

// SetUser of the chatroomUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.ChatroomUsers.
func (o *ChatroomUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"chatroom_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, chatroomUserPrimaryKeyColumns),
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
		o.R = &chatroomUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			ChatroomUsers: ChatroomUserSlice{o},
		}
	} else {
		related.R.ChatroomUsers = append(related.R.ChatroomUsers, o)
	}

	return nil
}

// ChatroomUsers retrieves all the records using an executor.
func ChatroomUsers(mods ...qm.QueryMod) chatroomUserQuery {
	mods = append(mods, qm.From("\"chatroom_users\""))
	return chatroomUserQuery{NewQuery(mods...)}
}

// FindChatroomUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChatroomUser(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ChatroomUser, error) {
	chatroomUserObj := &ChatroomUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"chatroom_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, chatroomUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chatroom_users")
	}

	return chatroomUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChatroomUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chatroom_users provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(chatroomUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	chatroomUserInsertCacheMut.RLock()
	cache, cached := chatroomUserInsertCache[key]
	chatroomUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			chatroomUserAllColumns,
			chatroomUserColumnsWithDefault,
			chatroomUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(chatroomUserType, chatroomUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chatroomUserType, chatroomUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"chatroom_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"chatroom_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into chatroom_users")
	}

	if !cached {
		chatroomUserInsertCacheMut.Lock()
		chatroomUserInsertCache[key] = cache
		chatroomUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChatroomUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChatroomUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	chatroomUserUpdateCacheMut.RLock()
	cache, cached := chatroomUserUpdateCache[key]
	chatroomUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			chatroomUserAllColumns,
			chatroomUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update chatroom_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"chatroom_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, chatroomUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chatroomUserType, chatroomUserMapping, append(wl, chatroomUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update chatroom_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for chatroom_users")
	}

	if !cached {
		chatroomUserUpdateCacheMut.Lock()
		chatroomUserUpdateCache[key] = cache
		chatroomUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q chatroomUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for chatroom_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for chatroom_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChatroomUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatroomUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"chatroom_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, chatroomUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in chatroomUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all chatroomUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChatroomUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chatroom_users provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(chatroomUserColumnsWithDefault, o)

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

	chatroomUserUpsertCacheMut.RLock()
	cache, cached := chatroomUserUpsertCache[key]
	chatroomUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			chatroomUserAllColumns,
			chatroomUserColumnsWithDefault,
			chatroomUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			chatroomUserAllColumns,
			chatroomUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert chatroom_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(chatroomUserPrimaryKeyColumns))
			copy(conflict, chatroomUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"chatroom_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(chatroomUserType, chatroomUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chatroomUserType, chatroomUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert chatroom_users")
	}

	if !cached {
		chatroomUserUpsertCacheMut.Lock()
		chatroomUserUpsertCache[key] = cache
		chatroomUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChatroomUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChatroomUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChatroomUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chatroomUserPrimaryKeyMapping)
	sql := "DELETE FROM \"chatroom_users\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from chatroom_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for chatroom_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q chatroomUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no chatroomUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chatroom_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chatroom_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChatroomUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(chatroomUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatroomUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"chatroom_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chatroomUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chatroomUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chatroom_users")
	}

	if len(chatroomUserAfterDeleteHooks) != 0 {
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
func (o *ChatroomUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChatroomUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChatroomUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChatroomUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatroomUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"chatroom_users\".* FROM \"chatroom_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chatroomUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChatroomUserSlice")
	}

	*o = slice

	return nil
}

// ChatroomUserExists checks if the ChatroomUser row exists.
func ChatroomUserExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"chatroom_users\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chatroom_users exists")
	}

	return exists, nil
}
