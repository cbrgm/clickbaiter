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
	generate  *Clickbaiter
	templates []TemplateFunc
}

func NewGenerator() *ClickbaitGenerator {
	gen := &ClickbaitGenerator{
		generate:  NewClickbaiter(),
		templates: []TemplateFunc{},
	}
	gen.AddTemplates([]TemplateFunc{
		func() string {
			return fmt.Sprintf("This %s Is %s Using %s", gen.NounPersonType(), gen.BadPredicate(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("What This %s %s %s Will %s", gen.NounPersonType(), gen.VerbPast(), gen.Item(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("%ss hate This: %s %s About %s That Will %s", gen.NounPersonType(), gen.Number(), gen.GenericWord(), gen.Item(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("This %s Wanted to Make a Point About %s and Succeeded!", gen.NounPersonType(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("%ss Are Mad! %ss Are %s Using %s", gen.NounPersonType(), gen.NounPersonType(), gen.BadPredicate(), gen.Item())
		},
		func() string {
			return fmt.Sprintf("%ss Explain Why We're Seeing a Dramatic Increase In %s Software", gen.NounPersonType(), gen.VerbIng())
		},
		func() string {
			return fmt.Sprintf("%s Thanks %ss For %s Great %s! It Will %s", gen.NounPerson(), gen.NounPersonType(), gen.VerbIng(), gen.GenericWord(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("Why %s Becomes Important! What %s Is Talking About With %ss Will %s", gen.Item(), gen.NounPerson(), gen.NounPersonType(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("%s: %ss Are %s %s. What Happened Next Will %s", gen.Attention(), gen.NounPersonType(), gen.VerbIng(), gen.Item(), gen.Reaction())
		},
		func() string {
			return fmt.Sprintf("%s: %s As A Service is On The Rise! Read On and It Will %s", gen.Attention(), gen.Item(), gen.Reaction())
		},
	})
	return gen
}

func (cbg *ClickbaitGenerator) RandomSentence() string {
	f := randomFunc(cbg.templates)
	return f()
}

func (cbg *ClickbaitGenerator) AddTemplate(s TemplateFunc) {
	cbg.templates = append(cbg.templates, s)
}

func (cbg *ClickbaitGenerator) AddTemplates(se []TemplateFunc) {
	for _, s := range se {
		cbg.templates = append(cbg.templates, s)
	}
}

func (cbg *ClickbaitGenerator) UseHashtags(b bool) {
	cbg.generate.useHashtags = b
}

func (cbg *ClickbaitGenerator) NounPerson() string {
	return cbg.generate.NounPerson()
}

func (cbg *ClickbaitGenerator) NounPersonType() string {
	return cbg.generate.NounPersonType()
}

func (cbg *ClickbaitGenerator) GoodPredicate() string {
	return cbg.generate.GoodPredicate()
}

func (cbg *ClickbaitGenerator) BadPredicate() string {
	return cbg.generate.BadPredicate()
}

func (cbg *ClickbaitGenerator) GenericWord() string {
	return cbg.generate.GenericWord()
}

func (cbg *ClickbaitGenerator) VerbPast() string {
	return cbg.generate.VerbPast()
}

func (cbg *ClickbaitGenerator) VerbIng() string {
	return cbg.generate.VerbIng()
}

func (cbg *ClickbaitGenerator) Reaction() string {
	return cbg.generate.Reaction()
}

func (cbg *ClickbaitGenerator) Item() string {
	return cbg.generate.Item()
}

func (cbg *ClickbaitGenerator) Number() string {
	return cbg.generate.Number()
}

func (cbg *ClickbaitGenerator) Attention() string {
	return cbg.generate.Attention()
}
