package blockchain

// Output means where the money go to
type TxOutput struct {
	Value  int
	PubKey string
}

// Input means where the money come from
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
