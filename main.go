package main

import (
	"flag"
	"knights/battle_field"
	"log"
)

func main() {
	backward := flag.Bool("backward", false, "use backward policy when select next")
	random := flag.Bool("random", false, "use random policy when select next")
	knights := flag.Int("knights", 5, "number of knights")
	witches := flag.Int("witches", 5, "number of witches")

	flag.Parse()

	b := battle_field.NewBuilder()

	if *backward {
		b.WithBackwardPolicy()
	}
	if *random {
		b.WithRandomPolicy()
	}
	b.WithKnights(*knights)
	b.WithWitches(*witches)

	battle, err := b.Build()
	if err != nil {
		log.Fatal(err)
	}
	battle.Fight()
}
