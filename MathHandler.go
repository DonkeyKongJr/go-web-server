package main

func AddTwoNumbers(firstNumber float64, secondNumber float64) float64 {
	return firstNumber + secondNumber;
}

func SubtractTwoNumbers(firstNumber float64, secondNumber float64) float64 {
	return firstNumber - secondNumber;
}

func MultiplyTwoNumbers(firstNumber float64, secondNumber float64) float64 {
	return firstNumber * secondNumber;
}

func SquareNumber(number float64) float64 {
	return number * number;
}

func DivideNumbers(number1 float64, number2 float64) float64 {
	if number1 == 0 {
		panic("Divide by 0 is not allowed")
	}
	return number1 / number2;
}

func ModuloNumber(number1 int64, number2 int64) int64 {
	return number1 % number2;
}

func FactorialNumber(number int64) (result int64) {
	if (number > 0) {
		result = number * FactorialNumber(number-1)
		return result
	}
	return 1
}

