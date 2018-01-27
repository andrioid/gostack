package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andrioid/gostack/graphql"
	"github.com/andrioid/gostack/module"
	"github.com/andrioid/gostack/places"
	"github.com/gorilla/mux"
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
	viper.SetDefault("firebase.apiKey", "insertyours")
	viper.SetDefault("firebase.authDomain", "insertyours")
	viper.SetDefault("firebase.databaseURL", "insertyours")
	viper.SetDefault("firebase.projectId", "insertyours")
	viper.SetDefault("firebase.storageBucket", "insertyours")
	viper.SetDefault("firebase.messagingSenderId", "insertyours")
	viper.SetDefault("firebase.serviceAccountFile", "service-account-key.json")

	if err := viper.SafeWriteConfigAs(".gostack.toml"); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfigAs(".gostack.toml")
		}
	}

	// Database
	dbType := viper.GetString("db.type")
	dbOptions := viper.GetString("db.options")
	fmt.Printf("db type: '%v'\n", dbType)
	db, err := gorm.Open(dbType, dbOptions)
	defer db.Close()
	if err != nil {
		panic("failed to connect to database")
	}

	// Firebase
	// https://firebase.google.com/docs/auth/admin/verify-id-tokens
	// https://firebase.google.com/docs/admin/setup
	//opt := option.WithCredentialsFile(viper.GetString("firebase.serviceAccountFile"))
	//app, err := firebase.NewApp(context.Background(), nil, opt)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	// TODO: app isn't accessable to graphql package, so maybe create middleware for token verification

	// Modules
	var modules []module.Module
	placesModule, _ := places.New(db)
	modules = append(modules, placesModule)
	graphql.CreateSchema(modules)

	r := mux.NewRouter()
	r.HandleFunc("/graphql", graphql.HTTPHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build")))
	http.Handle("/", r)

	fmt.Printf("[serve] started %v\n", dbType)
	// http.HandleFunc("/graphql", graphql.HTTPHandler)
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
