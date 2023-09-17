package main

// maps are unordered lists of key-value pairs
// maps are reference types
// maps are dynamically resizable
// maps are created with make(map[Type]Type) or map[Type]Type{}
// maps are not comparable
// maps are not safe for concurrent use
// maps are initialized with a key type and a value type
// maps are accessed with square brackets []
// maps are a collection of unordered key-value pairs
func main() {
	salarios := map[string]int{"Wesley": 1000, "Amanda": 5000, "João": 2000}
	delete(salarios, "Wesley")
	salarios["Wes"] = 3000
	
	for nome, salario := range salarios {
		println(nome, salario)
	}

	sal := make(map[string]int)
	sal1 := map[string]int{}
	
	sal["Wesley"] = 1000
	sal1["Amanda"] = 2000

	println(sal["Wesley"])
	println(sal1["Amanda"])
}