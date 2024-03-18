function main() {
  if [ -z "$1" ]; then
    echo "one for you and one for me"
  else
    echo "one for $1 and one for me"
  fi
}

main "$@"