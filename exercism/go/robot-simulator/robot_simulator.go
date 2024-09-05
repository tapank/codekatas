package robot

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N = 0
	E = 1
	S = 2
	W = 3
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	return ""
}

// Step 2
// Define Action type here.
type Action rune

func StartRobot(command chan Command, action chan Action) {
	for a := range command {
		action <- Action(a)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	advance := func(extent Rect, robot Step2Robot) (NewE, NewN RU) {
		NewE, NewN = robot.Pos.Easting, robot.Pos.Northing
		switch robot.Dir {
		case N:
			NewN++
		case E:
			NewE++
		case S:
			NewN--
		case W:
			NewE--
		}

		if NewE >= extent.Min.Easting && NewN >= extent.Min.Northing && NewE <= extent.Max.Easting && NewN <= extent.Max.Northing {
			return NewE, NewN
		}
		return robot.Pos.Easting, robot.Pos.Northing
	}

	for a := range action {
		switch a {
		case 'R':
			robot.Dir = (robot.Dir + 1) % 4
		case 'L':
			robot.Dir = (robot.Dir + 3) % 4
		case 'A':
			robot.Pos.Easting, robot.Pos.Northing = advance(extent, robot)
		}
	}
	report <- robot
}

// Step 3
// Define Action3 type here.
type Action3 rune

func StartRobot3(name, script string, action chan Action3, log chan string) {
	panic("Please implement the StartRobot3 function")
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	panic("Please implement the Room3 function")
}
