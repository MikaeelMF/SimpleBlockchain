package main

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()
	Run(bc)
}
