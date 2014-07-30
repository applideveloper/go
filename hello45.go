package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

/*
Switch

みなさんは、 switch がどのようなものだったかを知っていたでしょうか。 Goは違います。

Goでのswitchは、 case の最後で自動的にbreakします。 もし、brakeせずに通したい場合は、 fallthrough 文をcaseの最後に記述します。
*/
