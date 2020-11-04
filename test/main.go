package main

import (
	"fmt"
	"log"
)

type pacet struct {
	str      string
	ykazatel []*pacet
}

var mass []string

func checkForDublicate(str string) bool {
	for _, v := range mass {
		if v == str {
			return true
		}
	}
	return false
}

func obhod(l pacet) {

	if checkForDublicate(l.str) {
		log.Fatal("ЕСТЬ ЦИКЛ")
	}

	mass = append(mass, l.str)

	for _, v := range l.ykazatel {
		obhod(*v)
	}

}
func main() {
	var pacs [6]pacet
	for i := 0; i < 6; i++ {
		pacs[i] = pacet{str: string(i + 49)}
	}
	pacs[0].ykazatel = append(pacs[0].ykazatel, &pacs[1])
	pacs[0].ykazatel = append(pacs[0].ykazatel, &pacs[2])
	pacs[1].ykazatel = append(pacs[1].ykazatel, &pacs[3])
	pacs[1].ykazatel = append(pacs[1].ykazatel, &pacs[4])
	pacs[2].ykazatel = append(pacs[2].ykazatel, &pacs[5])

	obhod(pacs[0])

	fmt.Print("ЦИКЛЫ ОТСУТСТВУЮТ")
}
//git rm -r myfolder -f
//git commit -m "myfolder is deleted"
//git push