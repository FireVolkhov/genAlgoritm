package genome

import "math"
import "../../app/core"

type Gen struct {
	Name string
	EnterGens int
	Key byte
}

func GetRandomGen () *Gen {
	return &*genome[core.RandomInt(0, len(genome) - 1)]
}

func GetGenByName (name string) *Gen {
	gen, ok := genomeMap[name]

	if (ok) {
		return &*gen
	} else {
		panic("Gen `" + name + "` not found")
	}
}

func Calc(name string, args []float64) float64 {
	switch name {
	case "ABS":
		return math.Abs(args[0])
	case "MOD":
		return math.Mod(args[0], args[1])
	case "MIN":
		return math.Min(args[0], args[1])
	case "MAX":
		return math.Max(args[0], args[1])

	case "COS":
		return math.Cos(args[0])
	case "SIN":
		return math.Sin(args[0])
	case "TAN":
		return math.Tan(args[0])
	case "ACOS":
		return math.Acos(args[0])
	case "ASIN":
		return math.Asin(args[0])
	case "ATAN":
		return math.Atan(args[0])

	case "POW":
		return math.Pow(args[0], args[1])
	case "LOG":
		return math.Log(args[0])

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

var genome []*Gen
var genomeMap map[string]*Gen

func init () {
	genome = []*Gen{
		&Gen{"ABS", 1, byte(0)},
		&Gen{"MOD", 2, byte(1)},
		&Gen{"MIN", 2, byte(2)},
		&Gen{"MAX", 2, byte(3)},

		&Gen{"COS", 1, byte(10)},
		&Gen{"SIN", 1, byte(11)},
		&Gen{"TAN", 1, byte(12)},
		&Gen{"ACOS", 1, byte(13)},
		&Gen{"ASIN", 1, byte(14)},
		&Gen{"ATAN", 1, byte(15)},

		&Gen{"POW", 2, byte(20)},
		&Gen{"LOG", 1, byte(21)},

		&Gen{"IF", 4, byte(30)},

		&Gen{"ADD", 2, byte(40)},
		&Gen{"SUB", 2, byte(41)},
		&Gen{"MULT", 2, byte(42)},
		&Gen{"DIV", 2, byte(43)},

		&Gen{"MINUS_ONE", 0, byte(50)},
		&Gen{"NULL", 0, byte(51)},
		&Gen{"ELER", 0, byte(52)},
		&Gen{"ONE", 0, byte(53)},
		&Gen{"GOLD_MEMBER", 0, byte(54)},
		&Gen{"NEPER", 0, byte(55)},
		&Gen{"PI", 0, byte(56)},
	}

	genomeMap = make(map[string]*Gen)

	for _, gen := range genome {
		genomeMap[gen.Name] = gen
	}
}
