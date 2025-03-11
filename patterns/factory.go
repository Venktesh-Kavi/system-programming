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
	eatIceCream(IceCreamOption{ss[0], ss[1]})
}

type IceCreamFactory interface {
	createIceCream()
}

type IceCreamOption struct {
	flavour string
	variety string
}

type IceCreamKind int

const (
	NORMAL IceCreamKind = iota
	GELLATO
	CREAMY
)

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

func eatIceCream(option IceCreamOption) {
	switch option.variety {
	case NORMAL:
		fmt.Println("Normal Icecream chosen")
	case GELLATO:
		fmt.Println("Gelator Icecream chosen")
	case CREAMY:
		fmt.Println("Creamy Icecream chosen")
	default:
		fmt.Println("Unknown variety")
	}
}
