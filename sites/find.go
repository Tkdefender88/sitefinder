package sites

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

var (
	sites map[string]string
)

type SiteFinder struct {
	sites map[string]string
}

func New() *SiteFinder {
	return &SiteFinder{
		sites: make(map[string]string),
	}
}

func (s *SiteFinder) Load() error {
	b, err := ioutil.ReadFile("sites.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &s.sites); err != nil {
		return err
	}

	return nil
}

func (s *SiteFinder) Find(skey string) (string, error) {
	v, ok := s.sites[strings.ToLower(skey)];
	if !ok {
		return "", errors.New("Site not found")
	}
	return v, nil
}

func (s *SiteFinder) Add(name, value string) {
	s.sites[strings.ToLower(name)] = value
}

func (s *SiteFinder) Save() error {
	b, err := json.Marshal(s.sites)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("sites.json", b, 0644); err != nil {
		return err
	}
	return nil
}
