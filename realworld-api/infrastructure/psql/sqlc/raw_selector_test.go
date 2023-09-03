package sqlc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v2"
	. "github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_RawSelector_Select(t *testing.T) {
	type testStruct struct {
		Value  string             `db:"value"`
		Value2 int                `db:"value2"`
		Time   pgtype.Timestamptz `db:"time"`
		UUID   uuid.UUID          `db:"uid"`
	}

	type args struct {
		builder squirrel.SelectBuilder
	}
	type mock_query struct {
		args_query   string
		args_params  []any
		returns_rows *pgxmock.Rows
		returns_err  error
	}
	type mocks struct {
		query mock_query
	}
	type wants struct {
		result []testStruct
		err    error
	}

	var (
		dummyErr  = errors.New("dummy error")
		now       = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		uid1      = uuid.New()
		uid2      = uuid.New()
		testData1 = struct {
			args  args
			mocks mocks
			wants wants
		}{
			args: args{
				builder: squirrel.
					Select("*").
					From("test_table").
					Where(squirrel.Eq{
						"value":  "test",
						"value2": 1,
					}),
			},
			mocks: mocks{
				query: mock_query{
					args_query:  `SELECT \* FROM test_table WHERE value = .* AND value2 = .*`,
					args_params: []any{"test", 1},
					returns_rows: pgxmock.
						NewRows([]string{"value", "value2", "time", "uid"}).
						AddRow("value1", 1, now, uid1).
						AddRow("value2", 2, now, uid2),
					returns_err: nil,
				},
			},
			wants: wants{
				result: []testStruct{
					{
						Value:  "value1",
						Value2: 1,
						Time: pgtype.Timestamptz{
							Time:  now,
							Valid: true,
						},
						UUID: uid1,
					},
					{
						Value:  "value2",
						Value2: 2,
						Time: pgtype.Timestamptz{
							Time:  now,
							Valid: true,
						},
						UUID: uid2,
					},
				},
				err: nil,
			},
		}
	)

	tests := []struct {
		name  string
		args  args
		mocks mocks
		wants wants
	}{
		{
			name:  "valid_params_should_call_expected_query_and_return_expected_result",
			args:  testData1.args,
			mocks: testData1.mocks,
			wants: testData1.wants,
		},
		{
			name: "valid_params_with_query_returns_error_should_call_expected_query_and_return_expected_error",
			args: testData1.args,
			mocks: mocks{
				query: mock_query{
					args_query:  testData1.mocks.query.args_query,
					args_params: testData1.mocks.query.args_params,
					returns_err: dummyErr,
				},
			},
			wants: wants{
				err: dummyErr,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			db.ExpectQuery(tt.mocks.query.args_query).
				WithArgs(tt.mocks.query.args_params...).
				WillReturnRows(tt.mocks.query.returns_rows).
				WillReturnError(tt.mocks.query.returns_err)

			got, err := NewRawSelector[testStruct]().Select(context.Background(), db, tt.args.builder)
			assert.Equal(t, tt.wants.result, got)
			if !assert.ErrorIs(t, err, tt.wants.err) {
				t.Logf("error: %v", err)
			}
		})
	}
}
