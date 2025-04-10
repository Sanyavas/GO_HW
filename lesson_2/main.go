package main

import "fmt"

func FibonacciIterative(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація без використання рекурсії
	if n <= 1 {
		return n
	}
	prev, next := 0, 1

	for i := 2; i <= n; i++ {
		prev, next = next, prev+next
	}
	return next
}

func FibonacciRecursive(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація з використанням рекурсії
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	// Функція повертає `true` якщо число `n` - просте.
	// Інакше функція повертає `false`
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsBinaryPalindrome(n int) bool {
	// Функція повертає `true` якщо число `n` у бінарному вигляді є паліндромом
	// Інакше функція повертає `false`
	//
	// Приклади:
	// Число 7 (111) - паліндром, повертаємо `true`
	// Число 5 (101) - паліндром, повертаємо `true`
	// Число 6 (110) - не є паліндромом, повертаємо `false`
	binary_dig := fmt.Sprintf("%b", n)
	fmt.Printf("Binary of %d is %s\n", n, binary_dig)

	for i := 0; i < len(binary_dig)/2; i++ {
		if binary_dig[i] != binary_dig[len(binary_dig)-1-i] {
			return false
		}
	}
	return true
}

func ValidParentheses(s string) bool {
	// Функція повертає `true` якщо у вхідній стрічці дотримані усі правила високристання дужок
	// Правила:
	// 1. Допустимі дужки `(`, `[`, `{`, `)`, `]`, `}`
	// 2. У кожної відкритої дужки є відповідна закриваюча дужка того ж типу
	// 3. Закриваючі дужки стоять у правильному порядку
	//    "[{}]" - правильно
	//    "[{]}" - не правильно
	// 4. Кожна закриваюча дужка має відповідну відкриваючу дужку

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch) // додаємо відкриту дужку
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false // або стек пустий, або верхня дужка не відповідає
			}
			stack = stack[:len(stack)-1] // видаляємо відкриту дужку
		}
	}

	return len(stack) == 0 // якщо стек порожній — усі дужки закриті
}

func Increment(num string) int {
	// Функція на вхід отримує стрічку яка складається лише з символів `0` та `1`
	// Тобто стрічка містить певне число у бінарному вигляді
	// Потрібно повернути число на один більше
	return 0
}

func main() {
	fmt.Println(FibonacciIterative(9))
	fmt.Println(FibonacciRecursive(8))
	fmt.Println(IsPrime(7))
	fmt.Println(IsBinaryPalindrome(334))
	fmt.Println(ValidParentheses("{({}[][[)]])}"))
}
