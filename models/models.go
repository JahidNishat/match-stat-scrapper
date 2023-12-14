package models

type EspnStat struct {
	Date   string
	Opp    string
	Result string
	Min    string
	FG     string
	FT     string
	ThreeP string
	REB    string
	AST    string
	BLK    string
	STL    string
	PF     string
	TO     string
	PTS    string
}

type EuroBasketStat struct {
	Date        string
	Team        string
	AgainstTeam string
	Result      string
	Min         string
	Pts         string
	TwoFGP      string
	ThreeFGP    string
	FT          string
	RO          string
	RD          string
	RT          string
	AS          string
	PF          string
	BS          string
	ST          string
	TO          string
	RNK         string
}

type BnxtStat struct {
	GameDate string
	Game     string
	Result   string
	PTS      string
	Min      string
	TwoP     string
	ThreeP   string
	FgP      string
	FtP      string
	Dr       string
	Or       string
	Tot      string
	Fp       string
	Df       string
	Ast      string
	St       string
	To       string
	Bs       string
	Br       string
	Eff      string
}

type NblStat struct {
	Date string
	Opp  string
	Min  string
	FgP  string
	FtP  string
	Reb  string
	Ast  string
	Blk  string
	Stl  string
	To   string
	Pf   string
	Pts  string
}

type BleagueStat struct {
	Day       string
	VS        string
	HA        string
	WL        string
	Min       string
	Pts       string
	FgP       string
	TwoFgP    string
	ThreeFgP  string
	FtP       string
	EfgP      string
	TsP       string
	Or        string
	Dr        string
	Tr        string
	As        string
	Ast       string
	To        string
	St        string
	Bs        string
	Bsr       string
	F         string
	Fd        string
	Eff       string
	PlusMinus string
}

type BritishBasketBallStat struct {
	Team    string
	Date    string
	Min     string
	FgP     string
	ThreePP string
	FtP     string
	Off     string
	Def     string
	Reb     string
	Ast     string
	Stl     string
	Blk     string
	Pf      string
	To      string
	Pts     string
}

type B3leagueStat struct {
	Date   string
	Pts    string
	ThreeP string
	TwoP   string
	FtP    string
	Pf     string
	Oreb   string
	Dreb   string
	Reb    string
	Tov    string
	Ast    string
	Stl    string
	Blk    string
	Eff    string
	Min    string
}
