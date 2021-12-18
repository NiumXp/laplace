package internal

import "fmt"

type Position [2]uint

func (p *Position) String() string {
	return fmt.Sprintf("(%d, %d)", p[0], p[1])
}
