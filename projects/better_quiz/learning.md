### Reading from a Stdin

- os.Stdin provides a file as output by accessing the /dev/stdin os file.
- Tried reading it with io.ReadAll(f), even after the ans/prompt is received from the user. io.ReadAll was trying to read from os.Stdin. 
- Resorted bufio.NewScanner() with line breaks as the read terminating condition.

### Timer
- Timer starts a timer internal to the run time. On timeout it fires the current time via the channel, configured internally 

### Channels

``` go
<- chan //(producing channel)
chan <- //(consuming channel)
```
Channels are a way for go routines to communicate with each other.


