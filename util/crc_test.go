package util

import (
	"fmt"
	"testing"
)

func Test_RedisClusterSlotOf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"the RedisClusterSlotOf test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 64; i++ {
				key := fmt.Sprintf("ad:idfv:idfa%d", i)
				slot := RedisClusterSlotOf(key)
				fmt.Printf("%s\t%d\n", key, slot)
			}
		})
	}
}
