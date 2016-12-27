package gohelper

func Ngo(n int, f func(chan bool)) int {
	c := make(chan bool, n*3)
	for i := 0; i < n; i++ {
		c <- true
	}
	f(c)
	t := -n
	for i := 0; i < n; i++ {
		if <-c {
			t++
		}
	}
	return t
}
