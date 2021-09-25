package parser

import (
	"fmt"
	"testing"

	"github.com/inf0rmer/deckdiff/pkg/mtg"
	"github.com/stretchr/testify/assert"
)

func TestDeckboxParser(t *testing.T) {
	input := `<!DOCTYPE html>
	<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
	<head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8" />
			<meta name="description" content="Card set print view" />
			<title>Deckbox.org card set print view</title>
	</head>
	<body>
					4 Blood Crypt<br/>4 Bloodstained Mire<br/>1 Forest<br/>1 Mountain<br/>1 Nurturing Peatland<br/>1 Overgrown Tomb<br/>1 Stomping Ground<br/>1 Swamp<br/>4 Urza's Saga<br/>3 Verdant Catacombs<br/>2 Wooded Foothills<br/>4 Dragon's Rage Channeler<br/>1 Expedition Map<br/>4 Inquisition of Kozilek<br/>1 Kolaghan's Command<br/>2 Lightning Bolt<br/>4 Mishra's Bauble<br/>1 Nihil Spellbomb<br/>1 Pyrite Spellbomb<br/>4 Ragavan, Nimble Pilferer<br/>1 Shadowspear<br/>4 Tarmogoyf<br/>2 Thoughtseize<br/>4 Unholy Heat<br/>4 Wrenn and Six<br/>
	
			<p><strong>Sideboard:</strong></p>
			2 Alpine Moon<br/>1 Chalice of the Void<br/>2 Collective Brutality<br/>2 Engineered Explosives<br/>1 Ghost Quarter<br/>1 Kolaghan's Command<br/>1 Lurrus of the Dream-Den<br/>1 Pithing Needle<br/>2 Terminate<br/>1 Torpor Orb<br/>1 Void Mirror<br/>
	
	</body>
	</html>
`
	expected := mtg.NewDecklist(
		[]*mtg.Card{
			mtg.NewCard("Blood Crypt", 4, mtg.None),
			mtg.NewCard("Bloodstained Mire", 4, mtg.None),
			mtg.NewCard("Forest", 1, mtg.None),
			mtg.NewCard("Mountain", 1, mtg.None),
			mtg.NewCard("Nurturing Peatland", 1, mtg.None),
			mtg.NewCard("Overgrown Tomb", 1, mtg.None),
			mtg.NewCard("Stomping Ground", 1, mtg.None),
			mtg.NewCard("Swamp", 1, mtg.None),
			mtg.NewCard("Urza's Saga", 4, mtg.None),
			mtg.NewCard("Verdant Catacombs", 3, mtg.None),
			mtg.NewCard("Wooded Foothills", 2, mtg.None),
			mtg.NewCard("Dragon's Rage Channeler", 4, mtg.None),
			mtg.NewCard("Expedition Map", 1, mtg.None),
			mtg.NewCard("Inquisition of Kozilek", 4, mtg.None),
			mtg.NewCard("Kolaghan's Command", 1, mtg.None),
			mtg.NewCard("Lightning Bolt", 2, mtg.None),
			mtg.NewCard("Mishra's Bauble", 4, mtg.None),
			mtg.NewCard("Nihil Spellbomb", 1, mtg.None),
			mtg.NewCard("Pyrite Spellbomb", 1, mtg.None),
			mtg.NewCard("Ragavan, Nimble Pilferer", 4, mtg.None),
			mtg.NewCard("Shadowspear", 1, mtg.None),
			mtg.NewCard("Tarmogoyf", 4, mtg.None),
			mtg.NewCard("Thoughtseize", 2, mtg.None),
			mtg.NewCard("Unholy Heat", 4, mtg.None),
			mtg.NewCard("Wrenn and Six", 4, mtg.None),
		},
		[]*mtg.Card{
			mtg.NewCard("Alpine Moon", 2, mtg.None),
			mtg.NewCard("Chalice of the Void", 1, mtg.None),
			mtg.NewCard("Collective Brutality", 2, mtg.None),
			mtg.NewCard("Engineered Explosives", 2, mtg.None),
			mtg.NewCard("Ghost Quarter", 1, mtg.None),
			mtg.NewCard("Kolaghan's Command", 1, mtg.None),
			mtg.NewCard("Lurrus of the Dream-Den", 1, mtg.None),
			mtg.NewCard("Pithing Needle", 1, mtg.None),
			mtg.NewCard("Terminate", 2, mtg.None),
			mtg.NewCard("Torpor Orb", 1, mtg.None),
			mtg.NewCard("Void Mirror", 1, mtg.None),
		},
		nil,
	)
	actual, err := NewDeckboxParser().Parse(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestDeckboxParserErrorWithInvalidDecklist(t *testing.T) {
	input := `<!DOCTYPE html>
	<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
	<head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8" />
			<meta name="description" content="Card set print view" />
			<title>Deckbox.org card set print view</title>
	</head>
	<body>
					4 Blood Crypt<br/>4 Bloodstained Mire<br/>1 Forest<br/>1 Mountain<br/>1 Nurturing Peatland<br/>1 Overgrown Tomb<br/>1 Stomping Ground<br/>1 Swamp<br/>4 Urza's Saga<br/>3 Verdant Catacombs<br/>2 Wooded Foothills<br/>4 Dragon's Rage Channeler<br/>1 Expedition Map<br/>4 Inquisition of Kozilek<br/>1 Kolaghan's Command<br/>2 Lightning Bolt<br/>4 Mishra's Bauble<br/>1 Nihil Spellbomb<br/>1 Pyrite Spellbomb<br/>4 Ragavan, Nimble Pilferer<br/>1 Shadowspear<br/>4 Tarmogoyf<br/>2 Thoughtseize<br/>4 Unholy Heat<br/>4 Wrenn and Six<br/>
	
			<p><str
`

	_, err := NewDeckboxParser().Parse(input)

	assert.Errorf(t, err, fmt.Sprintf("Decklist is invalid: %s", input))
}
