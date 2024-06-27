package puppy

import (
	dog "github.com/bruckZachew/dummer"
)

func Bark() string {
	return dog.Bark("woof!")
}

func Barks() string {
	return dog.Bark("woof! woof! woof!")
}
