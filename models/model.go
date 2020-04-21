package models

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
	"github.com/likiiiiii/foxpot_backend/utils"
	"github.com/olivere/elastic/v7"
	"github.com/oschwald/geoip2-golang"
)

// DB 数据库
var (
	DB    *gorm.DB
	GeoDB *geoip2.Reader
	ESCli *elastic.Client
)

// InitDB 初始化数据库
func InitDB() (err error) {
	// ES
	ESCli, err = elastic.NewClient(
		elastic.SetURL("https://"+utils.Config.ES.Address),
		elastic.SetSniff(false),
		elastic.SetHttpClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}),
	)
	if err != nil {
		return err
	}

	// GeoIP2
	GeoDB, err = geoip2.Open(utils.Config.GeoIP2.CityPath)
	if err != nil {
		return err
	}

	// Sqlite
	DB, err = gorm.Open(utils.Config.DB.Type, utils.Config.DB.DSN)
	if err != nil {
		return err
	}

	// >>>>>>>>>
	// 配置连接池
	// <<<<<<<<<
	DB.DB().SetMaxOpenConns(utils.Config.DB.MaxOpenConn)
	DB.DB().SetMaxIdleConns(utils.Config.DB.MaxIdleConn)
	DB.DB().SetConnMaxLifetime(time.Duration(utils.Config.DB.MaxLifeTime) * time.Second)

	if !DB.HasTable(&User{}) {
		DB.AutoMigrate(&User{})
	}
	return nil
}
