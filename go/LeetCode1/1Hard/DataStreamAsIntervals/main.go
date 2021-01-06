package main

import "fmt"

type SummaryRanges struct {
	intervals   [][]int
	breakPoints map[int]int
}

func isIn(n int, arr []int) bool {
	if n >= arr[0] && n <= arr[1] {
		return true
	}
	return false
}

func Constructor() SummaryRanges {
	ret := SummaryRanges{}
	ret.intervals = [][]int{}
	ret.breakPoints = map[int]int{}
	return ret
}

func (this *SummaryRanges) AddNum(val int) {
	v, exist := this.breakPoints[val]
	if exist {
		// merge two
		this.intervals[v][1] = this.intervals[v+1][1]
		head := this.intervals[:v+1]
		tail := [][]int{}
		if v+2 < len(this.intervals) {
			tail = this.intervals[v+2:]
		}
		this.intervals = append(head, tail...)

		// del
		delete(this.breakPoints, val)
		return
	}
	if len(this.intervals) == 0 {
		this.intervals = append(this.intervals, []int{val, val})
		return
	}

	for i, v := range this.intervals {
		if isIn(val, v) {
			//do nothing
			return
		} else {

			// branch break
			if i > 0 && val < v[0]-1 {
				// test breakpoints no done before

				// new interval insertion
				newInterval := []int{val, val}
				// update breakpoints

				head := this.intervals[:i]
				tail := this.intervals[i:]
				newTotal := [][]int{}
				newTotal = append(newTotal, head...)
				newTotal = append(newTotal, newInterval)
				newTotal = append(newTotal, tail...)
				this.intervals = newTotal
				// update breakpoints for this is after one interval
				if val == v[0]-2 {
					// matched
					this.breakPoints[val+1] = i
				}
				if val == this.intervals[i-1][1]+2 {
					// second break
					this.breakPoints[val-1] = i - 1
				}
				return
			}
			if i == 0 && val < v[0]-1 {
				newInterval := []int{val, val}
				this.intervals = append([][]int{newInterval}, this.intervals...)
				//update breakpoints
				if val == this.intervals[1][0]-2 {
					this.breakPoints[val+1] = 0
				}
				return
			}
			if i == len(this.intervals)-1 && val > v[1]+1 {
				newInterval := []int{val, val}
				this.intervals = append(this.intervals, newInterval)
				//update breakpoints
				if val == v[1]+2 {
					this.breakPoints[val-1] = i
				}
				return
			}

			if val == v[1]+1 {
				this.intervals[i][1] = val
				if i < len(this.intervals)-1 && this.intervals[i+1][0]-2 == val {
					this.breakPoints[val+1] = i
				}
				return
			}
			if val == v[0]-1 {
				this.intervals[i][0] = val
				if i > 0 && this.intervals[i-1][1]+2 == val {
					this.breakPoints[val-1] = i - 1
				}
				return
			}

		}
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	// ret := [][]int{}
	// copy(this.intervals, ret)
	return this.intervals
}

func main() {
	obj := Constructor()
	inputArr := []int{6, 6, 0, 4, 8, 7, 6, 4, 7, 5}
	for _, i := range inputArr {
		obj.AddNum(i)
		fmt.Printf("%v", obj.GetIntervals())
	}

}
