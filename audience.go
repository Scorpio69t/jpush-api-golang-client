package jpush

import (
	"log"
)

type AudienceType string

const (
	TAG             AudienceType = "tag"              // 标签OR
	TAG_AND         AudienceType = "tag_and"          // 标签AND
	TAG_NOT         AudienceType = "tag_not"          // 标签NOT
	ALIAS           AudienceType = "alias"            // 别名
	REGISTRATION_ID AudienceType = "registration_id"  // 注册ID
	SEGMENT         AudienceType = "segment"          // 用户分群 ID
	ABTEST          AudienceType = "abtest"           // A/B Test ID
	LIVEACTIVITYID  AudienceType = "live_activity_id" // 实时活动标识
)

func (a AudienceType) String() string {
	return string(a)
}

type Audience struct {
	Object   interface{}
	audience map[AudienceType]interface{}
}

func (a *Audience) Interface() interface{} {
	return a.Object
}

// All set all audiences
func (a *Audience) All() {
	a.Object = "all"
}

// SetID set audiences by id
func (a *Audience) SetID(ids []string) {
	a.set(REGISTRATION_ID, ids)
}

// SetTag set audiences by tag
func (a *Audience) SetTag(tags []string) {
	a.set(TAG, tags)
}

// SetTagAnd set audiences by tag_and
func (a *Audience) SetTagAnd(tags []string) {
	a.set(TAG_AND, tags)
}

// SetTagNot set audiences by tag_not
func (a *Audience) SetTagNot(tags []string) {
	a.set(TAG_NOT, tags)
}

// SetAlias set audiences by alias
func (a *Audience) SetAlias(aliases []string) {
	a.set(ALIAS, aliases)
}

// SetSegment set audiences by segment
func (a *Audience) SetSegment(segments []string) {
	a.set(SEGMENT, segments)
}

// SetABTest set audiences by abtest
func (a *Audience) SetABTest(abtests []string) {
	a.set(ABTEST, abtests)
}

// SetLiveActivityID set audiences by live_activity_id
func (a *Audience) SetLiveActivityID(liveActivityID string) {
	a.set(LIVEACTIVITYID, liveActivityID)
}

// set audiences
func (a *Audience) set(key AudienceType, v interface{}) {
	switch a.Object.(type) {
	case string:
		log.Printf("audience already set all")
		return // do nothing
	default:
	}

	if a.audience == nil {
		a.audience = make(map[AudienceType]interface{})
		a.Object = a.audience
	}

	a.audience[key] = v
}
