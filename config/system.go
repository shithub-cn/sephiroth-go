package config

type DbType string
type OssType string

const (
	MySql    DbType = "mysql"
	Postgres DbType = "postgres"
)

const (
	aliyun OssType = "aliyun"
)

type System struct {
	DbType        DbType  `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       OssType `mapstructure:"storage-type" json:"storage-type" yaml:"storage-type"` // Oss类型
	RouterPrefix  string  `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr          int     `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP  int     `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int     `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool    `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseRedis      bool    `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
}

type Cors struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CorsWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CorsWhitelist struct {
	AllowOrigin      string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
	AllowMethods     string `mapstructure:"allow-methods" json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     string `mapstructure:"allow-headers" json:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers" json:"expose-headers" yaml:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
}
