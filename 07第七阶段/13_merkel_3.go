package main
import (
	"fmt"
	"encoding/hex"
	"crypto/sha256"
)

type MerkelTree struct{
	RootNode *MerkelNode
}
type MerkelNode struct{
	Left *MerkelNode
	Right *MerkelNode
	Data []byte
}
func min(a,b int) int  {
	if (a > b) {
		return b
	}
	return a
	
}

func NewMerkelNode(left, right *MerkelNode,data []byte) *MerkelNode {
	mnode := MerkelNode{}
	if left == nil && right == nil {
		mnode.Data = data
	}else{
		prevhashes := append(left.Data, right.Data...)
		firsthash := sha256.Sum256(prevhashes)
		hash := sha256.Sum256(firsthash[:])
		mnode.Data = hash[:]
	}
	mnode.Left = left
	mnode.Right = right
	return &mnode
}

//构建默克尔树
func NewMerkelTree(data [][]byte) *MerkelTree {
	var nodes  []MerkelNode
	for _,datum := range data{
		node := NewMerkelNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	j := 0
	for nSize := len(data);nSize > 1;nSize = (nSize+1)/2{
		for i := 0;i < nSize;i += 2{
			i2 := min(i+1, nSize-1)	
			node := NewMerkelNode(&nodes[j+i], &nodes[j+i2],nil)
			nodes = append(nodes, *node)
		}
		j+= nSize
	}
	mTree := MerkelTree{&(nodes[len(nodes)- 1])}
	return &mTree
}
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
func main() {
	// https://www.blockchain.com/btc/block/00000000000090ff2791fe41d80509af6ffbd6c5b10294e29cdf1b603acab92c
	//传递hash
	data1, _ := hex.DecodeString("6b6a4236fb06fead0f1bd7fc4f4de123796eb51675fb55dc18c33fe12e33169d")
	data2, _ := hex.DecodeString("2af6b6f6bc6e613049637e32b1809dd767c72f912fef2b978992c6408483d77e")
	data3, _ := hex.DecodeString("6d76d15213c11fcbf4cc7e880f34c35dae43f8081ef30c6901f513ce41374583")
	data4, _ := hex.DecodeString("08c3b50053b010542dca85594af182f8fcf2f0d2bfe8a806e9494e4792222ad2")
	data5, _ := hex.DecodeString("612d035670b7b9dad50f987dfa000a5324ecb3e08745cfefa10a4cefc5544553")
	//MerkelRoot : c66ee6e01c2332b92e71e17b6c6c3d4e926f6330a06acbb0e203bf7d97d12249
	ReverseBytes(data1)
	ReverseBytes(data2)
	ReverseBytes(data3)
	ReverseBytes(data4)
	ReverseBytes(data5)

	hehe := [][]byte{
		data1,
		data2,
		data3,
		data4,
		data5,
	}
	merkelroot := NewMerkelTree(hehe)
	ReverseBytes(merkelroot.RootNode.Data)
	fmt.Printf("%x\n", merkelroot.RootNode.Data)
}
