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