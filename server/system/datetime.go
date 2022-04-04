package system

import (
	"math"
	"net/http"
	"time"
)

var (
	cst *time.Location
)

// Location 默认时区
const Location = "Asia/Shanghai"

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

type Datetime struct {
	*option
}

type Option func(opt *option)

type option struct {
	location  string
	cstLayout string
}

func WithLocation(location string) Option {
	return func(opt *option) {
		opt.location = location
	}
}

func WithCSTLayout(cst string) Option {
	return func(opt *option) {
		opt.cstLayout = cst
	}
}

func NewDatetime(opts ...Option) *Datetime {
	var err error
	var o = &option{
		location:  Location,
		cstLayout: CSTLayout,
	}

	for _, opt := range opts {
		opt(o)
	}

	var datetime = new(Datetime)
	datetime.option = o
	cst, err = time.LoadLocation(o.location)
	if err != nil {
		time.Local = cst // 设置时区
	}

	return datetime
}

// RFC3339ToCSTLayout convert rfc3339 value to China standard time layout
// 2020-11-08T08:18:46+08:00 => 2020-11-08 08:18:46
func (t *Datetime) RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(t.cstLayout), nil
}

// CSTLayoutString 格式化时间
// 返回 "2006-01-02 15:04:05" 格式的时间
func (t *Datetime) CSTLayoutString() string {
	return time.Now().In(cst).Format(t.cstLayout)
}

// ParseCSTInLocation 格式化时间
func (t *Datetime) ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(t.cstLayout, date, cst)
}

// CSTLayoutStringToUnix 返回 unix 时间戳
// 2020-01-24 21:11:11 => 1579871471
func (t *Datetime) CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(t.cstLayout, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}
	return stamp.Unix(), nil
}

// GMTLayoutString 格式化时间
// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func (t *Datetime) GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// ParseGMTInLocation 格式化时间
func (t *Datetime) ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

// SubInLocation 计算时间差
func (t *Datetime) SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(cst).Sub(ts).Seconds())
}