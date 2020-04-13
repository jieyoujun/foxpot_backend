package models

import (
	"context"

	"github.com/likiiiiii/foxpot_backend/utils"
	"github.com/olivere/elastic/v7"
)

// GetAllTypeData 获取最近1分钟所有数据
func GetAllTypeData(ctx context.Context) (*elastic.SearchResult, error) {
	return ESCli.Search().
		Index(utils.Config.ES.IndexName).
		Query(elastic.NewRangeQuery("@timestamp").
			Gte("now-1m").
			Lte("now")).
		Size(10000).
		Do(ctx)
}

// GetAllTypeCtr 获取所有数据的统计数据
func GetAllTypeCtr(ctx context.Context) (*elastic.SearchResult, error) {
	return ESCli.Search().
		Index(utils.Config.ES.IndexName).
		Query(elastic.NewMatchAllQuery()).
		Aggregation("ctr_all_time", elastic.NewTermsAggregation().
			Field("type.keyword").
			SubAggregation("ctr_smaller_time", elastic.NewRangeAggregation().
				Field("@timestamp").
				AddRangeWithKey("7d", "now-7d", "now").
				AddRangeWithKey("1d", "now-1d", "now").
				AddRangeWithKey("1h", "now-1h", "now").
				AddRangeWithKey("1m", "now-1m", "now"))).
		Size(0).
		Do(ctx)
}
