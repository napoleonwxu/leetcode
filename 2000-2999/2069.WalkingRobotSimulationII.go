type Robot struct {
	step int
	locs []Loc
}

type Loc struct {
	x   int
	y   int
	dir string
}

func Constructor(width int, height int) Robot {
	robot := Robot{step: 0, locs: make([]Loc, 0, 2*(width+height-2))}
	robot.locs = append(robot.locs, Loc{x: 0, y: 0, dir: "South"})
	for i := 1; i < width; i++ {
		robot.locs = append(robot.locs, Loc{x: i, y: 0, dir: "East"})
	}
	for j := 1; j < height; j++ {
		robot.locs = append(robot.locs, Loc{x: width - 1, y: j, dir: "North"})
	}
	for i := width - 2; i >= 0; i-- {
		robot.locs = append(robot.locs, Loc{x: i, y: height - 1, dir: "West"})
	}
	for j := height - 2; j > 0; j-- {
		robot.locs = append(robot.locs, Loc{x: 0, y: j, dir: "South"})
	}
	return robot
}

func (this *Robot) Move(num int) {
	this.step += num
}

func (this *Robot) GetPos() []int {
	loc := this.locs[this.step%len(this.locs)]
	return []int{loc.x, loc.y}
}

func (this *Robot) GetDir() string {
	if this.step == 0 {
		return "East"
	}
	return this.locs[this.step%len(this.locs)].dir
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */