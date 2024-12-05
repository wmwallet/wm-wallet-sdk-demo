package untils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUUID() string {
	now := time.Now()
	timePart := now.Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	randomPart := fmt.Sprintf("%06d", rand.Intn(1000000)) // 保证是6位数字
	return timePart + randomPart
}
