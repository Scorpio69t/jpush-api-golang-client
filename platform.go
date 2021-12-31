package jpush

import "errors"

type PlatformType string

const (
	IOS      PlatformType = "ios"
	ANDROID  PlatformType = "android"
	WINPHONE PlatformType = "winphone"
)

func (p PlatformType) String() string {
	return string(p)
}

type Platform struct {
	Os      interface{}
	osArray []PlatformType
}

// All set all platforms
func (p *Platform) All() {
	p.Os = "all"
}

// Add add platform
func (p *Platform) Add(os PlatformType) error {
	if p.osArray == nil {
		p.osArray = make([]PlatformType, 0)
	}

	switch p.Os.(type) {
	case string:
		return errors.New("platform already set all")
	default:
	}

	// check if already added
	for _, v := range p.osArray {
		if v == os {
			return nil
		}
	}

	switch os {
	case IOS:
		fallthrough
	case ANDROID:
		fallthrough
	case WINPHONE:
		p.osArray = append(p.osArray, os)
		p.Os = p.osArray
	default:
		return errors.New("invalid platform")
	}

	return nil
}

// AddIOS add ios platform
func (p *Platform) AddIOS() {
	p.Add(IOS)
}

// AddAndroid add android platform
func (p *Platform) AddAndroid() {
	p.Add(ANDROID)
}

// AddWinphone add winphone platform
func (p *Platform) AddWinphone() {
	p.Add(WINPHONE)
}

// Remove remove platform
func (p *Platform) Remove(os PlatformType) error {
	if p.osArray == nil {
		return errors.New("platform not set")
	}

	for i, v := range p.osArray {
		if v == os {
			p.osArray = append(p.osArray[:i], p.osArray[i+1:]...)
			if len(p.osArray) == 0 {
				p.Os = nil
			} else {
				p.Os = p.osArray
			}
			return nil
		}
	}

	return errors.New("platform not found")
}
