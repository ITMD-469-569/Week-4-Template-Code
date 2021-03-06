package main

import (
	"fmt"
)

type Song struct {
	name   string
	artist string
	next   *Song // reference to the next song
}

type Playlist struct {
	name             string
	head             *Song // first song in the playlist
	currentlyPlaying *Song
}

func (playlist *Playlist) addSong(name string, artist string) error {
	// implement using addSong method from lab 1

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

func (playlist *Playlist) playNewRecommendedSong(artists map[string][]*Song) error {
	// loops through the playlist, starting from the first song
	// after that song plays (Just print it to the console), play a new song from the same artist from the artist map
	// once the artist's recommended songs are done playing, go onto the next song in the playlist and do the same

}

func (playlist *Playlist) recommendSong(artistName string, artists map[string][]*Song) {

	// given the artists name and the global map of songs
	// play a recommended song by the same artist
	// songs are recommended if they are by the same artist

}

func main() {
	_, _, _, _, artists := createArtistMap()

	playlist := &Playlist{name: "MyCoolPlayList"}
	playlist.addSong("Blowin' In the Wind", "Bob Dylan")
	playlist.addSong("You Can't Always Get What You Want", "Rolling Stones")
	playlist.addSong("A Day in the Life", "The Beatles")
	playlist.addSong("Heart-Shaped Box", "Nirvana")

	playlist.playNewRecommendedSong(artists)

}

func createArtistMap() ([]*Song, []*Song, []*Song, []*Song, map[string][]*Song) {
	artists := make(map[string][]*Song)

	bobDylanSongs := make([]*Song, 0)
	song0 := &Song{name: "Like a Rolling Stone", artist: "Bob Dylan"}
	song1 := &Song{name: "Gimme Shelter", artist: "Bob Dylan"}
	song2 := &Song{name: "Hurricane", artist: "Bob Dylan"}
	bobDylanSongs = append(bobDylanSongs, song0, song1, song2)

	artists["Bob Dylan"] = bobDylanSongs

	rollingStonesSongs := make([]*Song, 0)
	song3 := &Song{name: "Paint it Black", artist: "Rolling Stones"}
	song4 := &Song{name: "Gimme Shelter", artist: "Rolling Stones"}
	song5 := &Song{name: "Angie", artist: "Rolling Stones"}
	rollingStonesSongs = append(rollingStonesSongs, song3, song4, song5)

	artists["Rolling Stones"] = rollingStonesSongs

	beatlesSongs := make([]*Song, 0)
	song6 := &Song{name: "Hey Jude", artist: "The Beatles"}
	song7 := &Song{name: "Come Together", artist: "The Beatles"}
	song8 := &Song{name: "Here Comes the Sun", artist: "The Beatles"}
	beatlesSongs = append(beatlesSongs, song6, song7, song8)

	artists["The Beatles"] = beatlesSongs

	nirvanaSongs := make([]*Song, 0)
	song9 := &Song{name: "Smells Like Teen Spirit", artist: "Nirvana"}
	song10 := &Song{name: "Come as You Are", artist: "Nirvana"}
	song11 := &Song{name: "About a Girl", artist: "Nirvana"}
	nirvanaSongs = append(nirvanaSongs, song9, song10, song11)

	artists["Nirvana"] = nirvanaSongs

	return bobDylanSongs, rollingStonesSongs, beatlesSongs, nirvanaSongs, artists
}
