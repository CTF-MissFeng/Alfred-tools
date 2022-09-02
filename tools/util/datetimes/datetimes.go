package datetimes

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"tools/util/consts"
)

var(
	layouts = []string{
		"2006-01-02 15:04:05.999 MST",
		"2006-01-02 15:04:05.999 -0700",
		time.RFC3339,
		time.RFC3339Nano,
		time.UnixDate,
		time.RubyDate,
		time.RFC1123Z,
	}
	moreLayouts = []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999",
	}
	RegexpTimestamp = regexp.MustCompile(`^[1-9]{1}\d+$`)
)

// ProcessNow 输出当前时间戳
func ProcessNow() {
	now := time.Now()
	// 预先添加Unix时间戳
	secs := fmt.Sprintf("%d", now.Unix())
	consts.Workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(consts.TimeIcon).
		Arg(secs).
		Valid(true)
	// 处理所有时间布局
	ProcessTimestamp(now)
}

// ProcessTimestamp 所有时间布局
func ProcessTimestamp(timestamp time.Time) {
	for _, layout := range layouts {
		v := timestamp.Format(layout)
		consts.Workflow.NewItem(v).
			Subtitle(layout).
			Icon(consts.TimeIcon).
			Arg(v).
			Valid(true)
	}
}

// ProcessTimeStr 解析时间字符串
func ProcessTimeStr(timestr string) error {
	timestamp := time.Time{}
	layoutMatch := ""
	layoutMatch, timestamp, ok := matchedLayout(layouts, timestr)
	if !ok {
		layoutMatch, timestamp, ok = matchedLayout(moreLayouts, timestr)
		if !ok {
			return errors.New("未找到匹配的时间布局")
		}
	}
	secs := fmt.Sprintf("%d", timestamp.Unix())
	consts.Workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(consts.TimeIcon).
		Arg(secs).
		Valid(true)

	for _, layout := range layouts {
		if layout == layoutMatch {
			continue
		}
		v := timestamp.Format(layout)
		consts.Workflow.NewItem(v).
			Subtitle(layout).
			Icon(consts.TimeIcon).
			Arg(v).
			Valid(true)
	}

	return nil
}

func matchedLayout(layouts []string, timestr string) (matched string, timestamp time.Time, ok bool) {
	for _, layout := range layouts {
		v, err := time.Parse(layout, timestr)
		if err == nil {
			return layout, v, true
		}
	}
	return
}