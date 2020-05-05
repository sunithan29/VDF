// VDF
//
// Squareroot test 
//
//

package sqrt

import (
	"math/big"
)

// VDFSqrt --
type VDFSquare_rt struct {
	p *big.Int
}

// NewVDFSqrt -- creates new VDF.
func NewVDF_Square_rt(p *big.Int) *VDFSquare_rt {
	return &VDFSquare_rt{
		p: p,
	}
}

// Euler's criterion is a formula for determining whether
//  an integer is a quadratic residue modulo a prime.
// A^((p-1)/2) % p, if it equals 1 then A is a quadratic
//  residue module a prime number

func (v *VDFSquare_rt) quadraticResidue(x *big.Int) bool {
	// p-1/2
	m := new(big.Int).Set(x)
	t := new(big.Int).Exp(m, new(big.Int).Div(new(big.Int).Sub(v.p, big.NewInt(1)), big.NewInt(2)), v.p)
	return t.Cmp(big.NewInt(1)) == 0
}

// Delay -- eval the delay.
func (v *VDFSquare_rt) Delay(t int64, x *big.Int) *big.Int {
	var r *big.Int
	m := new(big.Int).Set(x)

	for i := int64(0); i < t; i++ {
		if !v.quadraticResidue(m) {
			m = m.Neg(m).Mod(m, v.p)
		}
		r = m.ModSqrt(m, v.p)
	}
	return r
}

// Verify -- verify the result y.
func (v *VDFSquare_rt) Verify(t int64, x *big.Int, y *big.Int) bool {
	var r *big.Int
	n := new(big.Int).Set(y)

	if !v.quadraticResidue(x) {
		x = big.NewInt(0).Mod(big.NewInt(0).Neg(x), v.p)
	}
	for i := int64(0); i < t; i++ {
		r = n.Exp(n, big.NewInt(2), v.p)
	}
	return r.Cmp(x) == 0
}
