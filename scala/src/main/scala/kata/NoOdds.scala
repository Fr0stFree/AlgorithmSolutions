package kata

object NoOdds extends App {
    def noOdds(values: List[Int]): List[Int] = values.filter(_ % 2 == 0)

    assert(noOdds(List(0, 1, 2, 3)) == List(0, 2))
}