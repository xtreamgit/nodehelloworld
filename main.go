package main


import "fmt"


type Dog struct {
	Breed	string
	Weight	int
	Sound	string
}


func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound,)
	fmt.Println(d.Sound)
}

func main() {
	chihuahua := Dog{"chihuahua", 8, "Woof"}
	//terrier := Dog{"terrier", 21, "Bark"}
	//greyhound := Dog{"", 34, "Baowao"}

	fmt.Println(chihuahua)
	chihuahua.Speak()

	chihuahua.Sound = "Arf"
	chihuahua.Speak()

	chihuahua.SpeakThreeTimes()
	chihuahua.SpeakThreeTimes()
	chihuahua.SpeakThreeTimes()
}