package functions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

type FunctionsImpl struct {
	shell *shell.Shell
	cache map[string]*Function
}

func New(shell *shell.Shell) FunctionInterface {
	impl := FunctionsImpl{
		shell: shell,
		cache: make(map[string]*Function),
	}

	return impl
}

func (fi FunctionsImpl) Store(f *Function) (string, error) {
	b, err := f.Marshal()
	if err != nil {
		return "", err
	}
	cid, err := fi.shell.Add(bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	fi.cache[cid] = f

	return cid, nil
}

func (fi FunctionsImpl) GetAndExecute(path string) error {
	err := fi.shell.Get(path, os.TempDir()+path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(os.TempDir() + path)
	if err != nil {
		return err
	}

	f := Function{}

	err = f.Unmarshal(data)
	if err != nil {
		return err
	}

	return fi.Execute(&f)
}

func (fi FunctionsImpl) Execute(function *Function) error {
	ts := fmt.Sprint(time.Now().Unix())
	b := make([]byte, function.file.Len())

	_, err := function.file.Read(b)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(os.TempDir()+"/"+ts, b, 0777)
	if err != nil {
		return err
	}

	out, err := exec.Command(os.TempDir() + ts).Output()
	if err != nil {
		return err
	}

	fmt.Print(string(out))
	return nil
}
