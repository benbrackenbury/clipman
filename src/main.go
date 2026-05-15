package main

import (
	"github.com/benbrackenbury/clipman/src/store"
	"github.com/benbrackenbury/clipman/src/transmit"
)

func main() {
	store := store.NewSQLiteStore("clipboard.db")
	defer store.Close()
	transmit.Transmit(store)
}
