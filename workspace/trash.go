package main

import "fmt"

type myInterface interface {
	PlayMusic(string)
	StopMusic()
}

type Player string

func (p Player) PlayMusic(sng string) {
	fmt.Printf("it's time to play that very good song - %s\n", sng)
}

func (p Player) StopMusic() {
	fmt.Println("stop music")
}

type Recorder string

func (r Recorder) PlayMusic(sng string) {
	fmt.Printf("it's time to play that very good song - %s\n", sng)
}

func (r Recorder) StopMusic() {
	fmt.Println("stop music")
}

func Sonic(r myInterface, songs []string) {
	for _, song := range songs {
		r.PlayMusic(song)
	}
	r.StopMusic()
}

func main() {
	player := Player("")
	var experience myInterface = player
	listOfSongs := []string{"mother", "i", "love"}
	Sonic(experience, listOfSongs)
	experience = Recorder("")
	Sonic(experience, listOfSongs)
}
