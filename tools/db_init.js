// A file to initialize mongodb database.
// Creates database and collections, and adds initial documents,
// also creates indexes.

// run `mongo <database> db_init.js`

try {
	// Check if database kails exists
	// drop it and create a new clean
	// database.
	db._dropDatabase("kails");
	db._createDatabase("kails");
	db._useDatabase("kails");

} catch(e) {
	if (e == "[ArangoError 1228: database not found]") {
		db._createDatabase("kails");
		db._useDatabase("kails");
	} else {
		throw "Some unknown error" + e + "happened";
	}
}

db._create("users");
db._create("languages");
db._create("topics");
db._create("privileges");
db._createEdgeCollection("relations");

// Setup graph stuff
var graph_module = require("org/arangodb/general-graph");

// Define relationships
relation = graph_module._directedRelation("relations", "users", "users");

// Create graph
var graph = graph_module._create("relations");

// Add vertex collection
graph._addVertexCollection("users");

// Add edges to graph
graph._extendEdgeDefinitions(relation);



// User = user
// Password = password
default_user = {
	"Username" : "user",
	"Email" : "user@email.com",
	"Password" : "RurG60nC/Kx9N0MumOq74K7tNwebAjWC9AYXhJOqTFY=",
	"Salt" : "Li6QELxiH4vcqg==",
	"InterfaceLanguage" : "es-MX",
	"StudyLanguage": "english",
	"Since" : new Date(),
	"Level": 1,
	"Experience": 35,
	"Lessons" : {
		"1": {
			"Unlocked": true,
			// last review
			"LastReview":  new Date(),
		},
		"2": {
			"Unlocked": true,
			// last review
			"LastReview":  "2013-10-24T06:00:00.000Z",
		}
	},
	"Topics" : [ "sports", "entertainment", "vehicles"]
};
db.users.save(default_user);

// Username = other
// Password = password
other_user = {
	"Username" : "other",
	"Email" : "other@email.com",
	"Password" : "RurG60nC/Kx9N0MumOq74K7tNwebAjWC9AYXhJOqTFY=",
	"Salt" : "Li6QELxiH4vcqg==",
	"InterfaceLanguage" : "en-US",
	"StudyLanguage": "spanish",
	"Level": 1,
	"Experience": 35,
	"Since" : new Date(),
	"Lessons" : {
		"1": {
			"Unlocked": true,
			// last review
			"LastReview":  new Date(),
			// SRS info
			"Bucket": 1
		}
	},
};
db.users.save(other_user);


english_program = {
	"Type": "program",
	"Language": "english",
	"Lessons": [
		{
			"Id": 1,
			"Description": "some sort of description",
		},{
			"Id": 2,
			"Description": "some other sort of description",
		},
	]
};
db.languages.save(english_program);

// English language collection
// contains lessons and words
lesson_1 = {
	"Id": 1,
	"Type": "lesson",
	"Language": "english",
	// description
	"Description": "some sort of description",
	"Version": 1,
	"Words": [
		{
			"Word": "word1",
			"Definition": "blah blah blah",
			"Class": "verb",
			// rename this to challenge?
			// fuck it, so much to change T.T
			"Sentences": [
				{
					// .......Dropping the thought here.........
					// perhaps it's better "question/answer"
					// instead of native/translation
					// .........................................
					"Native": "english sentence here",
					"Translation": "translation here"
				},{
					"Native": "other english sentence here",
					"Translation": "other translation here"
				},
			]
		},
	]
};

// lesson_2
lesson_2 = {
	"Id": 2,
	"Type": "lesson",
	"Language": "english",
	// description
	"Desc": "some sort of description",
	"Version": 1,
	"Words": [
		{
			"Word": "word1",
			"Definition": "blah blah blah",
			"Class": "verb",
			// rename this to challenge?
			// fuck it, so much to change T.T
			"Sentences": [
				{
					// .......Dropping the thought here.........
					// perhaps it's better "question/answer"
					// instead of native/translation
					// .........................................
					"Native": "blah blah",
					"Translation": "blah blah"
				},{
					"Native": "other english sentence here",
					"Translation": "other translation here"
				},
			]
		},
	]
};
db.languages.save(lesson_1);
db.languages.save(lesson_2);

word1 = {
	"Word": "word1",
	"Type": "word",
	"Lesson": 1,
	"Class": "verb",
	"Language": "english",
	"Definition": "blah blah blah blah",
	"Sentences": [
		{
			"Native": "english sentence here",
			"Translation": "translation here"
		},{
			"Native": "other english sentence here",
			"Translation": "other translation here"
		},
	],

};
db.languages.save(word1);

topic = {
	"Id": 1,
	"Name": "sports",
	"Subtopics": ["soccer", "baseball", "basketball", "football"],
	"NoUser": 0
};
db.topics.save(topic);

// Add indexes.
// user collection indexes
db.users.ensureFulltextIndex("Username");
db.users.ensureUniqueConstraint("Name");
db.users.ensureUniqueConstraint("Email");

// languages collection indexes
db.languages.ensureHashIndex("Language");
db.languages.ensureHashIndex("Lesson");
db.languages.ensureHashIndex("Type");
db.languages.ensureHashIndex("Id");
db.languages.ensureHashIndex("Word");

print("done!");
