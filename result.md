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

**4. Sharing properly**
1. foo.c
- used POSIX mutexes to prevent race contdition (only one thread can modify the variable)
- each thread locks and unlocks the variable in question to ensure that it is not changes otherwise
- balanced increments/decrements give 0 (changing one function repeats a different number)

2. foo.go
- a server goroutine handles all modifications on data by listnening to a channel for requests (3 options and done messege)

