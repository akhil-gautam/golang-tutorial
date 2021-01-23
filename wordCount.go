package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
    result := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words{
	  _, ok := result[v]
	  if ok {
	    result[v] = result[v]+1
	  } else {
	    result[v] = 1
	  }
	}
	return result
}

func main() {
	WordCount("This is so cool")
}
