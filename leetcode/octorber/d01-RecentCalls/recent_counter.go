package main

type RecentCounter struct {
	Requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Requests: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.Requests = append(this.Requests, t)
	start := t - 3000
	for i, v := range this.Requests {
		if v >= start {
			return len(this.Requests[i:])
		}
	}
	return 0
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
