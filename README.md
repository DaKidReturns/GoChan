# GoChan

Download images from 4chan  

### Features
1. Download all the images from a 4chan forum.
2. No dependincies, no extra modules required.
3. You can import this package into your go project and use it

### Building From the source (If you just want an image scrapper and nothing more)
1. Ensure your go version is >= 1.16
2. Clone the repository `git clone https://github.com/DaKidReturns/GoChan`
3. Go into the folder GoChan `$ cd GoChan/example` where you will find the example code as `main.go`
4. Run `go build -o gochan`
5. Now that you have compiled your source you can use it by `$ ./gochan [4chan-urlname]`

### Using the module for your project
1. Ensure your go version is >= 1.16
2. Run the command `go get github.com/DaKidReturns/GoChan`
3. Use it in your project by the import statement `import "github.com/DaKidReturns/GoChan"` it will be imported as `gochan` and not as `GoChan`.
If you prefer the latter then prefix the import statement with `GoChan` like so, ```import GoChan "github.com/DaKidReturns/GoChan"``` or 
	```go
	import(
	 	GoChan "github.com/DaKidReturns/GoChan"
	 )
	 ```

### Future Plans
1. [ ] Add a progress bar to show the currently downloading image
2. [ ] Currently it saves the pictures to the Directory `(HOME)/Pictures/Gochan` and takes only one command line argument  
3. [ ] I plan to do something with concurrency in golang, like download multiple files at once. (after I learn concurrency in golang)
4. [ ] Create better Documentation
5. [ ] Add a Logging feature
6. [ ] Add an option for saving to custom directories  

### Why did I do this?
1. I wanted to learn golang
2. I wanted to download all the images from a forum on 4chan [the one that contains a lot of technical drawings][4chanLink]


So I figured why not create a web scaper for it ;-)

[LINKS]:()
[4chanLink]:https://boards.4chan.org/hr/thread/3828834
