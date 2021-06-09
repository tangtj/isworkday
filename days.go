package main

import (
	"embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

//go:embed config
var Config embed.FS

//var year *Years

var years map[int]*Years

var loadLocker sync.RWMutex

func init() {
	loadLocker = sync.RWMutex{}
	years = make(map[int]*Years)
	//year = new(Years)
	//for i := 0; i < 12; i++ {
	//	year[i] = make(map[int]Day)
	//}
}

func GetByNow() (*Day, error) {
	now := time.Now()
	year, month, day := now.Date()
	loadLocker.RLock()
	y, ok := years[year]
	loadLocker.RUnlock()
	if !ok {
		//从文件加载
		if t, err := loadConfig(year); err != nil {
			log.Printf("加载配置文件异常 err：%s", err)
			return nil, errors.New("日期配置加载异常")
		} else {
			y = t
		}
	}
	m := y[month]
	d, ok := m[day]
	if ok {
		return &d, nil
	}
	w := time.Date(year, month, day, 8, 0, 0, 0, time.Local).Weekday()
	today := &Day{
		Year:    year,
		Month:   int(month),
		Day:     day,
		Weekday: int(w),
		Date:    now.Format("2006-01-02"),
		Work:    true,
	}
	if w == time.Sunday || w == time.Saturday {
		today.Work = false
		return today, nil
	}
	return today, nil
}

func loadConfig(year int) (*Years, error) {
	defer loadLocker.Unlock()
	loadLocker.Lock()
	if y, ok := years[year]; ok {
		return y, nil
	}

	path := "config/" + strconv.Itoa(year) + ".json"
	if f, err := Config.Open(path); err != nil {
		return nil, err
	} else {
		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		ret := make(map[int][]Day)
		if err := json.Unmarshal(bytes, &ret); err != nil {
			return nil, err
		}
		ty := new(Years)
		for month, days := range ret {
			m := make(map[int]Day)
			for _, d := range days {
				d.Year = year
				d.Month = month
				d.Weekday = int(time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.Local).Weekday())
				m[d.Day] = d
			}
			ty[month] = m
		}
		return ty, nil
	}
}
