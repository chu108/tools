package number

import (
	"math/rand"
	"time"
)

type Rand struct {
}

func NewRand() *Rand {
	return &Rand{}
}

//返回范围内随机数
func (*Rand) RandRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// String returns a random string ['a', 'z'] and ['0', '9'] in the specified length.
func (*Rand) String(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)
	letter := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
