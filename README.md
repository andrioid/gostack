# Gostack experiment
If I was starting a green-field project, I would consider something like this.

The idea is to be modern, but as lean as possible.

## What I want

### Uses ENV vars and TOML for configuration
Handled by Cobra and Viper from SPF13

### Can use whatever database golang/sql can
I'm considering sqlx or gorm for this purpose. I might also skip data mapping all together.

My ideal production environment would be postgres + some sort of cache (redis maybe).

### GraphQL for API
Handled by go-graphql/go-graphql

### Use an external JWT auth provider
- Going to get it working with Firebase first
- Maybe add auth0 later

### Organized by business-domain into packages
Just my wanting to keep business logic as a first-class concern. So technical sub-directories are to be kept at a bare minimum.

## Structure
The current directory vision is:
- /awesomeproject/accounts: A domain package of awesomeproject. Fulfills gostack/module interface
- /awesomeproject/main.go: Calls gostack.NewHandler(...modules)
    - This will iterate over modules, generate query and mutation types, hook up /graphql handler and such

How inter-linked your modules are, is up to you. Gostack doesn't care. But, if you need to access other packages, you need to do it from your own program and not from gostack.


## TODO

- Parse for configuration file (use rootCmd lookup thingy)
- Iterate over modules to create graphql handler
- Try adding some data types and querying them from GraphiQL
- Hook up Postgres or Sqlite
- Setup http middleware
    - Move CORS shit in there
    - Enforce Firebase Auth here
- Login screen
    - Prompt login before GraphiQL if client has no token
    - Store client token in local storage or some shit