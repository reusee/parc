package parc

import "testing"

func BenchmarkMatch(b *testing.B) {
	grammar := NewGrammar("S")
	grammar.Rule("S", Named("S", Alt(
		Named("Sa", Seq(
			Named("S1", N("A")),
			Named("S2", N("S")),
			Named("S3", T(ByteEq('d'))),
		)),
		Named("Sb", Seq(
			Named("S4", N("B")),
			Named("S5", N("S")),
		)),
		Named("Sc", Seq()),
	)))
	grammar.Rule("A", Named("A", T(ByteIn([]byte{'a', 'c'}))))
	grammar.Rule("B", Named("B", T(ByteIn([]byte{'a', 'b'}))))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grammar.Match([]byte("aab"))
	}
}
