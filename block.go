package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
	"crypto/sha256"
)

type Block struct {
	Version uint64
	PrevHash []byte
	MerkelRoot []byte
	TimeStamp uint64
	Difficulty uint64
	Nonce uint64
	Hash  []byte
	Data   []byte
}
func Uint64ToByte(num uint64)[]byte{
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, //随便填写的无效值
		Nonce:      0, //同上
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()

	return &block
}
//3. 生成哈希
func (block *Block) SetHash() {
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	//将二维的切片数组链接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	//2. sha256
	//func Sum256(data []byte) [Size]byte {
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}