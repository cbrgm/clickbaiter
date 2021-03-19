package clickbaiter

import (
	"fmt"
	"math/rand"
)

type TemplateFunc func() string

func randomFunc(f []TemplateFunc) TemplateFunc {
	return f[rand.Intn(len(f))]
}

type ClickbaitGenerator struct {
	words     Clickbaiter
	sentences []TemplateFunc
}

func NewGenerator() ClickbaitGenerator {
	gen := ClickbaitGenerator{
		words:     NewClickbaiter(),
		sentences: []TemplateFunc{},
	}
	gen.AddTemplates([]TemplateFunc{
		func() string {
			return fmt.Sprintf("This %s is %s Using %s", gen.NounPersonType(), gen.BadPredicate(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("What This %s %s %s will %s", gen.NounPersonType(), gen.VerbPast(), gen.Item(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("%ss hate This: %s %s about %s that will %s", gen.NounPersonType(), gen.Number(), gen.GenericWord(), gen.Item(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("This %s Wanted to Make a Point About %s and Succeeded!", gen.NounPersonType(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("%ss are mad! %ss are %s Using %s", gen.NounPersonType(), gen.NounPersonType(), gen.BadPredicate(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("%ss Explain Why We're Seeing a Dramatic Increase in %s Software", gen.NounPersonType(), gen.VerbIng())
		},
		func() string {
			return fmt.Sprintf("%s Thanks %ss For %s Great %s! It will %s", gen.NounPerson(), gen.NounPersonType(), gen.VerbIng(), gen.GenericWord(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("Why %s Becomes Important! What %s is Talking About with %ss Will %s", gen.Item(), gen.NounPerson(), gen.NounPersonType(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("%s: %ss are %s %s. What happened next will %s", gen.Attention(), gen.NounPersonType(), gen.VerbIng(), gen.Item(), gen.Reaction())
		},


	})
	return gen
}

func (cbg *ClickbaitGenerator) RandomSentence() string {
	f := randomFunc(cbg.sentences)
	return f()
}

func (cbg *ClickbaitGenerator) AddTemplate(s TemplateFunc) {
	cbg.sentences = append(cbg.sentences, s)
}

func (cbg *ClickbaitGenerator) AddTemplates(se []TemplateFunc) {
	for _, s := range se {
		cbg.sentences = append(cbg.sentences, s)
	}
}

func (cbg *ClickbaitGenerator) NounPerson() string {
	return cbg.words.NounPerson()
}

func (cbg *ClickbaitGenerator) NounPersonType() string {
	return cbg.words.NounPersonType()
}

func (cbg *ClickbaitGenerator) GoodPredicate() string {
	return cbg.words.GoodPredicate()
}

func (cbg *ClickbaitGenerator) BadPredicate() string {
	return cbg.words.BadPredicate()
}

func (cbg *ClickbaitGenerator) GenericWord() string {
	return cbg.words.GenericWord()
}

func (cbg *ClickbaitGenerator) VerbPast() string {
	return cbg.words.VerbPast()
}

func (cbg *ClickbaitGenerator) VerbIng() string {
	return cbg.words.VerbIng()
}

func (cbg *ClickbaitGenerator) Reaction() string {
	return cbg.words.Reaction()
}

func (cbg *ClickbaitGenerator) Item() string {
	return cbg.words.Item()
}

func (cbg *ClickbaitGenerator) Number() string {
	return cbg.words.Number()
}

func (cbg *ClickbaitGenerator) Attention() string {
	return cbg.words.Attention()
}


