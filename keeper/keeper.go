package keeper

type Keeper struct {
	num int
	min int
	max int
}

func (k *Keeper) MakeGuess(val int) int{
	if val > k.num {
		return 1
	}
	if val < k.num {
		return -1
	}
	return 0
}
