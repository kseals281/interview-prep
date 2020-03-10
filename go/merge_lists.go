package main

/*******************************************************
In order to win the prize for most cookies sold, my friend Alice and I are going to merge our Girl Scout
Cookies orders and enter as one unit.

We have our lists of orders sorted numerically already, in lists. Write a function to merge our lists of orders into
one sorted list.
*********************************************************/

func mergeCookieLists(mine, alice []int) (sorted []int) {
	for len(mine) != 0 && len(alice) != 0 {
		if mine[0] < alice[0] {
			sorted = append(sorted, mine[0])
			mine = mine[1:]
		} else {
			sorted = append(sorted, alice[0])
			alice = alice[1:]
		}
	}
	if len(mine) > 0 {
		sorted = append(sorted, mine...)
	} else {
		sorted = append(sorted, alice...)
	}

	return
}
