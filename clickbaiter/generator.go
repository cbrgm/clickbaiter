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
			return fmt.Sprintf("This %s is Taking on %s Using %s", gen.NounPersonType(), gen.BadPredicate(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("This %s is Taking on %s Using %s", gen.NounPersonType(), gen.BadPredicate(), gen.Item())
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

func (cbg *ClickbaitGenerator) Item() string {
	return cbg.words.Item()
}
