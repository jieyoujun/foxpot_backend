package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GetAttackMapData 给球数据
func GetAttackMapData(c *gin.Context) {
	session := sessions.Default(c)
	sr, err := models.GetAllTypeData(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    200,
			"message": "攻击地图数据查询失败",
		})
		return
	}
	var (
		data  []models.AttackMapData
		datum models.AttackMapData
	)
	for _, hit := range sr.Hits.Hits {
		err := json.Unmarshal(hit.Source, &datum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    200,
				"message": "反序列化失败",
				"data":    datum,
			})
			return
		}
		// 源IP
		if utils.IsInternal(datum.SrcIP) {
			// TODO
			// 优化查询公网IP速度, 考虑缓存, 先放到session
			if extIP, ok := session.Get("ext_ip").(string); ok {
				datum.SrcIP = extIP
			} else {
				extIP, err = utils.GetExternalIPByHTTP()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    200,
						"message": "公网IP查询失败",
						"data":    datum,
					})
					return
				}
				session.Set("ext_ip", extIP)
				datum.SrcIP = extIP
			}
		}
		srcInfo, err := models.GetGeoIP2Info(datum.SrcIP)
		if err != nil {
			datum.SrcLat = 0.0
			datum.SrcLng = 0.0
			datum.SrcRegion = "未知"
		} else {
			datum.SrcLat = srcInfo.Latitude
			datum.SrcLng = srcInfo.Longitude
			datum.SrcRegion = srcInfo.Region
		}
		// 目的IP
		if utils.IsInternal(datum.DstIP) {
			if extIP, ok := session.Get("ext_ip").(string); ok {
				datum.DstIP = extIP
			} else {
				extIP, err = utils.GetExternalIPByHTTP()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    200,
						"message": "公网IP查询失败",
						"data":    datum,
					})
					return
				}
				session.Set("ext_ip", extIP)
				datum.DstIP = extIP
			}
		}
		dstInfo, err := models.GetGeoIP2Info(datum.DstIP)
		if err != nil {
			datum.DstLat = 0.0
			datum.DstLng = 0.0
			datum.DstRegion = "未知"
		} else {
			datum.DstLat = dstInfo.Latitude
			datum.DstLng = dstInfo.Longitude
			datum.DstRegion = dstInfo.Region
		}
		data = append(data, datum)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "攻击地图数据查询成功",
		"data":    data,
	})
}

// GetAttackMapCtr 给球统计数据
func GetAttackMapCtr(c *gin.Context) {
	sr, err := models.GetAllTypeCtr(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    200,
			"message": "攻击地图统计数据查询失败",
		})
		return
	}
	ar := struct {
		CtrAllTime struct {
			Buckets []struct {
				Key            string `json:"key"`
				Count          int    `json:"doc_count"`
				CtrSmallerTime struct {
					Buckets []struct {
						Key   string `json:"key"`
						Count int    `json:"doc_count"`
					}
				} `json:"ctr_smaller_time"`
			} `json:"buckets"`
		} `json:"ctr_all_time"`
	}{}
	bs, err := json.Marshal(sr.Aggregations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    200,
			"message": "攻击地图统计数据序列化失败",
			"data":    sr.Aggregations,
		})
		return
	}
	err = json.Unmarshal(bs, &ar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    200,
			"message": "攻击地图统计数据反序列化失败",
			"data":    string(bs),
		})
		return
	}
	var ctr []models.AttackMapCtr
	for _, b := range ar.CtrAllTime.Buckets {
		ctr = append(ctr, models.AttackMapCtr{
			SourceType: b.Key,
			CtrAllTime: b.Count,
			Ctr7d:      b.CtrSmallerTime.Buckets[0].Count,
			Ctr1d:      b.CtrSmallerTime.Buckets[1].Count,
			Ctr1h:      b.CtrSmallerTime.Buckets[2].Count,
			Ctr1m:      b.CtrSmallerTime.Buckets[3].Count,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "攻击地图统计数据查询成功",
		"data":    ctr,
	})
}
