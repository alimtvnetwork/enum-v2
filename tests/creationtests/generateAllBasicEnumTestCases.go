package creationtests

import (
	"fmt"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/constants"
)

func generateAllBasicEnumTestCases() {
	length := len(simpleEnumCollectionTestCases)
	tab := constants.Tab

	fmt.Println("var allBasicEnumsCollection = [...]enuminf.BasicEnumer{")

	for i := 0; i < length; i++ {
		item := simpleEnumCollectionTestCases[i]
		typeName := item.TypeName()
		name := item.Name()
		fullInvalidName := codestack.JoinPackageNameWithRelative(
			typeName,
			name)

		fmt.Println(
			tab,
			fullInvalidName+".AsBasicEnumContractsBinder(),",
		)
	}

	fmt.Println(
		"}",
	)
}
