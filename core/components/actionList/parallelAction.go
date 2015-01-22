package actionList

func init() {
	// factory.RegisterType((*ParallelAction)(nil), func() interface{} {
	//     return &ParallelAction{}
	// })
}

type actionTracker struct {
	action action
	done   bool
}
type ParallelAction struct {
	actInList
	actions []*actionTracker
}

func InParallel(actions ...action) action {
	p := ParallelAction{}
	for _, act := range actions {
		if act.Next() != nil {
			panic("actions can't have next")
		}
		p.actions = append(p.actions, &actionTracker{action: act})
	}
	return &p
}

func (p *ParallelAction) Run(obj GameObject) action {
	next := true
	for _, act := range p.actions {
		if act.done {
			continue
		}
		if ret := act.action.Run(obj); ret != nil {
			next = false
		} else {
			act.done = true
		}
	}
	if next {
		return p.Next()
	}
	return p
}

func (p *ParallelAction) Unmarshal(data interface{}) {
	m := data.(map[string]interface{})

	actions := []*actionTracker{}
	for _, v := range m["Actions"].([]interface{}) {
		actionData := v.(map[string]interface{})

		typename, _ := actionData["Type"]
		action := factoryFunc(typename.(string)).(action)

		SerializeInPlace(action, actionData)
		actions = append(actions, &actionTracker{action: action})
	}
	p.actions = append(p.actions, actions...)
}
