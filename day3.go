package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strconv"
	"strings"
)

type claim struct {
	id      string
	offsetX int
	offsetY int
	lengthX int
	lengthY int
}

func main() {
	rawClaimInputs := utils.ReadArgumentFile()
	claims, _ := mapRawClaimsToStruct(rawClaimInputs)
	claimedFabric := make(map[int]int)
	// Part 1
	for _, claim := range claims {
		claimedFabric = claimFabric(claimedFabric, claim)
	}
	fmt.Printf("size: %d\n", len(claimedFabric))
	overlap := 0
	for _, claimedInch := range claimedFabric {
		if claimedInch > 1 {
			overlap++
		}
	}

	fmt.Printf("overlap: %d\n", overlap)

	// Part 2, sloppy
	for _, claim := range claims {
		alone := true;
		for x := claim.offsetX; x < (claim.lengthX + claim.offsetX); x++ {
			for y := claim.offsetY; y < (claim.lengthY + claim.offsetY); y++ {
				slotHashKey := (x+y)*(x+y+1)/2 + x
				if val := claimedFabric[slotHashKey]; val != 1 {
					alone = false;
				}
			}
		}
		if alone {
			fmt.Printf("alone: %s\n", claim.id)
		}
	}
}

func mapRawClaimsToStruct(rawClaims []string) ([]claim, error) {
	claims := make([]claim, len(rawClaims))
	for i, rawClaim := range rawClaims {
		// noise removal
		rawClaim := strings.Replace(rawClaim, "@ ", "", -1)
		rawClaim = strings.Replace(rawClaim, ":", "", -1)
		splitInput := strings.Split(rawClaim, " ")
		if len(splitInput) != 3 {
			return nil, fmt.Errorf("Expected 3 split input, got %d", len(splitInput))
		}
		id := splitInput[0]
		offsetX, offsetY, err := parseOffset(splitInput[1])
		if err != nil {
			return nil, err
		}
		lengthX, lengthY, err := parseLength(splitInput[2])
		if err != nil {
			return nil, err
		}
		claims[i] = claim{
			id:      id,
			offsetX: offsetX,
			offsetY: offsetY,
			lengthX: lengthX,
			lengthY: lengthY,
		}
	}
	return claims, nil
}

func parseOffset(input string) (int, int, error) {
	// expects "<int>,<int>" format
	splitInput := strings.Split(input, ",")
	if len(splitInput) != 2 {
		return 0, 0, fmt.Errorf("Expected 2 split input, got %d", len(splitInput))
	}
	offsetX, err := strconv.Atoi(splitInput[0])
	if err != nil {
		return 0, 0, nil
	}
	offsetY, err := strconv.Atoi(splitInput[1])
	if err != nil {
		return 0, 0, nil
	}
	return offsetX, offsetY, nil
}

func parseLength(input string) (int, int, error) {
	// expects "<int>x<int>" format
	splitInput := strings.Split(input, "x")
	if len(splitInput) != 2 {
		return 0, 0, fmt.Errorf("Expected 2 split input, got %d", len(splitInput))
	}
	lengthX, err := strconv.Atoi(splitInput[0])
	if err != nil {
		return 0, 0, nil
	}
	lengthY, err := strconv.Atoi(splitInput[1])
	if err != nil {
		return 0, 0, nil
	}
	return lengthX, lengthY, nil
}

func claimFabric(claimedFabric map[int]int, claim claim) map[int]int {
	for x := claim.offsetX; x < (claim.lengthX + claim.offsetX); x++ {
		for y := claim.offsetY; y < (claim.lengthY + claim.offsetY); y++ {
			slotHashKey := (x+y)*(x+y+1)/2 + x
			if _, ok := claimedFabric[slotHashKey]; ok {
				claimedFabric[slotHashKey]++
			} else {
				claimedFabric[slotHashKey] = 1
			}
		}
	}
	return claimedFabric
}
