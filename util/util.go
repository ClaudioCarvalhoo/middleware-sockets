package util

import "strings"
import "math"

func TrimString(text string) string {
	text = strings.TrimRight(text, "\n")
	text = strings.TrimRight(text, "\r\n")
	return text
}

func StdDev(numbers []float64, mean float64) float64 {
    total := 0.0
    for _, number := range numbers {
        total += math.Pow(number-mean, 2)
    }
    variance := total / float64(len(numbers)-1)
    return math.Sqrt(variance)
}