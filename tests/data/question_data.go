package data

import (
	"net/http"

	serverModels "gitlab.com/banter-bus/banter-bus-management-api/internal/server/models"
)

// AddQuestion is the test data for add questions to a game
var AddQuestion = []struct {
	TestDescription string
	Game            string
	Payload         interface{}
	Expected        int
}{
	{
		"Add a question to quibly and to round pair",
		"quibly",
		&serverModels.NewQuestion{
			Content: "this is another question?",
			Round:   "pair",
		}, http.StatusCreated,
	},
	{
		"Add a question to quibly and to round answer, language de",
		"quibly",
		&serverModels.NewQuestion{
			Content:      "what is the funniest thing ever told?",
			LanguageCode: "de",
			Round:        "answers",
		}, http.StatusCreated,
	},
	{
		"Add a question to quibly and to round group",
		"quibly",
		&serverModels.NewQuestion{
			Content: "what does ATGM stand for?",
			Round:   "group",
		}, http.StatusCreated,
	},
	{
		"Add a question to drawlosseum, language ur",
		"drawlosseum",
		&serverModels.NewQuestion{
			Content:      "camel",
			LanguageCode: "ur",
		}, http.StatusCreated,
	},
	{
		"Add another question to drawlosseum",
		"drawlosseum",
		&serverModels.NewQuestion{
			Content: "pencil",
		}, http.StatusCreated,
	},
	{
		"Add yet another question to drawlosseum",
		"drawlosseum",
		&serverModels.NewQuestion{
			Content: "food fight",
		}, http.StatusCreated,
	},
	{
		"Add a question to fibbing it, round opinion new group bike group, language en",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content:      "do you love bikes?",
			LanguageCode: "en",
			Round:        "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "question",
			},
		}, http.StatusCreated,
	},
	{
		"Add another question to fibbing it, round opinion new group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "how much does liam love bikes?",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "question",
			},
		}, http.StatusCreated,
	},
	{
		"Add an answer to fibbing it, round opinion existing group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "super love",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "answer",
			},
		}, http.StatusCreated,
	},
	{
		"Add an answer to fibbing it, round free_form existing group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is love?",
			Round:   "free_form",
			Group: &serverModels.Group{
				Name: "bike_group",
			},
		}, http.StatusCreated,
	},
	{
		"Add an answer to fibbing it, round free_form new group horse group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is the fastest horse?",
			Round:   "free_form",
			Group: &serverModels.Group{
				Name: "horse_group",
			},
		}, http.StatusCreated,
	},
	{
		"Add an answer to fibbing it, round free_form existing group horse group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is the second horse called?",
			Round:   "free_form",
			Group: &serverModels.Group{
				Name: "horse_group",
			},
		}, http.StatusCreated,
	},
	{
		"Add an answer to fibbing it, round likely",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "to never eat a vegetable again?",
			Round:   "likely",
		}, http.StatusCreated,
	},
	{
		"Add question to quibly, invalid round",
		"quibly",
		&serverModels.NewQuestion{
			Content: "This is another question?",
			Round:   "invalid",
		}, http.StatusBadRequest,
	},
	{
		"Add question to quibly, invalid2 round",
		"quibly",
		&serverModels.NewQuestion{
			Content: "This is another question?",
			Round:   "invalid2",
		}, http.StatusBadRequest,
	},
	{
		"Add an answer to fibbing it, invalid round free_form",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is the fastest horse?",
			Round:   "invalid_free_form",
			Group: &serverModels.Group{
				Name: "horse_group",
			},
		}, http.StatusBadRequest,
	},
	{
		"Add an answer to fibbing it, invalid language code",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content:      "What is the fastest horse?",
			LanguageCode: "deed",
			Round:        "opinion",
			Group: &serverModels.Group{
				Name: "horse_group",
				Type: "answer",
			},
		}, http.StatusBadRequest,
	},
	{
		"Add an answer to fibbing it, round opinion invalid answers type",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "super love",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "answers",
			},
		}, http.StatusBadRequest,
	},
	{
		"Add an answer to fibbing it, round opinion invalid questions type",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "super love",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "questions",
			},
		}, http.StatusBadRequest,
	},
	{
		"Add an answer to fibbing it, round opinion invalid type",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "super love",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "type",
			},
		}, http.StatusBadRequest,
	},
	{
		"game does not exist but missing content",
		"quibly v3",
		&serverModels.NewQuestion{}, http.StatusBadRequest,
	},
	{
		"game does not exist",
		"quibly_v2",
		&serverModels.NewQuestion{
			Content: "super love",
		}, http.StatusNotFound,
	},
	{
		"another game does not exist",
		"quibly v3",
		&serverModels.NewQuestion{
			Content: "super love",
		}, http.StatusNotFound,
	},
	{
		"Add a question that already exists to quibly and to round pair",
		"quibly",
		&serverModels.NewQuestion{
			Content: "this is also question?",
			Round:   "pair",
		}, http.StatusConflict,
	},
	{
		"Add a question that already exists to quibly and to round answer",
		"quibly",
		&serverModels.NewQuestion{
			Content: "pink mustard",
			Round:   "answers",
		}, http.StatusConflict,
	},
	{
		"Add a question that already exists to quibly and to round answer",
		"quibly",
		&serverModels.NewQuestion{
			Content:      "german",
			LanguageCode: "de",
			Round:        "answers",
		}, http.StatusConflict,
	},
	{
		"Add a question that already exists to quibly and to round group",
		"quibly",
		&serverModels.NewQuestion{
			Content: "what does ATGM stand for?",
			Round:   "group",
		}, http.StatusConflict,
	},
	{
		"Add a question that already exists to drawlosseum",
		"drawlosseum",
		&serverModels.NewQuestion{
			Content: "horse",
		}, http.StatusConflict,
	},
	{
		"Add another question that already exists to drawlosseum",
		"drawlosseum",
		&serverModels.NewQuestion{
			Content: "pencil",
		}, http.StatusConflict,
	},
	{
		"Add yet another question that already exists to drawlosseum",
		"drawlosseum",
		&serverModels.NewQuestion{
			Content: "food fight",
		}, http.StatusConflict,
	},
	{
		"Add a question to fibbing it that already exists, round opinion new group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "do you love bikes?",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "question",
			},
		}, http.StatusConflict,
	},
	{
		"Add another question to fibbing it that already exists, round opinion new group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "how much does liam love bikes?",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "question",
			},
		}, http.StatusConflict,
	},
	{
		"Add an answer to fibbing it that already exists, round opinion existing group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "super love",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "bike_group",
				Type: "answer",
			},
		}, http.StatusConflict,
	},
	{
		"Add an answer to fibbing it that already exists, round free_form existing group bike group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is love?",
			Round:   "free_form",
			Group: &serverModels.Group{
				Name: "bike_group",
			},
		}, http.StatusConflict,
	},
	{
		"Add an answer to fibbing it that already exists, round free_form new group horse group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is the fastest horse?",
			Round:   "free_form",
			Group: &serverModels.Group{
				Name: "horse_group",
			},
		}, http.StatusConflict,
	},
	{
		"Add an answer to fibbing it that already exists, round free_form existing group horse group",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What is the second horse called?",
			Round:   "free_form",
			Group: &serverModels.Group{
				Name: "horse_group",
			},
		}, http.StatusConflict,
	},
	{
		"Add an answer to fibbing it tthat already exists, round likely",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "to never eat a vegetable again?",
			Round:   "likely",
		}, http.StatusConflict,
	},
	{
		"Add a question to fibbing it that already exists",
		"fibbing_it",
		&serverModels.NewQuestion{
			Content: "What do you think about horses?",
			Round:   "opinion",
			Group: &serverModels.Group{
				Name: "horse_group",
				Type: "question",
			},
		}, http.StatusConflict,
	},
}

// RemoveQuestion is the test data for removing questions from a game
var RemoveQuestion = []struct {
	TestDescription string
	Game            string
	ID              string
	Expected        int
}{
	{
		"Remove a question from quibly and round pair",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		http.StatusOK,
	},
	{
		"Remove a question from drawlossuem",
		"drawlosseum",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		http.StatusOK,
	},
	{
		"Remove a question from fibbing it",
		"fibbing_it",
		"714464a5-337f-4ce7-a4df-2b00764e5c5b",
		http.StatusOK,
	},
	{
		"Remove a question from fibbing it that was already removed",
		"fibbing_it",
		"714464a5-337f-4ce7-a4df-2b00764e5c5b",
		http.StatusNotFound,
	},
	{
		"Remove a question from quibly that was already removed",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		http.StatusNotFound,
	},
	{
		"Remove a question that was already removed from drawlossuem",
		"drawlosseum",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		http.StatusNotFound,
	},
	{
		"Remove a question that doesn't exist from fibbing_it",
		"fibbing_it",
		"invalid-id",
		http.StatusNotFound,
	},
}

// AddTranslationQuestion is the test data for adding translations to questions.
var AddTranslationQuestion = []struct {
	TestDescription string
	Game            string
	LanguageCode    string
	ID              string
	Payload         interface{}
	Expected        int
}{
	{
		"Update question in quibly and round pair, new language fr",
		"quibly",
		"fr",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		&serverModels.QuestionTranslation{
			Content: "this is a question?",
		},
		http.StatusCreated,
	},
	{
		"Update question in quibly and round pair, replace existing language de",
		"quibly",
		"de",
		"bf64d60c-62ee-420a-976e-bfcaec77ad8b",
		&serverModels.QuestionTranslation{
			Content: "le german?",
		},
		http.StatusCreated,
	},
	{
		"Update question in drawlosseum",
		"drawlosseum",
		"hi",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		&serverModels.QuestionTranslation{
			Content: "ऊंट",
		},
		http.StatusCreated,
	},
	{
		"Update question in fibbing it, round opinion",
		"fibbing_it",
		"it",
		"580aeb14-d907-4a22-82c8-f2ac544a2cd1",
		&serverModels.QuestionTranslation{
			Content: "Cosa ne pensi dei cavalli?",
		},
		http.StatusCreated,
	},
	{
		"Update question in fibbing it, round opinion and answers section",
		"fibbing_it",
		"de",
		"aa9fe2b5-79b5-458d-814b-45ff95a617fc",
		&serverModels.QuestionTranslation{
			Content: "Liebe",
		}, http.StatusCreated,
	},
	{
		"Missing content",
		"quibly",
		"en",
		"a9c00e19-d41e-4b15-a8bd-ec921af9123d",
		&serverModels.NewQuestion{}, http.StatusBadRequest,
	},
	{
		"Update question in fibbing it but invalid language code",
		"fibbing_it",
		"ittt",
		"3e2889f6-56aa-4422-a7c5-033eafa9fd39",
		&serverModels.QuestionTranslation{
			Content: "was ist Liebe?",
		}, http.StatusBadRequest,
	},
	{
		"game does not exist",
		"quibly v3",
		"de",
		"3e2889f6-56aa-4422-a7c5-033eafa9fd39",
		&serverModels.QuestionTranslation{
			Content: "was ist Liebe?",
		}, http.StatusNotFound,
	},
	{
		"Question doesn't exist",
		"fibbing_it",
		"de",
		"9f64d60c-62ee-420a-976e-bfcaec77ad8b",
		&serverModels.QuestionTranslation{
			Content: "was ist Liebe?",
		}, http.StatusNotFound,
	},
}

// RemoveTranslationQuestion is the test data for removing questions from game.
var RemoveTranslationQuestion = []struct {
	TestDescription string
	Game            string
	ID              string
	LanguageCode    string
	Expected        int
}{
	{
		"Delete a question quibly from round pair",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		"en",
		http.StatusOK,
	},
	{
		"Delete a question quibly from round pair, language ur",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		"ur",
		http.StatusOK,
	},
	{
		"Delete a question quibly from round answers",
		"quibly",
		"bf64d60c-62ee-420a-976e-bfcaec77ad8b",
		"en",
		http.StatusOK,
	},
	{
		"Delete a question quibly from round group, language fr",
		"quibly",
		"4b4dd325-04fd-4aa4-9382-2874dcfd5cae",
		"fr",
		http.StatusOK,
	},
	{
		"Delete a question drawlosseum",
		"drawlosseum",
		"815464a5-337f-4ce7-a4df-2b00764e5c6c",
		"en",
		http.StatusOK,
	},
	{
		"Delete another question drawlosseum",
		"drawlosseum",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		"en",
		http.StatusOK,
	},
	{
		"Delete a question to fibbing it, round opinion from group horse group",
		"fibbing_it",
		"3e2889f6-56aa-4422-a7c5-033eafa9fd39",
		"en",
		http.StatusOK,
	},
	{
		"Delete a answer to fibbing it, round opinion from group horse group",
		"fibbing_it",
		"03a462ba-f483-4726-aeaf-b8b6b03ce3e2",
		"en",
		http.StatusOK,
	},
	{
		"Delete a question quibly from round pair that was already deleted",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		"en",
		http.StatusNotFound,
	},
	{
		"Delete a question drawlosseum that was already deleted",
		"drawlosseum",
		"815464a5-337f-4ce7-a4df-2b00764e5c6c",
		"en",
		http.StatusNotFound,
	},
	{
		"Delete a question already removed from fibbing it, round free_form from group bike group",
		"fibbing_it",
		"3e2889f6-56aa-4422-a7c5-033eafa9fd39",
		"en",
		http.StatusNotFound,
	},
	{
		"Delete another  already removed from fibbing it, round likely",
		"fibbing_it",
		"03a462ba-f483-4726-aeaf-b8b6b03ce3e2",
		"en",
		http.StatusNotFound,
	},
}

// EnableQuestion test data used to test enable endpoint
var EnableQuestion = []struct {
	TestDescription string
	Game            string
	ID              string
	Expected        int
}{
	{
		"Enable a question, quibly and round pair",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		http.StatusOK,
	},
	{
		"Enable a question, quibly and round answers",
		"quibly",
		"4b4dd325-04fd-4aa4-9382-2874dcfd5cae",
		http.StatusOK,
	},
	{
		"Enable a question, fibbing_it and round opinion",
		"fibbing_it",
		"7799e38a-758d-4a1b-a191-99c59440af76",
		http.StatusOK,
	},
	{
		"Enable an answer, fibbing_it and round opinion",
		"fibbing_it",
		"03a462ba-f483-4726-aeaf-b8b6b03ce3e2",
		http.StatusOK,
	},
	{
		"Enable a question, fibbing_it and round free_form",
		"fibbing_it",
		"580aeb14-d907-4a22-82c8-f2ac544a2cd1",
		http.StatusOK,
	},
	{
		"Enable a question, fibbing_it and round likely",
		"fibbing_it",
		"d80f2d90-0fb0-462a-8fbd-1aa00b4e42a5",
		http.StatusOK,
	},
	{
		"Enable a question, drawlosseum",
		"drawlosseum",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		http.StatusOK,
	},
	{
		"Enable an already enabled question, drawlosseum",
		"drawlosseum",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		http.StatusOK,
	},
	{
		"Game does not exist",
		"quibly v3",
		"901464a5-337f-4ce7-a4df-2b00764e5d8d",
		http.StatusNotFound,
	},
}

// DisableQuestion test data used to test disable endpoint
var DisableQuestion = []struct {
	TestDescription string
	Game            string
	ID              string
	Expected        int
}{
	{
		"Disable a question, quibly and round pair",
		"quibly",
		"4d18ac45-8034-4f8e-b636-cf730b17e51a",
		http.StatusOK,
	},
	{
		"Disable a question, quibly and round answers",
		"quibly",
		"bf64d60c-62ee-420a-976e-bfcaec77ad8b",
		http.StatusOK,
	},
	{
		"Disable a question, fibbing_it and round opinion",
		"fibbing_it",
		"3e2889f6-56aa-4422-a7c5-033eafa9fd39",
		http.StatusOK,
	},
	{
		"Disable an answer, fibbing_it and round opinion",
		"fibbing_it",
		"03a462ba-f483-4726-aeaf-b8b6b03ce3e2",
		http.StatusOK,
	},
	{
		"Disable a question, fibbing_it and round likely",
		"fibbing_it",
		"d5aa9153-f48c-45cc-b411-fb9b2d38e78f",
		http.StatusOK,
	},
	{
		"Disable a question, drawlosseum",
		"drawlosseum",
		"101464a5-337f-4ce7-a4df-2b00764e5d8d",
		http.StatusOK,
	},
	{
		"Question does not exist",
		"quibly",
		"90aa9153-f48c-45cc-b411-fb9b2d38e78f",
		http.StatusNotFound,
	},
	{
		"Game does not exist",
		"quibly v3",
		"d5aa9153-f48c-45cc-b411-fb9b2d38e78f",
		http.StatusNotFound,
	},
}

// GetAllGroups is the data for the get groups tests
var GetAllGroups = []struct {
	TestDescription string
	Payload         *serverModels.GroupInput
	ExpectedGroups  []string
	ExpectedCode    int
}{
	{
		"Get all groups from questions from the opinion round in the Fibbing It game",
		&serverModels.GroupInput{
			GameParams: serverModels.GameParams{
				Name: "fibbing_it",
			},
			Round: "opinion",
		},
		[]string{
			"horse_group",
		},
		http.StatusOK,
	},

	{
		"Get all groups from questions from the free form round in the Fibbing It game",
		&serverModels.GroupInput{
			GameParams: serverModels.GameParams{
				Name: "fibbing_it",
			},
			Round: "free_form",
		},
		[]string{
			"bike_group",
			"cat_group",
		},
		http.StatusOK,
	},

	{
		"Try to get groups from a round in Fibbing It that does not have groups",
		&serverModels.GroupInput{
			GameParams: serverModels.GameParams{
				Name: "fibbing_it",
			},
			Round: "likely",
		},
		[]string{},
		http.StatusNotFound,
	},

	{
		"Try to get groups from a non-existent round",
		&serverModels.GroupInput{
			GameParams: serverModels.GameParams{
				Name: "fibbing_it",
			},
			Round: "genocide",
		},
		[]string{},
		http.StatusNotFound,
	},

	{
		"Try to get groups from a game that does not have groups",
		&serverModels.GroupInput{
			GameParams: serverModels.GameParams{
				Name: "quibly",
			},
			Round: "opinion",
		},
		[]string{},
		http.StatusNotFound,
	},
}
