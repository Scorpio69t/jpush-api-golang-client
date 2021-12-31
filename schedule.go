package jpush

import (
	"encoding/json"
	"time"
)

type Schecule struct {
	Cid     string                 `json:"cid"`     // 定时任务id
	Name    string                 `json:"name"`    // 定时任务名称
	Enabled bool                   `json:"enabled"` // 是否启用
	Trigger map[string]interface{} `json:"trigger"` // 定时任务触发条件
	Push    *PayLoad               `json:"push"`    // 定时任务推送内容
}

const (
	formatTime = "2006-01-02 15:04:05"
)

// NewSchedule 创建定时任务
func NewSchedule(cid, name string, enabled bool, push *PayLoad) *Schecule {
	return &Schecule{
		Cid:     cid,
		Name:    name,
		Enabled: enabled,
		Push:    push,
	}
}

// SetCid 设置定时任务id
func (s *Schecule) SetCid(cid string) {
	s.Cid = cid
}

// GetCid 获取定时任务id
func (s *Schecule) GetCid() string {
	return s.Cid
}

// SetName 设置定时任务名称
func (s *Schecule) SetName(name string) {
	s.Name = name
}

// GetName 获取定时任务名称
func (s *Schecule) GetName() string {
	return s.Name
}

// SetEnabled 设置定时任务是否启用
func (s *Schecule) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

// GetEnabled 获取定时任务是否启用
func (s *Schecule) GetEnabled() bool {
	return s.Enabled
}

// SetPayLoad 设置定时任务推送内容
func (s *Schecule) SetPayLoad(push *PayLoad) {
	s.Push = push
}

// SingleTrigger 单次触发
func (s *Schecule) SingleTrigger(t time.Time) {
	s.Trigger = map[string]interface{}{
		"single": map[string]interface{}{
			"time": t.Format(formatTime),
		},
	}
}

// PeriodicalTrigger 周期触发
func (s *Schecule) PeriodicalTrigger(start, end, t time.Time, timeUnit string, frequency int, point []string) {
	s.Trigger = map[string]interface{}{
		"periodical": map[string]interface{}{
			"start":     start.Format(formatTime),
			"end":       end.Format(formatTime),
			"time":      t.Format(formatTime),
			"time_unit": timeUnit,
			"frequency": frequency,
			"point":     point,
		},
	}
}

// Bytes 转换为字节数组
func (s *Schecule) Bytes() ([]byte, error) {
	return json.Marshal(s)
}
