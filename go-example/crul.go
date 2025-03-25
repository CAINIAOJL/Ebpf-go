/*package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
	//"golang.org/x/tools/go/analysis/passes/loopclosure"
)

func walkDir(path string,n *sync.WaitGroup, fileSizes chan <- int64) {
	defer n.Done()
	for _, entry := range drients(path) {
		if entry.Type().IsDir() {
			n.Add(1)
			subdir := filepath.Join(path, entry.Name())
			//log.Fatalf("This is dir: %v", path)
			fmt.Fprintf(os.Stdout, "This is a dir : %v\n", path)
			walkDir(subdir, n, fileSizes)
		} else {
			fileinfo, err := entry.Info()
			if err != nil {
				log.Fatalf("due: %v\n", err)
			}
			fmt.Fprintf(os.Stdout, "This is a file : %v\n", fileinfo.Name())
			fmt.Fprintf(os.Stdout, "This is a file, size is  : %f GB\n", float64(fileinfo.Size())/1e9)
			fileSizes <- fileinfo.Size()
		}
	}
}

func drients(path string) []os.DirEntry {
	entires, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		//log.Fatalf("due: %v\n", err)
		return nil
	}
	return entires
}

func printdisk(nfiles int64, nbytes int64) {
	fmt.Printf("%d files  %f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick <- chan time.Time
	tick = time.Tick(1 * time.Millisecond)

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	/*go func ()  {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)	
	}()*/

	/*for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func ()  {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	/*for size := range fileSizes {
		nfiles++
		nbytes +=size
	}*/
/*loop:
	for {
		select{
		case size, ok := <- fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printdisk(nfiles, nbytes)
		}
	}
	//fmt.Printf("%d files  %f GB\n", nfiles, float64(nbytes)/1e9)
	printdisk(nfiles, nbytes)
}*/

