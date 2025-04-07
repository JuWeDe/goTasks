package main

import (
    "testing"
)

func TestReverseNumber(t *testing.T) {
    result := reverseNumber(843)
    if result != 348 {
        t.Errorf("Ожидалось 348, получено %d", result)
    }
}

func TestIsRightTriangle(t *testing.T) {
    result := isRightTriangle(6, 8, 10)
    if result != "Прямоугольный" {
        t.Errorf("Ожидалось 'Прямоугольный', получено '%s'", result)
    }
}

func TestTimeFromSeconds(t *testing.T) {
    hours, minutes := timeFromSeconds(13257)
    if hours != 3 || minutes != 40 {
        t.Errorf("Ожидалось 3 часа 40 минут, получено %d часов %d минут", hours, minutes)
    }
}	
func TestAddStarsBetweenLetters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Строка с разным регистром и символами",
			input:    "Go-1.2.3",
			expected: "G*o*-*1*.*2*.*3",
		},
		{
			name:     "Кириллица и спецсимволы",
			input:    "Привет!Как дела?",
			expected: "П*р*и*в*е*т*!*К*а*к* *д*е*л*а*?",
		},
		{
			name:     "Пустая строка",
			input:    "",
			expected: "",
		},
		{
			name:     "Один символ",
			input:    "A",
			expected: "A",
		},
		{
			name:     "Эмодзи",
			input:    "😀👍🚀",
			expected: "😀*👍*🚀",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addStarsBetweenLetters(tt.input)
			if result != tt.expected {
				t.Errorf("Ожидалось '%s', получено '%s'", tt.expected, result)
			}
		})
	}
}

func TestSquareDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Число с нулями",
			input:    1001,
			expected: 1001,
		},
		{
			name:     "Однозначное число",
			input:    7,
			expected: 49,
		},
		{
			name:     "Число с повторяющимися цифрами",
			input:    222,
			expected: 444,
		},
		{
			name:     "Большое число",
			input:    987654321,
			expected: 816449362516941,
		},
		{
			name:     "Ноль",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := squareDigits(tt.input)
			if result != tt.expected {
				t.Errorf("Для %d ожидалось %d, получено %d", tt.input, tt.expected, result)
			}
		})
	}
}