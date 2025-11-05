package puzzles

type Day1 struct {
	Input []byte
}

func (d Day1) Solve(part int) {
	lines := d.Input
	switch part {
	case 1:
		d.part1(lines)
	}
}

func (d Day1) part1(lines []byte) {}
