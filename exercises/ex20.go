package main

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

}
