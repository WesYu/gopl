package conv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func conv(input string, unitType string) (string, error) {
	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return "", err
	}
	switch unitType {
	case "temp":
		f := Fahrenheit(v)
		c := Celsius(v)
		k := Kelvin(v)
		return fmt.Sprintf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
			f, f.ToC(), f.ToK(), c, c.ToF(), c.ToK(), k, k.ToF(), k.ToC()), nil
	case "length":
		m := Meter(v)
		f := Foot(v)
		i := Inch(v)
		return fmt.Sprintf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
			m, m.ToF(), m.ToI(), f, f.ToM(), f.ToI(), i, i.ToM(), i.ToF()), nil
	case "weight":
		k := Kilogram(v)
		p := Pound(v)
		o := Ounce(v)
		return fmt.Sprintf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
			k, k.ToP(), k.ToO(), p, p.ToK(), p.ToO(), o, o.ToK(), o.ToP()), nil
	default:
		return "", fmt.Errorf("unsupported unit type: %q", unitType)
	}
}

func Conv() {
	switch {
	case len(os.Args) < 2:
		fmt.Println("Usage: ./main [temp|length|weight] ...val, or")
		fmt.Println("Usage: ./main [temp|length|weight] and then put in numbers")
		fmt.Println("	and press q to quit")

	case len(os.Args) == 2:
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			s := input.Text()
			if s == "q" {
				break
			}
			result, err := conv(s, os.Args[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Print(result)
		}

	case len(os.Args) > 2:
		for _, arg := range os.Args[2:] {
			result, err := conv(arg, os.Args[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Print(result)
		}
	}
}
