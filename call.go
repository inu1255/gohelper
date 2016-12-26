package gohelper

func Ngo(n int, f func(chan bool)) {
	c := make(chan bool, n*3)
	for i := 0; i < n; i++ {
		c <- true
	}
	f(c)
	for i := 0; i < n; i++ {
		<-c
	}
}
