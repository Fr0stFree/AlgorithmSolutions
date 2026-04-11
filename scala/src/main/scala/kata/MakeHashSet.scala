package kata
import scala.collection.Searching._
import scala.collection.mutable.ArrayBuffer

class MyHashSet() {
  private val storage = ArrayBuffer[Int]()

  def add(key: Int): Unit = {
    storage.search(key) match
      case Found(foundIndex)              => ()
      case InsertionPoint(insertionPoint) => storage.insert(insertionPoint, key)
  }

  def remove(key: Int): Unit = {
    storage.search(key) match
      case Found(foundIndex)              => storage.remove(foundIndex)
      case InsertionPoint(insertionPoint) => ()
  }

  def contains(key: Int): Boolean = {
    storage.search(key) match
      case Found(foundIndex)              => true
      case InsertionPoint(insertionPoint) => false
  }
}

object MakeHashSet extends App {
  val set = new MyHashSet();
  set.add(1)     // set = [1]
  set.add(2)     // set = [1, 2]
  println(set.contains(1)) // return True
  println(set.contains(2)) // return True
  println(set.contains(3)) // return True

}
