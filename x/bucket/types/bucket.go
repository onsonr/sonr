package types


// AddObjects takes a list of fields and adds it to BucketDoc
func (o *BucketDoc) AddObjects(l ...string) {
	for _, v := range o.GetObjectDids() {
		o.ObjectDids = append(o.ObjectDids, v)

	}
}

// RemoveObjects takes a list of Object Dids
// and removes the matching label from the BucketDoc
func (o *BucketDoc) RemoveObjects(l ...string) {
	for _, v := range l {
		remove(o.ObjectDids, v)
	}
}

func remove(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
