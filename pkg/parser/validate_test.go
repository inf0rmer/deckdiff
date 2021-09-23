package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDecklistWithValidDeck(t *testing.T) {
	input := `4 Blood Crypt
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

	assert.True(t, ValidateDecklist(input))
}

func TestValidateDecklistWithInvalidDeck(t *testing.T) {
	input := `foobar`

	assert.False(t, ValidateDecklist(input))
}
