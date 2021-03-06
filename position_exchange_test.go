// Copyright (c) 2014-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package donna

import(`github.com/michaeldv/donna/expect`; `testing`)

func TestExchange000(t *testing.T) { // c4,d4 vs. d5,e6 (protected).
	p := NewGame(`Kg1,c4,d4`, `Kg8,d5,e6`).start()
	exchange := p.exchange(NewMove(p, C4, D5))
	expect.Eq(t, exchange, 0)
}

func TestExchange010(t *testing.T) { // c4,d4,e4 vs. c6,d5,e6 (protected).
	p := NewGame(`Kg1,Qb3,Nc3,a2,b2,c4,d4,e4,f2,g2,h2`, `Kg8,Qd8,Nf6,a7,b6,c6,d5,e6,f7,g7,h7`).start()
	exchange := p.exchange(NewMove(p, E4, D5))
	expect.Eq(t, exchange, 0)
}

func TestExchange020(t *testing.T) { // c4,d4,e4 vs. c6,d5,e6 (white wins a pawn).
	p := NewGame(`Kg1,Qb3,Nc3,Nf3,a2,b2,c4,d4,e4,f2,g2,h2`, `Kg8,Qd8,Nd7,Nf6,a7,b6,c6,d5,e6,f7,g7,h7`).start()
	exchange := p.exchange(NewMove(p, E4, D5))
	expect.Eq(t, exchange, valuePawn.midgame)
}
