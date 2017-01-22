package longint

import (
	"strconv"
)

// LongInt - структура, представляющая собой длинное целое положительное
// число (больше обычного int). В данной библиотеке реализована часть арифметики
// для таких чисел, а именно: сложение (Sum), вычитание (Sub) и умножение (Mul).
// Несмотря на то, что число положительно, вычитание может возвращать отрицательный результат.
type LongInt struct{
	number []int
}

// NewLongInt - конструктор, принимающий в качестве параметра целое число -
// размер предполагаемого числа.
func NewLongInt(size int) *LongInt {
	num := new(LongInt)
	num.number= make([]int, size)
	return num
}

// NewLongIntFromString - конструктор, принимающий в качестве параметра строку,
// из которой будет получено число. Пока нет проверки на некорректные символы.
func NewLongIntFromString(str string) *LongInt {
	length := len(str)
	if length >= 30000 {
		println("Your number is too large. It must be less or equal 30000 digits.")
		return &LongInt{}
	}
	if length == 0 {
		return &LongInt{}
	} else {
		num := LongInt{}
		num.number = make([]int, length)
		for i := 0; i < length; i++ {
			// Переводим символы типа char в int.
			num.number[i] = int(str[i] - '0')
		}
		return &num
	}
}

// Sum - функция сложения, принимающая в качестве параметров два числа LongInt и возвращающая их сумму.
func Sum(val1, val2 *LongInt) *LongInt {
	// Создаем копии переданных параметров, чтобы защитить их от изменения.
	num1 := copyLongInt(val1)
	num2 := copyLongInt(val2)

	// Отнимаем от длин 1 для более удобной записи обращения по индексу.
	maxSize := max(len(num1.number), len (num2.number)) - 1
	minSize := min(len(num1.number), len(num2.number)) - 1

	// Создаем результат с запасом на один разряд (на случай кейса вроде "9+9=18").
	result := NewLongInt(maxSize + 2)
	var tmp, i, k int

	// Если первое число больше второго, меняем их местами, т.к. в дальнейшем обращении
	// по индексам считаем, что num1 >= num2.
	if more(num2, num1) {
		num1, num2 = num2, num1
	}

	// Складываем только столько разрядов, сколько их у наименьшего числа.
	// Можно было дополнить второе число нулями до первого, но так показалось
	// экономней.
	for i = maxSize; i >= (maxSize - minSize); i = i - 1 {
		tmp = num1.number[i] + num2.number[i - (maxSize - minSize)] + result.number[i+1]
		k = i + 1
		for (tmp >= 10) {
			tmp = num1.number[k-1] + num2.number[k-1 - (maxSize - minSize)]
			result.number[k] = result.number[k] + (tmp-10)
			k--
			tmp = result.number[k] + 1
		}
		result.number[k] = result.number[k] + tmp - result.number[k]
	}

	// Оставшиеся разряды длинного первого числа суммируем к результату.
	for i = 0; i < (maxSize - minSize); i++ {
		result.number[i+1] = result.number[i+1] + num1.number[i]
	}

	// Если они есть, подчищаем лишние нули в начале результата, оставляя при этом,
	// если так вышло, результат, равный нулю.
	for (result.number[0] == 0 && len(result.number) > 1) {
		//0+0
		result.number = result.number[1:]
	}

	return result
}

// min - функция возвращающая минимальное число из двух переданных int.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max - функция возвращающая максимальное число из двух переданных int.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sub - функция вычитания, принимающая в качестве параметров два числа LongInt
// (уменьшаемое и вычитаемое, соответственно) и возвращающая их разность.
// Если результат получился отрицательным, в начале числа будет 0.
func Sub(val1, val2 *LongInt) *LongInt {
	// Создаем копии переданных параметров, чтобы защитить их от изменения.
	num1 := copyLongInt(val1)
	num2 := copyLongInt(val2)
	// Знак.
	sign := true

	if more(num2, num1) {
		// Меняем знак на отрицательный.
		sign = false
		// Если первое число больше второго, меняем их местами, т.к. в дальнейшем обращении
		// по индексам считаем, что num1 >= num2.
		num1, num2 = num2, num1
	}

	maxSize := max(len(num1.number), len (num2.number)) - 1
	minSize := min(len(num1.number), len(num2.number)) - 1

	// Создаем результат с запасом на один разряд.
	result := NewLongInt(maxSize +1)
	var tmp, i, k int

	for i = maxSize; i >= (maxSize - minSize); i = i - 1 {
		tmp = num1.number[i] - num2.number[i - (maxSize - minSize)]
		// Прикидываем, нужно ли заниммать у следующего разряда для вычитания.
		if tmp < 0 {
			k = i - 1
			result.number[i] = 10 + tmp
			for !(num1.number[k] >= 1) {
				// Модифицируем первое число, занимая у следующего разряда.
				// Если у него не хвататет, идем к следующему и т.д.
				// Числа точно хватит, тк в начале убеждаемся, что num1 >= num2.
				num1.number[k] = 9
				k--
			}
			// Нашли, у кого занять.
			num1.number[k] = num1.number[k] - 1
		} else {
			result.number[i] = tmp
		}
	}
	// Оставшиеся разряды длинного первого числа переносим в результат.
	for i = 0; i < (maxSize - minSize); i++ {
		result.number[i] = num1.number[i]
	}

	// Если они есть, подчищаем лишние нули в начале результата, оставляя при этом,
	// если так вышло, результат, равный нулю.
	for (result.number[0] == 0 && len(result.number) > 1) {
		result.number = result.number[1:]
	}

	// Дописываем 0, если знак отрицательный.
	if sign == false {
		result.number = append(result.number, 0)
		copy(result.number[1:], result.number)
		result.number[0] = 0
	}

	return result
}

// copyLongInt - констуктор копирования, возвращающий копию параметра типа LongInt.
func copyLongInt(num *LongInt) *LongInt {
	result := NewLongInt(len(num.number))
	for i:=0; i < len(num.number); i++ {
		result.number[i] = num.number[i]
	}
	return result
}

// more - функция, возвращающая наибольшее из переданных LongInt.
func more(val1, val2 *LongInt) bool {
	num1len := len(val1.number)
	num2len := len(val2.number)
	if num1len > num2len {
		return true
	} else if num2len > num1len {
		return false
	} else {
		for i := 0; i < num1len; i++ {
			if val1.number[i] > val2.number[i] {
				return true
			} else if val2.number[i] > val1.number[i] {
				return false
			}
		}
		return false
	}
}

// Mul - функция, принимающая в качестве параметра два числа LongInt,
// возвращает их произведение.
func Mul(val1, val2 *LongInt) *LongInt{
	// Создаем копии переданных параметров, чтобы защитить их от изменения.
	num1 := copyLongInt(val1)
	num2 := copyLongInt(val2)

	maxSize := max(len(num1.number), len (num2.number)) - 1

	// Создаем результат с запасом на два разряда (на случай кейса вроде "99*99=9801").
	result := NewLongInt(maxSize +2)

	// Если первое числоа больше второго, меняем их местам, чтобы уменьшить
	// число сложений (когда считаем в столбик, то если число сверху длиннее числа
	// снизу, то слоэений промежуточных результатов будет меньше, чем наоборот).
	if more(num2, num1) == true {
		num2, num1 = num1, num2
	}

	for i:=0; i < len(num2.number); i++ {
		// Берем разряд второго числа
		mul := num2.number[len(num2.number) - i - 1]
		// и умножаем на него первое,
		tmp := mulOnNum(num1,mul)
		// затем сдвигаем результат влево на нужное число нулей.
		tmp = addNZero(*tmp, i)
		// Результат суммируем к уже полученному.
		result = Sum(result,tmp)
	}
	return result

}

// mulOnNum - вспомогательная функция для Mul, возвращающая результат умножения
// переданного LongInt на int.
func mulOnNum(val *LongInt, mul int) *LongInt {
	// Случай умножения на 0.
	if mul == 0 {
		return NewLongInt(1)
	}
	// Создаем копию переданного параметра, чтобы защитить его от изменения.
	num := copyLongInt(val)
	size := len(num.number)

	var i, tmp int

	// Создаем результат с запасом на один разряд (кейс "9*2=18").
	result := NewLongInt(size+1)

	for i = size - 1; i >= 0; i-- {
		// Умножаем разряд первого числа на наш множитель.
		tmp = num.number[i] * mul
		// Конвертируем полученное целое в строку.
		str := strconv.Itoa(tmp)
		// Создаем из строки LongInt.
		a := NewLongIntFromString(str)
		// Сдвигаем результат влево на нужное число нулей.
		a = addNZero(*a, size - 1 - i)
		// Результат суммируем к уже полученному.
		result = Sum(result,a)
	}
	return result
}

// addNZero - функция, добавляющая к переданному LongInt заданное число нулей.
// Нужна для сдвига промежуточных результатов по разрядам во время умножения.
func addNZero(val LongInt, count int) *LongInt{
	for count > 0 {
		val.number = append(val.number, 0)
		count--
	}
	return &val
}

// String - метод, возвращающий строковое представление переданного в качестве параметра LongInt.
// Для красивого вывода стуктур типа LongInt.
func (num *LongInt) String() string {
	result := ""
	for i := 0; i < len(num.number); i++ {
		result = result + strconv.Itoa(num.number[i])
	}

	// Заменяем "0" на "-", если число отрицательное.
	if result[0] == 48 {
		return "-" + result[1:]
	}
	return result
}
