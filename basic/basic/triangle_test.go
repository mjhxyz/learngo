package main

import "testing"

func BenchmarkTriangle(b *testing.B) {
	s := struct {
		a, b, c int
	}{30000, 40000, 50000}

	for i := 0; i < b.N; i++ {
		actual := calcTriangle(s.a, s.b)
		if actual != s.c {
			b.Errorf("calcTriangle(%d,%d)返回为%d, 期望为: %d", s.a, s.b, actual, s.c)
		}
	}
}

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c { // 测试失败
			t.Errorf("calcTriangle(%d, %d); 结果为: %d, 期望为: %d,",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
