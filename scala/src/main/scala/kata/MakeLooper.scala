package kata

object MakeLooper extends App {
  def makeLooper(s: String): () => Char = LazyList.continually(s).flatten.iterator.next

  val looper = makeLooper("abc")
  println(looper())
  println(looper())
  println(looper())
  println(looper())
}
