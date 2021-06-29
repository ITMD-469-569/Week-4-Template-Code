package main

import "fmt"

type Song struct {
	name   string
	artist string
	next   *Song
}

type Playlist struct {
	name             string
	head             *Song // first song in the playlist
	currentlyPlaying *Song
}

func (playlist *Playlist) addSong(name string, artist string) error {
	// fill in

}

func (playlist *Playlist) listAllSongs() error {
	currentSong := playlist.head // start at the first song in the playlist

	if currentSong == nil {
		fmt.Println("The playlist is empty! You need to add some songs.")
		return nil
	}

	fmt.Println(currentSong.name)
	for currentSong.next != nil { // loop through the songs until the end is reached (the next song is nil)
		currentSong = currentSong.next
		fmt.Println(currentSong.name) // print out the song
	}

	return nil

}

func main() {

	playlist := &Playlist{name: "MyCoolPlayList"}
	playlist.addSong("SongName", "Artist")
	playlist.addSong("SongName2", "Artist2")
	playlist.listAllSongs()
}
