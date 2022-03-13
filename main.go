package main

func main() {
	bc := InitBlockchain()

	defer bc.db.Close()
	Run(bc)
}
