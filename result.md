**3. Sharing variable** 

The implementation in both foo.c and foo.go increments or decrements a hared variable using threads. 

1. foo.c 
- two threads are declared
- using the threads and created functions, two separate threads are started
- the threads are joined to wait for eachother

2. foo.go
- runtime.GOMAXPROCS(2) sets the maximum number of CPUs that can be executing simultaneously to 2 (to goroutines in parallell)
- go starts the functions as goroutines
- the sleep command makes the main function to sleep so the goroutines get time to execute (not proper, sync.WaitGroup)