package main

import "errors"

var (
	ErrNotNumber  = errors.New("Введённый символ не является целым числом")
	ErrOutOfRange = errors.New("Число вне допустимого диапазона")
)
