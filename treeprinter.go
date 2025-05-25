/*
Complete port of https://github.com/billvanyo/tree_printer/blob/master/src/main/java/tech/vanyo/treePrinter/TreePrinter.java
to Go. This is not my original implementation.
*/

package tree

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Printer renders a horizontal ASCII representation of a binary tree.
type Printer[T cmp.Ordered] struct {
	// Label converts a node value to the text that will be printed inside the
	// node. If Label is nil the default is fmt.Sprint(node.Value()).
	Label func(T) string

	// Out is the destination for rendered output. Defaults to os.Stdout.
	Out io.Writer

	// SquareBranches prints branches using Unicode box‑drawing characters
	// instead of classic / and \.
	SquareBranches bool

	// LRAgnostic prints a straight vertical branch (│) for single‑child nodes
	// instead of elbows that reveal whether the child sits on the left or the
	// right.
	LRAgnostic bool

	// HSpace is the minimum number of spaces between adjacent node labels in a
	// single tree. Must be positive. Default is 2.
	HSpace int

	// TSpace is the horizontal distance between columns of trees when printing
	// several trees next to each other. It is also the number of blank lines
	// inserted between successive rows of trees. Default is 1.
	TSpace int
}

// New returns a Printer with zero‑value defaults overridden by the supplied
// functional options.
func New[T cmp.Ordered](opts ...func(*Printer[T])) *Printer[T] {
	p := &Printer[T]{
		Out:            os.Stdout,
		HSpace:         2,
		TSpace:         1,
		SquareBranches: false,
		LRAgnostic:     false,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// PrintTree renders a single tree rooted at root.
func (p *Printer[T]) PrintTree(root Node[T]) {
	p.printTreeLines(p.buildTreeLines(root))
}

// PrintTrees renders several trees across the page, wrapping according to the
// supplied maximum line width.
func (p *Printer[T]) PrintTrees(roots []Node[T], lineWidth int) {
	if len(roots) == 0 {
		return
	}

	// Build all line slices and measure widths.
	allLines := make([][]treeLine, len(roots))
	widths := make([]int, len(roots))
	minOffsets := make([]int, len(roots))
	maxOffsets := make([]int, len(roots))
	for i, r := range roots {
		lines := p.buildTreeLines(r)
		allLines[i] = lines
		minOffsets[i] = minLeftOffset(lines)
		maxOffsets[i] = maxRightOffset(lines)
		widths[i] = maxOffsets[i] - minOffsets[i] + 1
	}

	next := 0
	for next < len(roots) {
		// Decide which trees fit on this row.
		sum := widths[next]
		end := next + 1
		for end < len(roots) && sum+p.TSpace+widths[end] < lineWidth {
			sum += p.TSpace + widths[end]
			end++
		}
		end-- // inclusive index

		// Determine tallest tree in this row.
		maxLines := 0
		for i := next; i <= end; i++ {
			if l := len(allLines[i]); l > maxLines {
				maxLines = l
			}
		}

		// Print row line‑by‑line.
		for line := 0; line < maxLines; line++ {
			for col := next; col <= end; col++ {
				ln := "" // actual text for this tree & line
				if line < len(allLines[col]) {
					tl := allLines[col][line]
					left := -minOffsets[col] + tl.leftOffset
					right := maxOffsets[col] - tl.rightOffset
					ln = spaces(left) + tl.line + spaces(right)
				} else {
					ln = spaces(widths[col])
				}
				io.WriteString(p.Out, ln)
				if col < end {
					io.WriteString(p.Out, spaces(p.TSpace))
				}
			}
			io.WriteString(p.Out, "\n")
		}

		// Vertical spacing between rows.
		for i := 0; i < p.TSpace; i++ {
			io.WriteString(p.Out, "\n")
		}

		next = end + 1
	}
}

// --------------------------- internals ---------------------------

type treeLine struct {
	line        string
	leftOffset  int
	rightOffset int
}

func (t treeLine) String() string {
	return fmt.Sprintf("(line:\"%s\" l:%d r:%d)", t.line, t.leftOffset, t.rightOffset)
}

func (p *Printer[T]) buildTreeLines(root Node[T]) []treeLine {

	if root == nil {
		return nil
	}

	rootLabel := p.label(root)
	leftLines := p.buildTreeLines(root.Left())
	rightLines := p.buildTreeLines(root.Right())

	fmt.Println("treeliunbes", "left", leftLines, "right", rightLines)
	// Determine spacing between left‑ and right‑subtree roots.
	leftCnt := len(leftLines)
	rightCnt := len(rightLines)
	minCnt := min(leftCnt, rightCnt)

	maxRootSpacing := 0
	for i := 0; i < minCnt; i++ {
		spacing := leftLines[i].rightOffset - rightLines[i].leftOffset
		if spacing > maxRootSpacing {
			maxRootSpacing = spacing
		}
	}
	rootSpacing := maxRootSpacing + p.HSpace
	if rootSpacing%2 == 0 {
		rootSpacing++
	}

	// strip ANSI escape sequences when measuring length
	rendered := stripANSI(rootLabel)
	renderedLen := len(rendered)

	// Build result lines slice.
	var result []treeLine
	result = append(result, treeLine{
		line:        rootLabel,
		leftOffset:  -(renderedLen - 1) / 2,
		rightOffset: renderedLen / 2,
	})

	// Track subtree offset adjustments.
	var leftAdj, rightAdj int

	switch {
	case len(leftLines) == 0 && len(rightLines) == 0:
		// No children. Done.
	case len(leftLines) == 0: // right‑only
		if p.SquareBranches {
			if p.LRAgnostic {
				result = append(result, treeLine{line: "│", leftOffset: 0, rightOffset: 0})
			} else {
				result = append(result, treeLine{line: "└┐", leftOffset: 0, rightOffset: 1})
				rightAdj = 1
			}
		} else {
			result = append(result, treeLine{line: "\\", leftOffset: 1, rightOffset: 1})
			rightAdj = 2
		}
	case len(rightLines) == 0: // left‑only
		if p.SquareBranches {
			if p.LRAgnostic {
				result = append(result, treeLine{line: "│", leftOffset: 0, rightOffset: 0})
			} else {
				result = append(result, treeLine{line: "┌┘", leftOffset: -1, rightOffset: 0})
				leftAdj = -1
			}
		} else {
			result = append(result, treeLine{line: "/", leftOffset: -1, rightOffset: -1})
			leftAdj = -2
		}
	default: // both children present
		if p.SquareBranches {
			adjust := rootSpacing/2 + 1
			horizontal := strings.Repeat("─", rootSpacing/2)
			branch := "┌" + horizontal + "┬" + horizontal + "┐"
			result = append(result, treeLine{line: branch, leftOffset: -adjust, rightOffset: adjust})
			leftAdj = -adjust
			rightAdj = adjust
		} else {
			if rootSpacing == 1 {
				result = append(result, treeLine{line: "/ \\", leftOffset: -1, rightOffset: 1})
				leftAdj = -2
				rightAdj = 2
			} else {
				for i := 1; i < rootSpacing; i += 2 {
					result = append(result, treeLine{
						line:        "/" + spaces(i) + "\\",
						leftOffset:  -((i + 1) / 2),
						rightOffset: (i + 1) / 2,
					})
				}
				leftAdj = -(rootSpacing/2 + 1)
				rightAdj = rootSpacing/2 + 1
			}
		}
	}

	// Merge subtree lines.
	maxCnt := max(len(leftLines), len(rightLines))
	for i := 0; i < maxCnt; i++ {
		if i >= len(leftLines) { // only right remains
			rl := rightLines[i]
			rl.leftOffset += rightAdj
			rl.rightOffset += rightAdj
			result = append(result, rl)
			continue
		}
		if i >= len(rightLines) { // only left remains
			ll := leftLines[i]
			ll.leftOffset += leftAdj
			ll.rightOffset += leftAdj
			result = append(result, ll)
			continue
		}

		ll := leftLines[i]
		rl := rightLines[i]
		adjSpacing := rootSpacing
		if rootSpacing == 1 {
			if p.SquareBranches {
				adjSpacing = 1
			} else {
				adjSpacing = 3
			}
		}
		line := ll.line + spaces(adjSpacing-ll.rightOffset+rl.leftOffset) + rl.line
		result = append(result, treeLine{
			line:        line,
			leftOffset:  ll.leftOffset + leftAdj,
			rightOffset: rl.rightOffset + rightAdj,
		})
	}
	return result
}

func (p *Printer[T]) printTreeLines(lines []treeLine) {
	fmt.Println("lines", lines)
	if len(lines) == 0 {
		return
	}
	minLeft := minLeftOffset(lines)
	maxRight := maxRightOffset(lines)
	for _, tl := range lines {
		left := -minLeft + tl.leftOffset
		right := maxRight - tl.rightOffset
		io.WriteString(p.Out, spaces(left)+tl.line+spaces(right)+"\n")
	}
}

func (p *Printer[T]) label(n Node[T]) string {
	if p.Label != nil {
		return p.Label(n.Value())
	}
	return fmt.Sprint(n.Value())
}

// --------------------------- helpers ---------------------------

func stripANSI(s string) string {
	// Matches ANSI escape sequences. Same pattern as in original Java port.
	re := regexp.MustCompile(`\x1b\[[0-9;]*[^0-9;]`)
	return re.ReplaceAllString(s, "")
}

func spaces(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(" ", n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minLeftOffset(lines []treeLine) int {
	min := 0
	for i, tl := range lines {
		if i == 0 || tl.leftOffset < min {
			min = tl.leftOffset
		}
	}
	return min
}

func maxRightOffset(lines []treeLine) int {
	maxVal := 0
	for i, tl := range lines {
		if i == 0 || tl.rightOffset > maxVal {
			maxVal = tl.rightOffset
		}
	}
	return maxVal
}

// --------------------------- convenience API ---------------------------

// Print renders a tree with default settings to Stdout. A convenience wrapper
// for quick debugging without having to create a Printer manually.
func Print[T cmp.Ordered](root Node[T]) {
	New[T]().PrintTree(root)
}
