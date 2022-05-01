package pattern

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const NumberPattern = "%d"

var once sync.Once

type NumberOptions struct {
	Padding   int
	Increment int
	Start     int

	current int
}

func NewNumberOptions() *NumberOptions {
	return &NumberOptions{
		Padding:   0,
		Increment: 1,
		Start:     1,
	}
}

func (n *NumberOptions) Iterate() string {
	once.Do(func() {
		n.current = n.Start
	})

	current := int(math.Abs(float64(n.current)))
	n.current += n.Increment

	var sign string
	if current < 0 {
		sign = "-"
	} else {
		sign = ""
	}

	count := getDigitsInNumber(current)
	count = int(math.Max(0, float64(n.Padding-count)))

	zeroes := strings.Repeat("0", count)

	return sign + zeroes + fmt.Sprint(current)
}

func getDigitsInNumber(number int) int {
	count := 0
	for number > 0 {
		number = number / 10
		count++
	}

	// If number == 0 then count is 0, but it's not true
	if count == 0 {
		return 1
	}

	return count
}
