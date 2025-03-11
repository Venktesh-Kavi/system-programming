package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in := os.Stdin
	buf := []byte{}
	in.Read(buf)
	ss := strings.Split(string(buf), ",")

	mapToIceCreamKind(ss[1])
	ic := eatIceCream(IceCreamOption{ss[0], ss[1]})
	fmt.Println(ic)
}

type IceCreamKind int
type ToppingKind int

const (
	NORMAL IceCreamKind = iota
	GELLATO
	CREAMY
)

const (
	CHOC ToppingKind = iota
	SHERBATH
	CARAMEL
)

type IceCream interface{}

type CommonIceCream struct {
	IceCream
	cone     bool
	size     string
	toppings []ToppingKind
}

type GellatoIceCream struct {
	CommonIceCream
	gellatoPull int
}

type NormalIceCream struct {
	CommonIceCream
}

type CreamyIceCream struct {
	CommonIceCream
	creamWip bool
}

type IceCreamOption struct {
	flavour string
	variety string
}

type IceCreamFactory interface {
	createIceCream()
}
type GellatoIceCreamFactory struct{}

func (g GellatoIceCreamFactory) createIceCream(flavour string, size string) GellatoIceCream {
	if flavour == "vanilla" {
		return GellatoIceCream{
			CommonIceCream: CommonIceCream{
				cone:     false,
				size:     size,
				toppings: []ToppingKind{CHOC},
			},
			gellatoPull: 3,
		}
	} else {
		return GellatoIceCream{
			CommonIceCream: CommonIceCream{
				cone:     false,
				size:     size,
				toppings: []ToppingKind{SHERBATH},
			},
			gellatoPull: 1,
		}
	}
}

func mapToIceCreamKind(flavour string) IceCreamKind {
	if flavour == "NORMAL" {
		return NORMAL
	} else if flavour == "GELLATO" {
		return GELLATO
	} else if flavour == "CREAMY" {
		return CREAMY
	} else {
		panic("unknown ice cream kind")
	}
}

func eatIceCream(option IceCreamOption) IceCream {
	switch mapToIceCreamKind(option.variety) {
	case NORMAL:
		fmt.Println("Normal Icecream chosen")
	case GELLATO:
		g := GellatoIceCreamFactory{}
		return g.createIceCream("vanilla", "M")
	case CREAMY:
		fmt.Println("Creamy Icecream chosen")
	default:
		fmt.Println("Unknown variety")
		panic("nothing to eat")
	}
	return nil
}
