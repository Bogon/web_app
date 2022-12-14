module webapp.io/logic/post

go 1.18

replace webapp.io/models => ../../models

replace webapp.io/pkg/snowflakeID => ../../pkg/snowflakeID

replace webapp.io/dao/mysql => ../../dao/mysql

replace webapp.io/settings => ../../settings

replace webapp.io/dao/redis => ../../dao/redis

require (
	go.uber.org/zap v1.17.0
	webapp.io/dao/mysql v0.0.0-00010101000000-000000000000
	webapp.io/dao/redis v0.0.0-00010101000000-000000000000
	webapp.io/models v0.0.0-00010101000000-000000000000
	webapp.io/pkg/snowflakeID v0.0.0-00010101000000-000000000000
)

require (
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	webapp.io/settings v0.0.0-00010101000000-000000000000 // indirect
)
