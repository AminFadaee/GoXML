package main

import (
	"kirby.lensreader.com/processors"
)

func main() {
	processors.GistEpub("books/foo.epub", "books/foo-gist.epub")
}
