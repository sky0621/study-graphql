// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiled

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

func testTodos(t *testing.T) {
	t.Parallel()

	query := Todos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTodosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
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

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTodosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Todos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTodosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TodoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTodosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TodoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Todo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TodoExists to return true, but got false.")
	}
}

func testTodosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	todoFound, err := FindTodo(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if todoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTodosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Todos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTodosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Todos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTodosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	todoOne := &Todo{}
	todoTwo := &Todo{}
	if err = randomize.Struct(seed, todoOne, todoDBTypes, false, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}
	if err = randomize.Struct(seed, todoTwo, todoDBTypes, false, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = todoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = todoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Todos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTodosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	todoOne := &Todo{}
	todoTwo := &Todo{}
	if err = randomize.Struct(seed, todoOne, todoDBTypes, false, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}
	if err = randomize.Struct(seed, todoTwo, todoDBTypes, false, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = todoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = todoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func todoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func todoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Todo) error {
	*o = Todo{}
	return nil
}

func testTodosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Todo{}
	o := &Todo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, todoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Todo object: %s", err)
	}

	AddTodoHook(boil.BeforeInsertHook, todoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	todoBeforeInsertHooks = []TodoHook{}

	AddTodoHook(boil.AfterInsertHook, todoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	todoAfterInsertHooks = []TodoHook{}

	AddTodoHook(boil.AfterSelectHook, todoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	todoAfterSelectHooks = []TodoHook{}

	AddTodoHook(boil.BeforeUpdateHook, todoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	todoBeforeUpdateHooks = []TodoHook{}

	AddTodoHook(boil.AfterUpdateHook, todoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	todoAfterUpdateHooks = []TodoHook{}

	AddTodoHook(boil.BeforeDeleteHook, todoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	todoBeforeDeleteHooks = []TodoHook{}

	AddTodoHook(boil.AfterDeleteHook, todoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	todoAfterDeleteHooks = []TodoHook{}

	AddTodoHook(boil.BeforeUpsertHook, todoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	todoBeforeUpsertHooks = []TodoHook{}

	AddTodoHook(boil.AfterUpsertHook, todoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	todoAfterUpsertHooks = []TodoHook{}
}

func testTodosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTodosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(todoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTodoToOneCustomerUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Todo
	var foreign Customer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, todoDBTypes, false, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
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

	slice := TodoSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Todo)(&slice), nil); err != nil {
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

func testTodoToOneSetOpCustomerUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Todo
	var b, c Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, todoDBTypes, false, strmangle.SetComplement(todoPrimaryKeyColumns, todoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Customer{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserTodos[0] != &a {
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

func testTodosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
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

func testTodosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TodoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTodosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Todos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	todoDBTypes = map[string]string{`ID`: `bigint`, `Task`: `character varying`, `UserID`: `bigint`}
	_           = bytes.MinRead
)

func testTodosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(todoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(todoAllColumns) == len(todoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, todoDBTypes, true, todoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTodosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(todoAllColumns) == len(todoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Todo{}
	if err = randomize.Struct(seed, o, todoDBTypes, true, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, todoDBTypes, true, todoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(todoAllColumns, todoPrimaryKeyColumns) {
		fields = todoAllColumns
	} else {
		fields = strmangle.SetComplement(
			todoAllColumns,
			todoPrimaryKeyColumns,
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

	slice := TodoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTodosUpsert(t *testing.T) {
	t.Parallel()

	if len(todoAllColumns) == len(todoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Todo{}
	if err = randomize.Struct(seed, &o, todoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Todo: %s", err)
	}

	count, err := Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, todoDBTypes, false, todoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Todo: %s", err)
	}

	count, err = Todos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
