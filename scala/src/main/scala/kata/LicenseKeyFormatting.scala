package kata

class LicenseKeyFormatting extends App {
  def licenseKeyFormatting(s: String, k: Int) = {
    s.filter(_ != '-')
      .toUpperCase
      .reverse
      .grouped(k)
      .mkString("-")
      .reverse
  }
  println(licenseKeyFormatting("A5F3Z-2e-9-w", k = 4))
}
