package custpack

func Reverse(s string) string {
	var resoult string

	for _, v := range s{
		resoult = string(v) + resoult
	}
	return resoult

}
