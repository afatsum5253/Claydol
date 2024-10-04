package utils

func ConstructPrompt(selectedText string) string {
	return `You are my Anki flashcards creator who will respond as JSON with the created flascards. You will create the front and back content of the flashcards using the following text:

Text start

` + selectedText + `

Text end

Rules:
Follow these rules strictly:
1) Only give the response as a valid JSON in this format: 

{ cards: [{modelName: "Basic", "front": "question", "back": ["list of bullet points"]]}

or

{ cards: [{modelName: "Cloze", "text": "The capital of Romania is {{c1::Bucharest}} and {{c2::Budapest}}"]}

The cards must be a mix of either Basic or Cloze type cards, choose the cards modelType based on the context.

Rules about Cloze cards: Keep the "Cloze" tag minimum number of keywords, if there are more keywords, breakdown the cloze tags with c1, c2 etc for the same sentence. if you want yo add more cloze tags, you can add do it using c1::"The cloze tags examples are {{c1::Bucharest}} and {{c2::Budapest}}" just increment the "c" when adding it to the same sentence. Prefer using this card model if the content relates to examples of the card topic.

Rules about Basic cards: The "Basic" card's back content will be parsed as bullet points for each element, so treat each element as bullet points but do not add any special characters like dashes, dots, colons, semi-colons etc to the content of front or back properties of the card that are not relevant to the context of the data.

2) The content of the cards should not be too long.

3) Only use information that is important in the context of the given text.

4) The front and back content of the cards must be in questions and answer format. The answers must be short and in bullet points catered to medical students as much as possible.`
}
