package cmd

import (
	"database/sql"
	"os"
	"text/template"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const (
	dialectPostgres = "postgres"
	dialectPgx      = "pgx"
)

var (
	paths      []string
	dialect    string
	connection string
)

var rootCmd = &cobra.Command{
	Use:     "testdata-loader",
	Short:   "ローカル環境用テストデータ読み込みツール",
	Example: "testdata-loader --dir [FIXTURES_DIRECTORY_PATH]",
	Run:     runRootCmd,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	d := dialect
	if d == dialectPostgres {
		d = dialectPgx
	}
	db, err := sql.Open(d, connection)
	if err != nil {
		cmd.PrintErrf("failed to open database: %v", err)
		os.Exit(1)
	}

	// NOTE(ryutah): New 内のオプションの並び順が違うと、うまくテンプレートが動作しなかったりするので注意
	// see: https://github.com/go-testfixtures/testfixtures/issues/87
	loader, err := testfixtures.New(
		testfixtures.Template(),
		testfixtures.TemplateDelims("{{", "}}"),
		testfixtures.TemplateFuncs(template.FuncMap{
			"Iterate": func(cnt uint) []uint {
				return lo.RangeFrom(uint(0), int(cnt))
			},
		}),
		testfixtures.Database(db),
		testfixtures.Dialect(dialect),
		testfixtures.DangerousSkipCleanupFixtureTables(),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.SkipResetSequences(),
		testfixtures.Paths(paths...),
	)
	if err != nil {
		cmd.PrintErrf("failed to init fixtures: %v", err)
		os.Exit(1)
	}
	if err := loader.Load(); err != nil {
		cmd.PrintErrf("failed to load fixtures: %v", err)
		os.Exit(1)
	}

	cmd.Println("success load fixtures")
}

func init() {
	rootCmd.PersistentFlags().StringArrayVarP(&paths, "path", "p", nil, "required: fixtures path.")
	rootCmd.PersistentFlags().StringVarP(&dialect, "dialect", "d", "postgres", "dialect of database. (default: postgres)")
	rootCmd.PersistentFlags().StringVarP(&connection, "connection_name", "c", "", "required: connection name of database.")

	if err := rootCmd.MarkPersistentFlagFilename("path"); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("path"); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("connection_name"); err != nil {
		panic(err)
	}
}
