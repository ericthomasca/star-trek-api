package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

// TOSData is the data for `Star Trek: The Original Series`.
var TOSData = models.Series{
	Title: "Star Trek: The Original Series",
	Seasons: []models.Season{
		{
			SeasonNumber: 1,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "The Man Trap",
					AirDate:       strPtr("1966-09-08"),
					StarDate:      floatPtr(1513.1),
				},
				{
					EpisodeNumber: 2,
					Title:         "Charlie X",
					AirDate:       strPtr("1966-09-15"),
					StarDate:      floatPtr(1533.6),
				},
				{
					EpisodeNumber: 3,
					Title:         "Where No Man Has Gone Before",
					AirDate:       strPtr("1966-09-22"),
					StarDate:      floatPtr(1312.4),
				},
				{
					EpisodeNumber: 4,
					Title:         "The Naked Time",
					AirDate:       strPtr("1966-09-29"),
					StarDate:      floatPtr(1704.2),
				},
				{
					EpisodeNumber: 5,
					Title:         "The Enemy Within",
					AirDate:       strPtr("1966-10-06"),
					StarDate:      floatPtr(1672.1),
				},
				{
					EpisodeNumber: 6,
					Title:         "Mudd's Women",
					AirDate:       strPtr("1966-10-13"),
					StarDate:      floatPtr(1329.8),
				},
				{
					EpisodeNumber: 7,
					Title:         "What Are Little Girls Made Of?",
					AirDate:       strPtr("1966-10-20"),
					StarDate:      floatPtr(2712.4),
				},
				{
					EpisodeNumber: 8,
					Title:         "Miri",
					AirDate:       strPtr("1966-10-27"),
					StarDate:      floatPtr(2713.5),
				},
				{
					EpisodeNumber: 9,
					Title:         "Dagger of the Mind",
					AirDate:       strPtr("1966-11-03"),
					StarDate:      floatPtr(2715.1),
				},
				{
					EpisodeNumber: 10,
					Title:         "The Corbomite Maneuver",
					AirDate:       strPtr("1966-11-10"),
					StarDate:      floatPtr(1512.2),
				},
				{
					EpisodeNumber: 11,
					Title:         "The Menagerie, Part I",
					AirDate:       strPtr("1966-11-17"),
					StarDate:      floatPtr(3012.4),
				},
				{
					EpisodeNumber: 12,
					Title:         "The Menagerie, Part II",
					AirDate:       strPtr("1966-12-08"),
					StarDate:      floatPtr(3013.1),
				},
				{
					EpisodeNumber: 13,
					Title:         "The Conscience of the King",
					AirDate:       strPtr("1966-12-01"),
					StarDate:      floatPtr(2817.6),
				},
				{
					EpisodeNumber: 14,
					Title:         "Balance of Terror",
					AirDate:       strPtr("1966-12-15"),
					StarDate:      floatPtr(1709.2),
				},
				{
					EpisodeNumber: 15,
					Title:         "Shore Leave",
					AirDate:       strPtr("1966-12-29"),
					StarDate:      floatPtr(3025.3),
				},
				{
					EpisodeNumber: 16,
					Title:         "The Galileo Seven",
					AirDate:       strPtr("1967-01-05"),
					StarDate:      floatPtr(2821.5),
				},
				{
					EpisodeNumber: 17,
					Title:         "The Squire of Gothos",
					AirDate:       strPtr("1967-01-12"),
					StarDate:      floatPtr(2124.5),
				},
				{
					EpisodeNumber: 18,
					Title:         "Arena",
					AirDate:       strPtr("1967-01-19"),
					StarDate:      floatPtr(3045.6),
				},
				{
					EpisodeNumber: 19,
					Title:         "Tomorrow Is Yesterday",
					AirDate:       strPtr("1967-01-26"),
					StarDate:      floatPtr(3113.2),
				},
				{
					EpisodeNumber: 20,
					Title:         "Court Martial",
					AirDate:       strPtr("1967-02-02"),
					StarDate:      floatPtr(2947.3),
				},
				{
					EpisodeNumber: 21,
					Title:         "The Return of the Archons",
					AirDate:       strPtr("1967-02-09"),
					StarDate:      floatPtr(3156.2),
				},
				{
					EpisodeNumber: 22,
					Title:         "Space Seed",
					AirDate:       strPtr("1967-02-16"),
					StarDate:      floatPtr(3141.9),
				},
				{
					EpisodeNumber: 23,
					Title:         "A Taste of Armageddon",
					AirDate:       strPtr("1967-02-23"),
					StarDate:      floatPtr(3192.1),
				},
				{
					EpisodeNumber: 24,
					Title:         "This Side of Paradise",
					AirDate:       strPtr("1967-03-02"),
					StarDate:      floatPtr(3417.3),
				},
				{
					EpisodeNumber: 25,
					Title:         "The Devil in the Dark",
					AirDate:       strPtr("1967-03-09"),
					StarDate:      floatPtr(3196.1),
				},
				{
					EpisodeNumber: 26,
					Title:         "Errand of Mercy",
					AirDate:       strPtr("1967-03-23"),
					StarDate:      floatPtr(3198.4),
				},
				{
					EpisodeNumber: 27,
					Title:         "The Alternative Factor",
					AirDate:       strPtr("1967-03-30"),
					StarDate:      floatPtr(3087.6),
				},
				{
					EpisodeNumber: 28,
					Title:         "The City on the Edge of Forever",
					AirDate:       strPtr("1967-04-06"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 29,
					Title:         "Operation: Annihilate!",
					AirDate:       strPtr("1967-04-13"),
					StarDate:      floatPtr(3287.2),
				},
			},
		},
		{
			SeasonNumber: 2,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Amok Time",
					AirDate:       strPtr("1967-09-15"),
					StarDate:      floatPtr(3372.7),
				},
				{
					EpisodeNumber: 2,
					Title:         "Who Mourns for Adonais?",
					AirDate:       strPtr("1967-09-22"),
					StarDate:      floatPtr(3468.1),
				},
				{
					EpisodeNumber: 3,
					Title:         "The Changeling",
					AirDate:       strPtr("1967-09-29"),
					StarDate:      floatPtr(3541.9),
				},
				{
					EpisodeNumber: 4,
					Title:         "Mirror, Mirror",
					AirDate:       strPtr("1967-10-06"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 5,
					Title:         "The Apple",
					AirDate:       strPtr("1967-10-13"),
					StarDate:      floatPtr(3715.3),
				},
				{
					EpisodeNumber: 6,
					Title:         "The Doomsday Machine",
					AirDate:       strPtr("1967-10-20"),
					StarDate:      floatPtr(4202.9),
				},
				{
					EpisodeNumber: 7,
					Title:         "Catspaw",
					AirDate:       strPtr("1967-10-27"),
					StarDate:      floatPtr(3018.2),
				},
				{
					EpisodeNumber: 8,
					Title:         "I, Mudd",
					AirDate:       strPtr("1967-11-03"),
					StarDate:      floatPtr(4513.3),
				},
				{
					EpisodeNumber: 9,
					Title:         "Metamorphosis",
					AirDate:       strPtr("1967-11-10"),
					StarDate:      floatPtr(3219.4),
				},
				{
					EpisodeNumber: 10,
					Title:         "Journey to Babel",
					AirDate:       strPtr("1967-11-17"),
					StarDate:      floatPtr(3842.3),
				},
				{
					EpisodeNumber: 11,
					Title:         "Friday's Child",
					AirDate:       strPtr("1967-12-01"),
					StarDate:      floatPtr(3497.2),
				},
				{
					EpisodeNumber: 12,
					Title:         "The Deadly Years",
					AirDate:       strPtr("1967-12-08"),
					StarDate:      floatPtr(3478.2),
				},
				{
					EpisodeNumber: 13,
					Title:         "Obsession",
					AirDate:       strPtr("1967-12-15"),
					StarDate:      floatPtr(3619.2),
				},
				{
					EpisodeNumber: 14,
					Title:         "Wolf in the Fold",
					AirDate:       strPtr("1967-12-22"),
					StarDate:      floatPtr(3614.9),
				},
				{
					EpisodeNumber: 15,
					Title:         "The Trouble With Tribbles",
					AirDate:       strPtr("1967-12-29"),
					StarDate:      floatPtr(4523.3),
				},
				{
					EpisodeNumber: 16,
					Title:         "The Gamesters of Triskelion",
					AirDate:       strPtr("1968-01-05"),
					StarDate:      floatPtr(3211.7),
				},
				{
					EpisodeNumber: 17,
					Title:         "A Piece of the Action",
					AirDate:       strPtr("1968-01-12"),
					StarDate:      floatPtr(4598.0),
				},
				{
					EpisodeNumber: 18,
					Title:         "The Immunity Syndrome",
					AirDate:       strPtr("1968-01-19"),
					StarDate:      floatPtr(4307.1),
				},
				{
					EpisodeNumber: 19,
					Title:         "A Private Little War",
					AirDate:       strPtr("1968-02-02"),
					StarDate:      floatPtr(4211.4),
				},
				{
					EpisodeNumber: 20,
					Title:         "Return to Tomorrow",
					AirDate:       strPtr("1968-02-09"),
					StarDate:      floatPtr(4768.3),
				},
				{
					EpisodeNumber: 21,
					Title:         "Patterns of Force",
					AirDate:       strPtr("1968-02-16"),
					StarDate:      floatPtr(2534.0),
				},
				{
					EpisodeNumber: 22,
					Title:         "By Any Other Name",
					AirDate:       strPtr("1968-02-23"),
					StarDate:      floatPtr(4657.5),
				},
				{
					EpisodeNumber: 23,
					Title:         "The Omega Glory",
					AirDate:       strPtr("1968-03-01"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 24,
					Title:         "The Ultimate Computer",
					AirDate:       strPtr("1968-03-08"),
					StarDate:      floatPtr(4729.4),
				},
				{
					EpisodeNumber: 25,
					Title:         "Bread and Circuses",
					AirDate:       strPtr("1968-03-15"),
					StarDate:      floatPtr(4040.7),
				},
				{
					EpisodeNumber: 26,
					Title:         "Assignment: Earth",
					AirDate:       strPtr("1968-03-29"),
					StarDate:      nil, // No StarDate
				},
			},
		},
		{
			SeasonNumber: 3,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Spock's Brain",
					AirDate:       strPtr("1968-09-20"),
					StarDate:      floatPtr(5431.4),
				},
				{
					EpisodeNumber: 2,
					Title:         "The Enterprise Incident",
					AirDate:       strPtr("1968-09-27"),
					StarDate:      floatPtr(5027.3),
				},
				{
					EpisodeNumber: 3,
					Title:         "The Paradise Syndrome",
					AirDate:       strPtr("1968-10-04"),
					StarDate:      floatPtr(4842.6),
				},
				{
					EpisodeNumber: 4,
					Title:         "And the Children Shall Lead",
					AirDate:       strPtr("1968-10-11"),
					StarDate:      floatPtr(5029.5),
				},
				{
					EpisodeNumber: 5,
					Title:         "Is There in Truth No Beauty?",
					AirDate:       strPtr("1968-10-18"),
					StarDate:      floatPtr(5630.7),
				},
				{
					EpisodeNumber: 6,
					Title:         "Spectre of the Gun",
					AirDate:       strPtr("1968-10-25"),
					StarDate:      floatPtr(4385.3),
				},
				{
					EpisodeNumber: 7,
					Title:         "Day of the Dove",
					AirDate:       strPtr("1968-11-01"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 8,
					Title:         "For the World Is Hollow and I Have Touched the Sky",
					AirDate:       strPtr("1968-11-08"),
					StarDate:      floatPtr(5476.3),
				},
				{
					EpisodeNumber: 9,
					Title:         "The Tholian Web",
					AirDate:       strPtr("1968-11-15"),
					StarDate:      floatPtr(5693.2),
				},
				{
					EpisodeNumber: 10,
					Title:         "Plato's Stepchildren",
					AirDate:       strPtr("1968-11-22"),
					StarDate:      floatPtr(5784.2),
				},
				{
					EpisodeNumber: 11,
					Title:         "Wink of an Eye",
					AirDate:       strPtr("1968-11-29"),
					StarDate:      floatPtr(5710.5),
				},
				{
					EpisodeNumber: 12,
					Title:         "The Empath",
					AirDate:       strPtr("1968-12-06"),
					StarDate:      floatPtr(5121.5),
				},
				{
					EpisodeNumber: 13,
					Title:         "Elaan of Troyius",
					AirDate:       strPtr("1968-12-20"),
					StarDate:      floatPtr(4372.5),
				},
				{
					EpisodeNumber: 14,
					Title:         "Whom Gods Destroy",
					AirDate:       strPtr("1969-01-03"),
					StarDate:      floatPtr(5718.3),
				},
				{
					EpisodeNumber: 15,
					Title:         "Let That Be Your Last Battlefield",
					AirDate:       strPtr("1969-01-10"),
					StarDate:      floatPtr(5730.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "The Mark of Gideon",
					AirDate:       strPtr("1969-01-17"),
					StarDate:      floatPtr(5423.4),
				},
				{
					EpisodeNumber: 17,
					Title:         "That Which Survives",
					AirDate:       strPtr("1969-01-24"),
					StarDate:      nil, // No StarDate
				},
				{
					EpisodeNumber: 18,
					Title:         "The Lights of Zetar",
					AirDate:       strPtr("1969-01-31"),
					StarDate:      floatPtr(5725.3),
				},
				{
					EpisodeNumber: 19,
					Title:         "Requiem for Methuselah",
					AirDate:       strPtr("1969-02-14"),
					StarDate:      floatPtr(5843.7),
				},
				{
					EpisodeNumber: 20,
					Title:         "The Way to Eden",
					AirDate:       strPtr("1969-02-21"),
					StarDate:      floatPtr(5832.3),
				},
				{
					EpisodeNumber: 21,
					Title:         "The Cloud Minders",
					AirDate:       strPtr("1969-02-28"),
					StarDate:      floatPtr(5818.4),
				},
				{
					EpisodeNumber: 22,
					Title:         "The Savage Curtain",
					AirDate:       strPtr("1969-03-07"),
					StarDate:      floatPtr(5906.4),
				},
				{
					EpisodeNumber: 23,
					Title:         "All Our Yesterdays",
					AirDate:       strPtr("1969-03-14"),
					StarDate:      floatPtr(5943.7),
				},
				{
					EpisodeNumber: 24,
					Title:         "Turnabout Intruder",
					AirDate:       strPtr("1969-06-03"),
					StarDate:      floatPtr(5928.5),
				},
			},
		},
	},
}
