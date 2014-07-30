package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

/*
Switch evaluation order

switch caseは、上から下へcaseを評価します。 caseの条件が一致すれば、そこで停止します。

(例えば、

switch i {
case 0:
case f():
}
というコードで、もし、caseの条件が i==0 であれば、 f の関数を呼び出しません)

注意：Go playgroundでの時間はいつも 2009-11-10 23:00:00 UTC で始まります。 この値の意味は、読者の演習のために残しておきます。
*/
