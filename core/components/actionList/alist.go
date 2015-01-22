package actionList

func init() {
	// factory.RegisterType((*ActionList)(nil), func() interface{} {
	// 	return &ActionList{}
	// })
}

type actInList struct {
	next action
}

func (a *actInList) Next() action {
	return a.next
}
func (a *actInList) SetNext(act action) {
	a.next = act
}

type action interface {
	SetNext(action)
	Next() action
	Run(GameObject) action
}

type ActionList struct {
	TypeId
	OwnerMngr
	Sequence action
}

func (a *ActionList) Enqueue(actions ...action) {
	tmp := []action{}
	for _, act := range actions {
		tmp = append(tmp, act)
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		act := tmp[i]
		act.SetNext(a.Sequence)
		a.Sequence = act
	}
}
func (a *ActionList) Run() {
	if a.Sequence != nil {
		a.Sequence = a.Sequence.Run(a.Owner())
	}
}
func (a *ActionList) IsFinished() bool {
	return a.Sequence == nil
}

func (a *ActionList) Unmarshal(data interface{}) {
	m := data.(map[string]interface{})

	actions := []action{}
	for _, v := range m["Sequence"].([]interface{}) {
		actionData := v.(map[string]interface{})

		typename, _ := actionData["Type"]
		action := factoryFunc(typename.(string)).(action)

		SerializeInPlace(action, actionData)
		actions = append(actions, action)
	}
	a.Enqueue(actions...)
}
