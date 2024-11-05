package utils
import "math"

// MagicSum takes a single integer n and returns n + n.
func MagicSum(n int) int {
    return n + n
}

// MagicPow takes a single integer n and returns n raised to the power of n.
func MagicPow(n int) int {
    return int(math.Pow(float64(n), float64(n)))
}

// Magicodd takes a single integer n and returns true if n is odd, false otherwise.
func MagicOdd(n int) bool {
    return n%2 != 0
}

// MagicGrade takes a single integer n and returns a grade as a string based on its value.
func MagicGrade(n int) string {
    switch n {
    case 0:
        return "Zero"
    case 1:
        return "Bad"
    case 2:
        return "Ok"
    case 3:
        return "Nice"
    case 4:
        return "Awesome"
    case 5:
        return "Exceptional"
    default:
        return "Unknown"
    }
}

// MagicName takes a single integer n and returns a slice of strings with the name "Mila" repeated n times.
func MagicName(n int) []string {
    names := make([]string, n)
    for i := 0; i < n; i++ {
        names[i] = "Joyce"
    }
    return names
}

// MagicTria takes a single integer n and returns the sum of all numbers from 1 to n.
func MagicTria(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// MagicChange takes a pointer to an integer n and changes its value to n * 2.
func MagicChange(n *int) {
	*n *= 3
}

// MagicNumber struct with a single field Number.
type MagicNumber struct {
	Number int
}

// Multiply method that multiplies Number by n and updates it.
func (m *MagicNumber) Multiply(n int) {
	m.Number *= n
}