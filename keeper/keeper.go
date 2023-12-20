package keeper

type Keeper struct {
	Num int
	Min int
	Max int
}

func (k *Keeper) MakeGuess(val int) int{
	return k.Num - val
}
