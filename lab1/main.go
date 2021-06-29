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
	// create song struct instance
	newSong := &Song{name: name, artist: artist}
	// the next field of this new song, will be populated when a new song gets added

	// need to find out what song is the last in the playlist, so we can put our new song as its next field
	// check if the playlist (linked list) is empty
	if playlist.head == nil {
		// its empty, add the song as the new head of the list
		playlist.head = newSong
	} else {
		// its not empty, traverse the playlist until the next song doesnt exist
		// this means we have reached the end, and we can put our new song there
		currentSong := playlist.head  //start at first song
		for currentSong.next != nil { // use the next field of the song struct, which is a pointer to the next song to traverse the playlist
			currentSong = currentSong.next
		}

		// when code execution reaches here, the next song in the playlist is nil, meaning we are at the end
		// set the next song to be our new song
		currentSong.next = newSong
	}

	return nil

}

func (playlist *Playlist) listAllSongs() error {
	currentSong := playlist.head

	if currentSong == nil {
		fmt.Println("The playlist is empty! You need to add some songs.")
		return nil
	}

	fmt.Println(currentSong.name)
	for currentSong.next != nil {
		currentSong = currentSong.next
		fmt.Println(currentSong.name)
	}

	return nil

}

func main() {

	// their lab is to create the listAllSongs and addSongs function
	playlist := &Playlist{name: "MyCoolPlayList"}
	playlist.addSong("SongName", "Artist")
	playlist.addSong("SongName2", "Artist2")
	playlist.listAllSongs()

}
