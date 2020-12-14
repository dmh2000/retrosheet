// prettier-ignore
const glfields = [
                                // Field(s)  Meaning
"date",                         //    1     Date in the form"yyyymmdd"
"gameNumber",                   //    2     Number of game:
                                //   "0" -- a single game
                                //   "1" -- the first game of a double (or triple) header
                                //                     including seperate admission doubleheaders
                                //   "2" -- the second game of a double (or triple) header
                                //                     including seperate admission doubleheaders
                                //   "3" -- the third game of a triple-header
                                //   "A" -- the first game of a double-header involving 3 teams
                                //   "B" -- the second game of a double-header involving 3 teams
"dayOfWeek",                    //            Day of week  ("Sun","Mon","Tue","Wed","Thu","Fri","Sat")
"visitorTeam",                  //     4     Visiting team and league
"visitorLeague",                //     5
"visitorGameNumber",            //     6     Visiting team game number
                                //           For this and the home team game number, ties are counted as
                                //           games and suspended games are counted from the starting
                                //           rather than the ending date.
"homeTeam",                     //     7     Home team and league
"homeLeague",                   //     8
"homeGameNumber",               //     9     Home team game number
"visitorScore",                 //    10     Visiting and home team score (unquoted)
"homeScore",                    //    11
"lengthInOuts",                 //    12     Length of game in outs (unquoted).  A full 9-inning game would
                                //           have a 54 in this field.  If the home team won without batting
                                //           in the bottom of the ninth, this field would contain a 51.
"dayOrNight",                   //    13     Day/night indicator ("D" or"N")
"completion",                   //    14     Completion information.  If the game was completed at a
                                //           later date (either due to a suspension or an upheld protest)
                                //           this field will include: 
                                //   "yyyymmdd,park,vs,hs,len" Where
                                //           yyyymmdd -- the date the game was completed
                                //           park -- the park ID where the game was completed
                                //           vs -- the visitor score at the time of interruption
                                //           hs -- the home score at the time of interruption
                                //           len -- the length of the game in outs at time of interruption
                                //           All the rest of the information in the record refers to the
                                //           entire game.
"forfeit",                      //    15     Forfeit information:
                                //   "V" -- the game was forfeited to the visiting team
                                //   "H" -- the game was forfeited to the home team
                                //   "T" -- the game was ruled a no-decision
"protest",                      //    16     Protest information:
                                //   "P" -- the game was protested by an unidentified team
                                //   "V" -- a disallowed protest was made by the visiting team
                                //   "H" -- a disallowed protest was made by the home team
                                //   "X" -- an upheld protest was made by the visiting team
                                //   "Y" -- an upheld protest was made by the home team
                                //           Note: two of these last four codes can appear in the field
                                //           (if both teams protested the game).
"parkId",                       //    17     Park ID
"attendance",                   //    18     Attendance (unquoted)
"duration",                     //    19     Time of game in minutes (unquoted)
"visitorLine",                  //    20       Visiting and home line scores.  For example:
"homeLine",                     //    21
                                //   "010000(10)0x"
                                //           Would indicate a game where the home team scored a run in
                                //           the second inning, ten in the seventh and didn't bat in the
                                //           bottom of the ninth.
// VISITOR STATS                                
                                // 22-38    Visiting team offensive statistics (unquoted) (in order):
"visitorAtBats",                // 22             at-bats
"visitorHits",                  // 23             hits
"visitorDbls",                  // 24             doubles
"visitorTriples",               // 25             triples
"visitorHomers",                // 26             homeruns
"visitorRBI",                   // 27             RBI
"visitorSacHits",               // 29             sacrifice hits.  This may include sacrifice flies for years
                                //                 prior to 1954 when sacrifice flies were allowed.
"visitorSacFlies",              // 29             sacrifice flies (since 1954)
"visitorHitByPitch",            // 30             hit-by-pitch
"visitorWalks",                 // 31             walks
"visitorIntentionalWalks",      // 32             intentional walks
"visitorStrikeouts",            // 33             strikeouts
"visitorSteals",                // 34             stolen bases
"visitorStealCaught",           // 35             caught stealing
"visitorGrounded",              // 36             grounded into double plays
"visitorCatcherIfc",            // 37             awarded first on catcher's interference
"visitorLeftOnBase",            // 39             left on base
                                // 39-43     Visiting team pitching statistics (unquoted)(in order):
"visitorPitchers",              // 39            pitchers used ( 1 means it was a complete game )
"visitorIndiERs",               // 40             individual earned runs
"visitorTeamERs",               // 41             team earned runs
"visitorWilds",                 //  42            wild pitches
"visitorBalks",                 // 43             balks
                                // 44-49    Visiting team defensive statistics (unquoted) (in order):
"visitorPutouts",               // 44            putouts.  Note: prior to 1931, this may not equal 3 times
                                //                 the number of innings pitched.  Prior to that, no
                                //                 putout was awarded when a runner was declared out for
                                //                 being hit by a batted ball.
"visitorAssists",               //  45           assists
"visitorErrors",                //  46            errors
"visitorPassedBalls",           //  47            passed balls
"visitorDoublePlays",           //  48            double plays
"visitorTriplePlays",           //  49            triple plays

// HOME STATS
                                // 50-66    Home team offensive statistics
                                // 67-71    Home team pitching statistics
                                // 72-77    Home team defensive statistics
                                // 78-79    Home plate umpire ID and name
"homeAtBats",                   // 50             at-bats
"homeHits",                     // 51            hits
"homeDbls",                     // 52             doubles
"homeTriples",                  // 53            triples
"homeHomers",                   // 54             homeruns
"homeRBI",                      // 55             RBI
"homeSacHits",                  // 56             sacrifice hits.  This may include sacrifice flies for years
                                //                 prior to 1954 when sacrifice flies were allowed.
"homeSacFlies",                 // 57             sacrifice flies (since 1954)
"homeHitByPitch",               // 58             hit-by-pitch
"homeWalks",                    // 59             walks
"homeIntentionalWalks",         // 60             intentional walks
"homeStrikeouts",               // 61             strikeouts
"homeSteals",                   // 62             stolen bases
"homeStealCaught",              // 63             caught stealing
"homeGrounded",                 // 64             grounded into double plays
"homeCatcherIfc",               // 65             awarded first on catcher's interference
"homeLeftOnBase",               // 66             left on base
                                //          home team pitching statistics (unquoted)(in order):
"homePitchers",                 // 67            pitchers used ( 1 means it was a complete game )
"homeIndiERs",                  // 69             individual earned runs
"homeTeamERs",                  // 69             team earned runs
"homeWilds",                    // 70            wild pitches
"homeBalks",                    // 71             balks
                                //          home team defensive statistics (unquoted) (in order):
"homePutouts",                  // 72            putouts.  Note: prior to 1931, this may not equal 3 times
                                //                 the number of innings pitched.  Prior to that, no
                                //                 putout was awarded when a runner was declared out for
                                //                 being hit by a batted ball.
"homeAssists",                  //  73           assists
"homeErrors",                   //  74            errors
"homePassedBalls",              //  75            passed balls
"homedoublePlays",              //  76            double plays
"homeTriplePlays",              //  77            triple plays

// UMPIRES
"umpireHomeId",                 // 78
"umpireHomeName",               // 79
"umpFirstBaseId",               // 80       1B umpire ID and name
"umpFirstBaseName",             // 81
"umpSecondBaseId",              // 82       2B umpire ID and name
"umpSecondBaseName",            // 83
"umpThirdBaseId",               // 84       3B umpire ID and name
"umpThirdBaseName",             // 85 
"umpLeftFieldId",               // 86        LF umpire ID and name
"umpLeftFieldName",             // 86
"umpRightFieldId",              // 88        RF umpire ID and name
"umpRightFieldName",            // 89
                                //           If any umpire positions were not filled for a particular game
                                //           the fields will be"","(none)".
// MANAGERS                                
"visitorManagerId",             // 90        Visiting team manager ID and name
"visitorManagerName",           // 91
"homeManagerId",                // 92        Home team manager ID and name
"homeMangerName",               // 93
// PITCHERS
"winningPitcherId",             // 94        Winning pitcher ID and name
"winningPitcherName",           // 95
"losingPitcherId",              // 96        Losing pitcher ID and name
"losingPitcherName",            // 97
"savingPitcherId",              // 98        Saving pitcher ID and name--"","(none)" if none awarded
"savingPitcherName",            // 99
"rbiBatterId",                  // 100       Game Winning RBI batter ID and name--"","(none)" if none
"rbiBatterName",                // 101
"visitorStarterId",             // 102       Visiting starting pitcher ID and name
"visitorStarterName",           // 103
"homeStarterId",                // 104       Home starting pitcher ID and name
"homeStarterName",              // 105

// 106-132   Visiting starting players ID, name and defensive position,
//           listed in the order (1-9) they appeared in the batting order.
"visitorLineup1Id",             // 106
"visitorLineup1Name",           // 107
"visitorLineup1Pos",            // 108
"visitorLineup2Id",             // 109
"visitorLineup2Name",           // 110
"visitorLineup2Pos",            // 111
"visitorLineup3Id",             // 112
"visitorLineup3Name",           // 113
"visitorLineup3Pos",            // 114
"visitorLineup4Id",             // 115
"visitorLineup4Name",           // 116
"visitorLineup4Pos",            // 117
"visitorLineup5Id",             // 118
"visitorLineup5Name",           // 119
"visitorLineup5Pos",            // 120
"visitorLineup6Id",             // 121
"visitorLineup6Name",           // 122
"visitorLineup6Pos",            // 123
"visitorLineup7Id",             // 124
"visitorLineup7Name",           // 125
"visitorLineup7Pos",            // 126
"visitorLineup8Id",             // 127
"visitorLineup8Name",           // 128
"visitorLineup8Pos",            // 129
"visitorLineup9Id",             // 130
"visitorLineup9Name",           // 131
"visitorLineup9Pos",            // 132

// 133-159   Home starting players ID, name and defensive position
//           listed in the order (1-9) they appeared in the batting order.
"homeLineup1Id",                // 133
"homeLineup1Name",              // 134
"homeLineup1Pos",               // 135
"homeLineup2Id",                // 136
"homeLineup2Name",              // 137
"homeLineup2Pos",               // 138
"homeLineup3Id",                // 139
"homeLineup3Name",              // 140
"homeLineup3Pos",               // 141
"homeLineup4Id",                // 142
"homeLineup4Name",              // 143
"homeLineup4Pos",               // 144
"homeLineup5Id",                // 145
"homeLineup5Name",              // 146
"homeLineup5Pos",               // 147
"homeLineup6Id",                // 148
"homeLineup6Name",              // 149
"homeLineup6Pos",               // 150
"homeLineup7Id",                // 151
"homeLineup7Name",              // 152
"homeLineup7Pos",               // 153
"homeLineup8Id",                // 154
"homeLineup8Name",              // 155
"homeLineup8Pos",               // 156
"homeLineup9Id",                // 157
"homeLineup9Name",              // 158
"homeLineup9Pos",               // 159

"additionalInfo",               //   160     Additional information.  This is a grab-bag of informational
                                //           items that might not warrant a field on their own.  The field 
                                //           is alpha-numeric. Some items are represented by tokens such as:
                                //   "HTBF" -- home team batted first.
                                //              Note: if"HTBF" is specified it would be possible to see
                                //              something like"01002000x" in the visitor's line score.
                                //           Changes in umpire positions during a game will also appear in 
                                //           this field.  These will be in the form:
                                //              umpchange,inning,umpPosition,umpid with the latter three
                                //              repeated for each umpire.
                                //           These changes occur with umpire injuries, late arrival of 
                                //           umpires or changes from completion of suspended games. Details
                                //           of suspended games are in field 14.
  
"acquisitionInfo",              //   161     Acquisition information:
                                //   "Y" -- we have the complete game
                                //   "N" -- we don't have any portion of the game
                                //   "D" -- the game was derived from box score and game story
                                //   "P" -- we have some portion of the game.  We may be missing
                                //                     innings at the beginning, middle and end of the game.
                                
                                // Missing fields will be NULL.
                                //
]

// matching types for each of those fields
const gltypes = [
  "string","string","string","string","string","number","string","string","number","number","number","number","string","string","string","string","string",
  "number","number","string","string","number","number","number","number","number","number","number","number","number","number","number","number","number",
  "number","number","number","number","number","number","number","number","number","number","number","number","number","number","number","number","number",
  "number","number","number","number","number","number","number","number","number","number","number","number","number","number","number","number","number",
  "number","number","number","number","number","number","number","number","number","string","string","string","string","string","string","string","string",
  "string","string","string","string","string","string","string","string","string","string","string","string","string","string","string","string","string",
  "string","string","string","string","string","number","string","string","number","string","string","number","string","string","number","string","string",
  "number","string","string","number","string","string","number","string","string","number","string","string","number","string","string","number","string",
  "string","number","string","string","number","string","string","number","string","string","number","string","string","number","string","string","number",
];
module.exports = {glfields,gltypes}
