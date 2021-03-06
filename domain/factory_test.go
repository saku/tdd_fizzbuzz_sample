package domain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tddbc/go_gotest/domain"
)

type testFactory struct {
	suite.Suite
}

func TestFactory(t *testing.T) {
	suite.Run(t, &testFactory{})
}

func (s *testFactory) TestFactoryCreate() {
	type testCaseIn struct {
		num uint
	}

	type testCaseOut struct {
		result *domain.FizzBuzz
		err    error
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "invalid low number",
			in: testCaseIn{
				num: 101,
			},
			expected: testCaseOut{
				err: fmt.Errorf("invalid number"),
			},
		},
		{
			name: "normal lowest val",
			in: testCaseIn{
				num: 1,
			},
			expected: testCaseOut{
				result: domain.NewFizzBuzz(1),
			},
		},
		{
			name: "normal highest val",
			in: testCaseIn{
				num: 100,
			},
			expected: testCaseOut{
				result: domain.NewFizzBuzz(100),
			},
		},
		{
			name: "invalid high number",
			in: testCaseIn{
				num: 101,
			},
			expected: testCaseOut{
				err: fmt.Errorf("invalid number"),
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			factory := domain.NewFizzBuzzFactory()

			fizzbuzz, err := factory.Create(c.in.num)

			assert.Equal(t, c.expected.result, fizzbuzz)
			assert.Equal(t, c.expected.err, err)
		})
	}
}
