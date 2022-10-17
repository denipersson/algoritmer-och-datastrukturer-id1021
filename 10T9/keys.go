package main

func get_code_from_char(char rune) int {

	code := -1

	switch char {
	case 'a':
		code = 0
	case 'b':
		code = 1
	case 'c':
		code = 2
	case 'd':
		code = 3
	case 'e':
		code = 4
	case 'f':
		code = 5
	case 'g':
		code = 6
	case 'h':
		code = 7
	case 'i':
		code = 8
	case 'j':
		code = 9
	case 'k':
		code = 10
	case 'l':
		code = 11
	case 'm':
		code = 12
	case 'n':
		code = 13
	case 'o':
		code = 14
	case 'p':
		code = 15
	case 'r':
		code = 16
	case 's':
		code = 17
	case 't':
		code = 18
	case 'u':
		code = 19
	case 'v':
		code = 20
	case 'x':
		code = 21
	case 'y':
		code = 22
	case 'z':
		code = 23
	case 'å':
		code = 24
	case 'ä':
		code = 25
	case 'Ã':
		code = 25
	case 'ö':
		code = 26
	}
	return code
}
func get_char_from_code(code int) rune {
	var char rune
	switch code {
	case 0:
		char = 'a'
	case 1:
		char = 'b'
	case 2:
		char = 'c'
	case 3:
		char = 'd'
	case 4:
		char = 'e'
	case 5:
		char = 'f'
	case 6:
		char = 'g'
	case 7:
		char = 'h'
	case 8:
		char = 'i'
	case 9:
		char = 'j'
	case 10:
		char = 'k'
	case 11:
		char = 'l'
	case 12:
		char = 'm'
	case 13:
		char = 'n'
	case 14:
		char = 'o'
	case 15:
		char = 'p'
	case 16:
		char = 'r'
	case 17:
		char = 's'
	case 18:
		char = 't'
	case 19:
		char = 'u'
	case 20:
		char = 'v'
	case 21:
		char = 'x'
	case 22:
		char = 'y'
	case 23:
		char = 'z'
	case 24:
		char = 'å'
	case 25:
		char = 'ä'
	case 26:
		char = 'ö'
	}
	return char
}
func index_from_key(key rune) int {
	var index int

	switch key {
	case '1':
		index = 0
	case '2':
		index = 3
	case '3':
		index = 6
	case '4':
		index = 9
	case '5':
		index = 12
	case '6':
		index = 15
	case '7':
		index = 18
	case '8':
		index = 21
	case '9':
		index = 24
	}

	return index
}
func key_from_char(char rune) string {
	code := get_code_from_char(char)
	key := "-1"
	if code < 3 {
		key = "1"
	} else if code < 6 {
		key = "2"
	} else if code < 9 {
		key = "3"
	} else if code < 12 {
		key = "4"
	} else if code < 15 {
		key = "5"
	} else if code < 18 {
		key = "6"
	} else if code < 21 {
		key = "7"
	} else if code < 24 {
		key = "8"
	} else {
		key = "9"
	}
	return key
}
