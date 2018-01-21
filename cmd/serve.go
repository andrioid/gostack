package cmd

import (
	"fmt"
	"net/http"

	"github.com/andrioid/gostack/graphql"
	"github.com/andrioid/gostack/places"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	// Listen on HTTP port
	// Handle /graphql
	// dbType := rootCmd.PersistentFlags().Lookup("db.type")
	viper.SetDefault("db.type", "sqlite3")
	viper.SetDefault("db.options", ".test.db")
	dbType := viper.GetString("db.type")
	dbOptions := viper.GetString("db.options")

	fmt.Printf("db type: '%v'\n", dbType)
	db, err := gorm.Open(dbType, dbOptions)
	if err != nil {
		panic("failed to connect to database")
	}
	placesModule, _ := places.New(db)
	placesModule.Hello()

	// Iterate over modules
	// Generate schema from query and mutation types
	// Generate http handler from schema for /graphql
	fmt.Printf("[serve] started %v\n", dbType)
	http.HandleFunc("/graphql", graphql.HTTPHandler)
	http.ListenAndServe(":8080", nil)
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
