package trac

// Represents an issue with the test.
type Problem struct {
	Id                 string              `json:"id,omitempty"`
	Type_              string              `json:"type,omitempty"`
	Identity           string              `json:"identity,omitempty"`
	Href               string              `json:"href,omitempty"`
	Description        string              `json:"description,omitempty"`
	Mutes              *Mutes              `json:"mutes,omitempty"`
	Investigations     *Investigations     `json:"investigations,omitempty"`
	ProblemOccurrences *ProblemOccurrences `json:"problemOccurrences,omitempty"`
	Locator            string              `json:"locator,omitempty"`
}

// Represents a paginated list of Problem entities.
type Problems struct {
	Count    int32     `json:"count,omitempty"`
	NextHref string    `json:"nextHref,omitempty"`
	PrevHref string    `json:"prevHref,omitempty"`
	Problem  []Problem `json:"problem,omitempty"`
}

// Represents a paginated list of ProblemOccurrence entities.
type ProblemOccurrences struct {
	Count             int32               `json:"count,omitempty"`
	Href              string              `json:"href,omitempty"`
	NextHref          string              `json:"nextHref,omitempty"`
	PrevHref          string              `json:"prevHref,omitempty"`
	ProblemOccurrence []ProblemOccurrence `json:"problemOccurrence,omitempty"`
	Passed            int32               `json:"passed,omitempty"`
	Failed            int32               `json:"failed,omitempty"`
	NewFailed         int32               `json:"newFailed,omitempty"`
	Ignored           int32               `json:"ignored,omitempty"`
	Muted             int32               `json:"muted,omitempty"`
}

// Represents an instance of a failed test in the specific build.
type ProblemOccurrence struct {
	Id                    string   `json:"id,omitempty"`
	Type_                 string   `json:"type,omitempty"`
	Identity              string   `json:"identity,omitempty"`
	Href                  string   `json:"href,omitempty"`
	Muted                 bool     `json:"muted,omitempty"`
	CurrentlyInvestigated bool     `json:"currentlyInvestigated,omitempty"`
	CurrentlyMuted        bool     `json:"currentlyMuted,omitempty"`
	LogAnchor             string   `json:"logAnchor,omitempty"`
	NewFailure            bool     `json:"newFailure,omitempty"`
	Details               string   `json:"details,omitempty"`
	AdditionalData        string   `json:"additionalData,omitempty"`
	Problem               *Problem `json:"problem,omitempty"`
	Mute                  *Mute    `json:"mute,omitempty"`
	Build                 *Build   `json:"build,omitempty"`
}

// Represents a test results counter (how many times this test was successful/failed/muted/ignored).
type TestCounters struct {
	Ignored   int32 `json:"ignored,omitempty"`
	Failed    int32 `json:"failed,omitempty"`
	Muted     int32 `json:"muted,omitempty"`
	Success   int32 `json:"success,omitempty"`
	All       int32 `json:"all,omitempty"`
	NewFailed int32 `json:"newFailed,omitempty"`
	Duration  int64 `json:"duration,omitempty"`
}

// Represents a relation between a test and the specific build.
type TestOccurrence struct {
	Id                    string           `json:"id,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Status                string           `json:"status,omitempty"`
	Ignored               bool             `json:"ignored,omitempty"`
	Duration              int32            `json:"duration,omitempty"`
	RunOrder              string           `json:"runOrder,omitempty"`
	NewFailure            bool             `json:"newFailure,omitempty"`
	Muted                 bool             `json:"muted,omitempty"`
	CurrentlyMuted        bool             `json:"currentlyMuted,omitempty"`
	CurrentlyInvestigated bool             `json:"currentlyInvestigated,omitempty"`
	Href                  string           `json:"href,omitempty"`
	IgnoreDetails         string           `json:"ignoreDetails,omitempty"`
	Details               string           `json:"details,omitempty"`
	Test                  *Test            `json:"test,omitempty"`
	Mute                  *Mute            `json:"mute,omitempty"`
	Build                 *Build           `json:"build,omitempty"`
	FirstFailed           *TestOccurrence  `json:"firstFailed,omitempty"`
	NextFixed             *TestOccurrence  `json:"nextFixed,omitempty"`
	Invocations           *TestOccurrences `json:"invocations,omitempty"`
	Metadata              *TestRunMetadata `json:"metadata,omitempty"`
	LogAnchor             string           `json:"logAnchor,omitempty"`
}

// Represents a paginated list of TestOccurrence entities.
type TestOccurrences struct {
	Count          int32            `json:"count,omitempty"`
	Href           string           `json:"href,omitempty"`
	NextHref       string           `json:"nextHref,omitempty"`
	PrevHref       string           `json:"prevHref,omitempty"`
	TestOccurrence []TestOccurrence `json:"testOccurrence,omitempty"`
	TestCounters   *TestCounters    `json:"testCounters,omitempty"`
	Failed         int32            `json:"failed,omitempty"`
	NewFailed      int32            `json:"newFailed,omitempty"`
	Muted          int32            `json:"muted,omitempty"`
	Passed         int32            `json:"passed,omitempty"`
	Ignored        int32            `json:"ignored,omitempty"`
}

// Represents a list of TypedValue entities.
type TestRunMetadata struct {
	Count       int32        `json:"count,omitempty"`
	TypedValues []TypedValue `json:"typedValues,omitempty"`
}

// Represents a test.
type Test struct {
	Id              string           `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Mutes           *Mutes           `json:"mutes,omitempty"`
	Investigations  *Investigations  `json:"investigations,omitempty"`
	TestOccurrences *TestOccurrences `json:"testOccurrences,omitempty"`
	ParsedTestName  *ParsedTestName  `json:"parsedTestName,omitempty"`
	Href            string           `json:"href,omitempty"`
	Locator         string           `json:"locator,omitempty"`
}

// Represents a paginated list of Test entities.
type Tests struct {
	Count          int32         `json:"count,omitempty"`
	MyTestCounters *TestCounters `json:"myTestCounters,omitempty"`
	NextHref       string        `json:"nextHref,omitempty"`
	PrevHref       string        `json:"prevHref,omitempty"`
	Test           []Test        `json:"test,omitempty"`
}

// Represents a muted test.
type Mute struct {
	Id         int32          `json:"id,omitempty"`
	Href       string         `json:"href,omitempty"`
	Assignment *Comment       `json:"assignment,omitempty"`
	Scope      *ProblemScope  `json:"scope,omitempty"`
	Target     *ProblemTarget `json:"target,omitempty"`
	Resolution *Resolution    `json:"resolution,omitempty"`
}

// Represents a paginated list of Mute entities.
type Mutes struct {
	Count    int32  `json:"count,omitempty"`
	NextHref string `json:"nextHref,omitempty"`
	PrevHref string `json:"prevHref,omitempty"`
	Href     string `json:"href,omitempty"`
	Mute     []Mute `json:"mute,omitempty"`
}

// Represents test metadata (package, method, class, and so on).
type ParsedTestName struct {
	TestPackage            string `json:"testPackage,omitempty"`
	TestSuite              string `json:"testSuite,omitempty"`
	TestClass              string `json:"testClass,omitempty"`
	TestShortName          string `json:"testShortName,omitempty"`
	TestNameWithoutPrefix  string `json:"testNameWithoutPrefix,omitempty"`
	TestMethodName         string `json:"testMethodName,omitempty"`
	TestNameWithParameters string `json:"testNameWithParameters,omitempty"`
}

// Represents a name-value-type relation.
type TypedValue struct {
	Name  string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// Represents an investigation resolution timestamp and details.
type Resolution struct {
	Type_ string `json:"type,omitempty"`
	Time  string `json:"time,omitempty"`
}
