package helpers

import "./golibs-helperfuncs/unique"


func DeterministicIDGen(seed string) string {
	id := unique.DeterministicUUID(seed)
	return ""
}

func RandomIDGen() string {
	panic("Not implemented")
	return ""
}
