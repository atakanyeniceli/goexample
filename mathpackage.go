package mathpackage

//Sum add Function
func Sum(temps ...*int) (sum int) {
	for _, temp := range temps {
		sum += *temp
	}
	return
}

//Subtract Function
func Sub(temps ...*int) (sub int) {
	sub = *temps[0]
	for _, temp := range temps {
		sub -= *temp
	}
	sub += *temps[0]
	return
}

//Divide Function
func Div(temps ...*float64) (div float64) {
	div = *temps[0]
	for _, temp := range temps {
		div /= *temp
	}
	div *= *temps[0]
	return
}
func Mul(temps ...*int) (mul int) {
	mul = *temps[0]
	for _, temp := range temps {
		mul *= *temp
	}
	mul /= *temps[0]
	return
}
