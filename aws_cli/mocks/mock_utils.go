package mocks

// Mock a simple composition to wrap around types and introspect calls, return values
type Mock struct {
	calls        map[string][]interface{}
	expectations map[string]struct {
		returnValues []interface{}
	}
}

// MethodExpectations introduced to chain together On and Return Calls
type MethodExpectations struct {
	Mock        // embedded struct
	method      string
	matchedArgs []interface{}
}

func NewMock() Mock {
	return Mock{
		calls: make(map[string][]interface{}),
	}
}

// record a call and return a preset value.
func (m Mock) Called(method string, args ...interface{}) {
	m.recordCall(method, args)
	m.calls[method] = append(m.calls[method], args)
}

func (m Mock) recordCall(method string, args ...interface{}) {
	m.calls[method] = append(m.calls[method], args)
}

// return pre-set values when a method is called.
func (m Mock) On(method string, matchedArgs ...interface{}) MethodExpectations {
	return MethodExpectations{
		m,
		method,
		matchedArgs,
	}
}

func (me MethodExpectations) Return(returnVals ...interface{}) Mock {
	me.expectations[me.method] = struct {
		returnValues []interface{}
	}{
		returnValues: returnVals,
	}
	return me.Mock
}

//
//type PreSignClient struct{}
//
//func (p PreSignClient) PerformOps(a, b int) {
//
//}
//
//type MockPreSignClient struct {
//	Mock // embedded struct.
//}
//
//func (mp MockPreSignClient) PerformOps(a, b int) {
//	mp.Called("PerformOps", a, b)
//}
//
//func main() {
//	mockClient := MockPreSignClient{NewMock()}
//	mockClient.On("PerformOps", 1, 5).Return(5)
//	mockClient.PerformOps(1, 5)
//}
