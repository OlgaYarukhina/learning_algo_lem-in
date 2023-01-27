package functions

func GetSolution (n int, als [][][]*Cell) ([][]*Cell, int) {

	var bestSol [][]*Cell                   // arr of solution, it can be 2 the same
	var h int                               // best solution h
	                          
	bestSol = als[0]
	h = calcHeight(n, als[0])
	
		for s := 1; s<len(als); s++ {          // s: all solutions      
			eachsh := calcHeight(n, als[s])    // check height each solution 
			if eachsh < h {
				h = eachsh
				bestSol = als[s]
			}
			if eachsh == h && len(als[s]) < len(bestSol){
				h = eachsh
				bestSol = als[s]
				}
			}
		return bestSol, h
	}

	func calcHeight (n int, eachS [][]*Cell) int {
		var highestP int
		h := 0

		for i:= 0; i<len(eachS); i++ {
				if len(eachS) == 1 && len(eachS[i]) < h {
					h = len (eachS[i]) + n                 // height if only 1 path in the solution
				} else {
					highestP = len(eachS[0])               // the highest path in the solution
					hAllPath := len(eachS[0])                
					for j:= 1; j<len(eachS); j++  {
					hAllPath = hAllPath + len(eachS[j])    // height of all pathes in one solution
						if len(eachS[j]) >= highestP {
							highestP = len(eachS[j])
						}
					}
					ht := (hAllPath + n)/len(eachS)           //height with ants
					if (hAllPath + n)%len(eachS) > 0 {
						ht = ht +1
					} 
						if ht > highestP {
							h = ht
						} else {
							h = highestP
						}
				}
		}
		return h
	}