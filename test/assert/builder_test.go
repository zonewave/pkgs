package assert

func (s *Suite) TestBuilder_NoError() {
	b := NewBuilder(WithTestifySuite(&s.Suite), WithGlobalStopBuildOption())
	b.
		NoError(nil).
		Equal(3, 3).
		Equal(5, 5)
}
