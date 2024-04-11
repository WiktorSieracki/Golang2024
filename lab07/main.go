package main

import "fmt"

type Czlowiek struct {
	name string
}

type Student struct {
	Czlowiek
	rok int8
}

type Magister struct {
	Student
	temat string
}

type Doktor struct {
	Magister
	temat string
}

func (m Magister) mojTemat() {
	fmt.Println("Moj temat był taki ",m.temat)
}


func (c Czlowiek) Gadaj() {
	fmt.Println("A kuku, tutaj ", c.name)
}

func (d Doktor) mojTemat(){
	fmt.Println("Jako magister napisalem temat ",d.Magister.temat)
	fmt.Println("Jako doktor napisalem temat ",d.temat)
}

func immatrykulacja(n string) Student {
	return Student{
		Czlowiek: nowy_czlowiek(n),
		rok: 1,
	}
}


func nowy_czlowiek(n string) Czlowiek {
	return Czlowiek{name: n}
}

func main() {
	var cz Czlowiek = nowy_czlowiek("Zyndarm")
	cz.Gadaj()

	var st Student = immatrykulacja("Dawid")
	st.Gadaj()

	ma := Magister{
		Student: immatrykulacja("Mat"),
		temat: "jeszcze nie mam", 
	}
	ma.mojTemat()

	do := Doktor{
		Magister: ma,
		temat: "tez nie mam :(",
	}
	do.mojTemat()

	Wypisz(cz,st,ma,do)
}

type Ktos interface {
	Gadaj()
}
type Wyksztalciuch interface{
	mojTemat()
}

func Wypisz(ktos ...Ktos){
    for _,kto := range ktos {
        kto.Gadaj()
		if czyMozeSiePochwalicTytulem(kto){
			kto.(Wyksztalciuch).mojTemat()
		}
    }
}

func czyMozeSiePochwalicTytulem (kto Ktos) bool {
	_,ok := kto.(Magister)
		return ok
}