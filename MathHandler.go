package main

func AddTwoNumbers(firstNumber float32, secondNumber float32) float32{
	return firstNumber + secondNumber;
}

func SubtractTwoNumbers(firstNumber float32, secondNumber float32) float32{
	return firstNumber - secondNumber;
}

func SquareNumber(number float32) float32{
	return number * number;
}

func DivideNumbers(number1 float32, number2 float32) float32{
	if number1 == 0 {
		panic("Divide by 0 is not allowed")
	}
	return number1 / number2;
}
