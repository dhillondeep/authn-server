module github.com/keratin/authn-server

require (
	github.com/airbrake/gobrake v3.5.0+incompatible
	github.com/dlclark/regexp2 v1.1.6 // indirect
	github.com/felixge/httpsnoop v1.0.0
	github.com/getsentry/sentry-go v0.3.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.3.0
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/handlers v1.3.0
	github.com/gorilla/mux v1.6.1
	github.com/gorilla/schema v1.1.0
	github.com/jmoiron/sqlx v0.0.0-20170430194603-d9bd385d68c0
	github.com/lib/pq v0.0.0-20180327071824-d34b9ff171c2
	github.com/mattn/go-sqlite3 v1.6.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/test-go/testify v1.1.4
	github.com/trustelem/zxcvbn v1.0.1
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/square/go-jose.v2 v2.3.1
)

go 1.13
