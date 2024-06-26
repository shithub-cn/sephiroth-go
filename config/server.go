package config

import (
	"sephiroth-go/config/datasource"
	"sephiroth-go/config/storage"
)

type Server struct {
	Jwt   Jwt              `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap   Zap              `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis datasource.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Mongo   Mongo            `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  datasource.Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  datasource.Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DbList []datasource.SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// storage
	Local     storage.Local     `mapstructure:"local" json:"local" yaml:"local"`
	AliyunOss storage.AliyunOss `mapstructure:"aliyun-storage" json:"aliyun-storage" yaml:"aliyun-storage"`
	Excel     Excel             `mapstructure:"excel" json:"excel" yaml:"excel"`
	DiskList  []DiskList        `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`
	Cors      Cors              `mapstructure:"cors" json:"cors" yaml:"cors"`
}
