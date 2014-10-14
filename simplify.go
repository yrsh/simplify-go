package simplify

//-statck------------------------------------------------
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) > 0 {
		ret := (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
		return ret
	} else {
		return 0
	}
}

//-------------------------------------------------------

type Point struct {
	X float64
	Y float64
}

func getSqDist(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return dx*dx + dy*dy
}

func getSqSegDist(p, p1, p2 Point) float64 {
	x := p1.X
	y := p1.Y
	dx := p2.X - x
	dy := p2.Y - y
	if dx != 0 || dy != 0 {
		t := ((p.X-x)*dx + (p.Y-y)*dy) / (dx*dx + dy*dy)
		if t > 1 {
			x = p2.X
			y = p2.Y
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}
	dx = p.X - x
	dy = p.Y - y
	return dx*dx + dy*dy
}

func simplifyRadialDist(points []Point, sqTolerance float64) []Point {
	prevPoint := points[0]
	newPoints := []Point{prevPoint}
	var point Point
	for i := 1; i < len(points); i++ {
		point = points[i]
		if getSqDist(point, prevPoint) > sqTolerance {
			newPoints = append(newPoints, point)
			prevPoint = point
		}
	}
	if prevPoint != point {
		newPoints = append(newPoints, point)
	}
	return newPoints
}

func simplifyDouglasPeucker(points []Point, sqTolerance float64) []Point {
	var l = len(points)
	markers := make([]int, l)
	first := 0
	last := l - 1
	var stack Stack
	var newPoints []Point
	i, index := 0, 0
	maxSqDist, sqDist := float64(0), float64(0)
	markers[first], markers[last] = 1, 1
	for last > 0 {
		maxSqDist = 0
		for i = first + 1; i < last; i++ {
			sqDist = getSqSegDist(points[i], points[first], points[last])
			if sqDist > maxSqDist {
				index = i
				maxSqDist = sqDist
			}
		}
		if maxSqDist > sqTolerance {
			markers[index] = 1
			stack.Push(first)
			stack.Push(index)
			stack.Push(index)
			stack.Push(last)
		}
		last = stack.Pop()
		first = stack.Pop()
	}
	for i = 0; i < l; i++ {
		if checkArrIndex(markers, i) {
			newPoints = append(newPoints, points[i])

		}
	}
	return newPoints
}

func checkArrIndex(arr []int, index int) bool {
	if index < len(arr) && index >= 0 {
		if arr[index] > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Simplify(points []Point, tolerance float64, highestQuality bool) []Point {
	if len(points) <= 1 {
		return points
	}
	sqTolerance := tolerance * tolerance
	var _points []Point
	if highestQuality {
		_points = points
	} else {
		_points = simplifyRadialDist(points, sqTolerance)
	}
	_points = simplifyDouglasPeucker(_points, sqTolerance)
	return _points
}
