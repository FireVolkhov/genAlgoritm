package genome

import "math"
import "../../app/core"

type Gen struct {
	Name string
	EnterGens int
}

func GetRandomGen () *Gen {
	return &genome[core.RandomInt(0, len(genome) - 1)]
}

func Calc(name string, args []float64) float64 {
	switch name {
	case "ABS":
		return math.Abs(args[0])
	case "IF":
		ifResult := args[0]

		if (ifResult < 0) {
			return args[1]

		} else if (0 < ifResult) {
			return args[3]

		} else {
			return args[2]
		}

	case "ADD":
		return args[0] + args[1]
	case "SUB":
		return args[0] - args[1]
	case "MULT":
		return args[0] * args[1]
	case "DIV":
		return args[0] / args[1]

	case "MINUS_ONE":
		return float64(-1)
	case "NULL":
		return float64(0)
	case "ELER":
		return float64(0.5772)
	case "ONE":
		return float64(1)
	case "GOLD_MEMBER":
		return float64(1.6180)
	case "NEPER":
		return float64(2.718)
	case "PI":
		return float64(3.1415)

	default:
		panic("Unknow gen " + name)
	}
}

var genome []Gen

func init () {
	genome = []Gen{
		Gen{"ABS", 1},
		Gen{"IF", 4},

		Gen{"ADD", 2},
		Gen{"SUB", 2},
		Gen{"MULT", 2},
		Gen{"DIV", 2},

		Gen{"MINUS_ONE", 0},
		Gen{"NULL", 0},
		Gen{"ELER", 0},
		Gen{"ONE", 0},
		Gen{"GOLD_MEMBER", 0},
		Gen{"NEPER", 0},
		Gen{"PI", 0},
	}
}
