module github.com/ishisora/go_todo_app

go 1.23

toolchain go1.23.1

require golang.org/x/sync v0.8.0

require (
	github.com/caarlos0/env/v11 v11.2.0
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-playground/validator v9.31.0+incompatible
)

require github.com/jmoiron/sqlx v1.4.0

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/go-sql-driver/mysql v1.8.1
	github.com/matryer/moq v0.5.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/tools v0.25.0 // indirect
)

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/go-cmp v0.6.0
	github.com/leodido/go-urn v1.4.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
