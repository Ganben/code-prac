package GasStation

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalTank, currTank := 0, 0
	starting := 0
	for i := 0; i < n; i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]
		if currTank < 0 {
			starting = i + 1
			currTank = 0
		}
	}
	if totalTank >= 0 {
		return starting
	} else {
		return -1
	}
}
