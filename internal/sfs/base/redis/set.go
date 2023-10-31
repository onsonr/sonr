package redis

import (
	"fmt"

	"github.com/sonrhq/sonr/internal/sfs/types"
)

type sfsSet struct {
	key         string
	backupItems []string
	isOnline    bool
}

// NewSet creates a new set
func NewSet(key string) types.SFSSet {
	return &sfsSet{
		key:         key,
		isOnline:    ifr.rdb != nil,
		backupItems: make([]string, 0),
	}
}

func (st *sfsSet) Add(item string) error {
	if st.isOnline {
		return addSetItem(st.key, item)
	}
	st.backupItems = append(st.backupItems, item)
	return nil
}

func (st *sfsSet) Remove(item string) error {
	if st.isOnline {
		_, err := delSetItem(st.key, item)
		return err
	}
	for i, v := range st.backupItems {
		if v == item {
			st.backupItems = append(st.backupItems[:i], st.backupItems[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (st *sfsSet) Contains(item string) (bool, error) {
	if !st.isOnline {
		for _, v := range st.backupItems {
			if v == item {
				return true, nil
			}
		}
		return false, nil
	}
	return existsInSet(st.key, item)
}

func (st *sfsSet) GetAll() ([]string, error) {
	if !st.isOnline {
		return st.backupItems, nil
	}
	return getSet(st.key)
}

func (st *sfsSet) Get(index int) (string, error) {
	if !st.isOnline {
		return st.backupItems[index], nil
	}
	sti, err := st.GetAll()
	if err != nil {
		return "", err
	}
	return sti[index], nil
}
