package hash

import "myTools/srack/vars"

func MakeTaskHash(k string) string {
	hash := MD5(k)
	return hash
}

func CheckTaskHash(hash string) bool {
	_, ok := vars.SuccessHash.Load(hash)
	return ok
}

func SetTaskHash(hash string) () {
	vars.SuccessHash.Store(hash, true)
}
