package api

type Team struct {
	abbr      string
	league    string
	city      string
	nickname  string
	firstYear string
	lastYear  string
}

type teamResolver struct {
	t *Team
}

func (r *teamResolver) Abbr() string {
	return r.t.abbr
}

func (r *teamResolver) League() string {
	return r.t.league
}

func (r *teamResolver) City() string {
	return r.t.city
}

func (r *teamResolver) Nickname() string {
	return r.t.nickname
}

func (r *teamResolver) FirstYear() string {
	return r.t.firstYear
}

func (r *teamResolver) LastYear() string {
	return r.t.lastYear
}

// get team by abbreviation
func (*rootResolver) TeamByAbbr(args struct{ Abbr string }) *teamResolver {
	// query the db for the team by the abbreviation
	return &teamResolver{&Team{args.Abbr, "MLB", "Cleveland", "Indians", "1993", "1996"}}
}

// get team by name
func (*rootResolver) TeamByName(args struct{ Name string }) *teamResolver {
	// query the db for the team by the abbreviation
	return &teamResolver{&Team{"WAS", "MLB", "Cleveland", args.Name, "1993", "1996"}}
}
