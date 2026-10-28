package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"
)

// 独立驱动验证事务边界，不访问学习数据库。
type groupTxDriver struct{}
type groupTxConn struct{ state *groupTxState }
type groupTxState struct {
	beginErr, commitErr   error
	committed, rolledBack bool
}

func (groupTxDriver) Open(string) (driver.Conn, error) { return nil, errors.New("use connector") }

type groupTxConnector struct{ state *groupTxState }

func (c groupTxConnector) Driver() driver.Driver { return groupTxDriver{} }
func (c groupTxConnector) Connect(context.Context) (driver.Conn, error) {
	return &groupTxConn{c.state}, nil
}
func (c *groupTxConn) Prepare(string) (driver.Stmt, error) { return nil, errors.New("not used") }
func (c *groupTxConn) Close() error                        { return nil }
func (c *groupTxConn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}
func (c *groupTxConn) BeginTx(context.Context, driver.TxOptions) (driver.Tx, error) {
	if c.state.beginErr != nil {
		return nil, c.state.beginErr
	}
	return c.state, nil
}
func (s *groupTxState) Commit() error   { s.committed = true; return s.commitErr }
func (s *groupTxState) Rollback() error { s.rolledBack = true; return nil }

func TestGroupTransactionBoundary(t *testing.T) {
	cause := errors.New("database failure")
	tests := []struct {
		name                                 string
		beginErr, actionErr, commitErr       error
		wantAction, wantCommit, wantRollback bool
	}{
		{"success", nil, nil, nil, true, true, false},
		{"insert_failure_rolls_back_gid", nil, cause, nil, true, false, true},
		{"begin_failure_skips_action", cause, nil, nil, false, false, false},
		{"commit_failure_propagates", nil, nil, cause, true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &groupTxState{beginErr: tt.beginErr, commitErr: tt.commitErr}
			db := sql.OpenDB(groupTxConnector{state})
			defer db.Close()
			called := false
			err := NewGroupRepository(db).WithTransaction(context.Background(), func(*sql.Tx) error {
				called = true
				return tt.actionErr
			})
			wantError := tt.beginErr != nil || tt.actionErr != nil || tt.commitErr != nil
			if (err != nil) != wantError || (wantError && !errors.Is(err, cause)) {
				t.Fatalf("unexpected error: %v", err)
			}
			if called != tt.wantAction || state.committed != tt.wantCommit || state.rolledBack != tt.wantRollback {
				t.Fatalf("action=%v commit=%v rollback=%v", called, state.committed, state.rolledBack)
			}
		})
	}
}
