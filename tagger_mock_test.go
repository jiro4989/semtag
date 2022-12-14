package main

type MockTagger struct {
	Tagger

	// mock data
	tags []string
	err  error
}

func (m *MockTagger) Tags() ([]string, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.tags, nil
}
