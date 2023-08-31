package logger

import "testing"

func TestInfo(t *testing.T) {
	Info("测试信息")
}

func TestWarn(t *testing.T) {
	Warn("测试信息")
}

func TestError(t *testing.T) {
	Error("测试信息")
}

func BenchmarkInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("测试信息")
		}
	})
}
