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
	verbPast       []string
	verbIng        []string
	genericWord    []string
	item           []string
	reaction       []string
	number         []string
	attention      []string
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
			"MontanaBlack",
			"Joe Biden",
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
			"Tech Evangelist",
			"Junior Dev",
			"Senior Dev",
			"Tester",
			"Tech Influencer",
			"Contributor",
			"Recruiter",
			"Product Owner",
			"Scrum Master",
			"Operator",
			"Admin",
			"Team Lead",
			"Mobile Dev",
			"Twitch Streamer",
			"Android Developer",
			"iOS Developer",
			"Cloud Native",
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
			"Deploying to Wrong Kubernetes Clusters",
			"Merging Wrong Branches",
			"Copying From Crappy Blog Posts",
			"Destroying Infra",
			"Nuking Virtual Machines",
			"Having a Merge Conflict",
			"Deleting some Important Code",
			"Failing to Commit",
			"Deleting a Feature Branch",
			"Copying from StackOverflow",
			"Destroying Test Frameworks",
			"Deleting User Accounts",
			"Shattering Unix Permissions",
		},
		verbPast: []string{
			"Wrote about",
			"Fixed in",
			"Coded in",
			"Hacked in",
			"Achieved in",
			"Edited in",
			"Mentioned in",
			"Commented in",
			"Merged into",
			"Added to",
			"Contributed to",
			"Committed to",
			"Revealed in",
			"Exposed to",
			"Discovered in",
			"Unmasked in",
			"Announced in",
			"Believed in",
		},
		verbIng: []string{
			"Writing",
			"Fixing",
			"Coding",
			"Hacking",
			"Achieving",
			"Editing",
			"Mentioning",
			"Commenting",
			"Merging",
			"Committing",
			"Adding",
			"Contributing",
			"Revealing",
			"Exposing",
			"Discovering",
			"Unmasking",
			"Announcing",
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
			"Keyboards",
			"Android",
			"Smartphones",
			"Arch Linux",
			"Ubuntu",
			"OpenBSD",
			"Kali Linux",
			"Windows 10",
			"Raspbian",
			"CLIs",
			"Touchpads",
			"Serverless",
			"Cloud Computing",
			"IoT",
			"Machine Learning",
			"Continuous Integration",
			"Continuous Delivery",
			"Continuous Deployment",
			"Cloud Computing",
			"Service Meshes",
			"Containers",
			"Deployments",
			"Go",
			"Rust",
			"Javascript",
			"Python",
			"Bash",
			"Pure CSS",
			"Java",
			"Kotlin",
			"Elixir",
			"TypeScript",
			"Dart",
			"Swift",
			"C#",
			"C",
			"Prometheus",
			"Kubernetes",
			"Slack",
			"Zoom",
			"Microsoft Teams",
			"Helm",
			"React",
			"Vue.js",
			"Svelte",
			"Vi",
			"Vim",
			"Atom",
			"Infrastructure As a Service",
			"Software As A Service",
			"Platform As A Service",
			"Container As A Service",
			"NoCode",
			"GitOps",
			"NoOps",
			"Hybrid Cloud",
			"Community Cloud",
			"Private Cloud",
			"Multi Cloud",
			"Distributed Cloud",
			"Waterfall Model",
			"Spotify Model",
			"Kanban",
			"Software Engineering",
			"Ansible",
			"Puppet",
			"Helm",
			"Terraform",
			"MacOS",
			"Domain Driven Design",
			"Bounded Contexts",
			"Ubiquitous Language",
		},
		reaction: []string{
			"Blow Your Mind",
			"Make You Cry",
			"Take Your Breath",
			"Leave You Speechless",
			"Make You Laugh",
			"Give you a Heart Attack",
			"Change Your Life",
			"Shock You",
			"Get You Excited",
			"Make You Fall in Love",
			"Make You Quit your Job",
			"Make You Change Your Workflow",
			"Get You Mad",
			"Fascinate You",
			"Make your Boss Love You",
			"Make Doctors Hate You",
			"Make Researchers Hate You",
			"Make You A Startup Founder",
			"Double Your Monthly Income",
			"Double Your Daily Income",
			"Double Your Annual Income",
		},
		number: []string{
			"Two",
			"Three",
			"Four",
			"Five",
			"Six",
			"Seven",
			"Eight",
			"Nine",
			"Ten",
			"Eleven",
			"Twelve",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
			"10",
			"11",
			"12",
		},
		attention: []string{
			"Breaking",
			"Hot",
			"Shocking",
			"Urgent",
			"Hey you",
			"It's time to talk about it",
			"Listen",
			"Attention",
			"Update",
			"Finally",
			"Everyone talks about it",
			"Trending",
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

func (cb *Clickbaiter) VerbPast() string {
	return randomChoice(cb.verbPast)
}

func (cb *Clickbaiter) VerbIng() string {
	return randomChoice(cb.verbIng)
}

func (cb *Clickbaiter) Reaction() string {
	return randomChoice(cb.reaction)
}

func (cb *Clickbaiter) Number() string {
	return randomChoice(cb.number)
}

func (cb *Clickbaiter) Attention() string {
	return randomChoice(cb.attention)
}

func randomChoice(s []string) string {
	return s[rand.Intn(len(s))]
}
