package kata

object RemoveMarked extends App {
  implicit class BetterList[A](l: List[A]) {
    def removeMarked(valuesList: List[A]) = l.filterNot(valuesList.contains)
  }

  println(BetterList(List(1, 1, 2 ,3 ,1 ,2 ,3 ,4)).removeMarked(List(1, 3)))
}