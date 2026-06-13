package main

import "errors"

var (
	ErrNotNumber  = errors.New("Введённый символ не является целым числом")
	ErrNotFloat   = errors.New("Введённый символ не является дробным числом")
	ErrOutOfRange = errors.New("Число вне допустимого диапазона")
)
