package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type NewsFeed struct {
	Type       int     `json:"type"`
	FirmName   string  `json:"firm_name"`
	Name       string  `json:"name"`
	Value      float64 `json:"value"`
	UpdateTime string  `json:"update_time"`
}

func main() {
	drsPre := "deal:drs:"
	amtPre := "deal:amt:"
	t := time.Now()
	drs := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 8, 10, 15, 21, 25, 30, 33, 35, 37, 40, 41, 42, 45, 47, 49, 50, 52}
	amt := []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5, 0.8, 1.0, 1.2, 1.5, 2.0, 3.0, 3.5, 4.1, 5.2, 6.0, 6.5, 7.0, 7.3, 7.8, 8.0, 9.2, 10}
	cmd1 := "mset"
	cmd2 := "mset"
	for h := 0; h < 16; h++ {
		min0 := time.Date(t.Year(), t.Month(), t.Day(), h, 0, 0, 0, time.Local).Format("200601021504")
		min30 := time.Date(t.Year(), t.Month(), t.Day(), h, 30, 0, 0, time.Local).Format("200601021504")
		cmd1 += fmt.Sprintf(" %s%s %d", drsPre, min0, drs[h*2])
		cmd1 += fmt.Sprintf(" %s%s %d", drsPre, min30, drs[h*2+1])
		cmd2 += fmt.Sprintf(" %s%s %.2f", amtPre, min0, amt[h*2])
		cmd2 += fmt.Sprintf(" %s%s %.2f", amtPre, min30, amt[h*2+1])
	}
	fmt.Println("period_30m:")
	fmt.Println(cmd1)
	fmt.Println(cmd2)

	cmd3 := "lpush newsfeed"
	cmd4 := "ltrim newsfeed 0 9"
	newsfeed := []*NewsFeed{{
		Type:       0,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市人口统计数据",
		Value:      132087,
		UpdateTime: time.Date(2020, 12, 10, 9, 0, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       1,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市近10年道路增长统计数据",
		Value:      10987,
		UpdateTime: time.Date(2020, 12, 10, 15, 0, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       0,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市人口统计数据",
		Value:      132087,
		UpdateTime: time.Date(2020, 12, 10, 15, 30, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       0,
		FirmName:   "河南省新乡市社保中心",
		Name:       "新乡市2020年参保人员统计数据",
		Value:      100087,
		UpdateTime: time.Date(2020, 12, 10, 16, 30, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       0,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市近10年工业经济增长统计数据",
		Value:      213761,
		UpdateTime: time.Date(2020, 12, 10, 16, 50, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       0,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市近10年工业经济增长统计数据",
		Value:      100087,
		UpdateTime: time.Date(2020, 12, 10, 19, 50, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       1,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市人口统计数据",
		Value:      132087,
		UpdateTime: time.Date(2020, 12, 11, 9, 50, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       1,
		FirmName:   "河南省新乡市社保中心",
		Name:       "新乡市2020年参保人员统计数据",
		Value:      100087,
		UpdateTime: time.Date(2020, 12, 11, 10, 0, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       0,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市近5年工业经济增长统计数据",
		Value:      213761,
		UpdateTime: time.Date(2020, 12, 11, 10, 10, 19, 0, time.Local).Format(time.RFC3339),
	}, {
		Type:       0,
		FirmName:   "河南省新乡市统计局",
		Name:       "新乡市近5年工业经济增长统计数据",
		Value:      100087,
		UpdateTime: time.Date(2020, 12, 11, 10, 35, 15, 0, time.Local).Format(time.RFC3339),
	}}
	for i := 0; i < 10; i++ {
		nf, err := json.Marshal(newsfeed[i])
		if err != nil {
			log.Fatal("marshal newsfeed failed")
		}
		cmd3 += fmt.Sprintf(" '%s'", string(nf))
	}
	fmt.Println("newsfeed:")
	fmt.Println(cmd3)
	fmt.Println(cmd4)
}
