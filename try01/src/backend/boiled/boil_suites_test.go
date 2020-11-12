// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiled

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Customers", testCustomers)
	t.Run("Todos", testTodos)
}

func TestDelete(t *testing.T) {
	t.Run("Customers", testCustomersDelete)
	t.Run("Todos", testTodosDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Customers", testCustomersQueryDeleteAll)
	t.Run("Todos", testTodosQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Customers", testCustomersSliceDeleteAll)
	t.Run("Todos", testTodosSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Customers", testCustomersExists)
	t.Run("Todos", testTodosExists)
}

func TestFind(t *testing.T) {
	t.Run("Customers", testCustomersFind)
	t.Run("Todos", testTodosFind)
}

func TestBind(t *testing.T) {
	t.Run("Customers", testCustomersBind)
	t.Run("Todos", testTodosBind)
}

func TestOne(t *testing.T) {
	t.Run("Customers", testCustomersOne)
	t.Run("Todos", testTodosOne)
}

func TestAll(t *testing.T) {
	t.Run("Customers", testCustomersAll)
	t.Run("Todos", testTodosAll)
}

func TestCount(t *testing.T) {
	t.Run("Customers", testCustomersCount)
	t.Run("Todos", testTodosCount)
}

func TestHooks(t *testing.T) {
	t.Run("Customers", testCustomersHooks)
	t.Run("Todos", testTodosHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Customers", testCustomersInsert)
	t.Run("Customers", testCustomersInsertWhitelist)
	t.Run("Todos", testTodosInsert)
	t.Run("Todos", testTodosInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("TodoToCustomerUsingUser", testTodoToOneCustomerUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CustomerToUserTodos", testCustomerToManyUserTodos)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("TodoToCustomerUsingUserTodos", testTodoToOneSetOpCustomerUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CustomerToUserTodos", testCustomerToManyAddOpUserTodos)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Customers", testCustomersReload)
	t.Run("Todos", testTodosReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Customers", testCustomersReloadAll)
	t.Run("Todos", testTodosReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Customers", testCustomersSelect)
	t.Run("Todos", testTodosSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Customers", testCustomersUpdate)
	t.Run("Todos", testTodosUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Customers", testCustomersSliceUpdateAll)
	t.Run("Todos", testTodosSliceUpdateAll)
}
