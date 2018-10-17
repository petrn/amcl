/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

/* MiotCL Fp^12 functions */
/* FP12 elements are of the form a+i.b+i^2.c */

package XXX

//import "fmt"

type FP24 struct {
	a *FP8
	b *FP8
	c *FP8
}

/* Constructors */
func NewFP24fp8(d *FP8) *FP24 {
	F:=new(FP24)
	F.a=NewFP8copy(d)
	F.b=NewFP8int(0)
	F.c=NewFP8int(0)
	return F
}

func NewFP24int(d int) *FP24 {
	F:=new(FP24)
	F.a=NewFP8int(d)
	F.b=NewFP8int(0)
	F.c=NewFP8int(0)
	return F
}

func NewFP24fp8s(d *FP8,e *FP8,f *FP8) *FP24 {
	F:=new(FP24)
	F.a=NewFP8copy(d)
	F.b=NewFP8copy(e)
	F.c=NewFP8copy(f)
	return F
}

func NewFP24copy(x *FP24) *FP24 {
	F:=new(FP24)
	F.a=NewFP8copy(x.a)
	F.b=NewFP8copy(x.b)
	F.c=NewFP8copy(x.c)
	return F
}

/* reduce all components of this mod Modulus */
func (F *FP24) reduce() {
	F.a.reduce()
	F.b.reduce()
	F.c.reduce()
}
/* normalise all components of this */
func (F *FP24) norm() {
	F.a.norm()
	F.b.norm()
	F.c.norm()
}
/* test x==0 ? */
func (F *FP24) iszilch() bool {
	//F.reduce()
	return (F.a.iszilch() && F.b.iszilch() && F.c.iszilch())
}

/* Conditional move */
func (F *FP24) cmove(g *FP24,d int) {
	F.a.cmove(g.a,d)
	F.b.cmove(g.b,d)
	F.c.cmove(g.c,d)
}

/* Constant time select from pre-computed table */
func (F *FP24) selector(g []*FP24,b int32) {

	m:=b>>31
	babs:=(b^m)-m

	babs=(babs-1)/2

	F.cmove(g[0],teq(babs,0))  // conditional move
	F.cmove(g[1],teq(babs,1))
	F.cmove(g[2],teq(babs,2))
	F.cmove(g[3],teq(babs,3))
	F.cmove(g[4],teq(babs,4))
	F.cmove(g[5],teq(babs,5))
	F.cmove(g[6],teq(babs,6))
	F.cmove(g[7],teq(babs,7))
 
 	invF:=NewFP24copy(F) 
	invF.conj()
	F.cmove(invF,int(m&1))
}

/* test x==1 ? */
func (F *FP24) Isunity() bool {
	one:=NewFP8int(1)
	return (F.a.Equals(one) && F.b.iszilch() && F.c.iszilch())
}
/* return 1 if x==y, else 0 */
func (F *FP24) Equals(x *FP24) bool {
	return (F.a.Equals(x.a) && F.b.Equals(x.b) && F.c.Equals(x.c))
}

/* extract a from this */
func (F *FP24) geta() *FP8 {
	return F.a
}
/* extract b */
func (F *FP24) getb() *FP8 {
	return F.b
}
/* extract c */
func (F *FP24) getc() *FP8 {
	return F.c
}
/* copy this=x */
func (F *FP24) Copy(x *FP24) {
	F.a.copy(x.a)
	F.b.copy(x.b)
	F.c.copy(x.c)
}
/* set this=1 */
func (F *FP24) one() {
	F.a.one()
	F.b.zero()
	F.c.zero()
}
/* this=conj(this) */
func (F *FP24) conj() {
	F.a.conj()
	F.b.nconj()
	F.c.conj()
}

/* Granger-Scott Unitary Squaring */
func (F *FP24) usqr() {
	A:=NewFP8copy(F.a)
	B:=NewFP8copy(F.c)
	C:=NewFP8copy(F.b)
	D:=NewFP8int(0)

	F.a.sqr()
	D.copy(F.a); D.add(F.a)
	F.a.add(D)

	F.a.norm();
	A.nconj()

	A.add(A)
	F.a.add(A)
	B.sqr()
	B.times_i()

	D.copy(B); D.add(B)
	B.add(D)
	B.norm();

	C.sqr()
	D.copy(C); D.add(C)
	C.add(D)
	C.norm();

	F.b.conj()
	F.b.add(F.b)
	F.c.nconj()

	F.c.add(F.c)
	F.b.add(B)
	F.c.add(C)
	F.reduce()

}

/* Chung-Hasan SQR2 method from http://cacr.uwaterloo.ca/techreports/2006/cacr2006-24.pdf */
func (F *FP24)  sqr() {
	A:=NewFP8copy(F.a)
	B:=NewFP8copy(F.b)
	C:=NewFP8copy(F.c)
	D:=NewFP8copy(F.a)

	A.sqr()
	B.mul(F.c)
	B.add(B); B.norm()
	C.sqr()
	D.mul(F.b)
	D.add(D)

	F.c.add(F.a)
	F.c.add(F.b); F.c.norm()
	F.c.sqr()

	F.a.copy(A)

	A.add(B)
	A.norm();
	A.add(C)
	A.add(D)
	A.norm();

	A.neg()
	B.times_i();
	C.times_i()

	F.a.add(B)

	F.b.copy(C); F.b.add(D)
	F.c.add(A)
	F.norm()
}

/* FP24 full multiplication this=this*y */
func (F *FP24) Mul(y *FP24) {
	z0:=NewFP8copy(F.a)
	z1:=NewFP8int(0)
	z2:=NewFP8copy(F.b)
	z3:=NewFP8int(0)
	t0:=NewFP8copy(F.a)
	t1:=NewFP8copy(y.a)

	z0.mul(y.a)
	z2.mul(y.b)

	t0.add(F.b); t0.norm()
	t1.add(y.b); t1.norm() 

	z1.copy(t0); z1.mul(t1)
	t0.copy(F.b); t0.add(F.c); t0.norm()

	t1.copy(y.b); t1.add(y.c); t1.norm()
	z3.copy(t0); z3.mul(t1)

	t0.copy(z0); t0.neg()
	t1.copy(z2); t1.neg()

	z1.add(t0)
	//z1.norm();
	F.b.copy(z1); F.b.add(t1)

	z3.add(t1)
	z2.add(t0)

	t0.copy(F.a); t0.add(F.c); t0.norm()
	t1.copy(y.a); t1.add(y.c); t1.norm()
	t0.mul(t1)
	z2.add(t0)

	t0.copy(F.c); t0.mul(y.c)
	t1.copy(t0); t1.neg()

	F.c.copy(z2); F.c.add(t1)
	z3.add(t1)
	t0.times_i()
	F.b.add(t0)
	z3.norm()
	z3.times_i()
	F.a.copy(z0); F.a.add(z3)
	F.norm()
}

/* Special case of multiplication arises from special form of ATE pairing line function */
func (F *FP24) smul(y *FP24,twist int ) {
	if twist==D_TYPE {
		z0:=NewFP8copy(F.a)
		z2:=NewFP8copy(F.b)
		z3:=NewFP8copy(F.b)
		t0:=NewFP8int(0)
		t1:=NewFP8copy(y.a)
		
		z0.mul(y.a)
		z2.pmul(y.b.real());
		F.b.add(F.a)
		t1.real().add(y.b.real())

		t1.norm(); F.b.norm()
		F.b.mul(t1)
		z3.add(F.c); z3.norm()
		z3.pmul(y.b.real())

		t0.copy(z0); t0.neg()
		t1.copy(z2); t1.neg()

		F.b.add(t0)
	//F.b.norm();

		F.b.add(t1)
		z3.add(t1); z3.norm()
		z2.add(t0)

		t0.copy(F.a); t0.add(F.c); t0.norm()
		t0.mul(y.a)
		F.c.copy(z2); F.c.add(t0)

		z3.times_i()
		F.a.copy(z0); F.a.add(z3)
	}
	if twist==M_TYPE {
		z0:=NewFP8copy(F.a)
		z1:=NewFP8int(0)
		z2:=NewFP8int(0)
		z3:=NewFP8int(0)
		t0:=NewFP8copy(F.a)
		t1:=NewFP8int(0)
		
		z0.mul(y.a)
		t0.add(F.b)
		t0.norm()

		z1.copy(t0); z1.mul(y.a)
		t0.copy(F.b); t0.add(F.c)
		t0.norm()

		z3.copy(t0) //z3.mul(y.c);
		z3.pmul(y.c.getb())
		z3.times_i()

		t0.copy(z0); t0.neg()

		z1.add(t0)
		F.b.copy(z1) 
		z2.copy(t0)

		t0.copy(F.a); t0.add(F.c)
		t1.copy(y.a); t1.add(y.c)

		t0.norm()
		t1.norm()
	
		t0.mul(t1)
		z2.add(t0)

		t0.copy(F.c)
			
		t0.pmul(y.c.getb())
		t0.times_i()

		t1.copy(t0); t1.neg()

		F.c.copy(z2); F.c.add(t1)
		z3.add(t1)
		t0.times_i()
		F.b.add(t0)
		z3.norm()
		z3.times_i()
		F.a.copy(z0); F.a.add(z3)
	}
	F.norm()
}

/* this=1/this */
func (F *FP24) Inverse() {
	f0:=NewFP8copy(F.a)
	f1:=NewFP8copy(F.b)
	f2:=NewFP8copy(F.a)
	f3:=NewFP8int(0)

	F.norm()
	f0.sqr()
	f1.mul(F.c)
	f1.times_i()
	f0.sub(f1); f0.norm()

	f1.copy(F.c); f1.sqr()
	f1.times_i()
	f2.mul(F.b)
	f1.sub(f2); f1.norm()

	f2.copy(F.b); f2.sqr()
	f3.copy(F.a); f3.mul(F.c)
	f2.sub(f3); f2.norm()

	f3.copy(F.b); f3.mul(f2)
	f3.times_i()
	F.a.mul(f0)
	f3.add(F.a)
	F.c.mul(f1)
	F.c.times_i()

	f3.add(F.c); f3.norm()
	f3.inverse()
	F.a.copy(f0); F.a.mul(f3)
	F.b.copy(f1); F.b.mul(f3)
	F.c.copy(f2); F.c.mul(f3)
}

/* this=this^p using Frobenius */
func (F *FP24) frob(f *FP2, n int) {
	f2:=NewFP2copy(f)
	f3:=NewFP2copy(f)

	f2.sqr()
	f3.mul(f2)

	f3.mul_ip(); f3.norm()

	for i:=0;i<n;i++ {
		F.a.frob(f3);
		F.b.frob(f3);
		F.c.frob(f3);

		F.b.qmul(f); F.b.times_i2()
		F.c.qmul(f2); F.c.times_i2(); F.c.times_i2()
	}
}

/* trace function */
func (F *FP24) trace() *FP8 {
	t:=NewFP8int(0)
	t.copy(F.a)
	t.imul(3)
	t.reduce()
	return t;
}

/* convert from byte array to FP24 */
func FP24_fromBytes(w []byte) *FP24 {
	var t [int(MODBYTES)]byte
	MB:=int(MODBYTES)

	for i:=0;i<MB;i++ {t[i]=w[i]}
	a:=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+MB]}
	b:=FromBytes(t[:])
	c:=NewFP2bigs(a,b)

	for i:=0;i<MB;i++ {t[i]=w[i+2*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+3*MB]}
	b=FromBytes(t[:])
	d:=NewFP2bigs(a,b)

	ea:=NewFP4fp2s(c,d)

	for i:=0;i<MB;i++ {t[i]=w[i+4*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+5*MB]}
	b=FromBytes(t[:])
	c=NewFP2bigs(a,b)

	for i:=0;i<MB;i++ {t[i]=w[i+6*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+7*MB]}
	b=FromBytes(t[:])
	d=NewFP2bigs(a,b)

	eb:=NewFP4fp2s(c,d)

	e:=NewFP8fp4s(ea,eb)


	for i:=0;i<MB;i++ {t[i]=w[i+8*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+9*MB]}
	b=FromBytes(t[:])
	c=NewFP2bigs(a,b)

	for i:=0;i<MB;i++ {t[i]=w[i+10*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+11*MB]}
	b=FromBytes(t[:])
	d=NewFP2bigs(a,b)

	fa:=NewFP4fp2s(c,d)

	for i:=0;i<MB;i++ {t[i]=w[i+12*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+13*MB]}
	b=FromBytes(t[:])
	c=NewFP2bigs(a,b)

	for i:=0;i<MB;i++ {t[i]=w[i+14*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+15*MB]}
	b=FromBytes(t[:])
	d=NewFP2bigs(a,b)

	fb:=NewFP4fp2s(c,d)

	f:=NewFP8fp4s(fa,fb)

	for i:=0;i<MB;i++ {t[i]=w[i+16*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+17*MB]}
	b=FromBytes(t[:]);
		
	c=NewFP2bigs(a,b)

	for i:=0;i<MB;i++ {t[i]=w[i+18*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+19*MB]}
	b=FromBytes(t[:])
	d=NewFP2bigs(a,b)

	ga:=NewFP4fp2s(c,d)


	for i:=0;i<MB;i++ {t[i]=w[i+20*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+21*MB]}
	b=FromBytes(t[:]);
		
	c=NewFP2bigs(a,b)

	for i:=0;i<MB;i++ {t[i]=w[i+22*MB]}
	a=FromBytes(t[:])
	for i:=0;i<MB;i++ {t[i]=w[i+23*MB]}
	b=FromBytes(t[:])
	d=NewFP2bigs(a,b)

	gb:=NewFP4fp2s(c,d)

	g:=NewFP8fp4s(ga,gb)
	
	return NewFP24fp8s(e,f,g)
}

/* convert this to byte array */
func (F *FP24) ToBytes(w []byte) {
	var t [int(MODBYTES)]byte
	MB:=int(MODBYTES)
	F.a.a.geta().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i]=t[i]}
	F.a.a.geta().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+MB]=t[i]}
	F.a.a.getb().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+2*MB]=t[i]}
	F.a.a.getb().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+3*MB]=t[i]}

	F.a.b.geta().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+4*MB]=t[i]}
	F.a.b.geta().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+5*MB]=t[i]}
	F.a.b.getb().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+6*MB]=t[i]}
	F.a.b.getb().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+7*MB]=t[i]}

	F.b.a.geta().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+8*MB]=t[i]}
	F.b.a.geta().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+9*MB]=t[i]}
	F.b.a.getb().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+10*MB]=t[i]}
	F.b.a.getb().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+11*MB]=t[i]}

	F.b.b.geta().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+12*MB]=t[i]}
	F.b.b.geta().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+13*MB]=t[i]}
	F.b.b.getb().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+14*MB]=t[i]}
	F.b.b.getb().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+15*MB]=t[i]}

	F.c.a.geta().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+16*MB]=t[i]}
	F.c.a.geta().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+17*MB]=t[i]}
	F.c.a.getb().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+18*MB]=t[i]}
	F.c.a.getb().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+19*MB]=t[i]}

	F.c.b.geta().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+20*MB]=t[i]}
	F.c.b.geta().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+21*MB]=t[i]}
	F.c.b.getb().GetA().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+22*MB]=t[i]}
	F.c.b.getb().GetB().ToBytes(t[:])
	for i:=0;i<MB;i++ {w[i+23*MB]=t[i]}

}

/* convert to hex string */
func (F *FP24) ToString() string {
	return ("["+F.a.toString()+","+F.b.toString()+","+F.c.toString()+"]")
}

/* this=this^e */ 
func (F *FP24) Pow(e *BIG) *FP24 {
	sf:=NewFP24copy(F)
	sf.norm()
	
	e1:=NewBIGcopy(e)
	e1.norm()
	e3:=NewBIGcopy(e1)
	e3.pmul(3)
	e3.norm()

	w:=NewFP24copy(sf)
	//z:=NewBIGcopy(e)
	//r:=NewFP12int(1)

	nb:=e3.nbits()
	for i:=nb-2;i>=1;i-- {
		w.usqr()
		bt:=e3.bit(i)-e1.bit(i)
		if bt==1 {
			w.Mul(sf)
		}
		if bt==-1 {
			sf.conj()
			w.Mul(sf)
			sf.conj()	
		}
	}
	w.reduce()
	return w
/*
	for true {
		bt:=z.parity()
		z.fshr(1)
		if bt==1 {r.Mul(w)}
		if z.iszilch() {break}
		w.usqr()
	}
	r.reduce();
	return r; */
}

/* constant time powering by small integer of max length bts */
func (F *FP24) pinpow(e int,bts int) {
	var R []*FP24
	R=append(R,NewFP24int(1))
	R=append(R,NewFP24copy(F))

	for i:=bts-1;i>=0;i-- {
		b:=(e>>uint(i))&1
		R[1-b].Mul(R[b])
		R[b].usqr()
	}
	F.Copy(R[0])
}

/* Fast compressed FP8 power of unitary FP24 */
func (F *FP24) Compow(e *BIG,r *BIG) *FP8 {
	q:=NewBIGints(Modulus)
	//r:=NewBIGints(CURVE_Order)
	f:=NewFP2bigs(NewBIGints(Fra),NewBIGints(Frb))

	m:=NewBIGcopy(q)
	m.Mod(r)

	a:=NewBIGcopy(e)
	a.Mod(m)

	b:=NewBIGcopy(e)
	b.div(m)

	g1:=NewFP24copy(F);
	c:=g1.trace()

	if b.iszilch() {
		c=c.xtr_pow(e)
		return c
	}

	g2:=NewFP24copy(F)
	g2.frob(f,1)
	cp:=g2.trace()

	g1.conj()
	g2.Mul(g1)
	cpm1:=g2.trace()
	g2.Mul(g1)
	cpm2:=g2.trace()

	c=c.xtr_pow2(cp,cpm1,cpm2,a,b)
	return c
}

/* p=q0^u0.q1^u1.q2^u2.q3^u3.. */
// Bos & Costello https://eprint.iacr.org/2013/458.pdf
// Faz-Hernandez & Longa & Sanchez  https://eprint.iacr.org/2013/158.pdf
// Side channel attack secure 

func pow8(q []*FP24,u []*BIG) *FP24 {
	var g1 []*FP24
	var g2 []*FP24
	var w1 [NLEN*int(BASEBITS)+1]int8
	var s1 [NLEN*int(BASEBITS)+1]int8
	var w2 [NLEN*int(BASEBITS)+1]int8
	var s2 [NLEN*int(BASEBITS)+1]int8
	var t []*BIG
	r:=NewFP24int(0)
	p:=NewFP24int(0)
	mt:=NewBIGint(0)
	var bt int8
	var k int

	for i:=0;i<8;i++ {
		t=append(t,NewBIGcopy(u[i]))
	}

	g1=append(g1,NewFP24copy(q[0]))	// q[0]
	g1=append(g1,NewFP24copy(g1[0])); g1[1].Mul(q[1])	// q[0].q[1]
	g1=append(g1,NewFP24copy(g1[0])); g1[2].Mul(q[2])	// q[0].q[2]
	g1=append(g1,NewFP24copy(g1[1])); g1[3].Mul(q[2])	// q[0].q[1].q[2]
	g1=append(g1,NewFP24copy(g1[0])); g1[4].Mul(q[3])	// q[0].q[3]
	g1=append(g1,NewFP24copy(g1[1])); g1[5].Mul(q[3])	// q[0].q[1].q[3]
	g1=append(g1,NewFP24copy(g1[2])); g1[6].Mul(q[3])	// q[0].q[2].q[3]
	g1=append(g1,NewFP24copy(g1[3])); g1[7].Mul(q[3])	// q[0].q[1].q[2].q[3]

	Fra:=NewBIGints(Fra)
	Frb:=NewBIGints(Frb)
	X:=NewFP2bigs(Fra,Frb)	

// Use Frobenius
	for i:=0;i<8;i++ {
		g2=append(g2,NewFP24copy(g1[i])); g2[i].frob(X,4)		
	}

// Make them odd
	pb1:=1-t[0].parity()
	t[0].inc(pb1)
//	t[0].norm();

	pb2:=1-t[4].parity()
	t[4].inc(pb2)
//	t[4].norm();

// Number of bits
	mt.zero()
	for i:=0;i<8;i++ {
		t[i].norm()
		mt.or(t[i])
	}

	nb:=1+mt.nbits();
	
// Sign pivot 
	s1[nb-1]=1
	s2[nb-1]=1
	for i:=0;i<nb-1;i++ {
		t[0].fshr(1)
		s1[i]=2*int8(t[0].parity())-1
		t[4].fshr(1)
		s2[i]=2*int8(t[4].parity())-1

	}

// Recoded exponents
	for i:=0; i<nb; i++ {
		w1[i]=0
		k=1
		for j:=1; j<4; j++ {
			bt=s1[i]*int8(t[j].parity())
			t[j].fshr(1)
			t[j].dec(int(bt)>>1)
			t[j].norm()
			w1[i]+=bt*int8(k)
			k*=2
		}
		w2[i]=0
		k=1
		for j:=5; j<8; j++ {
			bt=s2[i]*int8(t[j].parity())
			t[j].fshr(1)
			t[j].dec(int(bt)>>1)
			t[j].norm()
			w2[i]+=bt*int8(k)
			k*=2
		}
	}
// Main loop
	p.selector(g1,int32(2*w1[nb-1]+1))  
	r.selector(g2,int32(2*w2[nb-1]+1)) 
	p.Mul(r)
	for i:=nb-2;i>=0;i-- {
		p.usqr()
		r.selector(g1,int32(2*w1[i]+s1[i]))
		p.Mul(r)
		r.selector(g2,int32(2*w2[i]+s2[i]))
		p.Mul(r)
	}

// apply correction
	r.Copy(q[0]); r.conj()   
	r.Mul(p)
	p.cmove(r,pb1)
	r.Copy(q[4]); r.conj()   
	r.Mul(p)
	p.cmove(r,pb2)

	p.reduce()
	return p;
}

/*
func pow4(q []*FP12,u []*BIG) *FP12 {
	var a [4]int8
	var g []*FP12
	var s []*FP12
	c:=NewFP12int(1)
	p:=NewFP12int(0)
	var w [NLEN*int(BASEBITS)+1]int8
	var t []*BIG
	mt:=NewBIGint(0)

	for i:=0;i<4;i++ {
		t=append(t,NewBIGcopy(u[i]))
	}

	s=append(s,NewFP12int(0))
	s=append(s,NewFP12int(0))

	g=append(g,NewFP12copy(q[0])); s[0].Copy(q[1]); s[0].conj(); g[0].Mul(s[0])
	g=append(g,NewFP12copy(g[0]))
	g=append(g,NewFP12copy(g[0]))
	g=append(g,NewFP12copy(g[0]))
	g=append(g,NewFP12copy(q[0])); g[4].Mul(q[1])
	g=append(g,NewFP12copy(g[4]))
	g=append(g,NewFP12copy(g[4]))
	g=append(g,NewFP12copy(g[4]))

	s[1].Copy(q[2]); s[0].Copy(q[3]); s[0].conj(); s[1].Mul(s[0])
	s[0].Copy(s[1]); s[0].conj(); g[1].Mul(s[0])
	g[2].Mul(s[1])
	g[5].Mul(s[0])
	g[6].Mul(s[1])
	s[1].Copy(q[2]); s[1].Mul(q[3])
	s[0].Copy(s[1]); s[0].conj(); g[0].Mul(s[0])
	g[3].Mul(s[1])
	g[4].Mul(s[0])
	g[7].Mul(s[1])

// if power is even add 1 to power, and add q to correction 

	for i:=0;i<4;i++ {
		if t[i].parity()==0 {
			t[i].inc(1); t[i].norm()
			c.Mul(q[i])
		}
		mt.add(t[i]); mt.norm()
	}
	c.conj()
	nb:=1+mt.nbits()

// convert exponent to signed 1-bit window 
	for j:=0;j<nb;j++ {
		for i:=0;i<4;i++ {
			a[i]=int8(t[i].lastbits(2)-2)
			t[i].dec(int(a[i])); t[i].norm();
			t[i].fshr(1)
		}
		w[j]=(8*a[0]+4*a[1]+2*a[2]+a[3])
	}
	w[nb]=int8(8*t[0].lastbits(2)+4*t[1].lastbits(2)+2*t[2].lastbits(2)+t[3].lastbits(2))
	p.Copy(g[(w[nb]-1)/2])

	for i:=nb-1;i>=0;i-- {
		m:=w[i]>>7
		j:=(w[i]^m)-m  // j=abs(w[i]) 
		j=(j-1)/2
		s[0].Copy(g[j]); s[1].Copy(g[j]); s[1].conj()
		p.usqr()
		p.Mul(s[m&1]);
	}
	p.Mul(c)  // apply correction 
	p.reduce()
	return p;
}
*/

