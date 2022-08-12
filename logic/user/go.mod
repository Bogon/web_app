module webapp.io/logic/user

go 1.18

replace webapp.io/dao/mysql/user => ../../dao/mysql/user

replace webapp.io/pkg/snowflakeID => ../../pkg/snowflakeID

require (
	webapp.io/dao/mysql/user v0.0.0-00010101000000-000000000000
	webapp.io/pkg/snowflakeID v0.0.0-00010101000000-000000000000
)

require github.com/bwmarrin/snowflake v0.3.0 // indirect
