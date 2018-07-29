package specs

import (
	"fmt"
)

// MaterialVoxel TODO
type MaterialVoxel struct {
	_ int // mock to fix memory alignment issue.
}

// Validate ...
func (m MaterialVoxel) Validate() error {
	return fmt.Errorf("not implemented")
}
