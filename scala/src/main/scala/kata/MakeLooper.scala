package kata

object MakeLooper extends App {
  def makeLooper(s: String): () => Char = {
    val cycle = LazyList.continually(s).flatten.iterator
    () => cycle.next()
  }

  val looper = makeLooper("abc")
  println(looper())
  println(looper())
  println(looper())
  println(looper())
}
