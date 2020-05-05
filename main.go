//
// Verifiable Delay Function (VDF)
// 

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"sqrt"
	"math/big"
	"strconv"
)

var (
	U_flag string
	T_flag string

	prime256 = "70858336831223790464553084758144107463372421605095026016236146050634307855407"
)

func init() {
	flag.StringVar(&U_flag, "u", "", "u value")
	flag.StringVar(&T_flag, "t", "1000", "t value")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " --u=[u value] --t=[t value]")
	fmt.Println("For example: " + os.Args[0] + " --u=5 --t=1500")
}

func square_rt_VDF(t int64, u *big.Int, p *big.Int) {
	vdf := sqrt.NewVDF_Square_rt(p)
	// VDF Delay.
	start := time.Now()
	vdf_delay := vdf.Delay(t, x)
	current := time.Now()
	delay := current.Sub(start).Seconds()
	fmt.Println("-->")
	fmt.Println("Type:\tSloth Quadratic Residue")
	fmt.Printf("Value:\t%v\n", y.String())
	fmt.Printf("Delay:\t%.5f secs\n", delay)

	// VDF Verify.
	start = time.Now()
	vdf_verify := vdf.Verify(t, x, vdf_delay)
	current = time.Now()
	verify := current.Sub(start).Seconds()
	fmt.Printf("Verify:\t%.5f secs\n", verify)
	fmt.Printf("Verify Result:\t%v\n", vdf_verify)
	fmt.Printf("Delay/Verify:\t%.5f\n", delay/verify)
}

func main() {
	flag.Parse()
	if U_flag == "" {
		usage()
		os.Exit(0)
	}
	t, _ := strconv.ParseInt(V_flag, 10, 64)
	u, _ := new(big.Int).SetString(U_flag, 0)
	p, _ := new(big.Int).SetString(prime256, 0)
	fmt.Printf("T:%v, u:%v, P:%v\n", t, u.Int64(), p.Int64())

	square_rt_VDF(t, u, p)
}

