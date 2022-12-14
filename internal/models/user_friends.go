// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// UserFriend is an object representing the database table.
type UserFriend struct {
	ID           int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	RequesterID  int       `boil:"requester_id" json:"requester_id" toml:"requester_id" yaml:"requester_id"`
	TargetID     int       `boil:"target_id" json:"target_id" toml:"target_id" yaml:"target_id"`
	RelationType int  `boil:"relation_type" json:"relation_type,omitempty" toml:"relation_type" yaml:"relation_type,omitempty"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userFriendR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userFriendL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserFriendColumns = struct {
	ID           string
	RequesterID  string
	TargetID     string
	RelationType string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	RequesterID:  "requester_id",
	TargetID:     "target_id",
	RelationType: "relation_type",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var UserFriendTableColumns = struct {
	ID           string
	RequesterID  string
	TargetID     string
	RelationType string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "user_friends.id",
	RequesterID:  "user_friends.requester_id",
	TargetID:     "user_friends.target_id",
	RelationType: "user_friends.relation_type",
	CreatedAt:    "user_friends.created_at",
	UpdatedAt:    "user_friends.updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var UserFriendWhere = struct {
	ID           whereHelperint
	RequesterID  whereHelperint
	TargetID     whereHelperint
	RelationType whereHelperint
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	ID:           whereHelperint{field: "\"user_friends\".\"id\""},
	RequesterID:  whereHelperint{field: "\"user_friends\".\"requester_id\""},
	TargetID:     whereHelperint{field: "\"user_friends\".\"target_id\""},
	RelationType: whereHelperint{field: "\"user_friends\".\"relation_type\""},
	CreatedAt:    whereHelpertime_Time{field: "\"user_friends\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"user_friends\".\"updated_at\""},
}

// UserFriendRels is where relationship names are stored.
var UserFriendRels = struct {
	Requester string
	Target    string
}{
	Requester: "Requester",
	Target:    "Target",
}

// userFriendR is where relationships are stored.
type userFriendR struct {
	Requester *User `boil:"Requester" json:"Requester" toml:"Requester" yaml:"Requester"`
	Target    *User `boil:"Target" json:"Target" toml:"Target" yaml:"Target"`
}

// NewStruct creates a new relationship struct
func (*userFriendR) NewStruct() *userFriendR {
	return &userFriendR{}
}

func (r *userFriendR) GetRequester() *User {
	if r == nil {
		return nil
	}
	return r.Requester
}

func (r *userFriendR) GetTarget() *User {
	if r == nil {
		return nil
	}
	return r.Target
}

// userFriendL is where Load methods for each relationship are stored.
type userFriendL struct{}

var (
	userFriendAllColumns            = []string{"id", "requester_id", "target_id", "relation_type", "created_at", "updated_at"}
	userFriendColumnsWithoutDefault = []string{"requester_id", "target_id"}
	userFriendColumnsWithDefault    = []string{"id", "relation_type", "created_at", "updated_at"}
	userFriendPrimaryKeyColumns     = []string{"id"}
	userFriendGeneratedColumns      = []string{}
)

type (
	// UserFriendSlice is an alias for a slice of pointers to UserFriend.
	// This should almost always be used instead of []UserFriend.
	UserFriendSlice []*UserFriend

	userFriendQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userFriendType                 = reflect.TypeOf(&UserFriend{})
	userFriendMapping              = queries.MakeStructMapping(userFriendType)
	userFriendPrimaryKeyMapping, _ = queries.BindMapping(userFriendType, userFriendMapping, userFriendPrimaryKeyColumns)
	userFriendInsertCacheMut       sync.RWMutex
	userFriendInsertCache          = make(map[string]insertCache)
	userFriendUpdateCacheMut       sync.RWMutex
	userFriendUpdateCache          = make(map[string]updateCache)
	userFriendUpsertCacheMut       sync.RWMutex
	userFriendUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single userFriend record from the query.
func (q userFriendQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserFriend, error) {
	o := &UserFriend{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_friends")
	}

	return o, nil
}

// All returns all UserFriend records from the query.
func (q userFriendQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserFriendSlice, error) {
	var o []*UserFriend

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserFriend slice")
	}

	return o, nil
}

// Count returns the count of all UserFriend records in the query.
func (q userFriendQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_friends rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userFriendQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_friends exists")
	}

	return count > 0, nil
}

// Requester pointed to by the foreign key.
func (o *UserFriend) Requester(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RequesterID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Target pointed to by the foreign key.
func (o *UserFriend) Target(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TargetID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadRequester allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userFriendL) LoadRequester(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserFriend interface{}, mods queries.Applicator) error {
	var slice []*UserFriend
	var object *UserFriend

	if singular {
		var ok bool
		object, ok = maybeUserFriend.(*UserFriend)
		if !ok {
			object = new(UserFriend)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserFriend)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserFriend))
			}
		}
	} else {
		s, ok := maybeUserFriend.(*[]*UserFriend)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserFriend)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserFriend))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userFriendR{}
		}
		args = append(args, object.RequesterID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userFriendR{}
			}

			for _, a := range args {
				if a == obj.RequesterID {
					continue Outer
				}
			}

			args = append(args, obj.RequesterID)

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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Requester = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.RequesterUserFriends = append(foreign.R.RequesterUserFriends, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RequesterID == foreign.ID {
				local.R.Requester = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.RequesterUserFriends = append(foreign.R.RequesterUserFriends, local)
				break
			}
		}
	}

	return nil
}

// LoadTarget allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userFriendL) LoadTarget(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserFriend interface{}, mods queries.Applicator) error {
	var slice []*UserFriend
	var object *UserFriend

	if singular {
		var ok bool
		object, ok = maybeUserFriend.(*UserFriend)
		if !ok {
			object = new(UserFriend)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserFriend)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserFriend))
			}
		}
	} else {
		s, ok := maybeUserFriend.(*[]*UserFriend)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserFriend)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserFriend))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userFriendR{}
		}
		args = append(args, object.TargetID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userFriendR{}
			}

			for _, a := range args {
				if a == obj.TargetID {
					continue Outer
				}
			}

			args = append(args, obj.TargetID)

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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Target = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.TargetUserFriends = append(foreign.R.TargetUserFriends, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TargetID == foreign.ID {
				local.R.Target = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.TargetUserFriends = append(foreign.R.TargetUserFriends, local)
				break
			}
		}
	}

	return nil
}

// SetRequester of the userFriend to the related item.
// Sets o.R.Requester to related.
// Adds o to related.R.RequesterUserFriends.
func (o *UserFriend) SetRequester(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_friends\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"requester_id"}),
		strmangle.WhereClause("\"", "\"", 2, userFriendPrimaryKeyColumns),
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

	o.RequesterID = related.ID
	if o.R == nil {
		o.R = &userFriendR{
			Requester: related,
		}
	} else {
		o.R.Requester = related
	}

	if related.R == nil {
		related.R = &userR{
			RequesterUserFriends: UserFriendSlice{o},
		}
	} else {
		related.R.RequesterUserFriends = append(related.R.RequesterUserFriends, o)
	}

	return nil
}

// SetTarget of the userFriend to the related item.
// Sets o.R.Target to related.
// Adds o to related.R.TargetUserFriends.
func (o *UserFriend) SetTarget(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_friends\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"target_id"}),
		strmangle.WhereClause("\"", "\"", 2, userFriendPrimaryKeyColumns),
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

	o.TargetID = related.ID
	if o.R == nil {
		o.R = &userFriendR{
			Target: related,
		}
	} else {
		o.R.Target = related
	}

	if related.R == nil {
		related.R = &userR{
			TargetUserFriends: UserFriendSlice{o},
		}
	} else {
		related.R.TargetUserFriends = append(related.R.TargetUserFriends, o)
	}

	return nil
}

// UserFriends retrieves all the records using an executor.
func UserFriends(mods ...qm.QueryMod) userFriendQuery {
	mods = append(mods, qm.From("\"user_friends\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_friends\".*"})
	}

	return userFriendQuery{q}
}

// FindUserFriend retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserFriend(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserFriend, error) {
	userFriendObj := &UserFriend{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_friends\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userFriendObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_friends")
	}

	return userFriendObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserFriend) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_friends provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userFriendColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userFriendInsertCacheMut.RLock()
	cache, cached := userFriendInsertCache[key]
	userFriendInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userFriendAllColumns,
			userFriendColumnsWithDefault,
			userFriendColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userFriendType, userFriendMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userFriendType, userFriendMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_friends\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_friends\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_friends")
	}

	if !cached {
		userFriendInsertCacheMut.Lock()
		userFriendInsertCache[key] = cache
		userFriendInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the UserFriend.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserFriend) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	userFriendUpdateCacheMut.RLock()
	cache, cached := userFriendUpdateCache[key]
	userFriendUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userFriendAllColumns,
			userFriendPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_friends, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_friends\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userFriendPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userFriendType, userFriendMapping, append(wl, userFriendPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_friends row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_friends")
	}

	if !cached {
		userFriendUpdateCacheMut.Lock()
		userFriendUpdateCache[key] = cache
		userFriendUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q userFriendQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_friends")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_friends")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserFriendSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userFriendPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_friends\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userFriendPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userFriend slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userFriend")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserFriend) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_friends provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(userFriendColumnsWithDefault, o)

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

	userFriendUpsertCacheMut.RLock()
	cache, cached := userFriendUpsertCache[key]
	userFriendUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userFriendAllColumns,
			userFriendColumnsWithDefault,
			userFriendColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userFriendAllColumns,
			userFriendPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_friends, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userFriendPrimaryKeyColumns))
			copy(conflict, userFriendPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_friends\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userFriendType, userFriendMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userFriendType, userFriendMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert user_friends")
	}

	if !cached {
		userFriendUpsertCacheMut.Lock()
		userFriendUpsertCache[key] = cache
		userFriendUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single UserFriend record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserFriend) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserFriend provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userFriendPrimaryKeyMapping)
	sql := "DELETE FROM \"user_friends\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_friends")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_friends")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userFriendQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userFriendQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_friends")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_friends")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserFriendSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userFriendPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_friends\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userFriendPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userFriend slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_friends")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserFriend) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserFriend(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserFriendSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserFriendSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userFriendPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_friends\".* FROM \"user_friends\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userFriendPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserFriendSlice")
	}

	*o = slice

	return nil
}

// UserFriendExists checks if the UserFriend row exists.
func UserFriendExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_friends\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_friends exists")
	}

	return exists, nil
}
