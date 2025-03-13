package mocks

import (
	"io"
	"strings"
)

type MockMeatTextService struct {
    meatTextReader io.ReadCloser
    meatAndFillerTextReader io.ReadCloser
}

func NewMockMeatTextService(meatText string, meatAndFillerText string) *MockMeatTextService {
    return &MockMeatTextService{
        meatTextReader: io.NopCloser(strings.NewReader(meatText)),
        meatAndFillerTextReader: io.NopCloser(strings.NewReader(meatAndFillerText)),
    }
}

func (m *MockMeatTextService) GetMeatText(hasFiller bool) (io.ReadCloser, error) {
    if hasFiller {
        return m.meatAndFillerTextReader, nil
    }
    return m.meatTextReader, nil
}
