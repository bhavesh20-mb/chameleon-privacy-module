package sanctions

import (
	"encoding/json"
	"os"
)

type SanctionsList struct {
	Blacklisted []string `json:"sanctioned"`
}

// LoadSanctionedAddresses reads a JSON file and parses it
func LoadSanctionedAddresses(path string) (*SanctionsList, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var list SanctionsList
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	return &list, nil
}

// IsSanctioned returns true if address is in the blacklist
func (s *SanctionsList) IsSanctioned(address string) bool {
	for _, a := range s.Blacklisted {
		if a == address {
			return true
		}
	}
	return false
}
