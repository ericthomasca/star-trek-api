package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

var DS9Data = models.Series{
	Title: "Star Trek: Deep Space Nine",
	Seasons: []models.Season{
		{
			SeasonNumber: 1,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Emissary: Part I",
					AirDate:       strPtr("1993-01-03"),
					StarDate:      floatPtr(46379.1),
				},
				{
					EpisodeNumber: 2,
					Title:         "Emissary: Part II",
					AirDate:       strPtr("1993-01-03"),
					StarDate:      floatPtr(46392.7),
				},
				{
					EpisodeNumber: 3,
					Title:         "Past Prologue",
					AirDate:       strPtr("1993-01-10"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 4,
					Title:         "A Man Alone",
					AirDate:       strPtr("1993-01-17"),
					StarDate:      floatPtr(46421.5),
				},
				{
					EpisodeNumber: 5,
					Title:         "Babel",
					AirDate:       strPtr("1993-01-24"),
					StarDate:      floatPtr(46423.7),
				},
				{
					EpisodeNumber: 6,
					Title:         "Captive Pursuit",
					AirDate:       strPtr("1993-01-31"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 7,
					Title:         "Q-Less",
					AirDate:       strPtr("1993-02-07"),
					StarDate:      floatPtr(46531.2),
				},
				{
					EpisodeNumber: 8,
					Title:         "Dax",
					AirDate:       strPtr("1993-02-14"),
					StarDate:      floatPtr(46910.1),
				},
				{
					EpisodeNumber: 9,
					Title:         "The Passenger",
					AirDate:       strPtr("1993-02-21"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 10,
					Title:         "Move Along Home",
					AirDate:       strPtr("1993-03-14"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 11,
					Title:         "The Nagus",
					AirDate:       strPtr("1993-03-21"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 12,
					Title:         "Vortex",
					AirDate:       strPtr("1993-04-18"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 13,
					Title:         "Battle Lines",
					AirDate:       strPtr("1993-04-25"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 14,
					Title:         "The Storyteller",
					AirDate:       strPtr("1993-05-02"),
					StarDate:      floatPtr(46729.1),
				},
				{
					EpisodeNumber: 15,
					Title:         "Progress",
					AirDate:       strPtr("1993-05-09"),
					StarDate:      floatPtr(46844.3),
				},
				{
					EpisodeNumber: 16,
					Title:         "If Wishes Were Horses",
					AirDate:       strPtr("1993-05-16"),
					StarDate:      floatPtr(46853.2),
				},
				{
					EpisodeNumber: 17,
					Title:         "The Forsaken",
					AirDate:       strPtr("1993-05-23"),
					StarDate:      floatPtr(46925.1),
				},
				{
					EpisodeNumber: 18,
					Title:         "Dramatis Personae",
					AirDate:       strPtr("1993-05-30"),
					StarDate:      floatPtr(46922.3),
				},
				{
					EpisodeNumber: 19,
					Title:         "Duet",
					AirDate:       strPtr("1993-06-13"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 20,
					Title:         "In the Hands of the Prophets",
					AirDate:       strPtr("1993-06-20"),
					StarDate:      nil, // No StarDate
				},
			},
		},
		{
			SeasonNumber: 2,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "The Homecoming: Part 1",
					AirDate:       strPtr("1993-09-26"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 2,
					Title:         "The Circle: Part 2",
					AirDate:       strPtr("1993-10-03"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 3,
					Title:         "The Siege: Part 3",
					AirDate:       strPtr("1993-10-10"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 4,
					Title:         "Invasive Procedures",
					AirDate:       strPtr("1993-10-17"),
					StarDate:      floatPtr(47182.1),
				},
				{
					EpisodeNumber: 5,
					Title:         "Cardassians",
					AirDate:       strPtr("1993-10-24"),
					StarDate:      floatPtr(47177.2),
				},
				{
					EpisodeNumber: 6,
					Title:         "Melora",
					AirDate:       strPtr("1993-10-31"),
					StarDate:      floatPtr(47229.1),
				},
				{
					EpisodeNumber: 7,
					Title:         "Rules of Acquisition",
					AirDate:       strPtr("1993-11-07"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 8,
					Title:         "Necessary Evil",
					AirDate:       strPtr("1993-11-14"),
					StarDate:      floatPtr(47282.5),
				},
				{
					EpisodeNumber: 9,
					Title:         "Second Sight",
					AirDate:       strPtr("1993-11-21"),
					StarDate:      floatPtr(47329.4),
				},
				{
					EpisodeNumber: 10,
					Title:         "Sanctuary",
					AirDate:       strPtr("1993-11-28"),
					StarDate:      floatPtr(47391.2),
				},
				{
					EpisodeNumber: 11,
					Title:         "Rivals",
					AirDate:       strPtr("1994-01-02"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 12,
					Title:         "The Alternate",
					AirDate:       strPtr("1994-01-09"),
					StarDate:      floatPtr(47391.7),
				},
				{
					EpisodeNumber: 13,
					Title:         "Armageddon Game",
					AirDate:       strPtr("1994-01-30"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 14,
					Title:         "Whispers",
					AirDate:       strPtr("1994-02-06"),
					StarDate:      floatPtr(47552.1),
				},
				{
					EpisodeNumber: 15,
					Title:         "Paradise",
					AirDate:       strPtr("1994-02-13"),
					StarDate:      floatPtr(47573.1),
				},
				{
					EpisodeNumber: 16,
					Title:         "Shadowplay",
					AirDate:       strPtr("1994-02-20"),
					StarDate:      floatPtr(47603.3),
				},
				{
					EpisodeNumber: 17,
					Title:         "Playing God",
					AirDate:       strPtr("1994-02-27"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 18,
					Title:         "Profit and Loss",
					AirDate:       strPtr("1994-03-20"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 19,
					Title:         "Blood Oath",
					AirDate:       strPtr("1994-03-27"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 20,
					Title:         "The Maquis: Part 1",
					AirDate:       strPtr("1994-04-24"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 21,
					Title:         "The Maquis: Part 2",
					AirDate:       strPtr("1994-05-01"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 22,
					Title:         "The Wire",
					AirDate:       strPtr("1994-05-08"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 23,
					Title:         "Crossover",
					AirDate:       strPtr("1994-05-15"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 24,
					Title:         "The Collaborator",
					AirDate:       strPtr("1994-05-22"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 25,
					Title:         "Tribunal",
					AirDate:       strPtr("1994-06-05"),
					StarDate:      floatPtr(47944.2),
				},
				{
					EpisodeNumber: 26,
					Title:         "The Jem'Hadar",
					AirDate:       strPtr("1994-06-12"),
					StarDate:      nil, // No StarDate
				},
			},
		},
		{
			SeasonNumber: 3,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "The Search: Part I",
					AirDate:       strPtr("1994-09-26"),
					StarDate:      floatPtr(48213.1),
				},
				{
					EpisodeNumber: 2,
					Title:         "The Search: Part II",
					AirDate:       strPtr("1994-10-03"),
					StarDate:      floatPtr(48213.1),
				},
				{
					EpisodeNumber: 3,
					Title:         "The House of Quark",
					AirDate:       strPtr("1994-10-10"),
					StarDate:      floatPtr(48224.2),
				},
				{
					EpisodeNumber: 4,
					Title:         "Equilibrium",
					AirDate:       strPtr("1994-10-17"),
					StarDate:      floatPtr(48231.7),
				},
				{
					EpisodeNumber: 5,
					Title:         "Second Skin",
					AirDate:       strPtr("1994-10-24"),
					StarDate:      floatPtr(48244.5),
				},
				{
					EpisodeNumber: 6,
					Title:         "The Abandoned",
					AirDate:       strPtr("1994-10-31"),
					StarDate:      floatPtr(48301.1),
				},
				{
					EpisodeNumber: 7,
					Title:         "Civil Defense",
					AirDate:       strPtr("1994-11-07"),
					StarDate:      floatPtr(48388.8),
				},
				{
					EpisodeNumber: 8,
					Title:         "Meridian",
					AirDate:       strPtr("1994-11-14"),
					StarDate:      floatPtr(48423.2),
				},
				{
					EpisodeNumber: 9,
					Title:         "Defiant",
					AirDate:       strPtr("1994-11-21"),
					StarDate:      floatPtr(48467.3),
				},
				{
					EpisodeNumber: 10,
					Title:         "Fascination",
					AirDate:       strPtr("1994-11-28"),
					StarDate:      floatPtr(48441.6),
				},
				{
					EpisodeNumber: 11,
					Title:         "Past Tense: Part I",
					AirDate:       strPtr("1995-01-02"),
					StarDate:      floatPtr(48481.2),
				},
				{
					EpisodeNumber: 12,
					Title:         "Past Tense: Part II",
					AirDate:       strPtr("1995-01-09"),
					StarDate:      floatPtr(48481.2),
				},
				{
					EpisodeNumber: 13,
					Title:         "Life Support",
					AirDate:       strPtr("1995-01-31"),
					StarDate:      floatPtr(48498.4),
				},
				{
					EpisodeNumber: 14,
					Title:         "Heart of Stone",
					AirDate:       strPtr("1995-02-06"),
					StarDate:      floatPtr(48521.5),
				},
				{
					EpisodeNumber: 15,
					Title:         "Destiny",
					AirDate:       strPtr("1995-02-13"),
					StarDate:      floatPtr(48543.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "Prophet Motive",
					AirDate:       strPtr("1995-02-20"),
					StarDate:      floatPtr(48555.5),
				},
				{
					EpisodeNumber: 17,
					Title:         "Visionary",
					AirDate:       strPtr("1995-02-27"),
					StarDate:      floatPtr(48576.7),
				},
				{
					EpisodeNumber: 18,
					Title:         "Distant Voices",
					AirDate:       strPtr("1995-04-10"),
					StarDate:      floatPtr(48592.2),
				},
				{
					EpisodeNumber: 19,
					Title:         "Through the Looking Glass",
					AirDate:       strPtr("1995-04-17"),
					StarDate:      floatPtr(48601.1),
				},
				{
					EpisodeNumber: 20,
					Title:         "Improbable Cause: Part 1",
					AirDate:       strPtr("1995-04-24"),
					StarDate:      floatPtr(48620.3),
				},
				{
					EpisodeNumber: 21,
					Title:         "The Die is Cast: Part 2",
					AirDate:       strPtr("1995-05-01"),
					StarDate:      floatPtr(48622.5),
				},
				{
					EpisodeNumber: 22,
					Title:         "Explorers",
					AirDate:       strPtr("1995-05-08"),
					StarDate:      floatPtr(48699.9),
				},
				{
					EpisodeNumber: 23,
					Title:         "Family Business",
					AirDate:       strPtr("1995-05-15"),
					StarDate:      floatPtr(48731.2),
				},
				{
					EpisodeNumber: 24,
					Title:         "Shakaar",
					AirDate:       strPtr("1995-05-22"),
					StarDate:      floatPtr(48764.8),
				},
				{
					EpisodeNumber: 25,
					Title:         "Facets",
					AirDate:       strPtr("1995-06-12"),
					StarDate:      floatPtr(48959.1),
				},
				{
					EpisodeNumber: 26,
					Title:         "The Adversary",
					AirDate:       strPtr("1995-06-19"),
					StarDate:      floatPtr(48962.5),
				},
			},
		},
		{
			SeasonNumber: 4,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "The Way of the Warrior: Part I",
					AirDate:       strPtr("1995-10-02"),
					StarDate:      floatPtr(49011.4),
				},
				{
					EpisodeNumber: 2,
					Title:         "The Way of the Warrior: Part II",
					AirDate:       strPtr("1995-10-02"),
					StarDate:      floatPtr(49011.4),
				},
				{
					EpisodeNumber: 3,
					Title:         "The Visitor",
					AirDate:       strPtr("1995-10-09"),
					StarDate:      floatPtr(49037.7),
				},
				{
					EpisodeNumber: 4,
					Title:         "Hippocratic Oath",
					AirDate:       strPtr("1995-10-16"),
					StarDate:      floatPtr(49066.5),
				},
				{
					EpisodeNumber: 5,
					Title:         "Indiscretion",
					AirDate:       strPtr("1995-10-23"),
					StarDate:      floatPtr(49122.4),
				},
				{
					EpisodeNumber: 6,
					Title:         "Rejoined",
					AirDate:       strPtr("1995-10-30"),
					StarDate:      floatPtr(49195.5),
				},
				{
					EpisodeNumber: 7,
					Title:         "Starship Down",
					AirDate:       strPtr("1995-11-06"),
					StarDate:      floatPtr(49263.5),
				},
				{
					EpisodeNumber: 8,
					Title:         "Little Green Men",
					AirDate:       strPtr("1995-11-13"),
					StarDate:      floatPtr(49201.3),
				},
				{
					EpisodeNumber: 9,
					Title:         "The Sword of Kahless",
					AirDate:       strPtr("1995-11-20"),
					StarDate:      floatPtr(49263.5),
				},
				{
					EpisodeNumber: 10,
					Title:         "Our Man Bashir",
					AirDate:       strPtr("1995-11-27"),
					StarDate:      floatPtr(49300.7),
				},
				{
					EpisodeNumber: 11,
					Title:         "Homefront: Part I",
					AirDate:       strPtr("1996-01-01"),
					StarDate:      floatPtr(49170.0),
				},
				{
					EpisodeNumber: 12,
					Title:         "Paradise Lost: Part II",
					AirDate:       strPtr("1996-01-08"),
					StarDate:      floatPtr(49482.3),
				},
				{
					EpisodeNumber: 13,
					Title:         "Crossfire",
					AirDate:       strPtr("1996-01-29"),
					StarDate:      floatPtr(49517.3),
				},
				{
					EpisodeNumber: 14,
					Title:         "Return to Grace",
					AirDate:       strPtr("1996-02-05"),
					StarDate:      floatPtr(49534.2),
				},
				{
					EpisodeNumber: 15,
					Title:         "Sons of Mogh",
					AirDate:       strPtr("1996-02-12"),
					StarDate:      floatPtr(49556.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "Bar Association",
					AirDate:       strPtr("1996-02-19"),
					StarDate:      floatPtr(49565.1),
				},
				{
					EpisodeNumber: 17,
					Title:         "Accession",
					AirDate:       strPtr("1996-02-26"),
					StarDate:      floatPtr(49600.7),
				},
				{
					EpisodeNumber: 18,
					Title:         "Rules of Engagement",
					AirDate:       strPtr("1996-04-08"),
					StarDate:      floatPtr(49665.3),
				},
				{
					EpisodeNumber: 19,
					Title:         "Hard Time",
					AirDate:       strPtr("1996-04-15"),
					StarDate:      floatPtr(49680.5),
				},
				{
					EpisodeNumber: 20,
					Title:         "Shattered Mirror",
					AirDate:       strPtr("1996-04-22"),
					StarDate:      floatPtr(49699.1),
				},
				{
					EpisodeNumber: 21,
					Title:         "The Muse",
					AirDate:       strPtr("1996-04-29"),
					StarDate:      floatPtr(49702.2),
				},
				{
					EpisodeNumber: 22,
					Title:         "For the Cause",
					AirDate:       strPtr("1996-05-06"),
					StarDate:      floatPtr(49729.8),
				},
				{
					EpisodeNumber: 23,
					Title:         "To the Death",
					AirDate:       strPtr("1996-05-13"),
					StarDate:      floatPtr(49904.2),
				},
				{
					EpisodeNumber: 24,
					Title:         "The Quickening",
					AirDate:       strPtr("1996-05-20"),
					StarDate:      floatPtr(49909.7),
				},
				{
					EpisodeNumber: 25,
					Title:         "Body Parts",
					AirDate:       strPtr("1996-06-10"),
					StarDate:      floatPtr(49930.3),
				},
				{
					EpisodeNumber: 26,
					Title:         "Broken Link",
					AirDate:       strPtr("1996-06-17"),
					StarDate:      floatPtr(49962.4),
				},
			},
		},
		{
			SeasonNumber: 5,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Apocalypse Rising",
					AirDate:       strPtr("1996-09-30"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 2,
					Title:         "The Ship",
					AirDate:       strPtr("1996-10-07"),
					StarDate:      floatPtr(50049.3),
				},
				{
					EpisodeNumber: 3,
					Title:         "Looking for par'Mach in All the Wrong Places",
					AirDate:       strPtr("1996-10-14"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 4,
					Title:         "...Nor the Battle to the Strong",
					AirDate:       strPtr("1996-10-21"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 5,
					Title:         "The Assignment",
					AirDate:       strPtr("1996-10-28"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 6,
					Title:         "Trials and Tribble-ations",
					AirDate:       strPtr("1996-11-04"),
					StarDate:      floatPtr(4523.7),
				},
				{
					EpisodeNumber: 7,
					Title:         "Let He Who Is Without Sin...",
					AirDate:       strPtr("1996-11-11"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 8,
					Title:         "Things Past",
					AirDate:       strPtr("1996-11-18"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 9,
					Title:         "The Ascent",
					AirDate:       strPtr("1996-11-25"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 10,
					Title:         "Rapture",
					AirDate:       strPtr("1996-12-30"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 11,
					Title:         "The Darkness and the Light",
					AirDate:       strPtr("1997-01-06"),
					StarDate:      floatPtr(50416.2),
				},
				{
					EpisodeNumber: 12,
					Title:         "The Begotten",
					AirDate:       strPtr("1997-01-27"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 13,
					Title:         "For the Uniform",
					AirDate:       strPtr("1997-02-03"),
					StarDate:      floatPtr(50485.2),
				},
				{
					EpisodeNumber: 14,
					Title:         "In Purgatory's Shadow: Part I",
					AirDate:       strPtr("1997-02-10"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 15,
					Title:         "By Inferno's Light: Part II",
					AirDate:       strPtr("1997-02-17"),
					StarDate:      floatPtr(50564.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "Doctor Bashir, I Presume?",
					AirDate:       strPtr("1997-02-24"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 17,
					Title:         "A Simple Investigation",
					AirDate:       strPtr("1997-03-31"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 18,
					Title:         "Business as Usual",
					AirDate:       strPtr("1997-04-07"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 19,
					Title:         "Ties of Blood and Water",
					AirDate:       strPtr("1997-04-14"),
					StarDate:      floatPtr(50712.5),
				},
				{
					EpisodeNumber: 20,
					Title:         "Ferengi Love Songs",
					AirDate:       strPtr("1997-04-21"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 21,
					Title:         "Soldiers of the Empire",
					AirDate:       strPtr("1997-04-28"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 22,
					Title:         "Children of Time",
					AirDate:       strPtr("1997-05-05"),
					StarDate:      floatPtr(50814.2),
				},
				{
					EpisodeNumber: 23,
					Title:         "Blaze of Glory",
					AirDate:       strPtr("1997-05-12"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 24,
					Title:         "Empok Nor",
					AirDate:       strPtr("1997-05-19"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 25,
					Title:         "In the Cards",
					AirDate:       strPtr("1997-06-09"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 26,
					Title:         "Call to Arms",
					AirDate:       strPtr("1997-06-16"),
					StarDate:      floatPtr(50975.2),
				},
			},
		},
		{
			SeasonNumber: 6,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "A Time to Stand",
					AirDate:       strPtr("1997-09-29"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 2,
					Title:         "Rocks and Shoals",
					AirDate:       strPtr("1997-10-06"),
					StarDate:      floatPtr(51107.2),
				},
				{
					EpisodeNumber: 3,
					Title:         "Sons and Daughters",
					AirDate:       strPtr("1997-10-13"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 4,
					Title:         "Behind the Lines",
					AirDate:       strPtr("1997-10-20"),
					StarDate:      floatPtr(51149.5),
				},
				{
					EpisodeNumber: 5,
					Title:         "Favor the Bold: Part I",
					AirDate:       strPtr("1997-10-27"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 6,
					Title:         "Sacrifice of Angels: Part II",
					AirDate:       strPtr("1997-11-03"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 7,
					Title:         "You Are Cordially Invited...",
					AirDate:       strPtr("1997-11-10"),
					StarDate:      floatPtr(51247.5),
				},
				{
					EpisodeNumber: 8,
					Title:         "Resurrection",
					AirDate:       strPtr("1997-11-17"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 9,
					Title:         "Statistical Probabilities",
					AirDate:       strPtr("1997-11-24"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 10,
					Title:         "The Magnificent Ferengi",
					AirDate:       strPtr("1998-01-01"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 11,
					Title:         "Waltz",
					AirDate:       strPtr("1998-01-08"),
					StarDate:      floatPtr(51413.6),
				},
				{
					EpisodeNumber: 12,
					Title:         "Who Mourns for Morn?",
					AirDate:       strPtr("1998-02-04"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 13,
					Title:         "Far Beyond the Stars",
					AirDate:       strPtr("1998-02-11"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 14,
					Title:         "One Little Ship",
					AirDate:       strPtr("1998-02-18"),
					StarDate:      floatPtr(51474.2),
				},
				{
					EpisodeNumber: 15,
					Title:         "Honor Among Thieves",
					AirDate:       strPtr("1998-02-25"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 16,
					Title:         "Change of Heart",
					AirDate:       strPtr("1998-03-04"),
					StarDate:      floatPtr(51597.2),
				},
				{
					EpisodeNumber: 17,
					Title:         "Wrongs Darker Than Death or Night",
					AirDate:       strPtr("1998-04-01"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 18,
					Title:         "Inquisition",
					AirDate:       strPtr("1998-04-08"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 19,
					Title:         "In the Pale Moonlight",
					AirDate:       strPtr("1998-04-15"),
					StarDate:      floatPtr(51721.3),
				},
				{
					EpisodeNumber: 20,
					Title:         "His Way",
					AirDate:       strPtr("1998-04-22"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 21,
					Title:         "The Reckoning",
					AirDate:       strPtr("1998-04-29"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 22,
					Title:         "Valiant",
					AirDate:       strPtr("1998-05-06"),
					StarDate:      floatPtr(51825.4),
				},
				{
					EpisodeNumber: 23,
					Title:         "Profit and Lace",
					AirDate:       strPtr("1998-05-13"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 24,
					Title:         "Time's Orphan",
					AirDate:       strPtr("1998-05-20"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 25,
					Title:         "The Sound of Her Voice",
					AirDate:       strPtr("1998-06-10"),
					StarDate:      floatPtr(51948.3),
				},
				{
					EpisodeNumber: 26,
					Title:         "Tears of the Prophets",
					AirDate:       strPtr("1998-06-17"),
					StarDate:      nil, // No StarDate
				},
			},
		},
		{
			SeasonNumber: 7,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Image in the Sand",
					AirDate:       strPtr("1998-09-30"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 2,
					Title:         "Shadows and Symbols",
					AirDate:       strPtr("1998-10-07"),
					StarDate:      floatPtr(52152.6),
				},
				{
					EpisodeNumber: 3,
					Title:         "Afterimage",
					AirDate:       strPtr("1998-10-14"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 4,
					Title:         "Take Me Out to the Holosuite",
					AirDate:       strPtr("1998-10-21"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 5,
					Title:         "Chrysalis",
					AirDate:       strPtr("1998-10-28"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 6,
					Title:         "Treachery, Faith, and the Great River",
					AirDate:       strPtr("1998-11-04"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 7,
					Title:         "Once More Unto the Breach",
					AirDate:       strPtr("1998-11-11"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 8,
					Title:         "The Siege of AR-558",
					AirDate:       strPtr("1998-11-18"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 9,
					Title:         "Covenant",
					AirDate:       strPtr("1998-11-25"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 10,
					Title:         "It's Only a Paper Moon",
					AirDate:       strPtr("1998-12-30"),
					StarDate:      floatPtr(52235.7),
				},
				{
					EpisodeNumber: 11,
					Title:         "Prodigal Daughter",
					AirDate:       strPtr("1999-01-06"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 12,
					Title:         "The Emperor's New Cloak",
					AirDate:       strPtr("1999-02-03"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 13,
					Title:         "Field of Fire",
					AirDate:       strPtr("1999-02-10"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 14,
					Title:         "Chimera",
					AirDate:       strPtr("1999-02-17"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 15,
					Title:         "Badda-Bing Badda-Bang",
					AirDate:       strPtr("1999-02-24"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 16,
					Title:         "Inter Arma Enim Silent Leges",
					AirDate:       strPtr("1999-03-03"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 17,
					Title:         "Penumbra",
					AirDate:       strPtr("1999-04-07"),
					StarDate:      floatPtr(52576.2),
				},
				{
					EpisodeNumber: 18,
					Title:         "’Til Death Do Us Part",
					AirDate:       strPtr("1999-04-14"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 19,
					Title:         "Strange Bedfellows",
					AirDate:       strPtr("1999-04-21"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 20,
					Title:         "The Changing Face of Evil",
					AirDate:       strPtr("1999-04-28"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 21,
					Title:         "When It Rains…",
					AirDate:       strPtr("1999-05-05"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 22,
					Title:         "Tacking Into the Wind",
					AirDate:       strPtr("1999-05-12"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 23,
					Title:         "Extreme Measures",
					AirDate:       strPtr("1999-05-19"),
					StarDate:      floatPtr(52645.7),
				},
				{
					EpisodeNumber: 24,
					Title:         "The Dogs of War",
					AirDate:       strPtr("1999-05-26"),
					StarDate:      floatPtr(52861.3),
				},
				{
					EpisodeNumber: 25,
					Title:         "What You Leave Behind: Part I",
					AirDate:       strPtr("1999-06-02"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 26,
					Title:         "What You Leave Behind: Part II",
					AirDate:       strPtr("1999-06-02"),
					StarDate:      nil, // No StarDate
				},
			},
		},
	},
}
