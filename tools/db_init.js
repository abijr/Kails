// A file to initialize mongodb database.
// Creates database and collections, and adds initial documents,
// also creates indexes.

use kails;

// User = user
// Password = password
default_user = {
	"name" : "user",
	"email" : "user@email.com",
	"pass" : BinData(0,"RurG60nC/Kx9N0MumOq74K7tNwebAjWC9AYXhJOqTFY="),
	"salt" : BinData(0,"Li6QELxiH4vcqg=="),
	"lang" : "",
	"since" : ISODate("2014-08-14T20:28:00.414Z"),
	"levels" : [ ],
};
db.users.save(default_user);

// English language collection
// contains levels and words
level_1 = {
	"level": 1,
	"type": "level",
	"name": "some sort of description",
	"version": 1,
	"word1": {
		"translation": "blah blah blah",
		"type": "verb",
		"sentences": [
			"english sentence": "translation",
			"other sentence": "with translation..."
		]
	},
	"word2": "same as above....",
	"word3": "etc...",
};
db.english.save(level_1);

word1 = {
	"word": "word1",
	"type": "word",
	"level": 1,
	"type": "verb",
	"translation": "blah blah blah blah",
	"sentences": [
		"english sentence": "translation",
		"other sentence": "translation..."
	],

};
db.english.save(word1);

// Add indexes.
// user collection indexes
db.users.ensureIndex({"name": 1}, {"unique": true});
db.users.ensureIndex({"email": 1}, {"unique": true});

// english collection indexes
db.english.ensureIndex({"level": 1});
db.english.ensureIndex({"type": 1});
db.english.ensureIndex({"word": 1}, {"sparse": true});
