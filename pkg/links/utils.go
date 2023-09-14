package links

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomShort(size uint64) string {
	str := make([]rune, size)

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}

func response(link *Link) ResponseBody {
	return ResponseBody{
		ID:    link.ID,
		Link:  link.Link,
		Short: link.Short,
	}
}
