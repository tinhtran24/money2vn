package utils

// Int2Vn return a Vietnamese string saying value of the parameter (int64).
// The first character is in upper case
func Int2Vn(number int64) string {

	vn := int2VnStr(number)
	r := []byte(vn)[0] - 32

	if r == 163 {
		// Special case: character 'â' is 2 bytes unicode. Get []byte(vn)[0] is incorrect
		return "Â" + vn[2:]
	}

	return string(r) + vn[1:]
}

// Internal function, used in both converting integer and float number
func int2VnStr(number int64) string {
	if number == 0 {
		return "không"
	}

	thousand := [...]string{"", " nghìn ", " triệu ", " tỷ ", " nghìn ", " triệu ", " tỷ "}

	digit := [...]string{"không", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín", "mười"}

	tail := [...]string{" mươi", " mốt", " hai", " ba", " tư", " lăm", " sáu", " bảy", " tám", " chín"}

	taild1 := [...]string{"mười", "mười một", "mười bốn"}
	var n, k, d1, d2, d3, dd int
	var vn, neg, s string

	if number < 0 {
		neg = "âm "
		number = -number
	}

	for number > 0 {
		n = int(number % 1000)

		if n > 0 {
			d3 = int(n / 100)

			if (number > 1000) || (d3 > 0) {
				s = digit[d3] + " trăm "
			} else {
				s = ""
			}

			dd = n % 100
			if dd > 0 {
				d1 = dd % 10
				d2 = int(dd / 10)
				switch d2 {
				case 0:
					if s != "" {
						s += "lẻ " + digit[d1]
					} else {
						s = digit[d1]
					}
				case 1:
					switch d1 {
					case 0:
						s += taild1[d1]
					case 1:
						s += taild1[d1]
					case 4:
						s += taild1[2]
					default:
						s += taild1[0] + tail[d1]
					}
				default:
					s += digit[d2] + tail[0] + tail[d1]
				}
			}
			vn = s + thousand[k] + vn
		}
		k++
		number = int64(number / 1000)
	}

	return neg + vn
}
