package types

func (a *AssetInfo) Equal(b *AssetInfo) bool {
	if a == nil && b == nil {
		return true
	}
	return false
}

func (c *ChainInfo) Equal(b *ChainInfo) bool {
	if c == nil && b == nil {
		return true
	}
	return false
}
