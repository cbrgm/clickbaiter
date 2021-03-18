package clickbaiter

import (
	"math/rand"
)

type Clickbaiter struct {
	nounSubject    string
	nounPerson     []string
	nounPersonType []string
	goodPredicate  []string
	badPredicate   []string
	genericWord    []string
	item           []string
}

func NewClickbaiter() Clickbaiter {
	return Clickbaiter{
		nounSubject: "Coworker",
		nounPerson: []string{
			"Elon Musk",
			"Bill Gates",
			"Angela Merkel",
			"Neil Armstrong",
			"Putin",
			"Trump",
			"Bob Ross",
		},
		nounPersonType: []string{
			"Colleague",
			"Employee",
			"Programmer",
			"Open Source Developer",
			"Maintainer",
			"Reviewer",
			"Issue Creator",
			"Team Member",
			"DevOps Guy",
			"SRE Engineer",
			"Frontend Dev",
			"Data Scientist",
			"CTO",
			"CEO",
			"Org Member",
			"Github Staff",
			"Tech Evangelist",
		},
		goodPredicate: []string{
			"Opening an Issue",
			"Reacting to Comment",
			"Doing Pair Programming",
			"Closing an Issue",
			"Finding True Love",
			"Coming Out",
			"Fixing a Bug",
			"Writing a Unit Test",
			"Writing an Integration Test",
			"Working Hard",
		},
		badPredicate: []string{
			"Deploying to Wrong Cluster Region",
			"Having a Merge Conflict",
			"Deleting Important Code",
			"Failing to Commit",
		},
		genericWord: []string{
			"Tips",
			"Things",
			"Secrets",
			"Shocking News",
			"Tricks",
			"Hints",
		},
		item: []string{
			"Keyboard",
			"Android",
			"Smartphone",
			"Arch Linux",
			"Console",
			"Touchpad",
			"Serverless",
			"Continuous Integration",
			"Continuous Delivery",
			"Continuous Deployment",
		},
	}
}

func (cb *Clickbaiter) NounPerson() string {
	return randomChoice(cb.nounPerson)
}

func (cb *Clickbaiter) NounPersonType() string {
	return randomChoice(cb.nounPersonType)
}

func (cb *Clickbaiter) GoodPredicate() string {
	return randomChoice(cb.goodPredicate)
}

func (cb *Clickbaiter) BadPredicate() string {
	return randomChoice(cb.badPredicate)
}

func (cb *Clickbaiter) GenericWord() string {
	return randomChoice(cb.genericWord)
}

func (cb *Clickbaiter) Item() string {
	return randomChoice(cb.item)
}

func randomChoice(s []string) string {
	return s[rand.Intn(len(s))]
}