package nonrepeating

import "testing"

func BenchmarkLengthOfNonRepeatingSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer() // 重新设置计算时钟

	for i := 0; i < b.N; i++ {
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Leng..Str(%s)结果为%d, 期望为:%d", "...", actual, ans)
		}
	}
}

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s string
		r int
	}{
		{"abcabc", 3},
		{"abbbb", 2},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		actual := LengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.r {
			t.Errorf("Leng..Str(%s)结果为%d, 期望为:%d", tt.s, actual, tt.r)
		}
	}
}
