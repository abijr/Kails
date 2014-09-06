// A file to initialize mongodb database.
// Creates database and collections, and adds initial documents,
// also creates indexes.

// run `mongo <database> db_init.js`

// User = user
// Password = password
default_user = {
	"name" : "user",
	"email" : "user@email.com",
	"pass" : BinData(0,"RurG60nC/Kx9N0MumOq74K7tNwebAjWC9AYXhJOqTFY="),
	"salt" : BinData(0,"Li6QELxiH4vcqg=="),
	"lang" : "spanish",
	"study": "english",
	"since" : ISODate("2014-08-14T20:28:00.414Z"),
	"levels" : {
		"1": {
			// last review
			"last":  ISODate("2014-08-14T20:28:00.414Z"),
		}
	},
};
db.users.save(default_user);

// English language collection
// contains levels and words
level_1 = {
	"id": 1,
	"type": "level",
	"lang": "english",
	// description
	"desc": "some sort of description",
	"version": 1,
	"words": [
		{
			"word": "word1",
			"translation": "blah blah blah",
			"class": "verb",
			// rename this to challenge?
			// fuck it, so much to change T.T
			"sentences": [
				{
					// .......Dropping the thought here.........
					// perhaps it's better "question/answer"
					// instead of native/translation
					// .........................................
					"native": "english sentence here",
					"translation": "translation here"
				},{
					"native": "other english sentence here",
					"translation": "other translation here"
				},
			]
		},
	]
};
db.languages.save(level_1);

word1 = {
	"word": "word1",
	"type": "word",
	"level": 1,
	"class": "verb",
	"lang": "english",
	"translation": "blah blah blah blah",
	"sentences": [
		{
			"native": "english sentence here",
			"translation": "translation here"
		},{
			"native": "other english sentence here",
			"translation": "other translation here"
		},
	],

};
db.languages.save(word1);

// Add indexes.
// user collection indexes
db.users.ensureIndex({"name": 1}, {"unique": true});
db.users.ensureIndex({"email": 1}, {"unique": true});

// languages collection indexes
db.languages.ensureIndex({"lang": 1});
db.languages.ensureIndex({"level": 1});
db.languages.ensureIndex({"type": 1});
db.languages.ensureIndex({"id": 1}, {"sparse": true});
db.languages.ensureIndex({"word": 1}, {"sparse": true});
