package main

func main() {
	bc := CreateBlockChain("zj")
	cli := CLI{bc: bc}
	cli.Run()
	//bc.AddBlock("11111111111111111111111")
	//bc.AddBlock("22222222222222222222222")

	//}
}
