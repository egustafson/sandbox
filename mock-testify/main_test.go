package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type exampleServiceMock struct {
	mock.Mock
}

// static check: struct exampleServiceMock implements ExampleService
var _ ExampleService = (*exampleServiceMock)(nil)

func (m *exampleServiceMock) Message(id int) string {
	// record the values passed in
	args := m.Called(id)
	// return the mock's value for args based on lookup
	return args.String(0)
}

// --  Test Case(s)  ---------------------------------------

func TestExampleService(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected string
	}{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
	}

	// create Mock and load input/output data
	testService := new(exampleServiceMock)
	for _, test := range tests {
		testService.On("Message", test.input).Return(test.expected)
	}

	// perform tests
	for _, test := range tests {
		r := testService.Message(test.input)
		assert.Equal(test.expected, r)
	}

}
