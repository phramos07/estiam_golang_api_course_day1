package main

import "fmt"

// interfaces cannot be instantiated
// interfaces cannot be instantiated -> which means NO ATTRIBUTESS!!!!!!!
// interfaces DESCRIBE METHODS THAT A STRUCT CAN HAVE
type MediaItem interface {
	Name() string
	Creator() string
	MediaType() string
}

type Book struct { // we don't have to tell the compiler that Book implements MediaItem
	Title  string
	Author string
	Pages  int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (b *Book) Name() string {
	return b.Title
}

func (b *Book) Creator() string {
	return b.Author
}

func (b *Book) MediaType() string {
	return "Book"
}

type Movie struct {
	Title    string
	Director string
	Duration int
}

func NewMovie(title string, director string, duration int) *Movie {
	return &Movie{
		Title:    title,
		Director: director,
		Duration: duration,
	}
}

func (m *Movie) Name() string {
	return m.Title
}

func (m *Movie) Creator() string {
	return m.Director
}

func (m *Movie) MediaType() string {
	return "Movie"
}

type MusicAlbum struct {
	Title   string
	Artist  string
	NTracks int
}

func NewMusicAlbum(title string, artist string, ntracks int) *MusicAlbum {
	return &MusicAlbum{
		Title:   title,
		Artist:  artist,
		NTracks: ntracks,
	}
}

func (ma *MusicAlbum) Name() string {
	return ma.Title
}

func (ma *MusicAlbum) Creator() string {
	return ma.Artist
}

func (ma *MusicAlbum) MediaType() string {
	return "Music Album"
}

/*
Exercise: Building a Media Library

Create a program that manages a media library containing different types of media items like books, movies, and music albums. Use interfaces to handle the common behaviors of these media items.

Define an interface MediaItem: Create an interface named MediaItem that declares methods for getting information about the item, such as its title and creator.

type MediaItem interface {
    Title() string
    Creator() string
    MediaType() string
}

Book, Movie, and MusicAlbum structs:

Book: title, author, and number of pages.
Movie: title, director, and duration.
MusicAlbum: title, artist, and number of tracks.

Methods for Book, Movie, and MusicAlbum:
For each media type, implement the methods declared in the MediaItem interface.
	function to display media library items:
	function that takes a slice of MediaItem interfaces and displays information about each item, including its title, creator, and media type.

Test on main:
	Create instances of Book, Movie, and MusicAlbum.
	Add them to a slice of MediaItem interfaces.
	Call the function to display the media library items.
*/

func main() {
	// create a slice of media items
	items := make([]MediaItem, 0)

	// add items to this slice
	book := NewBook("The Go Programming Language", "Donovan", 300)
	movie := NewMovie("Deep Impact", "James Cameron", 120)
	album := NewMusicAlbum("Thriller", "MJ", 9)

	objects := []MediaItem{book, movie, album}
	items = append(items, objects...)

	// print information on each
	for _, v := range items {
		fmt.Println(v)
	}
}
