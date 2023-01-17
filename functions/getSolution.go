package functions

func GetSolution (n int, als [][][]*Cell) []*Cell {

	var bestSol [][]*Cell                   // arr of solution, it can be 2 the same
	var h int                               // best solution h
	bestSol = als[0]
	h = calcHeight(n, als[0])

	


		for s := 1; s<len(als); s++ {        // s: all solutions
			var eachsh int                   // each solution 
			

		}



		return bestSol[0]

	}


	func calcHeight (n int, eachS [][]*Cell) int {
		var highestP int
		var h int

		for i:= 0; i<len(eachS); i++ {
			h = len(eachS[i])
				if len(eachS[i]) == 1 {
					h = len (eachS[i]) + n                 // height if only 1 path in the solution
				} else {
					highestP = len(eachS[i])               // the highest path in the solution
					hAllPath := len(eachS[i])                
					for j:= 1; j<len(eachS); j++  {
					hAllPath = hAllPath + len(eachS[j])    // height of all pathes in one solution
						if len(eachS[j]) > highestP {
							highestP = len(eachS[j])
						}
					}
					ht := (hAllPath + n)/len(eachS)           //height with ants

					if ht > highestP {
						h = ht
					} else {
						h = highestP
					}
				}
		}
		return h
	}