package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter flavour and variety")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	read := sc.Text()
	ss := strings.Split(read, ",")
	fmt.Println("### Preparing IceCream ###")
	ic, err := eatIceCream(IceCreamOption{ss[0], ss[1]})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ic.Stringify())
}

type IceCreamKind int
type ToppingKind int

const (
	NORMAL IceCreamKind = iota
	GELLATO
	CREAMY
	UNKNOWN
)

const (
	CHOC ToppingKind = iota
	SHERBATH
	CARAMEL
)

type IceCream interface {
	Stringify() string
} // created to just allow replacing CommonIceCream with a generic IceCream type.

func (c CommonIceCream) Stringify() string {
	return fmt.Sprintf("IceCream Prepared, toppings: %v, you are getting size: %s", c.toppings, c.size)
}

// Common set of parameters for all icecreams. Note one cannot replace CommonIceCream in place of Gellato, since they are composed go does not allow this.
type CommonIceCream struct {
	IceCream
	cone     bool
	size     string
	toppings []ToppingKind
}

type GellatoIceCream struct {
	CommonIceCream // embeds common icecream (composition), has access to all fields and methods on the embedded type.
	gellatoPull    int
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

func mapToIceCreamKind(flavour string) (IceCreamKind, error) {
	log.Println("received kind: ", flavour)
	switch flavour {
	case "NORMAL":
		return NORMAL, nil
	case "GELLATO":
		return GELLATO, nil
	case "CREAMY":
		return CREAMY, nil
	}
	return UNKNOWN, fmt.Errorf("unknown variety: %s", flavour)
}

func eatIceCream(option IceCreamOption) (IceCream, error) {
	switch kind, _ := mapToIceCreamKind(option.flavour); kind {
	case NORMAL:
		fmt.Println("Normal Icecream chosen")
	case GELLATO:
		g := GellatoIceCreamFactory{}
		return g.createIceCream("vanilla", "M"), nil
	case CREAMY:
		fmt.Println("Creamy Icecream chosen")
	}
	return nil, fmt.Errorf("unknown icecream kind: %s", option)
}
