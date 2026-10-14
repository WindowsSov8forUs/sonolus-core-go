package core

import "encoding/json"

type GameplayResult struct {
	Grade         Grade   `json:"grade"`
	ArcadeScore   float64 `json:"arcadeScore"`
	AccuracyScore float64 `json:"accuracyScore"`
	Combo         int     `json:"combo"`
	Perfect       int     `json:"perfect"`
	Great         int     `json:"great"`
	Good          int     `json:"good"`
	Miss          int     `json:"miss"`
	TotalCount    int     `json:"totalCount"`
}

type Sil struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Srl struct {
	Hash *Nilable[string] `json:"hash,omitempty"`
	URL  *Nilable[string] `json:"url,omitempty"`
}

// UnmarshalJSON keeps an explicit null (`{"hash":null}`) distinct from an
// absent key: encoding/json would otherwise set the pointer to nil for null.
func (s *Srl) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*s = Srl{}
	for key, target := range map[string]**Nilable[string]{"hash": &s.Hash, "url": &s.URL} {
		value, ok := raw[key]
		if !ok {
			continue
		}
		var field Nilable[string]
		if err := json.Unmarshal(value, &field); err != nil {
			return err
		}
		*target = &field
	}
	return nil
}

type Tag struct {
	Title Text `json:"title,omitempty"`
	Icon  Icon `json:"icon,omitempty"`
}
