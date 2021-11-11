package bc 

type BlockChain struct {
    Blocks []*Block
}

type Header struct {
    PrevHash []byte
    Hash []byte
    Timestamp string
    Version string
    Nonce int
    DifficultyTarget int
}
type Block struct {
    Header
    Transactions []byte
}

func CreateBlock(data string, prevHash []byte) *Block {
    header := Header{Nonce: 0, DifficultyTarget: 0, Version: "1", PrevHash: prevHash}
    block := &Block{header, []byte(data)} 
    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

func (chain *BlockChain) AddBlock(data string) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}