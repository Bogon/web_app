module webapp.io

go 1.18

replace webapp.io/settings => ./settings

replace webapp.io/logger => ./logger

replace webapp.io/dao/redis => ./dao/redis

replace webapp.io/appRoutes => ./appRoutes

replace webapp.io/controllers/userHanlder => ./controllers/userHandler

replace webapp.io/models => ./models

replace webapp.io/pkg/snowflakeID => ./pkg/snowflakeID

replace webapp.io/logic/user => ./logic/user

replace webapp.io/dao/mysql => ./dao/mysql

replace webapp.io/controllers/validatorHandler => ./controllers/validatorHandler

replace webapp.io/controllers/responseHandler => ./controllers/responseHandler

replace webapp.io/controllers/responseCode => ./controllers/responseCode

replace webapp.io/middlewares/jwtauth => ./middlewares/jwtauth

replace webapp.io/controllers/community => ./controllers/community

replace webapp.io/logic/community => ./logic/community

replace webapp.io/pkg/jwt => ./pkg/jwt

replace webapp.io/controllers/post => ./controllers/post

replace webapp.io/controllers/requestHandler => ./controllers/requestHandler

replace webapp.io/logic/post => ./logic/post

replace webapp.io/logic/voted => ./logic/voted

replace webapp.io/controllers/vote => ./controllers/vote

replace webapp.io/middlewares/ratelimit => ./middlewares/ratelimit

require (
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.2
	github.com/swaggo/swag v1.8.4
	go.uber.org/zap v1.22.0
	webapp.io/appRoutes v0.0.0-00010101000000-000000000000
	webapp.io/controllers/validatorHandler v0.0.0-00010101000000-000000000000
	webapp.io/dao/mysql v0.0.0-00010101000000-000000000000
	webapp.io/dao/redis v0.0.0-00010101000000-000000000000
	webapp.io/logger v0.0.0-00010101000000-000000000000
	webapp.io/pkg/snowflakeID v0.0.0-00010101000000-000000000000
	webapp.io/settings v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gin-contrib/pprof v1.4.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.8.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ratelimit v1.0.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.4 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220824171710-5757bc0c5503 // indirect
	golang.org/x/net v0.0.0-20220822230855-b0a4917ee28c // indirect
	golang.org/x/sys v0.0.0-20220825204002-c680a09ffe64 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	webapp.io/controllers/community v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/controllers/post v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/controllers/requestHandler v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/controllers/responseCode v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/controllers/responseHandler v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/controllers/userHanlder v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/controllers/vote v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/logic/community v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/logic/post v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/logic/user v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/logic/voted v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/middlewares/jwtauth v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/middlewares/ratelimit v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/models v0.0.0-00010101000000-000000000000 // indirect
	webapp.io/pkg/jwt v0.0.0-00010101000000-000000000000 // indirect
)
