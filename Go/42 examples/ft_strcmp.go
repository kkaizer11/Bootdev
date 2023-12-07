package main

func ft_strcmp(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	if s1 < s2 {
		return -1
	}
	return +1
}
