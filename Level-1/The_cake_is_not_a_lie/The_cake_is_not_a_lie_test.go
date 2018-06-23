package The_cake_is_not_a_lie

import "testing"

func Test_The_cake_is_not_a_lie(t *testing.T) {

	testDatas := []struct {
		x      string
		output int
	}{

		{"aa", 2},
		//{"aaaaaaaaaaaa", 12},

		//{"abccbaabccba", 2},
		//{"abcabcabcabc", 4},
		//{"abccbaabccbaabccbaabccba", 4},
		//{"abcabcabcabcabcabcabcabc", 8},
		//{"abababababababababababab", 12},
		//{"awwa", 1},
		//{"awwaawwa", 2},
		//{"awwaawwaawwa", 3},
		//{"awwaawwaawwaawwa", 4},
		//{"a", 1},
		//{"awwawaawwawa", 2},
	}

	for _,testData := range testDatas {

		result := answer1(testData.x)

		if result != testData.output {
			t.Errorf("Actual: output:%d not equal Expected: output:%d", result, testData.output)
		}
	}
}