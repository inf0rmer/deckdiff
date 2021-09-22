package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckboxParser(t *testing.T) {
	input := NewDeckboxParser().Parse(`<!DOCTYPE html>
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
`)

	expected := `4 Blood Crypt
4 Bloodstained Mire
1 Forest
1 Mountain
1 Nurturing Peatland
1 Overgrown Tomb
1 Stomping Ground
1 Swamp
4 Urza's Saga
3 Verdant Catacombs
2 Wooded Foothills
4 Dragon's Rage Channeler
1 Expedition Map
4 Inquisition of Kozilek
1 Kolaghan's Command
2 Lightning Bolt
4 Mishra's Bauble
1 Nihil Spellbomb
1 Pyrite Spellbomb
4 Ragavan, Nimble Pilferer
1 Shadowspear
4 Tarmogoyf
2 Thoughtseize
4 Unholy Heat
4 Wrenn and Six

2 Alpine Moon
1 Chalice of the Void
2 Collective Brutality
2 Engineered Explosives
1 Ghost Quarter
1 Kolaghan's Command
1 Lurrus of the Dream-Den
1 Pithing Needle
2 Terminate
1 Torpor Orb
1 Void Mirror`

	assert.Equal(t, input, expected)
}
