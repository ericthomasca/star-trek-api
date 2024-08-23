package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

var VOYData = models.Series{
	Title: "Star Trek: Voyager",
	Seasons: []models.Season{
		{
			SeasonNumber: 1,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Caretaker",
					AirDate:       strPtr("1995-01-16"),
					StarDate:      floatPtr(48315.6),
				},
				{
					EpisodeNumber: 2,
					Title:         "Caretaker",
					AirDate:       strPtr("1995-01-16"),
					StarDate:      floatPtr(48315.6),
				},
				{
					EpisodeNumber: 3,
					Title:         "Parallax",
					AirDate:       strPtr("1995-01-23"),
					StarDate:      floatPtr(48439.7),
				},
				{
					EpisodeNumber: 4,
					Title:         "Time and Again",
					AirDate:       strPtr("1995-01-30"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 5,
					Title:         "Phage",
					AirDate:       strPtr("1995-02-06"),
					StarDate:      floatPtr(48532.4),
				},
				{
					EpisodeNumber: 6,
					Title:         "The Cloud",
					AirDate:       strPtr("1995-02-13"),
					StarDate:      floatPtr(48546.2),
				},
				{
					EpisodeNumber: 7,
					Title:         "Eye of the Needle",
					AirDate:       strPtr("1995-02-20"),
					StarDate:      floatPtr(48579.4),
				},
				{
					EpisodeNumber: 8,
					Title:         "Ex Post Facto",
					AirDate:       strPtr("1995-02-27"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 9,
					Title:         "Emanations",
					AirDate:       strPtr("1995-03-13"),
					StarDate:      floatPtr(48623.5),
				},
				{
					EpisodeNumber: 10,
					Title:         "Prime Factors",
					AirDate:       strPtr("1995-03-20"),
					StarDate:      floatPtr(48642.5),
				},
				{
					EpisodeNumber: 11,
					Title:         "State of Flux",
					AirDate:       strPtr("1995-04-10"),
					StarDate:      floatPtr(48658.2),
				},
				{
					EpisodeNumber: 12,
					Title:         "Heroes and Demons",
					AirDate:       strPtr("1995-04-24"),
					StarDate:      floatPtr(48693.2),
				},
				{
					EpisodeNumber: 13,
					Title:         "Cathexis",
					AirDate:       strPtr("1995-05-01"),
					StarDate:      floatPtr(48734.2),
				},
				{
					EpisodeNumber: 14,
					Title:         "Faces",
					AirDate:       strPtr("1995-05-08"),
					StarDate:      floatPtr(48784.2),
				},
				{
					EpisodeNumber: 15,
					Title:         "Jetrel",
					AirDate:       strPtr("1995-05-15"),
					StarDate:      floatPtr(48832.1),
				},
				{
					EpisodeNumber: 16,
					Title:         "Learning Curve",
					AirDate:       strPtr("1995-05-22"),
					StarDate:      floatPtr(48846.5),
				},
			},
		},
		{
			SeasonNumber: 2,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "The 37's",
					AirDate:       strPtr("1995-08-28"),
					StarDate:      floatPtr(48975.1),
				},
				{
					EpisodeNumber: 2,
					Title:         "Initiations",
					AirDate:       strPtr("1995-09-04"),
					StarDate:      floatPtr(49005.3),
				},
				{
					EpisodeNumber: 3,
					Title:         "Projections",
					AirDate:       strPtr("1995-09-11"),
					StarDate:      floatPtr(48892.1),
				},
				{
					EpisodeNumber: 4,
					Title:         "Elogium",
					AirDate:       strPtr("1995-09-18"),
					StarDate:      floatPtr(48921.3),
				},
				{
					EpisodeNumber: 5,
					Title:         "Non Sequitur",
					AirDate:       strPtr("1995-09-25"),
					StarDate:      floatPtr(49011.0),
				},
				{
					EpisodeNumber: 6,
					Title:         "Twisted",
					AirDate:       strPtr("1995-10-02"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 7,
					Title:         "Parturition",
					AirDate:       strPtr("1995-10-09"),
					StarDate:      floatPtr(49068.5),
				},
				{
					EpisodeNumber: 8,
					Title:         "Persistence of Vision",
					AirDate:       strPtr("1995-10-30"),
					StarDate:      floatPtr(49037.2),
				},
				{
					EpisodeNumber: 9,
					Title:         "Tattoo",
					AirDate:       strPtr("1995-11-06"),
					StarDate:      floatPtr(49211.5),
				},
				{
					EpisodeNumber: 10,
					Title:         "Cold Fire",
					AirDate:       strPtr("1995-11-13"),
					StarDate:      floatPtr(49164.8),
				},
				{
					EpisodeNumber: 11,
					Title:         "Maneuvers",
					AirDate:       strPtr("1995-11-20"),
					StarDate:      floatPtr(48423.0),
				},
				{
					EpisodeNumber: 12,
					Title:         "Resistance",
					AirDate:       strPtr("1995-11-27"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 13,
					Title:         "Prototype",
					AirDate:       strPtr("1996-01-15"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 14,
					Title:         "Alliances",
					AirDate:       strPtr("1996-01-22"),
					StarDate:      floatPtr(49337.4),
				},
				{
					EpisodeNumber: 15,
					Title:         "Threshold",
					AirDate:       strPtr("1996-01-29"),
					StarDate:      floatPtr(49373.4),
				},
				{
					EpisodeNumber: 16,
					Title:         "Meld",
					AirDate:       strPtr("1996-02-05"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 17,
					Title:         "Dreadnought",
					AirDate:       strPtr("1996-02-12"),
					StarDate:      floatPtr(49447.0),
				},
				{
					EpisodeNumber: 18,
					Title:         "Death Wish",
					AirDate:       strPtr("1996-02-19"),
					StarDate:      floatPtr(49301.2),
				},
				{
					EpisodeNumber: 19,
					Title:         "Lifesigns",
					AirDate:       strPtr("1996-02-26"),
					StarDate:      floatPtr(49504.3),
				},
				{
					EpisodeNumber: 20,
					Title:         "Investigations",
					AirDate:       strPtr("1996-03-13"),
					StarDate:      floatPtr(49485.2),
				},
				{
					EpisodeNumber: 21,
					Title:         "Deadlock",
					AirDate:       strPtr("1996-03-18"),
					StarDate:      floatPtr(49548.7),
				},
				{
					EpisodeNumber: 22,
					Title:         "Innocence",
					AirDate:       strPtr("1996-04-08"),
					StarDate:      floatPtr(49578.2),
				},
				{
					EpisodeNumber: 23,
					Title:         "The Thaw",
					AirDate:       strPtr("1996-04-29"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 24,
					Title:         "Tuvix",
					AirDate:       strPtr("1996-05-06"),
					StarDate:      floatPtr(49655.2),
				},
				{
					EpisodeNumber: 25,
					Title:         "Resolutions",
					AirDate:       strPtr("1996-05-13"),
					StarDate:      floatPtr(49690.1),
				},
				{
					EpisodeNumber: 26,
					Title:         "Basics, Part I",
					AirDate:       strPtr("1996-05-20"),
					StarDate:      nil, // Unknown
				},
			},
		},
		{
			SeasonNumber: 3,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Basics, Part II",
					AirDate:       strPtr("1996-09-04"),
					StarDate:      floatPtr(50023.4),
				},
				{
					EpisodeNumber: 2,
					Title:         "Flashback",
					AirDate:       strPtr("1996-09-11"),
					StarDate:      floatPtr(50126.4),
				},
				{
					EpisodeNumber: 3,
					Title:         "The Chute",
					AirDate:       strPtr("1996-09-18"),
					StarDate:      floatPtr(50156.2),
				},
				{
					EpisodeNumber: 4,
					Title:         "The Swarm",
					AirDate:       strPtr("1996-09-25"),
					StarDate:      floatPtr(50252.3),
				},
				{
					EpisodeNumber: 5,
					Title:         "False Profits",
					AirDate:       strPtr("1996-10-02"),
					StarDate:      floatPtr(50074.3),
				},
				{
					EpisodeNumber: 6,
					Title:         "Remember",
					AirDate:       strPtr("1996-10-09"),
					StarDate:      floatPtr(50203.1),
				},
				{
					EpisodeNumber: 7,
					Title:         "Sacred Ground",
					AirDate:       strPtr("1996-10-30"),
					StarDate:      floatPtr(50063.2),
				},
				{
					EpisodeNumber: 8,
					Title:         "Future's End, Part I",
					AirDate:       strPtr("1996-11-06"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 9,
					Title:         "Future's End, Part II",
					AirDate:       strPtr("1996-11-13"),
					StarDate:      floatPtr(50312.6),
				},
				{
					EpisodeNumber: 10,
					Title:         "Warlord",
					AirDate:       strPtr("1996-11-20"),
					StarDate:      floatPtr(50348.1),
				},
				{
					EpisodeNumber: 11,
					Title:         "The Q and the Grey",
					AirDate:       strPtr("1996-11-27"),
					StarDate:      floatPtr(50384.2),
				},
				{
					EpisodeNumber: 12,
					Title:         "Macrocosm",
					AirDate:       strPtr("1996-12-11"),
					StarDate:      floatPtr(50425.1),
				},
				{
					EpisodeNumber: 13,
					Title:         "Fair Trade",
					AirDate:       strPtr("1997-01-08"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 14,
					Title:         "Alter Ego",
					AirDate:       strPtr("1997-01-15"),
					StarDate:      floatPtr(50460.3),
				},
				{
					EpisodeNumber: 15,
					Title:         "Coda",
					AirDate:       strPtr("1997-01-29"),
					StarDate:      floatPtr(50518.6),
				},
				{
					EpisodeNumber: 16,
					Title:         "Blood Fever",
					AirDate:       strPtr("1997-02-05"),
					StarDate:      floatPtr(50537.2),
				},
				{
					EpisodeNumber: 17,
					Title:         "Unity",
					AirDate:       strPtr("1997-02-12"),
					StarDate:      floatPtr(50614.2),
				},
				{
					EpisodeNumber: 18,
					Title:         "Darkling",
					AirDate:       strPtr("1997-02-19"),
					StarDate:      floatPtr(50693.2),
				},
				{
					EpisodeNumber: 19,
					Title:         "Rise",
					AirDate:       strPtr("1997-02-26"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 20,
					Title:         "Favorite Son",
					AirDate:       strPtr("1997-03-19"),
					StarDate:      floatPtr(50732.4),
				},
				{
					EpisodeNumber: 21,
					Title:         "Before and After",
					AirDate:       strPtr("1997-04-09"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 22,
					Title:         "Real Life",
					AirDate:       strPtr("1997-04-23"),
					StarDate:      floatPtr(50836.2),
				},
				{
					EpisodeNumber: 23,
					Title:         "Distant Origin",
					AirDate:       strPtr("1997-04-30"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 24,
					Title:         "Displaced",
					AirDate:       strPtr("1997-05-07"),
					StarDate:      floatPtr(50912.4),
				},
				{
					EpisodeNumber: 25,
					Title:         "Worst Case Scenario",
					AirDate:       strPtr("1997-05-14"),
					StarDate:      floatPtr(50953.4),
				},
				{
					EpisodeNumber: 26,
					Title:         "Scorpion, Part I",
					AirDate:       strPtr("1997-05-21"),
					StarDate:      floatPtr(50984.3),
				},
			},
		},
		{
			SeasonNumber: 4,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Scorpion, Part II",
					AirDate:       strPtr("1997-09-03"),
					StarDate:      floatPtr(51003.7),
				},
				{
					EpisodeNumber: 2,
					Title:         "The Gift",
					AirDate:       strPtr("1997-09-10"),
					StarDate:      floatPtr(51008.0),
				},
				{
					EpisodeNumber: 3,
					Title:         "Day of Honor",
					AirDate:       strPtr("1997-09-17"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 4,
					Title:         "Nemesis",
					AirDate:       strPtr("1997-09-24"),
					StarDate:      floatPtr(51082.4),
				},
				{
					EpisodeNumber: 5,
					Title:         "Revulsion",
					AirDate:       strPtr("1997-10-01"),
					StarDate:      floatPtr(51186.2),
				},
				{
					EpisodeNumber: 6,
					Title:         "The Raven",
					AirDate:       strPtr("1997-10-08"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 7,
					Title:         "Scientific Method",
					AirDate:       strPtr("1997-10-29"),
					StarDate:      floatPtr(51244.3),
				},
				{
					EpisodeNumber: 8,
					Title:         "Year of Hell, Part I",
					AirDate:       strPtr("1997-11-05"),
					StarDate:      floatPtr(51268.4),
				},
				{
					EpisodeNumber: 9,
					Title:         "Year of Hell, Part II",
					AirDate:       strPtr("1997-11-12"),
					StarDate:      floatPtr(51425.4),
				},
				{
					EpisodeNumber: 10,
					Title:         "Random Thoughts",
					AirDate:       strPtr("1997-11-19"),
					StarDate:      floatPtr(51367.2),
				},
				{
					EpisodeNumber: 11,
					Title:         "Concerning Flight",
					AirDate:       strPtr("1997-11-26"),
					StarDate:      floatPtr(51386.4),
				},
				{
					EpisodeNumber: 12,
					Title:         "Mortal Coil",
					AirDate:       strPtr("1997-12-17"),
					StarDate:      floatPtr(51449.2),
				},
				{
					EpisodeNumber: 13,
					Title:         "Hunters",
					AirDate:       strPtr("1998-02-11"),
					StarDate:      floatPtr(51501.4),
				},
				{
					EpisodeNumber: 14,
					Title:         "Prey",
					AirDate:       strPtr("1998-02-18"),
					StarDate:      floatPtr(51652.3),
				},
				{
					EpisodeNumber: 15,
					Title:         "Retrospect",
					AirDate:       strPtr("1998-02-25"),
					StarDate:      floatPtr(51658.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "The Killing Game, Part I",
					AirDate:       strPtr("1998-03-04"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 17,
					Title:         "The Killing Game, Part II",
					AirDate:       strPtr("1998-03-04"),
					StarDate:      floatPtr(51715.2),
				},
				{
					EpisodeNumber: 18,
					Title:         "Vis à Vis",
					AirDate:       strPtr("1998-04-04"),
					StarDate:      floatPtr(51762.4),
				},
				{
					EpisodeNumber: 19,
					Title:         "The Omega Directive",
					AirDate:       strPtr("1998-04-15"),
					StarDate:      floatPtr(51781.2),
				},
				{
					EpisodeNumber: 20,
					Title:         "Unforgettable",
					AirDate:       strPtr("1998-04-22"),
					StarDate:      floatPtr(51813.4),
				},
				{
					EpisodeNumber: 21,
					Title:         "One",
					AirDate:       strPtr("1998-05-13"),
					StarDate:      floatPtr(51929.3),
				},
				{
					EpisodeNumber: 22,
					Title:         "Living Witness",
					AirDate:       strPtr("1998-04-29"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 23,
					Title:         "Demon",
					AirDate:       strPtr("1998-05-06"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 24,
					Title:         "Hope and Fear",
					AirDate:       strPtr("1998-05-20"),
					StarDate:      floatPtr(51978.2),
				},
				{
					EpisodeNumber: 25,
					Title:         "The Gift",
					AirDate:       strPtr("1997-09-10"),
					StarDate:      floatPtr(51008.0),
				},
				{
					EpisodeNumber: 26,
					Title:         "Year of Hell, Part II",
					AirDate:       strPtr("1997-11-12"),
					StarDate:      floatPtr(51425.4),
				},
			},
		},
		{
			SeasonNumber: 5,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Night",
					AirDate:       strPtr("1998-10-14"),
					StarDate:      floatPtr(52081.2),
				},
				{
					EpisodeNumber: 2,
					Title:         "Drone",
					AirDate:       strPtr("1998-10-21"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 3,
					Title:         "Extreme Risk",
					AirDate:       strPtr("1998-10-28"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 4,
					Title:         "In the Flesh",
					AirDate:       strPtr("1998-11-04"),
					StarDate:      floatPtr(52136.4),
				},
				{
					EpisodeNumber: 5,
					Title:         "Once Upon a Time",
					AirDate:       strPtr("1998-11-11"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 6,
					Title:         "Timeless",
					AirDate:       strPtr("1998-11-18"),
					StarDate:      floatPtr(52143.6),
				},
				{
					EpisodeNumber: 7,
					Title:         "Infinite Regress",
					AirDate:       strPtr("1998-11-25"),
					StarDate:      floatPtr(52188.7),
				},
				{
					EpisodeNumber: 8,
					Title:         "Nothing Human",
					AirDate:       strPtr("1998-12-02"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 9,
					Title:         "Thirty Days",
					AirDate:       strPtr("1998-12-09"),
					StarDate:      floatPtr(52179.4),
				},
				{
					EpisodeNumber: 10,
					Title:         "Counterpoint",
					AirDate:       strPtr("1998-12-16"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 11,
					Title:         "Latent Image",
					AirDate:       strPtr("1999-01-20"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 12,
					Title:         "Bride of Chaotica!",
					AirDate:       strPtr("1999-01-27"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 13,
					Title:         "Gravity",
					AirDate:       strPtr("1999-02-03"),
					StarDate:      floatPtr(52438.9),
				},
				{
					EpisodeNumber: 14,
					Title:         "Bliss",
					AirDate:       strPtr("1999-02-10"),
					StarDate:      floatPtr(52542.3),
				},
				{
					EpisodeNumber: 15,
					Title:         "Dark Frontier, Part I",
					AirDate:       strPtr("1999-02-17"),
					StarDate:      floatPtr(52619.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "Dark Frontier, Part II",
					AirDate:       strPtr("1999-02-17"),
					StarDate:      floatPtr(52619.2),
				},
				{
					EpisodeNumber: 17,
					Title:         "The Disease",
					AirDate:       strPtr("1999-02-24"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 18,
					Title:         "Course: Oblivion",
					AirDate:       strPtr("1999-03-03"),
					StarDate:      floatPtr(52586.3),
				},
				{
					EpisodeNumber: 19,
					Title:         "The Fight",
					AirDate:       strPtr("1999-03-24"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 20,
					Title:         "Think Tank",
					AirDate:       strPtr("1999-03-31"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 21,
					Title:         "Juggernaut",
					AirDate:       strPtr("1999-04-26"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 22,
					Title:         "Someone to Watch Over Me",
					AirDate:       strPtr("1999-04-28"),
					StarDate:      floatPtr(52647.0),
				},
				{
					EpisodeNumber: 23,
					Title:         "11:59",
					AirDate:       strPtr("1999-05-05"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 24,
					Title:         "Relativity",
					AirDate:       strPtr("1999-05-12"),
					StarDate:      floatPtr(52861.274),
				},
				{
					EpisodeNumber: 25,
					Title:         "Warhead",
					AirDate:       strPtr("1999-05-19"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 26,
					Title:         "Equinox, Part I",
					AirDate:       strPtr("1999-05-26"),
					StarDate:      nil, // Unknown
				},
			},
		},
		{
			SeasonNumber: 6,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Equinox, Part II",
					AirDate:       strPtr("1999-09-22"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 2,
					Title:         "Survival Instinct",
					AirDate:       strPtr("1999-09-29"),
					StarDate:      floatPtr(53049.2),
				},
				{
					EpisodeNumber: 3,
					Title:         "Barge of the Dead",
					AirDate:       strPtr("1999-10-06"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 4,
					Title:         "Tinker, Tenor, Doctor, Spy",
					AirDate:       strPtr("1999-10-13"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 5,
					Title:         "Alice",
					AirDate:       strPtr("1999-10-20"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 6,
					Title:         "Riddles",
					AirDate:       strPtr("1999-11-03"),
					StarDate:      floatPtr(53263.2),
				},
				{
					EpisodeNumber: 7,
					Title:         "Dragon's Teeth",
					AirDate:       strPtr("1999-11-10"),
					StarDate:      floatPtr(53167.9),
				},
				{
					EpisodeNumber: 8,
					Title:         "One Small Step",
					AirDate:       strPtr("1999-11-17"),
					StarDate:      floatPtr(53292.7),
				},
				{
					EpisodeNumber: 9,
					Title:         "The Voyager Conspiracy",
					AirDate:       strPtr("1999-11-24"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 10,
					Title:         "Pathfinder",
					AirDate:       strPtr("1999-12-01"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 11,
					Title:         "Fair Haven",
					AirDate:       strPtr("2000-01-12"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 12,
					Title:         "Blink of an Eye",
					AirDate:       strPtr("2000-01-19"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 13,
					Title:         "Virtuoso",
					AirDate:       strPtr("2000-01-26"),
					StarDate:      floatPtr(53556.4),
				},
				{
					EpisodeNumber: 14,
					Title:         "Memorial",
					AirDate:       strPtr("2000-02-02"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 15,
					Title:         "Tsunkatse",
					AirDate:       strPtr("2000-02-09"),
					StarDate:      floatPtr(53447.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "Collective",
					AirDate:       strPtr("2000-02-16"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 17,
					Title:         "Spirit Folk",
					AirDate:       strPtr("2000-02-23"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 18,
					Title:         "Ashes to Ashes",
					AirDate:       strPtr("2000-03-01"),
					StarDate:      floatPtr(53679.4),
				},
				{
					EpisodeNumber: 19,
					Title:         "Child's Play",
					AirDate:       strPtr("2000-03-08"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 20,
					Title:         "Good Shepherd",
					AirDate:       strPtr("2000-03-15"),
					StarDate:      floatPtr(53753.2),
				},
				{
					EpisodeNumber: 21,
					Title:         "Live Fast and Prosper",
					AirDate:       strPtr("2000-04-19"),
					StarDate:      floatPtr(53849.2),
				},
				{
					EpisodeNumber: 22,
					Title:         "Muse",
					AirDate:       strPtr("2000-04-26"),
					StarDate:      floatPtr(53896.0),
				},
				{
					EpisodeNumber: 23,
					Title:         "Fury",
					AirDate:       strPtr("2000-05-03"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 24,
					Title:         "Life Line",
					AirDate:       strPtr("2000-05-10"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 25,
					Title:         "The Haunting of Deck Twelve",
					AirDate:       strPtr("2000-05-17"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 26,
					Title:         "Unimatrix Zero, Part I",
					AirDate:       strPtr("2000-05-24"),
					StarDate:      nil, // Unknown
				},
			},
		},
		{
			SeasonNumber: 7,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Unimatrix Zero, Part II",
					AirDate:       strPtr("2000-10-04"),
					StarDate:      floatPtr(54014.4),
				},
				{
					EpisodeNumber: 2,
					Title:         "Imperfection",
					AirDate:       strPtr("2000-10-11"),
					StarDate:      floatPtr(54058.6),
				},
				{
					EpisodeNumber: 3,
					Title:         "Drive",
					AirDate:       strPtr("2000-10-18"),
					StarDate:      floatPtr(54090.4),
				},
				{
					EpisodeNumber: 4,
					Title:         "Repression",
					AirDate:       strPtr("2000-10-25"),
					StarDate:      floatPtr(54129.4),
				},
				{
					EpisodeNumber: 5,
					Title:         "Critical Care",
					AirDate:       strPtr("2000-11-01"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 6,
					Title:         "Inside Man",
					AirDate:       strPtr("2000-11-08"),
					StarDate:      floatPtr(54208.3),
				},
				{
					EpisodeNumber: 7,
					Title:         "Body and Soul",
					AirDate:       strPtr("2000-11-15"),
					StarDate:      floatPtr(54238.3),
				},
				{
					EpisodeNumber: 8,
					Title:         "Nightingale",
					AirDate:       strPtr("2000-11-22"),
					StarDate:      floatPtr(54274.7),
				},
				{
					EpisodeNumber: 9,
					Title:         "Flesh and Blood, Part I",
					AirDate:       strPtr("2000-11-29"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 10,
					Title:         "Flesh and Blood, Part II",
					AirDate:       strPtr("2000-11-29"),
					StarDate:      floatPtr(54337.5),
				},
				{
					EpisodeNumber: 11,
					Title:         "Shattered",
					AirDate:       strPtr("2001-01-17"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 12,
					Title:         "Lineage",
					AirDate:       strPtr("2001-01-24"),
					StarDate:      floatPtr(54452.6),
				},
				{
					EpisodeNumber: 13,
					Title:         "Repentance",
					AirDate:       strPtr("2001-01-31"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 14,
					Title:         "Prophecy",
					AirDate:       strPtr("2001-02-07"),
					StarDate:      floatPtr(54518.2),
				},
				{
					EpisodeNumber: 15,
					Title:         "The Void",
					AirDate:       strPtr("2001-02-14"),
					StarDate:      floatPtr(54553.4),
				},
				{
					EpisodeNumber: 16,
					Title:         "Workforce, Part I",
					AirDate:       strPtr("2001-02-21"),
					StarDate:      floatPtr(54584.3),
				},
				{
					EpisodeNumber: 17,
					Title:         "Workforce, Part II",
					AirDate:       strPtr("2001-02-28"),
					StarDate:      floatPtr(54622.4),
				},
				{
					EpisodeNumber: 18,
					Title:         "Human Error",
					AirDate:       strPtr("2001-03-07"),
					StarDate:      nil, // Unknown
				},
				{
					EpisodeNumber: 19,
					Title:         "Q2",
					AirDate:       strPtr("2001-04-11"),
					StarDate:      floatPtr(54704.5),
				},
				{
					EpisodeNumber: 20,
					Title:         "Author, Author",
					AirDate:       strPtr("2001-04-18"),
					StarDate:      floatPtr(54732.3),
				},
				{
					EpisodeNumber: 21,
					Title:         "Friendship One",
					AirDate:       strPtr("2001-04-25"),
					StarDate:      floatPtr(54775.4),
				},
				{
					EpisodeNumber: 22,
					Title:         "Natural Law",
					AirDate:       strPtr("2001-05-02"),
					StarDate:      floatPtr(54827.7),
				},
				{
					EpisodeNumber: 23,
					Title:         "Homestead",
					AirDate:       strPtr("2001-05-09"),
					StarDate:      floatPtr(54868.6),
				},
				{
					EpisodeNumber: 24,
					Title:         "Renaissance Man",
					AirDate:       strPtr("2001-05-16"),
					StarDate:      floatPtr(54890.7),
				},
				{
					EpisodeNumber: 25,
					Title:         "Endgame",
					AirDate:       strPtr("2001-05-23"),
					StarDate:      floatPtr(54973.4),
				},
				{
					EpisodeNumber: 26,
					Title:         "Endgame",
					AirDate:       strPtr("2001-05-23"),
					StarDate:      floatPtr(54973.4),
				},
			},
		},
	},
}
