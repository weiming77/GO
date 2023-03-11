package main

import "fmt"

const SIZE=7

type  THashTable struct {
	array[SIZE]*TBucket
}

type TBucket {
	head *TBucketNode
}

type TBucketNode {
	key string
	next TBucketNode
}
func main (){

}