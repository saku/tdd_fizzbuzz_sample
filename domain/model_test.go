package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tddbc/go_gotest/domain"
)

type testFizzBuzz struct {
	suite.Suite
}

func TestFizzBuzz(t *testing.T) {
	suite.Run(t, &testFizzBuzz{})
}

func (s *testFizzBuzz) TestModel() {
	type testCaseIn struct {
		num uint
	}

	type testCaseOut struct {
		result string
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "normal val",
			in: testCaseIn{
				num: 1,
			},
			expected: testCaseOut{
				result: "1",
			},
		},
		{
			name: "fizz",
			in: testCaseIn{
				num: 3,
			},
			expected: testCaseOut{
				result: "fizz",
			},
		},
		{
			name: "buzz",
			in: testCaseIn{
				num: 5,
			},
			expected: testCaseOut{
				result: "buzz",
			},
		},
		{
			name: "fizzbuzz",
			in: testCaseIn{
				num: 15,
			},
			expected: testCaseOut{
				result: "fizzbuzz",
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			fizzbuzz := domain.NewFizzBuzz(c.in.num)

			assert.Equal(t, c.expected.result, fizzbuzz.Fizzbuzz())
		})
	}
}
